// This script relies on HTML having a "points" and "center" variables.
function on_loaded() {
  // Create map & tiles.
  var map = L.map("map").setView(center, 15);
  L.tileLayer("https://tile.openstreetmap.org/{z}/{x}/{y}.png", {
    attribution:
      '&copy; <a href="http://www.openstreetmap.org/copyright">OpenStreetMap</a>',
  }).addTo(map);
  L.control.scale().addTo(map);

  var group = new L.featureGroup();
  var p = [];
  var polyLineProperties = {
    weight: 4,
    interactive: false,
  };

  var first = points.shift();
  var last = points.pop();

  p.push([first.lat, first.lng]);
  group.addLayer(
    L.circleMarker([first.lat, first.lng], {
      color: "green",
      radius: 16,
    })
      .addTo(map)
      .bindTooltip(first.title),
  );

  // Add points with tooltip to map.
  points.forEach((pt) => {
    p.push([pt.lat, pt.lng]);
    group.addLayer(
      L.circleMarker([pt.lat, pt.lng], {
        opacity: 0,
        fill: false,
        radius: 4,
      })
        .addTo(map)
        .bindTooltip(pt.title),
    );
  });

  p.push([last.lat, last.lng]);
  group.addLayer(
    L.circleMarker([last.lat, last.lng], {
      color: "red",
      radius: 16,
    })
      .addTo(map)
      .bindTooltip(last.title),
  );

  L.polyline(p, polyLineProperties).addTo(map);

  map.fitBounds(group.getBounds());
}

document.addEventListener("DOMContentLoaded", on_loaded);
