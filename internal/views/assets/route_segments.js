let startMarker;
let endMarker;

function editMap(params) {
  document.addEventListener("DOMContentLoaded", () => {
    // Create map
    const map = L.map(params.elementID, {
      fadeAnimation: false,
    }).setView(params.center, 15);
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

    params.points.forEach((pt) => {
      p = [pt.lat, pt.lng];

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

    var last = params.points[params.points.length - 1];
    endMarker = L.circleMarker([last.lat, last.lng], {
      color: "red",
      fill: true,
      fillColor: "red",
      fillOpacity: 1,
      radius: 6,
    });
    group.addLayer(endMarker.addTo(map));

    var first = params.points[0];
    startMarker = L.circleMarker([first.lat, first.lng], {
      color: "green",
      fill: true,
      fillColor: "green",
      fillOpacity: 1,
      radius: 6,
    });
    group.addLayer(startMarker.addTo(map));

    map.fitBounds(group.getBounds(), { animate: false });
  });
}
function updateStart() {
  start = Number(document.getElementById("start").value);
  end = Number(document.getElementById("end").value);

  document.getElementById("start-show").textContent = start;
  startMarker.setLatLng(new L.LatLng(points[start].lat, points[start].lng));

  if (start > end) {
    document.getElementById("end").value = start;
    updateEnd();
    return;
  }

  updateInfo();
}
function updateEnd() {
  start = Number(document.getElementById("start").value);
  end = Number(document.getElementById("end").value);

  document.getElementById("end-show").textContent = end;
  endMarker.setLatLng(new L.LatLng(points[end - 1].lat, points[end - 1].lng));

  if (start > end) {
    document.getElementById("start").value = end;
    updateStart();
    return;
  }

  updateInfo();
}

function updateAll() {
  updateStart();
  updateEnd();
  updateInfo();
}

function updateInfo() {
  updateDistance();
  updateLines();
}

function updateDistance() {
  start = Number(document.getElementById("start").value);
  end = Number(document.getElementById("end").value);

  d = points[end - 1].distance - points[start].distance;
  document.getElementById("distance-show").textContent = d.toFixed(2) + " m";
}

function updateLines() {
  start = Number(document.getElementById("start").value);
  end = Number(document.getElementById("end").value);

  distance = points[end - 1].distance;

  points.forEach((pt, idx) => {
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
