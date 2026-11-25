import L from "leaflet";

export function streetLayer() {
  return L.tileLayer("https://tile.openstreetmap.org/{z}/{x}/{y}.png", {
    attribution:
      '&copy; <a href="http://www.openstreetmap.org/copyright">OpenStreetMap</a>',
    className: "map-tiles",
  });
}

export function aerialLayer() {
  return L.tileLayer(
    "https://server.arcgisonline.com/ArcGIS/rest/services/World_Imagery/MapServer/tile/{z}/{y}/{x}",
    {
      attribution: "Powered by Esri",
    },
  );
}

export function endpointMarker(point, color) {
  return L.circleMarker(point, {
    color,
    fill: true,
    fillColor: color,
    fillOpacity: 1,
    radius: 6,
  });
}
