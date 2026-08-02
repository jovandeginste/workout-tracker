import { html, LitElement, PropertyValues, TemplateResult } from "lit";
import { customElement, property } from "lit/decorators.js";
import {
  Chart,
  ScatterController,
  LineElement,
  PointElement,
  LinearScale,
  TimeScale,
  Tooltip,
  Legend,
} from "chart.js";
import "chartjs-adapter-date-fns";
import Zoom from "chartjs-plugin-zoom";
import { localized } from "@lit/localize";
import { initLocalize } from "../../locale.js";

initLocalize();

const MAX_TREND_BREAK_POINTS = 200;
const DAYS_PER_MONTH = 30.4375;

interface TrendDataPoint {
  date: string;
  speed: string;
  type: string;
}

interface Translations {
  speedUnit: string;
  [key: string]: string;
}

type TrendPeriod = "all" | "365" | "90" | "30" | "year";

interface ChartPoint {
  x: number;
  y: number;
}

interface TrendBreak {
  pivot: number;
  pValue: number;
}

interface Regression {
  slope: number;
  intercept: number;
  baseX: number;
}

@customElement("route-segment-stats")
@localized()
export class RouteSegmentStats extends LitElement {
  @property({
    converter: (v: string) => JSON.parse(v) as TrendDataPoint[],
  })
  data: TrendDataPoint[] = [];

  @property({
    attribute: "color-mode",
  })
  colorMode = "browser";

  @property()
  lang: string = null;

  @property({
    converter: (value: string) => JSON.parse(value) as Translations,
  })
  translations: Translations = null;

  @property({
    attribute: "trend-period",
  })
  trendPeriod: TrendPeriod = "365";

  @property({
    attribute: "trend-period-update-route",
  })
  trendPeriodUpdateRoute = "";

  @property({
    attribute: "trend-break-detection",
    type: Boolean,
  })
  trendBreakDetection = false;

  @property({
    attribute: "trend-break-detection-update-route",
  })
  trendBreakDetectionUpdateRoute = "";

  private chart: Chart | null = null;
  private hiddenTrendGroups = new Set<string>();

  public constructor() {
    super();
    Chart.register(
      ScatterController,
      LineElement,
      PointElement,
      LinearScale,
      TimeScale,
      Tooltip,
      Legend,
      Zoom,
    );
  }

  private isDark(): boolean {
    if (this.colorMode === "dark") return true;
    if (this.colorMode === "light") return false;
    return window.matchMedia("(prefers-color-scheme: dark)").matches;
  }

  private colorForType(type: string): string {
    const palette = [
      "#2563eb",
      "#16a34a",
      "#dc2626",
      "#9333ea",
      "#ea580c",
      "#0891b2",
      "#4f46e5",
      "#65a30d",
      "#be123c",
      "#7c3aed",
    ];

    let hash = 0;
    for (let i = 0; i < type.length; i++) {
      hash = (hash * 31 + type.charCodeAt(i)) % palette.length;
    }

    return palette[hash];
  }

  private trendCutoff(): number | null {
    const now = new Date();

    if (this.trendPeriod === "all") {
      return null;
    }

    if (this.trendPeriod === "year") {
      return new Date(now.getFullYear(), 0, 1).valueOf();
    }

    return now.valueOf() - Number(this.trendPeriod) * 24 * 60 * 60 * 1000;
  }

  private trendPeriodOptions(): { value: TrendPeriod; label: string }[] {
    return [
      { value: "all", label: this.translations?.trendPeriodAll || "All time" },
      {
        value: "365",
        label: this.translations?.trendPeriod365 || "Last 365 days",
      },
      {
        value: "90",
        label: this.translations?.trendPeriod90 || "Last 90 days",
      },
      {
        value: "30",
        label: this.translations?.trendPeriod30 || "Last 30 days",
      },
      {
        value: "year",
        label: this.translations?.trendPeriodYear || "This calendar year",
      },
    ];
  }

  private async updateTrendPeriod(event: Event): Promise<void> {
    const period = (event.target as HTMLSelectElement).value as TrendPeriod;
    this.trendPeriod = period;

    if (!this.trendPeriodUpdateRoute) {
      return;
    }

    const body = new URLSearchParams({
      route_segment_trend_period: period,
    });

    try {
      const response = await fetch(this.trendPeriodUpdateRoute, {
        method: "POST",
        headers: {
          "Content-Type": "application/x-www-form-urlencoded",
        },
        body,
      });

      if (response.ok) {
        return;
      }
    } catch (_err) {
      // Keep the selected chart period even if saving the preference fails.
    }

    if (this.trendPeriodUpdateRoute) {
      console.warn("Failed to save route segment trend period preference");
    }
  }

