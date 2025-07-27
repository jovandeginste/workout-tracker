import ApexCharts from "apexcharts";
import { formatDuration } from "../../helpers";

class WtStatistic extends HTMLElement {
  constructor() {
    super();
  }

  connectedCallback() {
    this.stats = JSON.parse(this.getAttribute("stats") || `{}`);
    this.preferredUnits = JSON.parse(
      this.getAttribute("preferred-units") || `{}`,
    );
    this.filterNoDuration = this.getAttribute("filter-no-duration") !== null;
    this.unit = this.getAttribute("unit");
    this.type = this.getAttribute("type");
    if (!this.stats || !this.stats.buckets) {
      console.warn("No stats provided for WtStatistic");
      return;
    }

    var theme = "light";
    if (
      window.matchMedia &&
      window.matchMedia("(prefers-color-scheme: dark)").matches
    ) {
      theme = "dark";
    }

    var options = {
      theme: { mode: theme },
      chart: {
        width: "99%",
        animations: { enabled: false },
        toolbar: { show: false },
        type: "bar",
      },
      dataLabels: { enabled: false },
      legend: { position: "top" },
      tooltip: {
        x: { format: "MMM 'yy" },
      },
      xaxis: { type: "datetime" },
    };

    if (this.type === "duration") {
      options.tooltip.y = [
        {
          formatter: function (val, _) {
            return formatDuration(val);
          },
        },
      ];
      options.yaxis = [
        {
          labels: {
            formatter: (val) => {
              return formatDuration(val);
            },
          },
        },
      ];
    } else if (this.type === "distance" || this.type === "speed") {
      options.tooltip.y = [
        {
          formatter: function (val, _) {
            return val + " " + this.preferredUnits[this.type];
          },
        },
      ];
      options.yaxis = [
        {
          labels: {
            formatter: (val) => {
              return val + " " + this.preferredUnits[this.type];
            },
          },
        },
      ];
    }

    const el = document.createElement("div");
    this.appendChild(el);
    new ApexCharts(el, {
      ...options,
      series: Object.entries(this.stats.buckets)
        .map((entry) => {
          const [_, value] = entry;
          return {
            name: value.localWorkoutType,
            data: Object.values(value.buckets)
              .filter((e) => !this.filterNoDuration || e.duration > 0)
              .map((e) => ({ x: e.bucket, y: e[this.type] })),
          };
        })
        .filter((e) => e.data.length > 0),
    }).render();
  }
}

customElements.define("wt-stat", WtStatistic);
