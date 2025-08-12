class WorkoutBreakdown extends HTMLElement {
  constructor() {
    super();

    this.mapElement = document.getElementById(this.getAttribute("map-id"));
    this.activeItem = null;
  }

  connectedCallback() {
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
      this.mapElement.setMarker(item);
    }
  }

  itemMouseOut() {
    if (!this.activeItem) {
      this.mapElement.clearMarker();
    }
  }

  setActiveItem(item) {
    if (this.activeItem) {
      this.activeItem.classList.remove(`active`);
    }

    this.activeItem = item;
    if (this.activeItem) {
      this.mapElement.scrollIntoView({ behavior: `smooth` });
      this.activeItem.classList.add(`active`);
      this.mapElement.setMarker(this.activeItem);
    } else {
      this.mapElement.clearMarker();
    }
  }
}

customElements.define("workout-breakdown", WorkoutBreakdown);