  private async updateTrendBreakDetection(event: Event): Promise<void> {
    this.trendBreakDetection =
      (event.target as HTMLSelectElement).value === "break";

    if (!this.trendBreakDetectionUpdateRoute) {
      return;
    }

    const body = new URLSearchParams({
      route_segment_trend_break_detection: String(this.trendBreakDetection),
    });

    try {
      const response = await fetch(this.trendBreakDetectionUpdateRoute, {
        method: "POST",
        headers: {
          "Content-Type": "application/x-www-form-urlencoded",
        },
        body,
      });

      if (response.ok) {
        return;
      }
    } catch (_err) {
      // Keep the selected chart setting even if saving the preference fails.
    }

    console.warn(
      "Failed to save route segment trend break detection preference",
    );
  }

  private average(points: ChartPoint[]): number {
    return points.reduce((sum, point) => sum + point.y, 0) / points.length;
  }

  private variance(points: ChartPoint[], average: number): number {
    if (points.length < 2) {
      return 0;
    }

    return (
      points.reduce((sum, point) => sum + (point.y - average) ** 2, 0) /
      (points.length - 1)
    );
  }

  private normalCDF(value: number): number {
    const sign = value < 0 ? -1 : 1;
    const x = Math.abs(value) / Math.sqrt(2);
    const t = 1 / (1 + 0.3275911 * x);
    const erf =
      1 -
      ((((1.061405429 * t - 1.453152027) * t + 1.421413741) * t - 0.284496736) *
        t +
        0.254829592) *
        t *
        Math.exp(-x * x);

    return 0.5 * (1 + sign * erf);
  }

  private trendBreakPValue(before: ChartPoint[], after: ChartPoint[]): number {
    const beforeAverage = this.average(before);
    const afterAverage = this.average(after);
    const beforeVariance = this.variance(before, beforeAverage);
    const afterVariance = this.variance(after, afterAverage);
    const standardError = Math.sqrt(
      beforeVariance / before.length + afterVariance / after.length,
    );

    if (standardError === 0) {
      return beforeAverage === afterAverage ? 1 : 0;
    }

    const testStatistic =
      Math.abs(beforeAverage - afterAverage) / standardError;

    return 2 * (1 - this.normalCDF(testStatistic));
  }

  private detectTrendBreak(points: ChartPoint[]): TrendBreak | null {
    const minimumPointsPerSide = 3;

    if (!this.trendBreakDetection) {
      return null;
    }

    if (points.length < minimumPointsPerSide * 2) {
      return null;
    }

    if (points.length > MAX_TREND_BREAK_POINTS) {
      return null;
    }

    let bestBreak: TrendBreak | null = null;
    let bestGroupSize = 0;

    for (
      let pivot = minimumPointsPerSide;
      pivot <= points.length - minimumPointsPerSide;
      pivot++
    ) {
      const before = points.slice(0, pivot);
      const after = points.slice(pivot);
      const pValue = this.trendBreakPValue(before, after);

      if (pValue > 0.05) {
        continue;
      }

      const groupSize = Math.min(before.length, after.length);

      if (
        groupSize > bestGroupSize ||
        (groupSize === bestGroupSize &&
          (!bestBreak || pValue < bestBreak.pValue))
      ) {
        bestGroupSize = groupSize;
        bestBreak = {
          pivot,
          pValue,
        };
      }
    }

    return bestBreak;
  }

  private canShowTrend(points: ChartPoint[]): boolean {
    if (points.length < 5) {
      return false;
    }

    const first = points[0].x;
    const last = points[points.length - 1].x;
    const months = Math.max(
      1,
      (last - first) / (DAYS_PER_MONTH * 24 * 60 * 60 * 1000),
    );

    return points.length / months >= 1;
  }

