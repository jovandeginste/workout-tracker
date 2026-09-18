import ApexCharts from "apexcharts";
import { formatDuration } from "../../helpers.js";

const DAY_MS = 24 * 60 * 60 * 1000;
const MAX_TREND_BREAK_POINTS = 200;
const DAYS_PER_MONTH = 30.4375;

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
    this.workoutType = this.getAttribute("workout-type") || "";

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
        height: 300,
        animations: { enabled: false },
        toolbar: { show: false },
        type: "line",
      },
      dataLabels: { enabled: false },
      legend: {
        position: "top",
        onItemClick: { toggleDataSeries: false },
        onItemHover: { highlightDataSeries: false },
      },
      stroke: { curve: "straight", width: 2 },
      markers: { size: 0 },
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
    const entries = Object.entries(this.stats.buckets).filter(
      ([type]) => !this.workoutType || type === this.workoutType,
    );

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

class RouteSegmentTrendStatistic extends HTMLElement {
  connectedCallback() {
    this.data = JSON.parse(this.getAttribute("data") || `[]`);
    this.since = this.getAttribute("since") || "1 year";
    this.workoutType = this.getAttribute("workout-type") || "";
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
      <option value="break">${this.translations.trendBreak || "Trend break"}</option>
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

    try {
      const response = await fetch(this.configUpdateRoute, {
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

  resetZoom() {
    this.chart?.resetSeries(false, true);
  }

  getChartOptions(series = []) {
    const theme =
      window.matchMedia &&
      window.matchMedia("(prefers-color-scheme: dark)").matches
        ? "dark"
        : "light";

    return {
      theme: { mode: theme },
      chart: {
        width: "99%",
        animations: { enabled: false },
        toolbar: { show: false },
        type: "line",
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

  colorForType(type) {
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

  getXAxisOptions() {
    const cutoff = this.trendCutoff();

    return {
      type: "datetime",
      min: cutoff ?? undefined,
      max: cutoff === null ? undefined : Date.now(),
    };
  }

  getSeriesData() {
    const entries = Object.entries(this.pointsByType()).filter(
      ([type]) => !this.workoutType || type === this.workoutType,
    );

    return entries.flatMap(([type, points]) => {
      const sortedPoints = points.sort((a, b) => a.x - b.x);
      const cutoff = this.trendCutoff();
      const visiblePoints = sortedPoints.filter(
        (point) => cutoff === null || point.x >= cutoff,
      );
      const trendSegments = this.trendSegmentsFor(visiblePoints);
      const typeLabel = this.translations[type] || type;
      const color = this.colorForType(type);

      const series = [
        {
          name: typeLabel,
          type: "scatter",
          color,
          data: visiblePoints,
        },
      ];

      trendSegments.forEach((trendData, index) => {
        if (trendData.length <= 1) {
          return;
        }

        series.push({
          name: `${typeLabel} ${this.translations.trend || "trend"}`,
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

  trendSegmentsFor(points) {
    if (!this.canShowTrend(points)) {
      return [];
    }

    if (!this.trendBreakDetection) {
      return [this.trendDataFor(points)];
    }

    const trendBreak = this.detectTrendBreak(points);
    if (!trendBreak) {
      return [this.trendDataFor(points)];
    }

    const before = points.slice(0, trendBreak.pivot);
    const after = points.slice(trendBreak.pivot);
    if (!this.canShowTrend(before) || !this.canShowTrend(after)) {
      return [this.trendDataFor(points)];
    }

    return [this.averageDataFor(before), this.averageDataFor(after)].filter(
      (segment) => segment.length > 0,
    );
  }

  trendDataFor(points) {
    if (!this.canShowTrend(points)) {
      return [];
    }

    const regression = this.linearRegression(points);

    return [
      {
        x: points[0].x,
        y:
          regression.slope * (points[0].x - regression.baseX) +
          regression.intercept,
      },
      {
        x: points[points.length - 1].x,
        y:
          regression.slope * (points[points.length - 1].x - regression.baseX) +
          regression.intercept,
      },
    ];
  }

  averageDataFor(points) {
    if (!this.canShowTrend(points)) {
      return [];
    }

    const y = this.average(points);

    return [
      { x: points[0].x, y },
      { x: points[points.length - 1].x, y },
    ];
  }

  canShowTrend(points) {
    if (points.length < 5) {
      return false;
    }

    const first = points[0].x;
    const last = points[points.length - 1].x;
    const months = Math.max(1, (last - first) / (DAYS_PER_MONTH * DAY_MS));

    return points.length / months >= 1;
  }

  detectTrendBreak(points) {
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

    let bestBreak = null;
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

  trendBreakPValue(before, after) {
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

  average(points) {
    return points.reduce((sum, point) => sum + point.y, 0) / points.length;
  }

  variance(points, average) {
    if (points.length < 2) {
      return 0;
    }

    return (
      points.reduce((sum, point) => sum + (point.y - average) ** 2, 0) /
      (points.length - 1)
    );
  }

  normalCDF(value) {
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

  trendCutoff() {
    if (this.since === "forever") {
      return null;
    }

    const match = this.since.match(/^(\d+)\s+(day|month|year)s?$/);
    if (!match) {
      return Date.now() - 365 * DAY_MS;
    }

    const count = Number(match[1]);
    const unit = match[2];
    const date = new Date();

    if (unit === "year") {
      date.setFullYear(date.getFullYear() - count);
    } else if (unit === "month") {
      date.setMonth(date.getMonth() - count);
    } else {
      date.setDate(date.getDate() - count);
    }

    return date.valueOf();
  }

  linearRegression(points) {
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
}

customElements.define("route-segment-trend-stat", RouteSegmentTrendStatistic);
