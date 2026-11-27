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
import { chartTheme, saveConfig, zoomOptions } from "../../chart";
import {
  ChartPoint,
  colorForType,
  trendCutoffFromPeriod,
  trendSegments,
} from "../../trend";

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

@customElement("route-segment-stats")
export class RouteSegmentStats extends LitElement {
  @property({
    attribute: "trend-data",
    converter: (v: string) => JSON.parse(v) as TrendDataPoint[],
  })
  trendData: TrendDataPoint[] = [];

  @property({
    attribute: "color-mode",
  })
  colorMode = "browser";

  @property({
    converter: (value: string) => JSON.parse(value) as Translations,
  })
  translations: Translations = null;

  @property({
    attribute: "trend-period",
  })
  trendPeriod: TrendPeriod = "365";

  @property({
    attribute: "config-update-route",
  })
  configUpdateRoute = "";

  @property({
    attribute: "trend-break-detection",
    type: Boolean,
  })
  trendBreakDetection = false;

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

    await saveConfig(
      this.configUpdateRoute,
      new URLSearchParams({ route_segment_trend_period: period }),
      "Failed to save route segment trend period preference",
    );
  }

  private async updateTrendBreakDetection(event: Event): Promise<void> {
    this.trendBreakDetection =
      (event.target as HTMLSelectElement).value === "break";

    await saveConfig(
      this.configUpdateRoute,
      new URLSearchParams({
        route_segment_trend_break_detection: String(this.trendBreakDetection),
      }),
      "Failed to save route segment trend break detection preference",
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

    if (!this.trendData || this.trendData.length === 0) return;

    const processedData = this.trendData.map((d) => ({
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
    const { fgColor, gridColor } = chartTheme(this.colorMode);

    const allTypes = Array.from(
      new Set(processedData.map((d) => d.type)),
    ).sort();

    const trendCutoff = trendCutoffFromPeriod(this.trendPeriod);
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

      const color = colorForType(type);

      const trendData = typeData.filter(
        (d: ChartPoint) => trendCutoff === null || d.x >= trendCutoff,
      );
      const segments = trendSegments(trendData, this.trendBreakDetection);

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
        ...segments.map((segment, index) => ({
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
          zoom: zoomOptions(true),
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
