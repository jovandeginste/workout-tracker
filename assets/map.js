var map;
var hoverMarker;

// This script relies on HTML having a "points" and "center" variables.
function on_loaded() {
  // Create map & tiles.
  map = L.map("map").setView(center, 15);
  L.tileLayer("https://tile.openstreetmap.org/{z}/{x}/{y}.png", {
    attribution:
      '&copy; <a href="http://www.openstreetmap.org/copyright">OpenStreetMap</a>',
  }).addTo(map);
  L.control.scale().addTo(map);

  var group = new L.featureGroup();
  var polyLineProperties = {
    weight: 4,
    interactive: false,
  };

  var prevPoint;
  // Add points with tooltip to map.
  points.forEach((pt) => {
    p = [pt.lat, pt.lng];

    if (prevPoint) {
      group.addLayer(
        L.circleMarker([pt.lat, pt.lng], {
          opacity: 0,
          fill: false,
          radius: 4,
        })
          .addTo(map)
          .bindTooltip(pt.title),
      );

      polyLineProperties["color"] = getColor(
        (pt.elevation - minElevation) / (maxElevation - minElevation),
      );
      L.polyline([prevPoint, p], polyLineProperties).addTo(map);
    }

    prevPoint = p;
  });

  var last = points[points.length - 1];
  group.addLayer(
    L.circleMarker([last.lat, last.lng], {
      color: "red",
      radius: 8,
    })
      .addTo(map)
      .bindTooltip(last.title),
  );

  var first = points[0];
  group.addLayer(
    L.circleMarker([first.lat, first.lng], {
      color: "green",
      radius: 8,
    })
      .addTo(map)
      .bindTooltip(first.title),
  );

  hoverMarker = L.circleMarker(first, {
    color: "blue",
    radius: 8,
  });

  hoverMarker.addTo(map); // Adding marker to the map
  map.fitBounds(group.getBounds());
}

function set_marker(title, lat, lon) {
  hoverMarker.bindTooltip(title);
  hoverMarker.setLatLng([lat, lon]);

  // Adding popup to the marker
  hoverMarker.openTooltip();
}
function clear_marker() {
  hoverMarker.closeTooltip();
}

function getColor(value) {
  //value from 0 to 1
  var hue = (240 + value * 120).toString(10);
  return ["hsl(", hue, ",100%,50%)"].join("");
}

document.addEventListener("DOMContentLoaded", on_loaded);
