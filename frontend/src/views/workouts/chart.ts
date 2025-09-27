import { html, LitElement, PropertyValues, TemplateResult } from "lit";
import { customElement, property } from "lit/decorators.js";
import { formatDuration } from "../../helpers.js";
import {
  CategoryScale,
  Chart,
  ChartDataCustomTypesPerDataset,
  ChartDatasetCustomTypesPerDataset,
  ChartOptions,
  Colors,
  Decimation,
  Filler,
  Legend,
  LinearScale,
  LineController,
  LineElement,
  PointElement,
  TimeScale,
  Tooltip,
} from "chart.js";
import "chartjs-adapter-date-fns";
import Zoom from "chartjs-plugin-zoom";
import { localized, msg } from "@lit/localize";
import { initLocalize } from "../../locale.js";
import { WorkoutData, WorkoutService } from "./service.js";

initLocalize();

@customElement("workout-chart")
@localized()
export class WorkoutChart extends LitElement {
  @property({
    attribute: "color-mode",
  })
  colorMode = "browser";

  // type can be "time" or "distance"
  @property()
  type = "time";

  @property({
    attribute: "map-id",
    converter: (value: string) => document.getElementById(value),
  })
  map = null;

  @property()
  tz: string = null;

  @property()
  lang: string = null;

  @property({
    converter: (value: string) => JSON.parse(value),
  })
  translations = null;

  private workoutService = new WorkoutService();
  private timeLabels: number[] = [];
  private preferredUnits = null;
  private data: WorkoutData = null;

  public constructor() {
    super();

    Chart.register(
      TimeScale,
      CategoryScale,
      LinearScale,
      PointElement,
      LineController,
      LineElement,
      Filler,
      Decimation,
      Colors,
      Tooltip,
      Legend,
      Zoom,
    );

    this.preferredUnits = this.workoutService.preferredUnits;
    this.data = this.workoutService.workoutData;

    const metricSettings = this.getMetricSettings();
    const datasets: ChartDatasetCustomTypesPerDataset[] = [];
    for (let metric of Object.keys(metricSettings)) {
      if (metric === "time") continue;
      if (metric === "duration") continue;
      if (metric === "distance") continue;
      if (this.data[metric] !== undefined) {
        datasets.push({
          type: "line",
          fill: metricSettings[metric].seriesType === "area" ? "start" : false,
          spanGaps: true,
          label: this.data[metric].label,
          hidden: metricSettings[metric].hiddenByDefault || false,
          data: this.data[metric].data,
          yAxisID: metric,
        });
      }
    }

    this.chartData.datasets = datasets;
  }

  private chart: Chart | null = null;
  private chartData: ChartDataCustomTypesPerDataset = {
    datasets: [],
    labels: [],
  };

  public zoomX(start: number | Date, end: number | Date) {
    if (!this.chart) {
      return;
    }

    if (start instanceof Date) {
      start = start.valueOf();
    }

    if (end instanceof Date) {
      end = end.valueOf();
    }

    if (this.type === "distance") {
      start = parseFloat(
        this.data["distance"].data[this.timeLabels.indexOf(start)],
      );
      end = parseFloat(
        this.data["distance"].data[this.timeLabels.indexOf(end)],
      );
    }

    this.chart.zoomScale("x", { min: start as number, max: end as number });
  }

  public resetZoom() {
    if (!this.chart) {
      return;
    }

    this.chart.resetZoom();
  }

  public willUpdate(cProps: PropertyValues<this>) {
    if (cProps.has("type")) {
      let labels: (number | Date)[];
      this.timeLabels = this.data["time"].data.map((t: number) =>
        new Date(t).valueOf(),
      );
      if (this.type === "time") {
        labels = this.timeLabels;
      } else if (this.type === "distance") {
        labels = this.data["distance"].data.map((d: string) => parseFloat(d));
      }

      this.chartData.labels = labels;
    }
  }

  public override updated(props: PropertyValues): void {
    super.updated(props);

    if (this.chart) {
      this.chart.options = this.getChartOptions();
      this.chart.update();
    } else {
      const canvas = this.querySelector("canvas");
      if (canvas && this.chartData) {
        this.chart = new Chart(canvas, {
          data: this.chartData,
          options: this.getChartOptions(),
        });
      }
    }
  }

  public render(): TemplateResult {
    this.style.display = "block";
    return html` <div class="border-b border-gray-200 dark:border-neutral-700">
        <div class="flex items-center justify-between px-4">
          <h3 class="font-semibold mb-0">
            <span>
              <span class="icon-decoration icon-[fa6-solid--gauge]"></span>
              ${msg("Average speed", { id: "translation.Average_speed" })}
            </span>
            /
            <span>
              <span class="icon-decoration icon-[fa6-solid--mountain]"></span>
              ${msg("Elevation", { id: "translation.Elevation" })}
            </span>
          </h3>
          <div class="flex space-x-4 text-sm">
            <nav class="flex space-x-8 px-4">
              <button
                class="tab-button ${this.type === "time" ? "active" : ""}"
                @click=${(e: Event) => {
                  e.preventDefault();
                  this.type = "time";
                }}
              >
                <span class="icon-decoration icon-[fa6-regular--clock]"></span>
                ${msg("Time", { id: "translation.Time" })}
              </button>
              <button
                class="tab-button ${this.type === "distance" ? "active" : ""}"
                @click=${(e: Event) => {
                  e.preventDefault();
                  this.type = "distance";
                }}
              >
                <span class="icon-decoration icon-[fa6-solid--road]"></span>
                ${msg("Distance", { id: "translation.Distance" })}
              </button>
            </nav>
          </div>
        </div>
      </div>
      <div class="p-4">
        <div class="h-[300px] md:h-[400px]">
          <canvas></canvas>
        </div>
      </div>`;
  }

