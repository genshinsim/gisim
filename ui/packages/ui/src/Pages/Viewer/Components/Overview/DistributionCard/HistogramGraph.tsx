import { Colors } from "@blueprintjs/core";
import { SummaryStat } from "@gcsim/types";
import { AxisBottom, AxisLeft, TickRendererProps } from "@visx/axis";
import { GridRows } from "@visx/grid";
import { Group } from "@visx/group";
import { scaleBand, scaleLinear } from "@visx/scale";
import { BoxPlot } from "@visx/stats";
import { Text } from "@visx/text";
import { useTooltip } from "@visx/tooltip";
import { range } from "lodash-es";
import { useMemo } from "react";
import { useTranslation } from "react-i18next";
import { DataColors, NoData } from "../../Util";
import { RenderTooltip, TooltipData, useTooltipHandles } from "./Tooltip";
import { VerticalLine } from "./VerticalLine";

type Props = {
  width: number;
  height: number;
  margin?: { left: number, right: number, top: number, bottom: number };

  data?: SummaryStat;
  barColor?: string;
  hoverColor?: string;
  accentColor?: string;
}

const defaultMargin = { left: 76, right: 1, top: 5, bottom: 16 };
const boxPlotHeight = 30

export const HistogramGraph = ({
      width,
      height,
      margin = defaultMargin,
      data,
      barColor = Colors.SEPIA3,
      hoverColor = Colors.SEPIA5,
      accentColor = Colors.SEPIA1,
    }: Props) => {
  const xMax = width - margin.left - margin.right;
  const yMax = height - margin.top - margin.bottom - boxPlotHeight;
  const numTicks = 7;
  const numHorizontalTicks = Math.min(width / 80, 12);
  
  const { i18n } = useTranslation();
  const { xScale, yScale, xLin, delta } = useScales(data, xMax, yMax);
  const tooltip = useTooltip<TooltipData>();
  const tooltipHandles = useTooltipHandles(
      tooltip.showTooltip, tooltip.hideTooltip, delta, margin);

  if (data?.histogram == null || delta == null) {
    return <NoData />;
  }

  return (
    <div className="relative">
      <svg
          width={width}
          height={height}
          onMouseMove={(e) => tooltipHandles.mouseHover(e, data)}
          onMouseLeave={() => tooltipHandles.mouseLeave()}>
        <Group left={margin.left} top={margin.top}>
          <GridRows
              scale={yScale}
              numTicks={numTicks}
              lineStyle={{ opacity: 0.5 }}
              stroke={DataColors.gray}
              width={xMax}
              height={height} />
          <AxisLeft
              hideAxisLine
              hideTicks
              scale={yScale}
              numTicks={numTicks}
              labelOffset={55}
              labelClassName="fill-gray-400 text-lg"
              tickClassName="fill-gray-400 font-mono text-xs"
              tickComponent={(props) => <TickLabel {...props} />}
                tickFormat={s =>
                  s.toLocaleString(i18n.language,
                      { notation: 'compact', maximumSignificantDigits: 3 })
                }
                label="# iterations" />
          <VerticalLine
              x={data?.mean}
              xScale={xLin}
              yMax={yMax}
              color={accentColor}
              className="opacity-75 fill-gray-400 font-mono" />
          <BoxPlot
              valueScale={xLin}
              min={data.min}
              max={data.max}
              top={yMax+boxPlotHeight}
              firstQuartile={data.q1}
              median={data.q2}
              thirdQuartile={data.q3}
              horizontal={true}
              boxWidth={10}
              fill={hoverColor}
              fillOpacity={0.1}
              stroke={hoverColor}
              strokeWidth={1}
              medianProps={{ style: { stroke: hoverColor } }} />
          <AxisBottom
              hideAxisLine
              top={yMax}
              scale={xScale}
              numTicks={numHorizontalTicks}
              labelClassName="fill-gray-400 text-lg"
              tickClassName="fill-gray-400 font-mono text-xs"
              tickStroke={DataColors.gray}
              tickComponent={(props) => <TickLabelHorizontal {...props} />}
                tickFormat={s => {
                  const histogramLength = data.histogram?.length ?? 1
                  const min = data?.min ?? 0
                  const max = data?.max ?? 0
                  const value = s / histogramLength * (max - min) + min
                  return value.toLocaleString(i18n.language,
                    { notation: 'compact', maximumSignificantDigits: 3 })
                  }
                }
                 />
          {data.histogram.map((c, i) => {
            const barWidth = xScale.bandwidth();
            const barHeight = yMax - yScale(c);
            const barX = xScale(i) ?? 0;
            const barY = yMax - barHeight;

            if (c <= 0 || barHeight < 0 || data.mean == null || data.min == null) {
              return null;
            }

            const x = tooltip.tooltipData?.x ?? 0;
            const temp = xLin.invert(x - margin.left);
            const idx = Math.max(Math.floor(delta * (temp - data.min)), 0);

            let fill = barColor;
            if (i === Math.floor(delta * (data.mean - data.min))) {
              fill = accentColor;
            }
            if (tooltip.tooltipData != null && i === idx) {
              fill = hoverColor;
            }

            return (
              <rect
                  key={"bin-" + i}
                  fill={fill}
                  x={barX}
                  y={barY}
                  width={barWidth}
                  height={barHeight}
              />
            );
          })}
        </Group>
      </svg>
      <RenderTooltip
          data={data}
          tooltipOpen={tooltip.tooltipOpen}
          tooltipData={tooltip.tooltipData}
          tooltipLeft={tooltip.tooltipLeft}
          tooltipTop={tooltip.tooltipTop}
          handles={tooltipHandles}
          showTooltip={tooltip.showTooltip}
          delta={delta}
          xLin={xLin}
          xScale={xScale}
          yScale={yScale}
          margin={margin}
      />
    </div>
  );
};

function useScales(data: SummaryStat | undefined, xMax: number, yMax: number) {
  const xScale = useMemo(() => scaleBand<number>({
    range: [0, xMax],
    domain: range(data?.histogram?.length ?? 1),
    paddingInner: 0.05,
  }), [data?.histogram?.length, xMax]);

  const xLin = useMemo(() => scaleLinear<number>({
    range: [0, xMax],
    domain: [data?.min ?? 0, data?.max ?? 1],
  }), [data?.max, data?.min, xMax]);

  const yScale = useMemo(() => {
    const max = Math.max(...(data?.histogram ?? [1000]));
    return scaleLinear<number>({
      range: [yMax, 0],
      domain: [0, max],
      clamp: true
    });
  }, [data?.histogram, yMax]);

  const delta = useMemo(() => {
    if (data?.histogram?.length == null || data?.max == null || data?.min == null) {
      return null;
    }

    if (data.histogram.length <= 1) {
      return 0;
    }

    return data.histogram.length / (data.max - data.min);
  }, [data?.histogram?.length, data?.max, data?.min]);

  return { xScale: xScale, yScale: yScale, xLin: xLin, delta: delta };
}

const TickLabel = (props: TickRendererProps) => {
  return (
    <Text x={props.x} y={props.y} dy="0.25em" textAnchor="end" className="cursor-default">
      {props.formattedValue}
    </Text>
  );
};

const TickLabelHorizontal = (props: TickRendererProps) => {
  return (
    <Text x={props.x} y={props.y} dy="0.25em" textAnchor="middle" className="cursor-default">
      {props.formattedValue}
    </Text>
  );
};