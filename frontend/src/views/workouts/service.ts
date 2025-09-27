interface WorkoutDataEntry {
  label: string;
  data: any[];
}

export type WorkoutData = Record<string, WorkoutDataEntry>;

export interface IntervalStats {
  totalDistance: number;
  totalTime: number;
  pauseTime: number;
  speed?: {
    avg: number;
    max: number;
    min: number;
  };
  elevation?: {
    gain: number;
    loss: number;
    min: number;
    max: number;
    avgSlope: number;
    maxSlope: number;
    minSlope: number;
    vam: number;
  };
  heartRate?: {
    avg: number;
    max: number;
    min: number;
  };
  cadence?: {
    avg: number;
    max: number;
    min: number;
  };
  temperature?: {
    avg: number;
    max: number;
    min: number;
  };
}

export interface PreferredUnits {
  distance: string;
  speed: string;
  elevation: string;
  heartRate: string;
  cadence: string;
  temperature: string;
  tempo: string;
}

// Singleton service to manage workout data
export class WorkoutService {
  private _preferredUnits: PreferredUnits;
  public get preferredUnits() {
    return this._preferredUnits;
  }

  private _workoutData: WorkoutData;
  public get workoutData() {
    return this._workoutData;
  }

  /**
   * Interval (in seconds) between datapoints, used to estimate pause time
   */
  private datapointInterval: number;

  public constructor() {
    if (globalThis.workoutService) {
      return globalThis.workoutService;
    }
    globalThis.workoutService = this;
    this.loadWorkout();
    this.updateDatapointInterval();

    console.log(this.getIntervalStats(), this._workoutData);
  }

  private loadWorkout() {
    const workoutDataJSON = document.getElementById("workout-data").textContent;
    this._workoutData = JSON.parse(workoutDataJSON);

    const preferredUnitsJSON = document.getElementById(
      "workout-preferred-units",
    ).textContent;
    this._preferredUnits = JSON.parse(preferredUnitsJSON);
  }

  private updateDatapointInterval() {
    // Get the 75th percentile of time intervals between datapoints
    const times = this._workoutData["time"].data;
    const intervals = [];
    for (let i = 1; i < times.length; i++) {
      intervals.push(Date.parse(times[i]) - Date.parse(times[i - 1]));
    }
    intervals.sort((a, b) => a - b);
    this.datapointInterval =
      (intervals[Math.floor(intervals.length * 0.75)] / 1000) * 2;
  }

  public getIntervalStats(sIdx: number = 0, eIdx?: number): IntervalStats {
    if (eIdx === undefined) {
      eIdx = this._workoutData["time"].data.length - 1;
    }

    const stats: IntervalStats = {
      totalDistance: 0,
      totalTime: 0,
      pauseTime: 0,
    };

    const timeData = this._workoutData["time"].data.map((t) => new Date(t));
    const speedData = this._workoutData["speed"].data;
    const distanceData = this._workoutData["distance"].data;
    for (let i = sIdx + 1; i <= eIdx; i++) {
      const duration =
        (timeData[i].getTime() - timeData[i - 1].getTime()) / 1000;
      if (duration <= this.datapointInterval && speedData[i] > 0) {
        stats.totalTime += duration;
        stats.totalDistance += distanceData[i] - distanceData[i - 1];
      } else {
        stats.pauseTime += duration;
      }
    }

    const elevationData = this._workoutData["elevation"]?.data;
    if (elevationData) {
      stats.elevation = {
        gain: 0,
        loss: 0,
        min: Math.min(...elevationData),
        max: Math.max(...elevationData),
        avgSlope: 0,
        maxSlope: 0,
        minSlope: 0,
        vam: 0,
      };
    }

    const slopeData = this._workoutData["slope"]?.data;
    if (slopeData && elevationData) {
      stats.elevation.maxSlope = Math.max(...slopeData);
      stats.elevation.minSlope = Math.min(...slopeData);
      stats.elevation.avgSlope = slopeData
        .map((s, i) => s * (i === 0 ? 0 : distanceData[i] - distanceData[i - 1]))
        .reduce((a, b) => a + b, 0) / stats.totalDistance;
    }

    const heartRateData = this._workoutData["heart-rate"]?.data;
    if (heartRateData) {
      const validData = heartRateData.filter((hr) => hr > 0);
      stats.heartRate = {
        avg: validData.reduce((a, b) => a + b, 0) / validData.length,
        max: Math.max(...validData),
        min: Math.min(...validData),
      };
    }

    const cadenceData = this._workoutData["cadence"]?.data;
    if (cadenceData) {
      const movingData = cadenceData.filter((c, i) => c > 0 && speedData[i] > 0);
      const validData = cadenceData.filter((c) => c !== null && c !== undefined);
      stats.cadence = {
        avg: movingData.reduce((a, b) => a + b, 0) / movingData.length,
        max: Math.max(...validData),
        min: Math.min(...validData),
      };
    }

    const temperatureData = this._workoutData["temperature"]?.data;
    if (temperatureData) {
      stats.temperature = {
        avg: temperatureData.filter((t) => t !== null && t !== undefined).reduce((a, b) => a + b, 0) / temperatureData.length,
        max: Math.max(...temperatureData),
        min: Math.min(...temperatureData),
      };
    }

    return stats;
  }
}
