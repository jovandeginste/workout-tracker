import * as L from "leaflet";
import "leaflet/dist/leaflet.css";
import { formatDuration } from "../helpers.js";

/*
interface Point {
  lat: number;
  lng: number;
  title: string;
  elevation: number;
}

interface Parameters {
  elementID: string;         // ID of the element to put the map in
  center: [number, number];  // Lat, long coordinate to center the map to
  points: Point[];           // Points of the route to show
  minElevation: number;
  maxElevation: number;
  maxSpeed: number;
  speedName: string;        // Name for speed layer
  elevationName: string;    // Name of elevation layer
  streetsName: string;
  aerialName: string;
}
*/
class WtMap extends HTMLElement {
  constructor() {
    super();

    this.hoverMarker = null;
  }

  connectedCallback() {
    this.style.display = "block";

    const mapConfig = JSON.parse(this.getAttribute("map-config") || "{}");
    this.config = {
      center: [mapConfig.CenterLat, mapConfig.CenterLng],
      minElevation: mapConfig.MinElevation,
      maxElevation: mapConfig.MaxElevation,
      maxSpeed: mapConfig.MaxSpeed,
      speedName: mapConfig.SpeedName,
      elevationName: mapConfig.ElevationName,
      slopeName: mapConfig.SlopeName,
      streetsName: mapConfig.StreetsName,
      aerialName: mapConfig.AerialName,
      showElevation: mapConfig.ShowElevation,
      showSlope: mapConfig.ShowSlope,
    };

    this.segmentLayerGroup = L.featureGroup();
    this.workout = JSON.parse(
      document.getElementById(this.getAttribute("data-el")).textContent,
    );
    this.preferredUnits = JSON.parse(
      document.getElementById(this.getAttribute("preferred-units-el"))
        .textContent,
    );
    this.trackGroup = new L.featureGroup();
    if (this.workout?.positions?.Data?.length !== 0) {
      this.makeMap();
    }
  }

  makeMap() {
    this.map = L.map(this, {
      fadeAnimation: false,
    }).setView(this.config.center, 15);

    const map = this.map;
    const layerStreet = this.getStreetLayer();
    const layerAerial = this.getAerialLayer();

    // Add features to the map
    const trackRenderer = L.canvas({ padding: 0.4 });
    const polyLineProperties = {
      renderer: trackRenderer,
      weight: 4,
      interactive: false,
    };

    let prevPoint;
    const hasSpeed = !!this.workout.speed?.Data?.length;
    // Add points with tooltip to map.
    let speedLayerGroup;
    if (hasSpeed) {
      speedLayerGroup = this.getSpeedLayerGroup(polyLineProperties);
    }

    const hasSlope = !!this.workout.slope?.Data?.length;
    // Add points with tooltip to map.
    let slopeLayerGroup;
    if (hasSlope) {
      slopeLayerGroup = this.getSlopeLayerGroup(polyLineProperties);
    }

    const elevationLayerGroup = new L.featureGroup();

    this.workout.position.Data.forEach((p, i) => {
      if (p === null) {
        return;
      }

      if (prevPoint) {
        const elevation = this.workout.elevation.Data[i] || 0;
        // Add invisible point to map to allow fitBounds to work
        this.trackGroup.addLayer(
          L.circleMarker(p, {
            renderer: trackRenderer,
            opacity: 0,
            fill: false,
            radius: 4,
          })
            .addTo(map)
            .bindTooltip(() => this.getTooltip(i)),
        );

        // Elevation
        polyLineProperties["color"] = this.getColor(
          (elevation - this.config.minElevation) /
            (this.config.maxElevation - this.config.minElevation),
        );
        L.polyline([prevPoint, p], polyLineProperties).addTo(
          elevationLayerGroup,
        );
      }

      prevPoint = p;
    });

    if (speedLayerGroup && this.config.showSpeed) {
      speedLayerGroup.addTo(map);
    } else if (slopeLayerGroup && this.config.showSlope) {
      slopeLayerGroup.addTo(map);
    } else {
      elevationLayerGroup.addTo(map);
    }

    const positions = this.workout.position.Data.filter((x) => x !== null);

    if (positions.length > 0) {
      let last = positions[positions.length - 1];
      this.trackGroup.addLayer(
        L.circleMarker(last, {
          color: "red",
          fill: true,
          fillColor: "red",
          fillOpacity: 1,
          radius: 6,
        })
          .addTo(map)
          .bindTooltip(this.getTooltip(positions.length - 1)),
      );

      let first = positions[0];
      this.trackGroup.addLayer(
        L.circleMarker(first, {
          color: "green",
          fill: true,
          fillColor: "green",
          fillOpacity: 1,
          radius: 6,
        })
          .addTo(map)
          .bindTooltip(this.getTooltip(0)),
      );

      if (!this.hoverMarker) {
        this.hoverMarker = L.circleMarker(first, {
          color: "blue",
          radius: 8,
        });
      }

      this.hoverMarker.addTo(map); // Adding marker to the map
    }

    const overlays = {
      [this.config.elevationName]: elevationLayerGroup,
    };
    if (speedLayerGroup) {
      overlays[this.config.speedName] = speedLayerGroup;
    }
    if (slopeLayerGroup) {
      overlays[this.config.slopeName] = slopeLayerGroup;
    }

    L.control.scale().addTo(map);
    L.control
      .layers(
        {
          [this.config.streetsName]: layerStreet,
          [this.config.aerialName]: layerAerial,
        },
        overlays,
      )
      .addTo(map);

    layerStreet.addTo(map);

    this.resetZoom();
  }

