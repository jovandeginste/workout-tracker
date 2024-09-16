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
let hoverMarker;
let map;

function fullMap() {
  d = document.getElementById("map-container");

  d.classList.toggle("small-size");
  d.classList.toggle("full-size");

  map.invalidateSize(true);
  return false;
}

function makeMap(params) {
  document.addEventListener("DOMContentLoaded", () => {
    // Create map
    map = L.map(params.elementID, {
      fadeAnimation: false,
    }).setView(params.center, 15);
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

    const speeds = params.points
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
    params.points.forEach((pt) => {
      p = [pt.lat, pt.lng];

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
        polyLineProperties["color"] = getColor(
          (pt.elevation - params.minElevation) /
            (params.maxElevation - params.minElevation),
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
          polyLineProperties["color"] = getColor(0.5 + zScore / 2);
        }
        L.polyline([prevPoint, p], polyLineProperties).addTo(speedLayerGroup);
      }

      prevPoint = p;
    });

    if (!hasSpeed || params.showElevation) {
      elevationLayerGroup.addTo(map);
    } else {
      speedLayerGroup.addTo(map);
    }

    var last = params.points[params.points.length - 1];
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

    var first = params.points[0];
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

    if (!hoverMarker) {
      hoverMarker = L.circleMarker(first, {
        color: "blue",
        radius: 8,
      });
    }

    hoverMarker.addTo(map); // Adding marker to the map

    L.control
      .layers(
        {
          [params.streetsName]: layerStreet,
          [params.aerialName]: layerAerial,
        },
        {
          [params.elevationName]: elevationLayerGroup,
          [params.speedName]: speedLayerGroup,
        },
      )
      .addTo(map);

    layerStreet.addTo(map);

    map.fitBounds(group.getBounds(), { animate: false });
  });
}

function set_marker(title, lat, lon) {
  if (!hoverMarker) return;

  if (title != null) {
    hoverMarker.bindTooltip(title);
  }

  hoverMarker.setLatLng([lat, lon]);

  // Adding popup to the marker
  hoverMarker.openTooltip();
}

function clear_marker() {
  if (!hoverMarker) return;
  hoverMarker.closeTooltip();
}

// Determine color for a value; value from 0 to 1
// Linearly interpolate between blue and green
function getColor(value) {
  value = Math.max(0, Math.min(1, value)); // Clamp to 0...1

  const lowColor = [50, 50, 255];
  const highColor = [50, 255, 50];
  const color = [0, 1, 2].map((i) =>
    Math.floor(value * (highColor[i] - lowColor[i]) + lowColor[i]),
  );
  return `rgb(${color.join(",")})`;
}
