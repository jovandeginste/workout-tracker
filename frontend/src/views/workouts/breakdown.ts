import { html, LitElement, PropertyValues, TemplateResult } from "lit";
import { customElement, property } from 'lit/decorators.js';

@customElement("workout-breakdown")
class WorkoutBreakdown extends LitElement {
  @property()
  intervalDistance = 1;

  @property()
  mapId = '';

  @property()
  workoutStatsId = '';

  @property()
  dataEl = '';

  @property()
  preferredUnitsEl = '';

  mapEl: HTMLElement | null = null;
  workoutStatsEl: HTMLElement | null = null;
  data: any = {};
  preferredUnits: any = {};
  availableMetrics: Record<string, string> = {
    distance: '',
    speed: '',
    duration: '',
    elevation: '',
    'heart-rate': '',
    cadence: '',
    temperature: '',
  };

  protected createRenderRoot() {
    return this;
  }

  willUpdate(changedProperties: PropertyValues<this>) {
    if (changedProperties.has('mapId')) {
      this.mapEl = document.getElementById(this.mapId);
    }

    if (changedProperties.has('workoutStatsId')) {
      this.workoutStatsEl = document.getElementById(this.workoutStatsId);
    }

    if (changedProperties.has('dataEl')) {
      const dataElement = document.getElementById(this.dataEl);
      if (dataElement) {
        this.data = JSON.parse(dataElement.textContent || '{}');
      }
    }

    if (changedProperties.has('preferredUnitsEl')) {
      const unitsElement = document.getElementById(this.preferredUnitsEl);
      if (unitsElement) {
        this.preferredUnits = JSON.parse(unitsElement.textContent || '{}');
      }

      this.availableMetrics = {
        distance: this.preferredUnits.distance || '',
        speed: this.preferredUnits.speed || '',
        duration: '',
        elevation: this.preferredUnits.elevation || '',
        'heart-rate': this.preferredUnits.heartRate || '',
        cadence: this.preferredUnits.cadence || '',
        temperature: this.preferredUnits.temperature || '',
      };
    }
  }

  tableHeader() {
    const header = html`<tr class="breakdown-header">
      <th></th>
      ${Object.keys(this.availableMetrics).map(metric => {
        if (this.data[metric] !== undefined) {
          const col = this.data[metric].Label;
          return html`<th>${col}</th>`;
        }
        return '';
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

    // Marks best and worst values by speed
    let fastest = [0, 0];
    let slowest = [0, 0];
    for (let i = 0; i < items.length; i++) {
      const values = items[i][1];
      if (values.speed && values.speed.length > 0) {
        const speed = values.speed.reduce((a, b) => a + b, 0) / values.speed.length;
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

  tableRow(distance, intervalValues) {
    return html`<tr class="breakdown-item">
      ${this.tableRecordCell(intervalValues)}
      ${Object.keys(this.availableMetrics).map(metric => {
        if (this.data[metric] === undefined) {
          return '';
        }

        try {
          return this.tableCell(distance, metric, intervalValues);
        } catch (e) {
          return html`<td>-</td>`;
        }
      })}
    </tr>`;
  }

  tableRecordCell(intervalValues) {
    if (intervalValues.best) {
      return html`<td class="text-right"><span class="text-green-500"><span class="icon-decoration icon-[fa6-solid--arrow-up-long]"></span></span></td>`;
    }

    if (intervalValues.worst) {
      return html`<td class="text-right"><span class="text-orange-600"><span class="icon-decoration icon-[fa6-solid--arrow-down-long]"></span></span></td>`;
    }

    return html`<td></td>`;
  }

  tableCell(distance, metric, intervalValues) {
    const displayDecimals = ["speed", "elevation", "temperature"];
    const mData = intervalValues[metric].filter((v: any) => v !== null);
    if (metric === "distance") {
      const lastDist = mData[mData.length - 1];
      if (lastDist < distance + this.intervalDistance - 0.05) {
        return html`<td>${lastDist.toFixed(2)} ${this.availableMetrics[metric]}</td>`;
      } else {
        return html`<td>${(distance + this.intervalDistance).toFixed(2)} ${this.availableMetrics[metric]}</td>`;
      }
    }

    if (metric === "duration") {
      if (mData.length === 0) {
        return html`<td>-</td>`;
      }

      const totalDuration = mData[mData.length - 1] - mData[0];
      const minutes = Math.floor(totalDuration / 60);
      let seconds: number | string = totalDuration % 60;
      if (seconds < 10) {
        seconds = `0${seconds}`;
      }

      return html`<td>${minutes}:${seconds} min/${this.preferredUnits.distance}</td>`;
    }

    if (mData.length === 0) {
      return html`<td>-</td>`;
    }

    const value = mData.reduce((a, b) => a + b, 0) / intervalValues[metric].length;
    if (displayDecimals.includes(metric)) {
      return html`<td>${value.toFixed(2)} ${this.availableMetrics[metric] || ""}</td>`;
    }

    return html`<td>${value.toFixed(0)} ${this.availableMetrics[metric] || ""}</td>`;
  }

  render() {
    return html`
      <table class="breakdown-table">
        <thead>${this.tableHeader()}</thead>
        <tbody class="whitespace-nowrap font-mono">${this.tableData()}</tbody>
      </table>
    `;
  }
  /*
  connectedCallback() {
    this.render();

    this.querySelectorAll(`.breakdown-item`).forEach((item) => {
      item.addEventListener("mouseover", () => this.itemMouseOver(item));
      item.addEventListener("mouseout", this.itemMouseOut.bind(this));
      item.addEventListener("click", () => this.itemClick(item));
    });
  }

  itemClick(item) {
    if (this.activeItem === item) {
      this.setActiveItem(null);
    } else {
      this.setActiveItem(item);
    }
  }

  itemMouseOver(item) {
    if (!this.activeItem) {
      this.mapEl.setMarker(item);
    }
  }

  itemMouseOut() {
    if (!this.activeItem) {
      this.mapEl.clearMarker();
    }
  }

  setActiveItem(item) {
    if (this.activeItem) {
      this.activeItem.classList.remove(`active`);
    }

    this.activeItem = item;
    if (this.activeItem) {
      this.mapEl.scrollIntoView({ behavior: `smooth` });
      this.activeItem.classList.add(`active`);
      this.mapEl.setMarker(this.activeItem);
    } else {
      this.mapEl.clearMarker();
    }
  }
  */
}
