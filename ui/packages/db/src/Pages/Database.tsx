import { Spinner } from "@blueprintjs/core";
import { db } from "@gcsim/types";
import axios from "axios";
import { useEffect, useReducer, useState } from "react";
import { Filter } from "../SharedComponents/Filter";
import {
  CharFilter,
  FilterContext,
  FilterDispatchContext,
  filterReducer,
  initialCharFilter,
  ItemFilterState,
} from "../SharedComponents/FilterComponents/Filter.utils";
import { ListView } from "../SharedComponents/ListView";

export function Database() {
  const [filter, dispatch] = useReducer(filterReducer, {
    charFilter: initialCharFilter,
    charIncludeCount: 0,
  });

  const [data, setData] = useState<db.IEntry[]>([]);

  const querydb = (query: DbQuery) => {
    axios(`/api/db?q=${encodeURIComponent(JSON.stringify(query))}`)
      .then(
        (resp: {
          data: {
            data: db.IEntry[];
          };
        }) => {
          if (resp.data) {
            // console.log("output: ", resp.data);
            setData(resp.data.data);
          } else {
            console.log("no data: ", resp.data);
          }
        }
      )
      .catch((err) => {
        console.log("error: ", err);
      });
  };

  useEffect(() => {
    const query = craftQuery(filter);
    // console.log("input", query);
    querydb(query);
  }, [filter]);

  return (
    <FilterContext.Provider value={filter}>
      <FilterDispatchContext.Provider value={dispatch}>
        <div className="flex flex-col  gap-4 m-8 my-4 ">
          <div className="flex flex-row justify-between items-center">
            <Filter />
            <div className="text-base  md:text-2xl">{`Showing ${
              data?.length ?? 0
            } Simulations `}</div>
            {/* <Sorter /> */}
          </div>
          {data ? <ListView data={data} /> : <Spinner />}
        </div>
      </FilterDispatchContext.Provider>
    </FilterContext.Provider>
  );
}

function craftQuery({ charFilter }: { charFilter: CharFilter }): DbQuery {
  const query: DbQuery["query"] = {};
  // sort all characters into included and excluded from the filter
  const includedChars: string[] = [];
  const excludedChars: string[] = [];
  for (const [charName, charState] of Object.entries(charFilter)) {
    if (charState.state === ItemFilterState.include) {
      includedChars.push(charName);
    } else if (charState.state === ItemFilterState.exclude) {
      excludedChars.push(charName);
    }
  }
  if (includedChars.length > 0) {
    query["summary.char_names"] = {};
    query["summary.char_names"]["$all"] = includedChars;
  }
  if (excludedChars.length > 0) {
    query["summary.char_names"] = query["summary.char_names"] ?? {};
    query["summary.char_names"]["$nin"] = excludedChars;
  }
  return { query, limit: 20 };
}

interface DbQuery {
  query: {
    "summary.char_names"?: {
      $all?: string[];
      $nin?: string[];
    };
  };
  limit?: number;
  sort?: unknown;
  skip?: number;
}
