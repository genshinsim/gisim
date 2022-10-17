import { pool } from "./Pages/Sim";
import { SimResults } from "./Pages/Viewer/SimResults";
import { ParsedResult } from "./types";
import { Aggregator, SimWorker } from "./Workers/common";

export class WorkerPool {
  private aggregator: Worker;
  private aggregatorReady: boolean;
  private workers: Worker[];
  private workersReady: boolean[];

  constructor() {
    this.aggregatorReady = false;
    this.aggregator = this.createAggregator();
    this.workers = [];
    this.workersReady = [];
  }

  public count(): number {
    return this.loaded().length;
  }

  public ready(): boolean {
    return this.aggregatorReady && this.count() > 0;
  }

  private loaded(): Worker[] {
    return this.workers.filter((_, i) => this.workersReady[i]);
  }

  private createAggregator(): Worker {
    this.aggregatorReady = false;
    const out = new Worker(
        new URL("./Workers/aggregator.ts", import.meta.url), { name: "aggregator" });
    out.onmessage = (ev) => {
      switch (ev.data.type as Aggregator.Response) {
        case Aggregator.Response.Ready:
          this.aggregatorReady = true;
          break;
      }
    };
    return out;
  }

  private createWorker(): Promise<number> {
    return new Promise((resolve, reject) => {
      const worker = new Worker(
          new URL("./Workers/worker.ts", import.meta.url),
          { name: "worker " + this.workers.length });
      const idx = this.workers.push(worker) - 1;
      this.workersReady.push(false);
      worker.onmessage = (ev) => {
        switch (ev.data.type as SimWorker.Response) {
          case SimWorker.Response.Ready:
            this.workersReady[idx] = true;
            resolve(idx);
            return;
          case SimWorker.Response.Failed:
            reject("Worker " + idx + " " + (ev.data as SimWorker.FailedResponse).reason);
            return;
          default:
            reject("Worker " + idx + " - unknown response: " + ev.data);
        }
      };
    });
  }

  public setWorkerCount(count: number, readycb: (count: number) => void) {
    console.log("loading workers", count, this);
    const diff = count - this.workers.length;

    if (diff < 0) {
      this.workersReady.splice(diff);
      this.workers.splice(diff).forEach((w) => w.terminate());
      console.log(pool);
      return readycb(count);
    }

    console.log("loading " + diff + " workers");
    for (let i = 0; i < diff; i++) {
      this.createWorker().then((w) => {
        console.log("worker " + w + " is now ready");
        readycb(this.count());
      });
    }
  }

  public validate(cfg: string): Promise<ParsedResult> {
    return new Promise((resolve, reject) => {
      this.aggregator.onmessage = (ev) => {
        switch (ev.data.type as Aggregator.Response) {
          case Aggregator.Response.Validate:
            resolve((ev.data as Aggregator.ValidateResponse).cfg);
            return;
          case Aggregator.Response.Failed:
            reject((ev.data as Aggregator.FailedResponse).reason);
            return;
          default:
            reject("unknown validate response: " + ev.data);
        }
      };
      this.aggregator.postMessage(Aggregator.ValidateRequest(cfg));
    });
  }

  public run(cfg: string, setResult: (result: SimResults) => void): Promise<boolean | void> {
    if (!this.ready()) {
      return Promise.reject("aggregators and/or workers are not ready!");
    }

    const startTime = Date.now() * 1_000_000;
    let result: SimResults | null = null;
    let maxIterations = 0;

    const initPromises = [];
    // 1. initialize aggregator
    initPromises.push(new Promise<boolean>((resolve, reject) => {
      this.aggregator.onmessage = (ev) => {
        switch (ev.data.type as Aggregator.Response) {
          case Aggregator.Response.Initialized:
            result = (ev.data as Aggregator.InitializeResponse).result;
            maxIterations = result?.max_iterations ?? 1000;
            resolve(true);
            return;
          case Aggregator.Response.Failed:
            reject((ev.data as Aggregator.FailedResponse).reason);
            return;
        }
      };
      this.aggregator.postMessage(Aggregator.InitializeRequest(cfg));
    }));

    // 2. initialize all workers
    this.loaded().forEach((worker) => {
      initPromises.push(new Promise<boolean>((resolve, reject) => {
        worker.onmessage = (ev) => {
          switch (ev.data.type as SimWorker.Response) {
            case SimWorker.Response.Initialized:
              resolve(true);
              return;
            case SimWorker.Response.Failed:
              reject((ev.data as Aggregator.FailedResponse).reason);
              return;
          }
        };
        worker.postMessage(SimWorker.InitializeRequest(cfg));
      }));
    });

    // 3. after all initializes complete, start execution
    return Promise.all(initPromises).then(() => {
      let completed = 0;
      let requested = 0;
      this.aggregator.onmessage = (ev) => {
        switch (ev.data.type as Aggregator.Response) {
          case Aggregator.Response.Result:
            const out = Object.assign({}, result);
            out.statistics = (ev.data as Aggregator.ResultResponse).result;
            setResult(out);
            if (completed >= maxIterations) {
              Promise.resolve(true);
            }
            return;
          case Aggregator.Response.Done:
            completed += 1;
            if (completed == 1 || completed % 10 == 0 || completed >= maxIterations) {
              this.aggregator.postMessage(Aggregator.FlushRequest(startTime));
            }
            return;
          case Aggregator.Response.Failed:
            throw (ev.data as Aggregator.FailedResponse).reason;
        }
      };

      this.loaded().forEach((worker) => {
        worker.onmessage = (ev) => {
          switch (ev.data.type as SimWorker.Response) {
            case SimWorker.Response.Done:
              const resp: SimWorker.RunResponse = ev.data;
              this.aggregator.postMessage(Aggregator.AddRequest(resp.result));
              if (requested < maxIterations) {
                worker.postMessage(SimWorker.RunRequest(requested++));
              }
              return;
            case SimWorker.Response.Failed:
              throw (ev.data as Aggregator.FailedResponse).reason;
          }
        };

        if (requested < maxIterations) {
          worker.postMessage(SimWorker.RunRequest(requested++));
        }
      });
    });
  }

  public cancel() {
    console.log("execution canceled");
    this.workers.forEach((worker) => {
      worker.onmessage = null;
    });

    // It is possible that there are N AddRequests in the aggregator queue that we have no control
    // over. Even if we set the onmessage here to null, the aggregator will still process through
    // all N requests. Since there is no way to clear the worker queue, recreating the worker is the
    // next best thing.
    //
    // Downside of this approach is any memory allocation/optimizations from previous runs will not
    // carry over, making executions after a cancel "less optimal".
    this.aggregator.terminate();
    this.aggregator = this.createAggregator();
  }

  public reset() {
    console.log("restarting all workers");
    const cnt = this.workers.length;
    this.setWorkerCount(0, () => {});
    this.setWorkerCount(cnt, () => {});

    this.aggregator.terminate();
    this.aggregator = this.createAggregator();
  }
}