  getStreetLayer() {
    return L.tileLayer("https://tile.openstreetmap.org/{z}/{x}/{y}.png", {
      attribution:
        '&copy; <a href="http://www.openstreetmap.org/copyright">OpenStreetMap</a>',
      className: "map-tiles",
    });
  }

  getAerialLayer() {
    return L.tileLayer(
      "https://server.arcgisonline.com/ArcGIS/rest/services/World_Imagery/MapServer/tile/{z}/{y}/{x}",
      {
        attribution: "Powered by Esri",
      },
    );
  }

  getSlopeLayerGroup(polyLineProperties = {}) {
    const slopeLayerGroup = new L.featureGroup();

    let prevPoint;
    const positions = this.workout.position.Data;
    positions.forEach((p, i) => {
      if (p === null) {
        return;
      }

      if (prevPoint) {
        const slope = this.workout.slope.Data[i] || 0;

        polyLineProperties["color"] = this.colorForSlope(slope);
        L.polyline([prevPoint, p], polyLineProperties).addTo(slopeLayerGroup);
      }
      prevPoint = p;
    });

    return slopeLayerGroup;
  }

  getSpeedLayerGroup(polyLineProperties = {}) {
    const MOVING_AVERAGE_LENGTH = 15;
    const movingSpeeds = [];
    const speeds = this.workout.speed?.Data.filter((x) => x !== null);

    const averageSpeed =
      speeds.reduce((a, x) => {
        return a + x;
      }, 0) / speeds.length;
    const stdevSpeed = Math.sqrt(
      speeds.reduce((a, x) => a + Math.pow(x - averageSpeed, 2), 0) /
        (speeds.length - 1),
    );

    const speedLayerGroup = new L.featureGroup();

    let prevPoint;
    this.workout.position.Data.forEach((p, i) => {
      if (p === null) {
        return;
      }
      if (prevPoint) {
        const speed = this.workout.speed.Data[i] || null;
        if (speed === null || speed < 0.1) {
          polyLineProperties["color"] = "rgb(0,0,0)"; // Pausing
        } else {
          if (movingSpeeds.length > MOVING_AVERAGE_LENGTH) {
            movingSpeeds.shift();
          }
          movingSpeeds.push(speed);
          const movingAverageSpeed =
            movingSpeeds.reduce((a, x) => a + x) / movingSpeeds.length;

          const zScore =
            ((movingAverageSpeed || averageSpeed) - averageSpeed) / stdevSpeed; // -1...1 is within one standard deviation
          polyLineProperties["color"] = this.getColor(0.5 + zScore / 2);
        }
        L.polyline([prevPoint, p], polyLineProperties).addTo(speedLayerGroup);
      }

      prevPoint = p;
    });

    return speedLayerGroup;
  }