  private linearRegression(points: ChartPoint[]): Regression {
    const count = points.length;
    const baseX = points[0].x;
    const sums = points.reduce(
      (acc, point) => {
        const x = point.x - baseX;
        acc.x += x;
        acc.y += point.y;
        acc.xy += x * point.y;
        acc.x2 += x * x;
        return acc;
      },
      { x: 0, y: 0, xy: 0, x2: 0 },
    );

    const denominator = count * sums.x2 - sums.x * sums.x;
    if (denominator === 0) {
      return { slope: 0, intercept: sums.y / count, baseX };
    }

    const slope = (count * sums.xy - sums.x * sums.y) / denominator;
    const intercept = (sums.y - slope * sums.x) / count;

    return { slope, intercept, baseX };
  }

  private trendLineData(points: ChartPoint[]): ChartPoint[] {
    if (!this.canShowTrend(points)) {
      return [];
    }

    const regression = this.linearRegression(points);
    const first = points[0];
    const last = points[points.length - 1];

    return [first, last].map((point) => ({
      x: point.x,
      y: regression.slope * (point.x - regression.baseX) + regression.intercept,
    }));
  }

  private averageLineData(points: ChartPoint[]): ChartPoint[] {
    if (!this.canShowTrend(points)) {
      return [];
    }

    const y = this.average(points);

    return [
      { x: points[0].x, y },
      { x: points[points.length - 1].x, y },
    ];
  }

  private trendSegments(points: ChartPoint[]): ChartPoint[][] {
    if (!this.canShowTrend(points)) {
      return [];
    }

    if (!this.trendBreakDetection) {
      return [this.trendLineData(points)].filter(
        (segment) => segment.length > 0,
      );
    }

    const trendBreak = this.detectTrendBreak(points);

    if (!trendBreak) {
      return [this.trendLineData(points)].filter(
        (segment) => segment.length > 0,
      );
    }

    const before = points.slice(0, trendBreak.pivot);
    const after = points.slice(trendBreak.pivot);
    if (!this.canShowTrend(before) || !this.canShowTrend(after)) {
      return [this.trendLineData(points)].filter(
        (segment) => segment.length > 0,
      );
    }

    return [this.averageLineData(before), this.averageLineData(after)].filter(
      (segment) => segment.length > 0,
    );
  }

  private syncHiddenTrendGroups(): void {
    if (!this.chart) {
      return;
    }

    this.chart.data.datasets.forEach((dataset, index) => {
      const groupedDataset = dataset as { trendGroup?: string };
      if (!groupedDataset.trendGroup) {
        return;
      }

      if (this.chart.isDatasetVisible(index)) {
        this.hiddenTrendGroups.delete(groupedDataset.trendGroup);
      } else {
        this.hiddenTrendGroups.add(groupedDataset.trendGroup);
      }
    });
  }

  public resetZoom(): void {
    this.chart?.resetZoom();
  }

