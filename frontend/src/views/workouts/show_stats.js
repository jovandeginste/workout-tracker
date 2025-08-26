import ApexCharts from "apexcharts";
import { formatDuration } from "../../helpers.js";

class WorkoutStats extends HTMLElement {
  constructor() {
    super();
    this.style.display = "block";
  }

  connectedCallback() {
    this.colorMode = this.getAttribute("color-mode") || "browser";
    this.mapElement = document.getElementById(this.getAttribute("map-id"));
    this.preferredUnits = JSON.parse(
      document.getElementById(this.getAttribute("preferred-units-el"))
        .textContent,
    );
    this.data = JSON.parse(
      document.getElementById(this.getAttribute("data-el")).textContent,
    );
    this.tz = this.getAttribute("tz");
    this.lang = this.getAttribute("lang");
    this.translations = JSON.parse(this.getAttribute("translations"));

    const metricSettings = {
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
        yaxis: { opposite: true },
      },
      "heart-rate": {
        formatter: (val) => `${val ?? "-"} ${this.preferredUnits.heartRate}`,
        formatterYaxis: true,
        hiddenByDefault: true,
        yaxis: {},
      },
      cadence: {
        formatter: (val) => `${val ?? "-"} ${this.preferredUnits.cadence}`,
        formatterYaxis: true,
        hiddenByDefault: true,
        yaxis: {},
      },
      temperature: {
        formatter: (val) =>
          `${val ?? "-"} ${this.preferredUnits.temperature || "°C"}`,
        formatterYaxis: true,
        hiddenByDefault: true,
        yaxis: {},
      },
      distance: {
        seriesType: "none",
        formatter: (val) => `${val ?? "-"} ${this.preferredUnits.distance}`,
        yaxis: { show: false },
        legend: false,
      },
      duration: {
        seriesType: "none",
        formatter: (val) => formatDuration(val),
        yaxis: { show: false },
        legend: false,
      },
    };

    const series = [];
    const yTooltips = [];
    const yaxis = [];
    for (let metric of Object.keys(metricSettings)) {
      if (metric === "time") continue;
      if (this.data[metric] !== undefined) {
        series.push({
          id: metric,
          name: this.data[metric].Label,
          type: metricSettings[metric].seriesType || "line",
          hidden: metricSettings[metric].hiddenByDefault || false,
          data: this.data[metric].Data.map((val, i) => ({
            x: new Date(this.data["time"].Data[i]),
            y: val,
          })),
        });

        yTooltips.push({
          formatter: metricSettings[metric].formatter,
        });

        const yaxisConfig = {};
        if (metricSettings[metric].yaxis) {
          Object.assign(yaxisConfig, metricSettings[metric].yaxis);
        }

        if (metricSettings[metric].formatterYaxis) {
          yaxisConfig.labels = {
            formatter:
              metricSettings[metric].labelFormatter ||
              metricSettings[metric].formatter,
          };
        }

        yaxis.push(yaxisConfig);
      }
    }

    let theme = this.colorMode;
    if (theme === "browser") {
      if (
        window.matchMedia &&
        window.matchMedia("(prefers-color-scheme: dark)").matches
      ) {
        theme = "dark";
      } else {
        theme = "light";
      }
    }

    let options = {
      theme: { mode: theme },
      chart: {
        height: "100%",
        background: "transparent",
        animations: { enabled: false },
        toolbar: { show: false },
        events: {
          mouseMove: (_event, _chartContext, config) => {
            if (config.dataPointIndex === -1) return;

            if (
              this.data["position"] !== undefined &&
              this.data["position"].Data[config.dataPointIndex]
            ) {
              let p = this.data["position"].Data[config.dataPointIndex];
              let el = document.createElement("div");
              el.setAttribute("data-lat", p[0]);
              el.setAttribute("data-lng", p[1]);

              this.mapElement.setMarker(el);
            }
          },
        },
      },
      legend: {
        position: "top",
        formatter: (seriesName, opts) => {
          if (series[opts.seriesIndex].type === "none") {
            return "";
          }

          return seriesName;
        },
        markers: { size: series.map((s) => (s.type !== `none` ? 12 : 0)) },
      },
      tooltip: {
        x: { format: "HH:mm" },
        y: yTooltips,
      },
      stroke: {
        width: 2,
        curve: "smooth",
      },
      plotOptions: {
        area: {
          fillTo: "end",
        },
      },
      series,
      xaxis: {
        labels: {
          formatter: (_val, ts, _opts) => {
            return new Date(ts).toLocaleTimeString(this.lang, {
              timeZone: this.tz,
            });
          },
        },
        type: "datetime",
      },
      yaxis,
    };

    this.chart = new ApexCharts(this, options);
    this.chart.render();
  }
}

customElements.define("workout-stats", WorkoutStats);
