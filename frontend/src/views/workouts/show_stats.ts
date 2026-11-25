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
@customElement("workout-stats")
export class WorkoutStats extends LitElement {
  @property({
    attribute: "color-mode",
  })
  colorMode = "browser";

  // type can be "time", "duration" or "distance"
  @property()
  type = "time";

  @property({
    attribute: "map-id",
    converter: (value: string) => document.getElementById(value),
  })
  map = null;

  @property({
    attribute: "preferred-units-el",
    converter: (id: string) =>
      JSON.parse(document.getElementById(id)?.textContent || "null"),
  })
  preferredUnits = null;

  @property({
    attribute: "data-el",
    converter: (id: string) =>
      JSON.parse(document.getElementById(id)?.textContent || "null"),
  })
  data = null;

  @property({
    converter: (value: string) => JSON.parse(value),
  })
  translations = null;

  private timeLabels: number[] = [];

  // xValues returns the values for the x-axis of the current type.
  private get xValues(): number[] {
    switch (this.type) {
      case "distance":
        return this.data["distance"].Data.map((d: string) => parseFloat(d));
      case "duration":
        return this.data["moving-duration"].Data.map((d: number) => Number(d));
      default:
        return this.timeLabels;
    }
  }

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

    if (this.type === "distance" || this.type === "duration") {
      const values = this.xValues;
      start = values[this.timeLabels.indexOf(start as number)];
      end = values[this.timeLabels.indexOf(end as number)];
    }

    this.chart.zoomScale("x", { min: start as number, max: end as number });
  }

  public resetZoom() {
    if (!this.chart) {
      return;
    }

    this.chart.resetZoom();
  }

  private get xRange(): { min: number; max: number } {
    const values = this.xValues;

    return { min: values[0], max: values[values.length - 1] };
  }

  private formatXValue(val: number): string {
    switch (this.type) {
      case "distance":
        return `${val % 1 ? val.toFixed(1) : val} ${
          this.preferredUnits.distance
        }`;
      case "duration":
        return formatDuration(val);
      default:
        return new Date(val).toTimeString().substr(0, 5);
    }
  }

  public willUpdate(cProps: PropertyValues<this>) {
    if (cProps.has("data") || cProps.has("type")) {
      this.timeLabels = this.data["time"].Data.map((t: number) =>
        new Date(t).valueOf(),
      );

      this.chartData.labels = this.xValues;
    }

    if (cProps.has("data") || !cProps.has("preferredUnits")) {
      const metricSettings = this.getMetricSettings();
      const datasets: ChartDatasetCustomTypesPerDataset[] = [];
      for (let metric of Object.keys(metricSettings)) {
        if (metric === "time") continue;
        if (metric === "duration") continue;
        if (metric === "distance") continue;
        if (this.data[metric] !== undefined) {
          datasets.push({
            type: "line",
            fill:
              metricSettings[metric].seriesType === "area" ? "start" : false,
            spanGaps: true,
            label: this.data[metric].Label,
            hidden: metricSettings[metric].hiddenByDefault || false,
            data: this.data[metric].Data,
            yAxisID: metric,
          });
        }
      }

      this.chartData.datasets = datasets;
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
    return html`
        <div>
            <nav>
              <a
                href="#"
                title=${this.translations?.["time-of-day"] ?? ""}
                class=${navPillClass(this.type === "time")}
                @click=${(e: Event) => {
                  e.preventDefault();
                  this.type = "time";
                }}
                ><span class="icon-[fa6-regular--clock]"></span></a>
              <a
                href="#"
                title=${this.translations?.["moving-duration"] ?? ""}
                class=${navPillClass(this.type === "duration")}
                @click=${(e: Event) => {
                  e.preventDefault();
                  this.type = "duration";
                }}
                ><span class="icon-[fa6-solid--stopwatch]"></span></a>
              <a
                href="#"
                title=${this.translations?.distance ?? ""}
                class=${navPillClass(this.type === "distance")}
                @click=${(e: Event) => {
                  e.preventDefault();
                  this.type = "distance";
                }}
                ><span class="icon-[fa6-solid--road]"></span></a
            </nav>
          </div>
        <div class="h-[300px] md:h-[400px]">
          <canvas></canvas>
        </div>
        `;
  }

  protected createRenderRoot() {
    return this;
  }

  private getChartOptions(): ChartOptions {
    const metricSettings = this.getMetricSettings();
    const { fgColor, gridColor } = chartTheme(this.colorMode);
    return {
      maintainAspectRatio: false,
      animation: false,
      scales: {
        x: {
          type: this.type === "time" ? "time" : "linear",
          time: this.type === "time" ? { unit: "minute" } : undefined,
          min: this.xRange.min,
          max: this.xRange.max,
          ticks: {
            color: fgColor,
            callback: (val: number) => this.formatXValue(val),
          },
          grid: { color: gridColor },
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
                    color: fgColor,
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
                  grid: { color: gridColor },
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
        const p = this.data["position"].Data[index];
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
          labels: { color: fgColor },
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

              return this.formatXValue(x);
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
        zoom: zoomOptions(),
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
      calories: {
        formatter: (val) => `${val ?? "-"} ${this.preferredUnits.calories}`,
        formatterYaxis: true,
        hiddenByDefault: true,
        yaxis: { min: 0 },
      },
      power: {
        formatter: (val) => `${val ?? "-"} ${this.preferredUnits.power}`,
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