  public override updated(_props: PropertyValues): void {
    super.updated(_props);

    if (!this.data || this.data.length === 0) return;

    const processedData = this.data.map((d) => ({
      ...d,
      type: d.type || "unknown",
    }));

    if (this.chart) {
      this.syncHiddenTrendGroups();
      this.chart.destroy();
      this.chart = null;
    }

    const canvas = this.querySelector("canvas") as HTMLCanvasElement;
    if (!canvas) return;

    const speedUnit = this.translations?.speedUnit || "";
    const dark = this.isDark();
    const fgColor = dark ? "#e4e4e7" : "#27272a";
    const gridColor = dark ? "#3f3f46" : "#d4d4d8";

    const allTypes = Array.from(
      new Set(processedData.map((d) => d.type)),
    ).sort();

    const trendCutoff = this.trendCutoff();
    const trendMax = new Date().valueOf();
    const datasets = allTypes.flatMap((type) => {
      const typeData = processedData
        .filter((d) => d.type === type)
        .map((d) => ({
          x: new Date(d.date).valueOf(),
          y: parseFloat(d.speed),
        }))
        .filter((d) => !isNaN(d.y))
        .sort((a, b) => a.x - b.x);

      const color = this.colorForType(type);

      const trendData = typeData.filter(
        (d: ChartPoint) => trendCutoff === null || d.x >= trendCutoff,
      );
      const trendSegments = this.trendSegments(trendData);

      return [
        {
          label: this.translations?.[type] || type,
          trendGroup: type,
          backgroundColor: color,
          borderColor: color,
          data: typeData,
          hidden: this.hiddenTrendGroups.has(type),
          showLine: false,
          borderWidth: 2,
          pointBorderWidth: 0,
          pointHoverBorderWidth: 0,
          pointRadius: 4,
        },
        ...trendSegments.map((segment, index) => ({
          label: `__trend:${type}:${index}`,
          trendGroup: type,
          backgroundColor: "transparent",
          borderColor: color,
          data: segment,
          hidden: this.hiddenTrendGroups.has(type),
          showLine: true,
          borderWidth: 2,
          pointRadius: 0,
          pointBorderWidth: 0,
          pointHoverBorderWidth: 0,
          pointHitRadius: 0,
        })),
      ];
    });

    this.chart = new Chart(canvas, {
      type: "scatter",
      data: {
        datasets: datasets as never[],
      },
      options: {
        maintainAspectRatio: false,
        animation: false,
        scales: {
          x: {
            type: "time",
            time: { unit: "month" },
            min: trendCutoff ?? undefined,
            max: trendCutoff === null ? undefined : trendMax,
            ticks: { color: fgColor },
            grid: { color: gridColor },
          },
          y: {
            ticks: {
              color: fgColor,
              callback: (val) => `${Number(val).toFixed(1)} ${speedUnit}`,
            },
            grid: { color: gridColor },
          },
        },
        plugins: {
          legend: {
            display: true,
            onClick: (_event, legendItem, legend) => {
              const datasetIndex = legendItem.datasetIndex;
              if (datasetIndex === undefined) {
                return;
              }

              const dataset = legend.chart.data.datasets[datasetIndex] as {
                trendGroup?: string;
              };
              const trendGroup = dataset.trendGroup;
              const visible = legend.chart.isDatasetVisible(datasetIndex);
              legend.chart.data.datasets.forEach((chartDataset, index) => {
                const groupedDataset = chartDataset as { trendGroup?: string };
                if (groupedDataset.trendGroup === trendGroup) {
                  legend.chart.setDatasetVisibility(index, !visible);
                }
              });
              if (visible) {
                this.hiddenTrendGroups.add(trendGroup);
              } else {
                this.hiddenTrendGroups.delete(trendGroup);
              }
              legend.chart.update();
            },
            labels: {
              color: fgColor,
              filter: (legendItem) =>
                !String(legendItem.text || "").startsWith("__trend:"),
            },
          },
          tooltip: {
            filter: (item) =>
              !String(item.dataset.label || "").startsWith("__trend:"),
            callbacks: {
              label: (item) => {
                const label = item.dataset.label || "";
                return `${label}: ${(item.raw as { y: number }).y.toFixed(
                  1,
                )} ${speedUnit}`;
              },
              title: (items) =>
                new Date(items[0].parsed.x).toLocaleDateString(),
            },
          },
          zoom: {
            limits: {
              x: { min: "original", max: "original" },
              y: { min: "original", max: "original" },
            },
            pan: {
              enabled: true,
              mode: "x",
            },
            zoom: {
              drag: {
                enabled: true,
              },
              wheel: {
                enabled: true,
              },
              pinch: {
                enabled: true,
              },
              mode: "x",
            },
          },
        },
      },
    });
  }

  public render(): TemplateResult {
    this.style.display = "grid";
    this.style.gridTemplateRows = "auto minmax(0, 1fr)";
    this.style.width = "100%";
    this.style.height = "100%";
    this.style.minHeight = "0";

    return html`
      <div class="mb-3 flex flex-wrap justify-end gap-3">
        <label for="trend-period" class="self-center">
          ${this.translations?.trendPeriod || "Trend period"}
        </label>
        <select
          id="trend-period"
          class="input-ellipsis"
          @change=${this.updateTrendPeriod}
        >
          ${this.trendPeriodOptions().map(
            (option) => html`
              <option
                value=${option.value}
                ?selected=${option.value === this.trendPeriod}
              >
                ${option.label}
              </option>
            `,
          )}
        </select>
        <select
          id="trend-break-detection"
          class="input-ellipsis"
          @change=${this.updateTrendBreakDetection}
        >
          <option value="trend" ?selected=${!this.trendBreakDetection}>
            ${this.translations?.trend || "Trend"}
          </option>
          <option value="break" ?selected=${this.trendBreakDetection}>
            ${this.translations?.trendBreak || "Trend break"}
          </option>
        </select>
        <button type="button" @click=${this.resetZoom}>
          ${this.translations?.resetZoom || "Reset zoom"}
        </button>
      </div>
      <div style="position: relative; min-height: 0; height: 100%;">
        <canvas style="height: 100%; width: 100%;"></canvas>
      </div>
    `;
  }

  protected createRenderRoot() {
    return this;
  }
}
