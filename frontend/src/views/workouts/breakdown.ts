import { html, LitElement, PropertyValues, TemplateResult } from "lit";
import { customElement, property } from "lit/decorators.js";
import { formatDuration } from "../../helpers.js";

@customElement("workout-breakdown")
class WorkoutBreakdown extends LitElement {
  private activeItem: HTMLElement | null = null;

  @property()
  intervalDistance = 1;

  @property()
  mapId = "";

  @property()
  chartId = "";

  @property()
  workoutStatsId = "";

  @property()
  dataEl = "";

  @property()
  preferredUnitsEl = "";

  mapEl: HTMLElement | null = null;
  chartEl: HTMLElement | null = null;
  workoutStatsEl: HTMLElement | null = null;
  data: any = {};
  preferredUnits: any = {};
  availableMetrics: Record<string, string> = {
    distance: "",
    duration: "",
    speed: "",
    elevation: "",
    "heart-rate": "",
    cadence: "",
    temperature: "",
  };

  protected createRenderRoot() {
    return this;
  }

  willUpdate(changedProperties: PropertyValues<this>) {
    if (changedProperties.has("mapId")) {
      this.mapEl = document.getElementById(this.mapId);
    }

    if (changedProperties.has("chartId")) {
      this.chartEl = document.getElementById(this.chartId);
    }

    if (changedProperties.has("workoutStatsId")) {
      this.workoutStatsEl = document.getElementById(this.workoutStatsId);
    }

    if (changedProperties.has("dataEl")) {
      const dataElement = document.getElementById(this.dataEl);
      if (dataElement) {
        this.data = JSON.parse(dataElement.textContent || "{}");
      }
    }

    if (changedProperties.has("preferredUnitsEl")) {
      const unitsElement = document.getElementById(this.preferredUnitsEl);
      if (unitsElement) {
        this.preferredUnits = JSON.parse(unitsElement.textContent || "{}");
      }

      this.availableMetrics = {
        distance: this.preferredUnits.distance || "",
        duration: "",
        speed: this.preferredUnits.speed || "",
        elevation: this.preferredUnits.elevation || "",
        "heart-rate": this.preferredUnits.heartRate || "",
        cadence: this.preferredUnits.cadence || "",
        temperature: this.preferredUnits.temperature || "",
      };
    }
  }

  tableHeader() {
    const header = html`<tr class="breakdown-header">
      <th></th>
      <th></th>
      ${Object.keys(this.availableMetrics).map((metric) => {
      if (this.data[metric] !== undefined) {
        const col = this.data[metric].Label;
        if (metric === "speed") {
          // TODO: localize "Tempo"
          return html`<th>${col}</th>
              <th>Tempo</th>`;
        }
        if (metric === "elevation") {
          // TODO: localize "Tempo"
          return html`<th colspan="2">${col}</th>`;
        }
        return html`<th>${col}</th>`;
      }
      return "";
    })}
    </tr>`;
    return header;
  }

  tableData() {
    let currentDistance = Math.floor(+this.data.distance.Data[0] || 0);
    let intervalValues = {};
    const items = [];
    for (let i = 0; i < this.data.time.Data.length; i++) {
      const distance = +this.data.distance.Data[i] || 0;
      if (distance >= currentDistance + this.intervalDistance) {
        items.push([currentDistance, intervalValues]);
        currentDistance += this.intervalDistance;
        intervalValues = {};
      }

      for (const metric of Object.keys(this.data)) {
        const value = this.data[metric].Data[i];
        if (value !== undefined) {
          if (!intervalValues[metric]) {
            intervalValues[metric] = [];
          }

          intervalValues[metric].push(value);
        }
      }
    }

    if (Object.keys(intervalValues).length !== 0) {
      items.push([currentDistance, intervalValues]);
    }

    // Marks best and worst values by speed
    let fastest = [0, 0];
    let slowest = [0, 0];
    for (let i = 0; i < items.length; i++) {
      const values = items[i][1];
      if (values.speed && values.speed.length > 0) {
        const speed =
          values.speed.reduce((a, b) => a + b, 0) / values.speed.length;
        if (speed > fastest[1]) {
          fastest = [i, speed];
        }

        if (speed < slowest[1] || slowest[1] === 0) {
          slowest = [i, speed];
        }
      }
    }
    items[fastest[0]][1].best = true;
    items[slowest[0]][1].worst = true;

    const rows: TemplateResult[] = [];
    for (const [distance, values] of items) {
      rows.push(this.tableRow(distance, values));
    }
    return rows;
  }

  tableRow(distance: number, intervalValues) {
    return html`<tr
      class="cursor-pointer"
      @click="${(e) => this.itemClick(e, intervalValues)}"
    >
      ${this.tableRecordCell(intervalValues)}
      <td>${distance / this.intervalDistance + 1}</td>
      ${Object.keys(this.availableMetrics).map((metric) => {
      if (this.data[metric] === undefined) {
        return "";
      }

      return this.tableCell(distance, metric, intervalValues);
    })}
    </tr>`;
  }