  protected createRenderRoot() {
    return this;
  }

  private getChartOptions(): ChartOptions {
    const metricSettings = this.getMetricSettings();
    return {
      maintainAspectRatio: false,
      animation: false,
      scales: {
        x: {
          type: this.type === "time" ? "time" : "linear",
          time: this.type === "time" ? { unit: "minute" } : undefined,
          min:
            this.type === "time"
              ? new Date(this.data["time"].data[0]).valueOf()
              : parseFloat(this.data["distance"].data[0]),
          max:
            this.type === "time"
              ? new Date(
                  this.data["time"].data[this.data["time"].data.length - 1],
                ).valueOf()
              : parseFloat(
                  this.data["distance"].data[
                    this.data["distance"].data.length - 1
                  ],
                ),
          ticks: {
            callback: (val: number) => {
              if (this.type === "distance") {
                return `${val % 1 ? val.toFixed(1) : val} ${
                  this.preferredUnits.distance
                }`;
              }

              return new Date(val as number).toTimeString().substr(0, 5);
            },
          },
        },
        ...Object.fromEntries(
          Object.keys(metricSettings)
            .map((metric) => {
              if (metricSettings[metric].yaxis === false) {
                return [];
              }

              return [
                metric,
                {
                  display: !metricSettings[metric].hiddenByDefault,
                  position: "left",
                  ...metricSettings[metric].yaxis,
                  ticks: {
                    callback: (val) => {
                      const settings = metricSettings[metric];
                      if (settings.formatterYaxis) {
                        return settings.labelFormatter
                          ? settings.labelFormatter(val as number)
                          : settings.formatter(val as number);
                      }
                      return val;
                    },
                  },
                },
              ];
            })
            .filter((e) => e.length > 0),
        ),
      },
      elements: {
        point: {
          radius: 0,
        },
      },
      interaction: {
        mode: "index",
        intersect: false,
      },
      onHover: (_, i) => {
        const index = i[0]?.index;
        const p = this.data["position"].data[index];
        if (p) {
          const el = document.createElement("div");
          el.setAttribute("data-lat", p[0]);
          el.setAttribute("data-lng", p[1]);

          this.map.setMarker(el);
        }
      },
      plugins: {
        decimation: {
          enabled: true,
          algorithm: "lttb",
        },
        legend: {
          display: true,
          onClick: (e, legendItem, legend) => {
            const meta = legend.chart.getDatasetMeta(legendItem.datasetIndex);
            meta.hidden = meta.hidden === null ? !legendItem.hidden : null;
            legend.chart.options.scales[meta.yAxisID].display =
              legendItem.hidden;
            legend.chart.update();
          },
        },
        tooltip: {
          callbacks: {
            title: (tooltipItems) => {
              if (!tooltipItems[0]) {
                return;
              }

              const x = tooltipItems[0].parsed.x;
              if (this.type === "distance") {
                return `${x.toFixed(2)} ${this.preferredUnits.distance}`;
              }
              return new Date(x).toTimeString().substr(0, 5);
            },
            label: (tooltipItem) => {
              const settings = metricSettings[tooltipItem.dataset.yAxisID];
              let value = tooltipItem.formattedValue;
              if (settings && settings.formatter) {
                value = settings.formatter(tooltipItem.raw);
              }

              return `${tooltipItem.dataset.label}: ${value}`;
            },
          },
        },
        zoom: {
          limits: {
            x: { min: "original", max: "original" },
            y: { min: "original", max: "original" },
          },
          zoom: {
            drag: {
              enabled: true,
            },
            wheel: {
              enabled: true,
            },
            mode: "x",
          },
        },
      },
    };
  }

  private getMetricSettings() {
    return {
      speed: {
        formatter: (val) => `${val ?? "-"} ${this.preferredUnits.speed}`,
        formatterYaxis: true,
        yaxis: { min: 0 },
      },
      elevation: {
        seriesType: "area",
        formatter: (val) =>
          `${val !== null ? val.toFixed(2) : "-"} ${
            this.preferredUnits.elevation
          }`,
        labelFormatter: (val) => `${val} ${this.preferredUnits.elevation}`,
        formatterYaxis: true,
        yaxis: { position: "right" },
      },
      "heart-rate": {
        formatter: (val) => `${val ?? "-"} ${this.preferredUnits.heartRate}`,
        formatterYaxis: true,
        hiddenByDefault: true,
      },
      cadence: {
        formatter: (val) => `${val ?? "-"} ${this.preferredUnits.cadence}`,
        formatterYaxis: true,
        hiddenByDefault: true,
        yaxis: { min: 0 },
      },
      temperature: {
        formatter: (val) =>
          `${val ?? "-"} ${this.preferredUnits.temperature || "°C"}`,
        formatterYaxis: true,
        hiddenByDefault: true,
      },
      distance: {
        seriesType: "none",
        formatter: (val) => `${val ?? "-"} ${this.preferredUnits.distance}`,
        yaxis: false,
        legend: false,
      },
      duration: {
        seriesType: "none",
        formatter: (val) => formatDuration(val),
        yaxis: false,
        legend: false,
      },
    };
  }
}
