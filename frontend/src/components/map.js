import * as L from "leaflet";
import "leaflet/dist/leaflet.css";

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
      streetsName: mapConfig.StreetsName,
      aerialName: mapConfig.AerialName,
      showElevation: mapConfig.ShowElevation,
    };

    this.points = JSON.parse(this.getAttribute("map-points") || "[]");
    if (this.points.length !== 0) {
      this.makeMap();
    }
  }

  makeMap() {
    const map = L.map(this, {
      fadeAnimation: false,
    }).setView(this.config.center, 15);
    const layerStreet = L.tileLayer(
      "https://tile.openstreetmap.org/{z}/{x}/{y}.png",
      {
        attribution:
          '&copy; <a href="http://www.openstreetmap.org/copyright">OpenStreetMap</a>',
        className: "map-tiles",
      },
    );

    const layerAerial = L.tileLayer(
      "https://server.arcgisonline.com/ArcGIS/rest/services/World_Imagery/MapServer/tile/{z}/{y}/{x}",
      {
        attribution: "Powered by Esri",
      },
    );

    L.control.scale().addTo(map);

    const speeds = this.points
      .filter((x) => x.speed !== null)
      .map((x) => x.speed);

    const averageSpeed =
      speeds.reduce((a, x) => {
        return a + x;
      }, 0) / speeds.length;
    const stdevSpeed = Math.sqrt(
      speeds.reduce((a, x) => a + Math.pow(x - averageSpeed, 2), 0) /
        (speeds.length - 1),
    );

    // Add features to the map
    const group = new L.featureGroup();
    const polyLineProperties = {
      weight: 4,
      interactive: false,
    };

    let prevPoint;
    // Add points with tooltip to map.
    const MOVING_AVERAGE_LENGTH = 15;
    const movingSpeeds = [];
    const speedLayerGroup = new L.featureGroup();
    const elevationLayerGroup = new L.featureGroup();

    var hasSpeed = false;
    this.points.forEach((pt) => {
      let p = [pt.lat, pt.lng];

      if (prevPoint) {
        // Add invisible point to map to allow fitBounds to work
        group.addLayer(
          L.circleMarker([pt.lat, pt.lng], {
            opacity: 0,
            fill: false,
            radius: 4,
          })
            .addTo(map)
            .bindTooltip(pt.title),
        );

        // Elevation
        polyLineProperties["color"] = this.getColor(
          (pt.elevation - this.config.minElevation) /
            (this.config.maxElevation - this.config.minElevation),
        );
        L.polyline([prevPoint, p], polyLineProperties).addTo(
          elevationLayerGroup,
        );

        // Speed
        if (pt.speed === null || pt.speed < 0.1) {
          polyLineProperties["color"] = "rgb(0,0,0)"; // Pausing
        } else {
          hasSpeed = true;
          if (movingSpeeds.length > MOVING_AVERAGE_LENGTH) {
            movingSpeeds.shift();
          }
          movingSpeeds.push(pt.speed);
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

    if (!hasSpeed || this.config.showElevation) {
      elevationLayerGroup.addTo(map);
    } else {
      speedLayerGroup.addTo(map);
    }

    var last = this.points[this.points.length - 1];
    group.addLayer(
      L.circleMarker([last.lat, last.lng], {
        color: "red",
        fill: true,
        fillColor: "red",
        fillOpacity: 1,
        radius: 6,
      })
        .addTo(map)
        .bindTooltip(last.title),
    );

    var first = this.points[0];
    group.addLayer(
      L.circleMarker([first.lat, first.lng], {
        color: "green",
        fill: true,
        fillColor: "green",
        fillOpacity: 1,
        radius: 6,
      })
        .addTo(map)
        .bindTooltip(first.title),
    );

    if (!this.hoverMarker) {
      this.hoverMarker = L.circleMarker(first, {
        color: "blue",
        radius: 8,
      });
    }

    this.hoverMarker.addTo(map); // Adding marker to the map

    L.control
      .layers(
        {
          [this.config.streetsName]: layerStreet,
          [this.config.aerialName]: layerAerial,
        },
        {
          [this.config.elevationName]: elevationLayerGroup,
          [this.config.speedName]: speedLayerGroup,
        },
      )
      .addTo(map);

    layerStreet.addTo(map);

    map.fitBounds(group.getBounds(), { animate: false });
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

  updateSize() {
    map.invalidateSize(true);
  };
}

customElements.define("wt-map", WtMap);
