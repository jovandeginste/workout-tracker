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
import pluginTrendlineLinear from "chartjs-plugin-trendline";
import { localized } from "@lit/localize";
import { initLocalize } from "../../locale.js";

initLocalize();

interface TrendDataPoint {
  date: string;
  speed: string;
}

interface Translations {
  averageSpeed: string;
  speedUnit: string;
  trend: string;
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

  private chart: Chart | null = null;

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
      pluginTrendlineLinear,
    );
  }

  private isDark(): boolean {
    if (this.colorMode === "dark") return true;
    if (this.colorMode === "light") return false;
    return window.matchMedia("(prefers-color-scheme: dark)").matches;
  }

  public override updated(_props: PropertyValues): void {
    super.updated(_props);

    if (!this.data || this.data.length === 0) return;

    if (this.chart) {
      this.chart.destroy();
      this.chart = null;
    }

    const canvas = this.querySelector("canvas") as HTMLCanvasElement;
    if (!canvas) return;

    const speedUnit = this.translations?.speedUnit || "";
    const dark = this.isDark();
    const dotColor = dark ? "#fef3c7" : "#b45309";
    const trendColor = dark ? "#fbbf24" : "#d97706";
    const fgColor = dark ? "#e4e4e7" : "#27272a";
    const gridColor = dark ? "#3f3f46" : "#d4d4d8";

    const scatterData = this.data
      .map((d) => ({
        x: new Date(d.date).valueOf(),
        y: parseFloat(d.speed),
      }))
      .filter((d) => !isNaN(d.y))
      .sort((a, b) => a.x - b.x);

    this.chart = new Chart(canvas, {
      type: "scatter",
      data: {
        datasets: [
          {
            label: this.translations?.averageSpeed || "Average speed",
            backgroundColor: dotColor,
            borderColor: dotColor,
            data: scatterData,
            trendlineLinear: {
              colorMin: trendColor,
              colorMax: trendColor,
              width: 2,
              lineStyle: "solid",
            },
          } as never,
        ],
      },
      options: {
        maintainAspectRatio: false,
        animation: false,
        scales: {
          x: {
            type: "time",
            time: { unit: "month" },
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
          legend: { display: false },
          tooltip: {
            callbacks: {
              label: (item) =>
                `${(item.raw as { y: number }).y.toFixed(1)} ${speedUnit}`,
              title: (items) =>
                new Date(items[0].parsed.x).toLocaleDateString(),
            },
          },
        },
      },
    });
  }

  public render(): TemplateResult {
    this.style.display = "block";
    return html`<div class="h-full"><canvas></canvas></div>`;
  }

  protected createRenderRoot() {
    return this;
  }
}