  tableRecordCell(intervalValues) {
    if (intervalValues.best) {
      return html`<td class="text-right">
        <span class="text-green-500"
          ><span class="icon-decoration icon-[fa6-solid--arrow-up-long]"></span
        ></span>
      </td>`;
    }

    if (intervalValues.worst) {
      return html`<td class="text-right">
        <span class="text-orange-600"
          ><span
            class="icon-decoration icon-[fa6-solid--arrow-down-long]"
          ></span
        ></span>
      </td>`;
    }

    return html`<td></td>`;
  }

  tableCell(distance: number, metric: string, intervalValues) {
    const displayDecimals = ["speed", "elevation", "temperature"];
    const mData = intervalValues[metric].filter((v: any) => v !== null);
    if (metric === "duration") {
      return html`<td>${formatDuration(mData[mData.length - 1])}</td>`;
    }

    if (metric === "distance") {
      const lastDist = +mData[mData.length - 1];
      if (lastDist < distance + this.intervalDistance - 0.05) {
        return html`<td>
          ${lastDist.toFixed(2)} ${this.availableMetrics[metric]}
        </td>`;
      } else {
        return html`<td>
          ${(distance + this.intervalDistance).toFixed(2)}
          ${this.availableMetrics[metric]}
        </td>`;
      }
    }

    if (metric === "elevation" && mData.length > 0) {
      const elevationChange = mData.reduce(
        (a, c) => {
          const elevationGain = c - a[0];
          if (elevationGain > 0) {
            return [c, a[1] + elevationGain, a[2]];
          }
          return [c, a[1], a[2] + Math.abs(elevationGain)];
        },
        [mData[0], 0, 0],
      );

      return html`<td>
          <span class="icon-decoration icon-[fa6-solid--chevron-up]"></span>
          ${elevationChange[1].toFixed(2)} ${this.availableMetrics[metric]}
        </td>
        <td>
          <span class="icon-decoration icon-[fa6-solid--chevron-down]"></span>
          ${elevationChange[2].toFixed(2)} ${this.availableMetrics[metric]}
        </td>`;
    }

    if (mData.length === 0) {
      return html`<td>-</td>`;
    }

    const value =
      mData.reduce((a, b) => a + b, 0) / intervalValues[metric].length;
    if (metric === "speed") {
      const pace = value > 0 ? 3600 / value : 0;
      const seconds = Math.round(pace % 60)
        .toString()
        .padStart(2, "0");
      // TODO: localize "min" unit
      return html`<td>
          ${value.toFixed(2)} ${this.availableMetrics[metric] || ""}
        </td>
        <td>
          ${Math.floor(pace / 60)}:${seconds}
          min/${this.preferredUnits.distance || ""}
        </td>`;
    }

    if (displayDecimals.includes(metric)) {
      return html`<td>
        ${value.toFixed(2)} ${this.availableMetrics[metric] || ""}
      </td>`;
    }

    return html`<td>
      ${value.toFixed(0)} ${this.availableMetrics[metric] || ""}
    </td>`;
  }

  render() {
    const totalDistance = +this.data.distance?.Data?.slice(-1)[0] || 0;
    const intervals = [1, 2, 5, 10, 25].filter((d) => d < totalDistance);

    return html`
      <table class="breakdown-table">
        <thead>
          ${this.tableHeader()}
        </thead>
        <tbody class="whitespace-nowrap font-mono">
          ${this.tableData()}
        </tbody>
      </table>
      <div class="flex justify-end py-3">
        <nav class="isolate inline-flex">
          ${intervals.map(interval => {
            return html`<a
              href="#"
              class="relative inline-flex items-center px-3 py-2 text-sm font-semibold text-gray-200 inset-ring inset-ring-gray-700 hover:bg-white/5 focus:z-20 focus:outline-offset-0 ${this.intervalDistance === interval ? 'bg-indigo-500 text-white' : ''}"
              @click=${(e: Event) => {
                e.preventDefault();
                this.setActiveItem(null);
                this.intervalDistance = interval;
              }}
              >${interval} ${this.preferredUnits.distance || ''}</a
            >`;
          })}
        </nav>
      </div>
    `;
  }

  private itemClick(e: Event, values: any) {
    if (this.activeItem === e.currentTarget) {
      this.setActiveItem(null);
    } else {
      this.setActiveItem(e.currentTarget as HTMLElement, values);
      this.mapEl.fitSegmentBounds();
    }
  }

  private setActiveItem(item: HTMLElement, values?: any) {
    if (this.activeItem) {
      this.activeItem.classList.remove(`active`);
    }

    this.activeItem = item;
    if (this.activeItem) {
      this.mapEl?.scrollIntoView({ behavior: `smooth` });
      this.activeItem.classList.add(`active`);
      this.mapEl?.setSegment("", values);
      this.chartEl?.chart.zoomX(
        new Date(values["time"][0]).getTime(),
        new Date(values["time"][values["time"].length - 1]).getTime(),
      );
    } else {
      this.mapEl.clearSegment();
      this.chartEl.chart.zoomX();
    }
  }
}
