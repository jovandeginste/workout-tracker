import { html, LitElement, PropertyValues, TemplateResult } from "lit";
import { customElement, property } from "lit/decorators.js";
import { formatDuration } from "../../helpers.js";
import {
  CategoryScale,
  Chart,
  ChartDataCustomTypesPerDataset,
  ChartDatasetCustomTypesPerDataset,
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

  // type can be "time" or "distance"
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

  @property()
  tz: string = null;

  @property()
  lang: string = null;

  @property({
    converter: (value: string) => JSON.parse(value),
  })
  translations = null;

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
  private chartData: ChartDataCustomTypesPerDataset;

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

    this.chart.zoomScale("x", { min: start, max: end });
  }

  public resetZoom() {
    if (!this.chart) {
      return;
    }

    this.chart.resetZoom();
  }

  public willUpdate(cProps: PropertyValues<this>) {
    if (!cProps.has("data") && !cProps.has("preferredUnits")) {
      return;
    }

    if (this.chart) {
      this.chart.destroy();
      this.chart = null;
    }

    let labels: (number | Date)[];
    if (this.type === "time") {
      labels = this.data["time"].Data.map((t: number) => new Date(t).valueOf());
    } else if (this.type === "distance") {
      labels = this.data["distance"].Data.map((d: string) => parseFloat(d));
    }

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
          label: this.data[metric].Label,
          hidden: metricSettings[metric].hiddenByDefault || false,
          data: this.data[metric].Data,
          yAxisID: metric,
        });
      }
    }

    this.chartData = {
      datasets,
      labels,
    };
  }

  public override updated(props: PropertyValues): void {
    super.updated(props);

    const canvas = this.querySelector("canvas");
    if (canvas && this.chartData) {
      const metricSettings = this.getMetricSettings();
      this.chart = new Chart(canvas, {
        data: this.chartData,
        options: {
          maintainAspectRatio: false,
          animation: false,
          scales: {
            x: {
              type: this.type === "time" ? "time" : "linear",
              time: this.type === "time" ? { unit: "minute" } : undefined,
              min:
                this.type === "time"
                  ? new Date(this.data["time"].Data[0]).valueOf()
                  : parseFloat(this.data["distance"].Data[0]),
              max:
                this.type === "time"
                  ? new Date(
                      this.data["time"].Data[this.data["time"].Data.length - 1],
                    ).valueOf()
                  : parseFloat(
                      this.data["distance"].Data[
                        this.data["distance"].Data.length - 1
                      ],
                    ),
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
            let p = this.data["position"].Data[index];
            let el = document.createElement("div");
            el.setAttribute("data-lat", p[0]);
            el.setAttribute("data-lng", p[1]);

            this.map.setMarker(el);
          },
          plugins: {
            decimation: {
              enabled: true,
              algorithm: "lttb",
            },
            legend: {
              display: true,
              onClick: (e, legendItem, legend) => {
                const meta = legend.chart.getDatasetMeta(
                  legendItem.datasetIndex,
                );
                meta.hidden = meta.hidden === null ? !legendItem.hidden : null;
                legend.chart.options.scales[meta.yAxisID].display =
                  legendItem.hidden;
                legend.chart.update();
              },
            },
            tooltip: {
              callbacks: {
                label: (tooltipItem) => {
                  const settings = metricSettings[tooltipItem.dataset.yAxisID];
                  let value = tooltipItem.formattedValue;
                  if (settings && settings.formatter) {
                    value = settings.formatter(tooltipItem.raw);
                  }

                  return value;
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
        },
      });
    }
  }

  public render(): TemplateResult {
    this.style.display = "block";
    return html`<canvas></canvas>`;
  }

  protected createRenderRoot() {
    return this;
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
        yaxis: { min: 0 },
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
