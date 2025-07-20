import * as L from "leaflet";
import "leaflet/dist/leaflet.css";

class CreateRouteSegmentMap extends HTMLElement {
  constructor() {
    super();

    this.startMarker = null;
    this.endMarker = null;
  }

  connectedCallback() {
    this.style.display = "block";

    this.center = JSON.parse(this.getAttribute("map-center") || "[0, 0]");
    this.points = JSON.parse(this.getAttribute("map-points") || "[]");
    if (this.points.length !== 0) {
      this.makeMap();
    }

    document.addEventListener("DOMContentLoaded", () => {
      this.updateInfo();
    });
  }

  makeMap() {
    const map = L.map(this, {
      fadeAnimation: false,
    }).setView(this.center, 15);
    L.tileLayer("https://tile.openstreetmap.org/{z}/{x}/{y}.png", {
      attribution:
        '&copy; <a href="http://www.openstreetmap.org/copyright">OpenStreetMap</a>',
      className: "map-tiles",
    }).addTo(map);
    L.control.scale().addTo(map);

    // Add features to the map
    const group = new L.featureGroup();
    const polyLineProperties = {
      weight: 4,
      interactive: false,
    };

    let prevPoint;
    this.points.forEach((pt) => {
      const p = [pt.lat, pt.lng];

      if (prevPoint) {
        let m = L.circleMarker([pt.lat, pt.lng], {
          opacity: 0,
          fill: false,
          radius: 4,
        });
        // Add invisible point to map to allow fitBounds to work
        group.addLayer(m.addTo(map).bindTooltip(pt.title));

        let l = L.polyline([prevPoint, p], polyLineProperties);
        l.addTo(map);
        pt.line = l;
      }

      prevPoint = p;
    });

    let last = this.points[this.points.length - 1];
    this.endMarker = L.circleMarker([last.lat, last.lng], {
      color: "red",
      fill: true,
      fillColor: "red",
      fillOpacity: 1,
      radius: 6,
    });
    group.addLayer(this.endMarker.addTo(map));

    let first = this.points[0];
    this.startMarker = L.circleMarker([first.lat, first.lng], {
      color: "green",
      fill: true,
      fillColor: "green",
      fillOpacity: 1,
      radius: 6,
    });
    group.addLayer(this.startMarker.addTo(map));

    map.fitBounds(group.getBounds(), { animate: false });
  }

  updateStart() {
    const start = Number(document.getElementById("start").value);
    const end = Number(document.getElementById("end").value);

    document.getElementById("start-show").textContent = start;
    this.startMarker.setLatLng(
      new L.LatLng(this.points[start].lat, this.points[start].lng),
    );

    if (start > end) {
      document.getElementById("end").value = start;
      this.updateEnd();
      return;
    }

    this.updateInfo();
  }

  updateEnd() {
    const start = Number(document.getElementById("start").value);
    const end = Number(document.getElementById("end").value);

    document.getElementById("end-show").textContent = end;
    this.endMarker.setLatLng(
      new L.LatLng(this.points[end - 1].lat, this.points[end - 1].lng),
    );

    if (start > end) {
      document.getElementById("start").value = end;
      this.updateStart();
      return;
    }

    this.updateInfo();
  }

  updateAll() {
    this.updateStart();
    this.updateEnd();
    this.updateInfo();
  }

  updateInfo() {
    this.updateDistance();
    this.updateLines();
  }

  updateDistance() {
    const start = Number(document.getElementById("start").value);
    const end = Number(document.getElementById("end").value);

    const d = this.points[end - 1].distance - this.points[start].distance;
    document.getElementById("distance-show").textContent = d.toFixed(2) + " m";
  }

  updateLines() {
    const start = Number(document.getElementById("start").value);
    const end = Number(document.getElementById("end").value);

    this.points.forEach((pt, idx) => {
      switch (true) {
        case pt.line == null:
          return;
        case idx < start:
          pt.line.setStyle({ color: "#FF0000" });
          return;
        case idx > end:
          pt.line.setStyle({ color: "#FF0000" });
          return;
        default:
          pt.line.setStyle({ color: "#00FF00" });
      }
    });
  }
}

customElements.define("create-route-segment-map", CreateRouteSegmentMap);