  getTooltip(i) {
    const tooltipDisplay = {
      time: "",
      distance: this.preferredUnits.distance,
      duration: "",
      speed: this.preferredUnits.speed,
      elevation: this.preferredUnits.elevation,
      "heart-rate": this.preferredUnits.heartRate,
      cadence: this.preferredUnits.cadence,
      temperature: this.preferredUnits.temperature,
      slope: "%",
    };

    let tooltip = `<ul>`;
    for (const [field, unit] of Object.entries(tooltipDisplay)) {
      if (this.workout[field]?.Data[i] !== undefined) {
        const label = this.workout[field].Label;
        const val = this.workout[field].Data[i];
        let formattedVal =
          typeof val === "number" && val % 1 !== 0 ? val.toFixed(2) : val;
        if (field === "duration") {
          formattedVal = formatDuration(val);
        } else if (field === "slope" && val !== null) {
          formattedVal = Math.round(100 * val);
        } else if (field === "time" && val !== null) {
          formattedVal = new Date(val).toTimeString().substr(0, 5);
        }

        if (val !== null) {
          tooltip += `<li><b>${label}</b>: ${formattedVal} ${unit}</li>`;
        }
      }
    }
    tooltip += `</ul>`;
    return tooltip;
  }

  // Determine color for a value; value from 0 to 1
  // Linearly interpolate between blue and green
  getColor(value) {
    value = Math.max(0, Math.min(1, value)); // Clamp to 0...1

    const lowColor = [50, 50, 255];
    const highColor = [50, 255, 50];
    const color = [0, 1, 2].map((i) =>
      Math.floor(value * (highColor[i] - lowColor[i]) + lowColor[i]),
    );
    return `rgb(${color.join(",")})`;
  }

  setMarker(obj) {
    const lat = obj.getAttribute("data-lat");
    const lng = obj.getAttribute("data-lng");
    const title = obj.getAttribute("data-title");

    if (!this.hoverMarker) return;

    if (title != null) {
      this.hoverMarker.bindTooltip(title);
    }

    this.hoverMarker.setLatLng([lat, lng]);

    // Adding popup to the marker
    this.hoverMarker.openTooltip();
  }

  clearMarker() {
    if (!this.hoverMarker) return;
    this.hoverMarker.closeTooltip();
  }

  setSegment(_title, data) {
    this.segmentLayerGroup.clearLayers();

    const positions = data.position.filter((p) => p !== null);
    if (positions.length < 2) {
      return;
    }

    for (let i = 1; i < positions.length; i++) {
      L.polyline([positions[i - 1], positions[i]], {
        color: "red",
      }).addTo(this.segmentLayerGroup);
    }

    this.segmentLayerGroup.addTo(this.map);
  }

  fitSegmentBounds() {
    this.map.fitBounds(this.segmentLayerGroup.getBounds(), {
      animate: false,
      padding: [20, 20],
    });
  }

  resetZoom() {
    this.map.fitBounds(this.trackGroup.getBounds(), { animate: false });
  }

  clearSegment() {
    this.segmentLayerGroup.clearLayers();
    this.resetZoom();
  }

  updateSize() {
    this.map.invalidateSize(true);
  }

  colorForSlope(slope) {
    // See https://github.com/alexgasconn/GPX-Analyzer/blob/bcf16608dda6c748d0fe6d87cd358ca8af4291e5/components/core/utils.py#L4
    switch (true) {
      case slope >= 0.18:
        return "#8B0000"; // Dark Red
      case slope >= 0.1:
        return "#FF8C00"; // Dark Orange
      case slope >= 0.02:
        return "#FFFF00"; // Yellow
      case slope >= 0.0:
        return "#ADFF2F"; // GreenYellow
      case slope >= -0.02:
        return "#ADD8E6"; // LightBlue
      case slope >= -0.1:
        return "#0000FF"; // Blue
      default:
        return "#00008B"; // Dark Blue
    }
  }
}

customElements.define("wt-map", WtMap);
