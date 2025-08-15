class WorkoutBreakdown extends HTMLElement {
  constructor() {
    super();

    this.activeItem = null;
    this.intervalDistance = 1;

    this.mapEl = document.getElementById(this.getAttribute("map-id"));
    this.chartEl = document.getElementById(this.getAttribute("workout-stats"));

    this.data = JSON.parse(
      document.getElementById(this.getAttribute("data-el")).textContent,
    );
    this.preferredUnits = JSON.parse(
      document.getElementById(this.getAttribute("preferred-units-el"))
        .textContent,
    );

    this.availableMetrics = {
      distance: this.preferredUnits.distance,
      speed: this.preferredUnits.speed,
      duration: "",
      elevation: this.preferredUnits.elevation,
      "heart-rate": this.preferredUnits.heartRate,
      cadence: this.preferredUnits.cadence,
      temperature: this.preferredUnits.temperature,
    };
  }

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

  render() {
    const header = this.renderHeader();
    const data = this.renderData();
    this.innerHTML = `
      <table class="breakdown-table">
        <thead></thead>
        <tbody class="whitespace-nowrap font-mono"></tbody>
      </table>
    `;

    this.querySelector("thead").appendChild(header);
    for (const item of data) {
      this.querySelector("tbody").appendChild(item);
    }
  }

  renderHeader() {
    const header = document.createElement("tr");
    header.classList.add("breakdown-header");
    header.appendChild(document.createElement("th"));

    for (const metric of Object.keys(this.availableMetrics)) {
      if (this.data[metric] !== undefined) {
        const col = this.data[metric].Label;
        if (col) {
          header.appendChild(document.createElement("th")).textContent = col;
        }
      }
    }

    return header;
  }

  renderData() {
    const items = [];
    let currentDistance = Math.floor(+this.data.distance.Data[0] || 0);
    let intervalValues = {};
    for (let i = 0; i < this.data.time.Data.length; i++) {
      const distance = +this.data.distance.Data[i] || 0;
      if (distance >= currentDistance + this.intervalDistance) {
        items.push(this.renderRow(currentDistance, intervalValues));

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

    items.push(this.renderRow(currentDistance, intervalValues));

    return items;
  }

  renderRow(distance, data) {
    const displayDecimals = ["speed", "elevation", "temperature"];

    const row = document.createElement("tr");
    row.classList.add("breakdown-item");
    row.appendChild(document.createElement("td")).textContent = ``;

    for (const metric of Object.keys(this.availableMetrics)) {
      const mData = data[metric].filter((v) => v !== null);
      if (metric === "distance") {
        const cell = row.appendChild(document.createElement("td"));
        const lastDist = +mData[mData.length - 1];
        if (lastDist < distance + this.intervalDistance - 0.05) {
          cell.textContent = `${lastDist.toFixed(2)} ${this.availableMetrics[metric]}`;
        } else {
          cell.textContent = `${distance + this.intervalDistance} ${this.availableMetrics[metric]}`;
        }
        continue;
      }

      if (metric === "duration") {
        const cell = row.appendChild(document.createElement("td"));
        if (mData.length === 0) {
          cell.textContent = "-";
          continue;
        }

        const totalDuration = mData[mData.length - 1] - mData[0];
        const minutes = Math.floor(totalDuration / 60);
        let seconds = totalDuration % 60;
        if (seconds < 10) {
          seconds = `0${seconds}`;
        }

        cell.textContent = `${minutes}:${seconds} min/${this.preferredUnits.distance}`;
        continue;
      }

      if (data[metric] !== undefined) {
        if (mData.length === 0) {
          row.appendChild(document.createElement("td")).textContent = "-";
          continue;
        }

        const value = mData.reduce((a, b) => a + b, 0) / data[metric].length;
        const cell = document.createElement("td");
        if (displayDecimals.includes(metric)) {
          cell.textContent = `${value.toFixed(2)} ${this.availableMetrics[metric] || ""}`;
        } else {
          cell.textContent = `${value.toFixed(0)} ${this.availableMetrics[metric] || ""}`;
        }

        row.appendChild(cell);
      }
    }

    return row;
  }
}

customElements.define("workout-breakdown", WorkoutBreakdown);
