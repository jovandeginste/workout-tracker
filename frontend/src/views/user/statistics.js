import ApexCharts from "apexcharts";
import { formatDuration } from "../../helpers.js";
import { colorForType, trendCutoffFromSince, trendSegments } from "../../trend";

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
    this.workoutTypes =
      JSON.parse(this.getAttribute("workout-types") || `[]`) || [];
    this.chartType = this.getAttribute("chart-type") === "bar" ? "bar" : "line";
    this.per = this.getAttribute("per") || "month";

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
    const theme = chartTheme().dark ? "dark" : "light";

    const isBar = this.chartType === "bar";

    const options = {
      theme: { mode: theme },
      chart: {
        width: "99%",
        height: 300,
        animations: { enabled: false },
        toolbar: { show: false },
        type: this.chartType,
      },
      dataLabels: { enabled: false },
      legend: {
        position: "top",
        onItemClick: { toggleDataSeries: false },
        onItemHover: { highlightDataSeries: false },
      },
      stroke: isBar ? { width: 0 } : { curve: "straight", width: 2 },
      markers: { size: 0 },
      tooltip: {
        x: { format: this.per === "year" ? "yyyy" : "MMM 'yy" },
      },
      xaxis: { type: "datetime" },
    };

    if (this.type === "durationSeconds") {
      this.setupDurationOptions(options);
    } else if (this.type.startsWith("local")) {
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

    const isElevation = this.type.includes("Up");

    const standardUnit =
      this.preferredUnits[
        this.type.includes("Distance")
          ? "distance"
          : isElevation
            ? "elevation"
            : "speed"
      ];
    const nauticalUnit = this.type.includes("Distance") ? "nm" : "kn";

    const formatStandardValue = isElevation
      ? (val) => Math.round(val) + " " + standardUnit
      : (val) => val.toFixed(1) + " " + standardUnit;

    const yaxis = [];

    series.forEach((s) => {
      if (!s.isNautical) {
        // Standard series
        yaxis.push({
          seriesName: firstStandardSeries.name,
          show: s.name === firstStandardSeries.name,
          title: { text: standardUnit },
          labels: {
            formatter: formatStandardValue,
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
    const entries = Object.entries(this.stats.buckets).filter(
      ([type]) =>
        this.workoutTypes.length === 0 || this.workoutTypes.includes(type),
    );

    return entries
      .map(([_, value]) => {
        const buckets = Object.values(value.buckets).filter(
          (e) => !this.filterNoDuration || e.duration > 0,
        );
        if (buckets.length === 0) return null;

        const isNautical = this.type.includes("Up")
          ? false
          : buckets[0].isNautical;

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
                unit: this.getValueUnit(e),
              };
            })
            .sort((a, b) => a.x.localeCompare(b.x)),
        };
      })
      .filter((e) => e !== null);
  }

  getValueUnit(bucket) {
    if (this.type.includes("Distance")) {
      return bucket.distanceUnit;
    }

    if (this.type.includes("Up")) {
      return this.preferredUnits.elevation;
    }

    return bucket.speedUnit;
  }
}

customElements.define("wt-stat", WtStatistic);
let statisticsScrollY = 0;

document.addEventListener("htmx:beforeSwap", (event) => {
  if (event.detail.target?.id === "statistics-content") {
    statisticsScrollY = window.scrollY;
  }
});

document.addEventListener("htmx:afterSwap", (event) => {
  if (event.detail.target?.id === "statistics-content") {
    requestAnimationFrame(() => window.scrollTo(0, statisticsScrollY));
  }
});

// Capture so the `All` checkbox is updated before HTMX serializes the form.
document.addEventListener(
  "change",
  (event) => {
    const input = event.target;
    if (!(input instanceof HTMLInputElement) || input.name !== "type") {
      return;
    }

    const group = input.closest(".workout-types");
    if (!group) {
      return;
    }

    const all = group.querySelector('input[name="type"][value=""]');

    if (!input.checked) {
      if (all && !group.querySelector('input[name="type"]:checked')) {
        all.checked = true;
      }
      return;
    }

    if (input.value === "") {
      group.querySelectorAll('input[name="type"]').forEach((other) => {
        if (other !== input) other.checked = false;
      });
    } else if (all) {
      all.checked = false;
    }
  },
  true,
);

class RouteSegmentTrendStatistic extends HTMLElement {
  connectedCallback() {
    this.data = JSON.parse(this.getAttribute("trend-data") || `[]`);
    this.since = this.getAttribute("since") || "1 year";
    this.workoutTypes =
      JSON.parse(this.getAttribute("workout-types") || `[]`) || [];
    this.trendBreakDetection =
      this.getAttribute("trend-break-detection") !== null;
    this.configUpdateRoute = this.getAttribute("config-update-route");
    this.speedUnit = this.getAttribute("speed-unit") || "";
    this.translations = JSON.parse(this.getAttribute("translations") || `{}`);

    if (!this.data || this.data.length === 0) {
      return;
    }

    this.appendControls();

    const el = document.createElement("div");
    this.appendChild(el);

    const series = this.getSeriesData();
    this.chart = new ApexCharts(el, {
      ...this.getChartOptions(series),
      series,
    });
    this.chart.render();
  }

  appendControls() {
    const controls = document.createElement("div");
    controls.className = "mb-3 flex flex-wrap justify-end gap-3";

    const selector = document.createElement("select");
    selector.className = "input-ellipsis";
    selector.innerHTML = `
      <option value="trend">${this.translations.trend || "Trend"}</option>
      <option value="break">${
        this.translations.trendBreak || "Trend break"
      }</option>
    `;
    selector.value = this.trendBreakDetection ? "break" : "trend";
    selector.addEventListener("change", () => {
      this.updateTrendBreakDetection(selector.value === "break");
    });
    controls.appendChild(selector);

    const resetButton = document.createElement("button");
    resetButton.type = "button";
    resetButton.textContent = this.translations.resetZoom || "Reset zoom";
    resetButton.addEventListener("click", () => this.resetZoom());
    controls.appendChild(resetButton);
    this.appendChild(controls);
  }

  async updateTrendBreakDetection(enabled) {
    this.trendBreakDetection = enabled;

    if (this.chart) {
      const series = this.getSeriesData();
      this.chart.updateOptions(
        { colors: this.seriesColors(series) },
        false,
        false,
      );
      this.chart.updateSeries(series);
    }

    if (!this.configUpdateRoute) {
      return;
    }

    const body = new URLSearchParams({
      route_segment_trend_break_detection: String(enabled),
    });

    await saveConfig(
      this.configUpdateRoute,
      body,
      "Failed to save route segment trend break detection preference",
    );
  }

  resetZoom() {
    this.chart?.resetSeries(false, true);
  }

  getChartOptions(series = []) {
    const theme = chartTheme().dark ? "dark" : "light";

    return {
      theme: { mode: theme },
      chart: {
        width: "99%",
        animations: { enabled: false },
        toolbar: { show: false },
        type: "line",
        height: 350,
        zoom: {
          enabled: true,
          type: "x",
          autoScaleYaxis: true,
        },
      },
      dataLabels: { enabled: false },
      colors: this.seriesColors(series),
      legend: {
        position: "top",
        onItemClick: { toggleDataSeries: false },
        onItemHover: { highlightDataSeries: false },
      },
      markers: { size: 4, strokeWidth: 0, hover: { sizeOffset: 1 } },
      stroke: { curve: "straight", width: 2 },
      tooltip: {
        x: { format: "MMM 'yy" },
        y: {
          formatter: (val) => `${Number(val).toFixed(1)} ${this.speedUnit}`,
        },
      },
      xaxis: this.getXAxisOptions(),
      yaxis: {
        labels: {
          formatter: (val) => `${Number(val).toFixed(1)} ${this.speedUnit}`,
        },
      },
    };
  }

  seriesColors(series) {
    return series.map((serie) => serie.color);
  }

  getXAxisOptions() {
    const cutoff = trendCutoffFromSince(this.since);

    return {
      type: "datetime",
      min: cutoff ?? undefined,
      max: cutoff === null ? undefined : Date.now(),
    };
  }

  getSeriesData() {
    const entries = Object.entries(this.pointsByType()).filter(
      ([type]) =>
        this.workoutTypes.length === 0 || this.workoutTypes.includes(type),
    );

    return entries.flatMap(([type, points]) => {
      const sortedPoints = points.sort((a, b) => a.x - b.x);
      const cutoff = trendCutoffFromSince(this.since);
      const visiblePoints = sortedPoints.filter(
        (point) => cutoff === null || point.x >= cutoff,
      );
      const segments = trendSegments(visiblePoints, this.trendBreakDetection);
      const typeLabel = this.translations[type] || type;
      const color = colorForType(type);

      const series = [
        {
          name: typeLabel,
          type: "scatter",
          color,
          data: visiblePoints,
        },
      ];

      segments.forEach((trendData) => {
        if (trendData.length <= 1) {
          return;
        }

        series.push({
          name: `${typeLabel} (${this.translations.trend || "trend"})`,
          type: "line",
          color,
          data: trendData,
        });
      });

      return series;
    });
  }

  pointsByType() {
    return this.data.reduce((result, item) => {
      const date = new Date(item.date).valueOf();
      const speed = parseFloat(item.speed);

      if (Number.isNaN(date) || Number.isNaN(speed)) {
        return result;
      }

      const type = item.type || "unknown";
      result[type] ||= [];
      result[type].push({ x: date, y: speed });

      return result;
    }, {});
  }
}

customElements.define("route-segment-trend-stat", RouteSegmentTrendStatistic);
