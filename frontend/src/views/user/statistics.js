import ApexCharts from "apexcharts";
import { formatDuration } from "../../helpers.js";

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

    const options = this.getChartOptions();
    const series = this.getSeriesData();

    const el = document.createElement("div");
    this.appendChild(el);

    new ApexCharts(el, {
      ...options,
      series,
    }).render();
  }

  getChartOptions() {
    const theme =
      window.matchMedia &&
      window.matchMedia("(prefers-color-scheme: dark)").matches
        ? "dark"
        : "light";

    const options = {
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

    if (this.type === "durationSeconds") {
      this.setupDurationOptions(options);
    } else if (
      this.type.startsWith("localDistance") ||
      this.type.startsWith("localAverageSpeed") ||
      this.type.startsWith("localMaxSpeed")
    ) {
      this.setupValueOptions(options);
    }

    return options;
  }

  setupDurationOptions(options) {
    options.tooltip.y = [
      {
        formatter: (val) => formatDuration(val),
      },
    ];
    options.yaxis = [
      {
        labels: {
          formatter: (val) => formatDuration(val),
        },
      },
    ];
  }

  setupValueOptions(options) {
    const series = this.getSeriesData();
    if (series.length === 0) return;

    const hasStandard = series.some((s) => !s.isNautical);
    const hasNautical = series.some((s) => s.isNautical);

    const firstStandardSeries = series.find((s) => !s.isNautical);
    const firstNauticalSeries = series.find((s) => s.isNautical);

    const standardUnit =
      this.preferredUnits[
        this.type.includes("Distance") ? "distance" : "speed"
      ];
    const nauticalUnit = this.type.includes("Distance") ? "nm" : "kn";

    const yaxis = [];

    series.forEach((s) => {
      if (!s.isNautical) {
        // Standard series
        yaxis.push({
          seriesName: firstStandardSeries.name,
          show: s.name === firstStandardSeries.name,
          title: { text: standardUnit },
          labels: {
            formatter: (val) => val.toFixed(1) + " " + standardUnit,
          },
        });
      } else {
        // Nautical series
        yaxis.push({
          seriesName: firstNauticalSeries.name,
          show: s.name === firstNauticalSeries.name,
          opposite: hasStandard,
          title: { text: nauticalUnit },
          labels: {
            formatter: (val) => val.toFixed(1) + " " + nauticalUnit,
          },
        });
      }
    });

    options.yaxis = yaxis;

    options.tooltip.y = series.map((s) => ({
      formatter: (val, { dataPointIndex, w, seriesIndex }) => {
        const unit = w.config.series[seriesIndex].data[dataPointIndex].unit;
        return val + (unit ? " " + unit : "");
      },
    }));
  }

  getSeriesData() {
    const unitType = this.type.includes("Distance")
      ? "distanceUnit"
      : "speedUnit";
    const entries = Object.entries(this.stats.buckets);

    return entries
      .map(([_, value]) => {
        const buckets = Object.values(value.buckets).filter(
          (e) => !this.filterNoDuration || e.duration > 0,
        );
        if (buckets.length === 0) return null;

        const isNautical = buckets[0].isNautical;

        return {
          name: value.localWorkoutType,
          isNautical: isNautical,
          data: buckets
            .map((e) => {
              const val = e[this.type];
              let numericValue = val;
              if (typeof val === "string") {
                numericValue = parseFloat(val.split(" ")[0]);
              }

              return {
                x: e.bucket,
                y: numericValue,
                unit: e[unitType],
              };
            })
            .sort((a, b) => a.x.localeCompare(b.x)),
        };
      })
      .filter((e) => e !== null);
  }
}

customElements.define("wt-stat", WtStatistic);
