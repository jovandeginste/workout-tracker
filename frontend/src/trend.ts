export interface ChartPoint {
  x: number;
  y: number;
}

export interface Regression {
  slope: number;
  intercept: number;
  baseX: number;
}

export interface TrendBreak {
  pivot: number;
  pValue: number;
}

export const DAY_MS = 24 * 60 * 60 * 1000;
export const MAX_TREND_BREAK_POINTS = 200;

const PALETTE = [
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

export function colorForType(type: string): string {
  let hash = 0;
  for (let i = 0; i < type.length; i++) {
    hash = (hash * 31 + type.charCodeAt(i)) % PALETTE.length;
  }

  return PALETTE[hash];
}

export function average(points: ChartPoint[]): number {
  return points.reduce((sum, point) => sum + point.y, 0) / points.length;
}

export function variance(points: ChartPoint[], avg: number): number {
  if (points.length < 2) {
    return 0;
  }

  return (
    points.reduce((sum, point) => sum + (point.y - avg) ** 2, 0) /
    (points.length - 1)
  );
}

export function normalCDF(value: number): number {
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

export function canShowTrend(points: ChartPoint[]): boolean {
  return points.length >= 3;
}

export function linearRegression(points: ChartPoint[]): Regression {
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

  return {
    slope,
    intercept: (sums.y - slope * sums.x) / count,
    baseX,
  };
}

export function trendLineData(points: ChartPoint[]): ChartPoint[] {
  if (!canShowTrend(points)) {
    return [];
  }

  const regression = linearRegression(points);
  const first = points[0];
  const last = points[points.length - 1];

  return [first, last].map((point) => ({
    x: point.x,
    y: regression.slope * (point.x - regression.baseX) + regression.intercept,
  }));
}

export function averageLineData(points: ChartPoint[]): ChartPoint[] {
  if (!canShowTrend(points)) {
    return [];
  }

  const y = average(points);

  return [
    { x: points[0].x, y },
    { x: points[points.length - 1].x, y },
  ];
}

export function trendBreakPValue(
  before: ChartPoint[],
  after: ChartPoint[],
): number {
  const beforeAverage = average(before);
  const afterAverage = average(after);
  const beforeVariance = variance(before, beforeAverage);
  const afterVariance = variance(after, afterAverage);
  const standardError = Math.sqrt(
    beforeVariance / before.length + afterVariance / after.length,
  );

  if (standardError === 0) {
    return beforeAverage === afterAverage ? 1 : 0;
  }

  const testStatistic = Math.abs(beforeAverage - afterAverage) / standardError;

  return 2 * (1 - normalCDF(testStatistic));
}

export function detectTrendBreak(
  points: ChartPoint[],
  enabled: boolean,
): TrendBreak | null {
  const minimumPointsPerSide = 3;

  if (
    !enabled ||
    points.length < minimumPointsPerSide * 2 ||
    points.length > MAX_TREND_BREAK_POINTS
  ) {
    return null;
  }

  let bestBreak: TrendBreak | null = null;
  let bestGroupSize = 0;

  for (
    let pivot = minimumPointsPerSide;
    pivot <= points.length - minimumPointsPerSide;
    pivot++
  ) {
    const before = points.slice(0, pivot);
    const after = points.slice(pivot);
    const pValue = trendBreakPValue(before, after);

    if (pValue > 0.05) {
      continue;
    }

    const groupSize = Math.min(before.length, after.length);

    if (
      groupSize > bestGroupSize ||
      (groupSize === bestGroupSize && (!bestBreak || pValue < bestBreak.pValue))
    ) {
      bestGroupSize = groupSize;
      bestBreak = { pivot, pValue };
    }
  }

  return bestBreak;
}

export function trendSegments(
  points: ChartPoint[],
  breakDetection: boolean,
): ChartPoint[][] {
  if (!canShowTrend(points)) {
    return [];
  }

  const trendLine = trendLineData(points);

  if (!breakDetection) {
    return [trendLine].filter((segment) => segment.length > 0);
  }

  const trendBreak = detectTrendBreak(points, breakDetection);

  if (!trendBreak) {
    return [trendLine].filter((segment) => segment.length > 0);
  }

  const before = points.slice(0, trendBreak.pivot);
  const after = points.slice(trendBreak.pivot);
  if (!canShowTrend(before) || !canShowTrend(after)) {
    return [trendLine].filter((segment) => segment.length > 0);
  }

  return [averageLineData(before), averageLineData(after)].filter(
    (segment) => segment.length > 0,
  );
}

export function trendCutoffFromSince(since: string): number | null {
  if (since === "forever") {
    return null;
  }

  const match = since.match(/^(\d+)\s+(day|month|year)s?$/);
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

export function trendCutoffFromPeriod(period: string): number | null {
  const now = new Date();

  if (period === "all") {
    return null;
  }

  if (period === "year") {
    return new Date(now.getFullYear(), 0, 1).valueOf();
  }

  return now.valueOf() - Number(period) * DAY_MS;
}
