import "leaflet";
import "simpleheat";
import "leaflet.heat";
import "leaflet.markercluster";

import "leaflet/dist/leaflet.css";
import "leaflet/dist/leaflet.css";
import "leaflet.markercluster/dist/leaflet.markercluster";
import "leaflet.markercluster/dist/MarkerCluster.css";
import "leaflet.markercluster/dist/MarkerCluster.Default.css";

import markerIcon from "leaflet/dist/images/marker-icon.png";
import markerIcon2x from "leaflet/dist/images/marker-icon-2x.png";
import markerShadow from "leaflet/dist/images/marker-shadow.png";

class WtHeatmap extends HTMLElement {
  constructor() {
    super();
  }

  connectedCallback() {
    this.apiWorkoutsCoordinatesRoute = JSON.parse(
      this.getAttribute("api-workouts-coordinates-route"),
    );
    this.apiWorkoutsCentersRoute = JSON.parse(
      this.getAttribute("api-workouts-centers-route"),
    );
    this.i18n = JSON.parse(this.getAttribute("i18n"));

    this.style.display = "block";

    var map = L.map(this, {
      fadeAnimation: false,
    });
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

    const iconDefault = L.icon({
      iconRetinaUrl: markerIcon2x,
      iconUrl: markerIcon,
      shadowUrl: markerShadow,
    });
    L.Marker.prototype.options.icon = iconDefault;

    L.control
      .layers({
        [this.i18n.streets]: layerStreet,
        [this.i18n.aerial]: layerAerial,
      })
      .addTo(map);

    let heatLayer = null;
    let heatMapData = null;
    let markers = null;
    const rerenderHeatMap = () => {
      if (heatMapData === null || markers === null) {
        console.log("data not ready")
        return
      }
      if (heatLayer !== null) {
        map.removeLayer(heatLayer);
      }
      const radius = L.DomUtil.get("radius").value;
      const blur = L.DomUtil.get("blur").value;
      const showMarkers = L.DomUtil.get("showMarkers").checked;
      heatLayer = L.heatLayer(heatMapData, { radius: Number(radius), blur: Number(blur) })
      heatLayer.addTo(map);
      if (showMarkers) {
        markers.addTo(map);
      } else {
        markers.removeFrom(map);
      }
    }

    let customControl = L.Control.extend({
      options: { position: 'topright' },

      onAdd: function () {
        const container = L.DomUtil.create('div', 'flex flex-col p-2');

        container.style.backgroundColor = 'white';
        container.innerHTML = `
        <div class="flex items-center"><label for="radius" class="w-12">Radius </label><input class="p-0" type="range" id="radius" value="10" min="5" max="30"/></div>
        <div class="flex items-center"><label for="blur" class="w-12">Blur </label><input class="p-0" type="range" id="blur" value="15" min="5" max="30"/></div>
        <div class="flex items-center"><input type="checkbox" id="showMarkers" name="showMarkers" checked /><label for="showMarkers">Show Markers</label></div>
        `;

        // Prevent map drag when clicking control
        L.DomEvent.disableClickPropagation(container);

        container.querySelectorAll("input")
          .forEach((element) => {
            element.oninput = L.Util.throttle(rerenderHeatMap, 50);
          })

        return container;
      }
    });

    map.addControl(new customControl());

    layerStreet.addTo(map);

    var clusterConfig = { showCoverageOnHover: false };

    fetch(this.apiWorkoutsCoordinatesRoute, {
      method: "GET",
      headers: {
        Accept: "application/json",
      },
    })
      .then((response) => response.json())
      .then((response) => {
        heatMapData = geoJson2heat(response.results);
        rerenderHeatMap();
      });

    fetch(this.apiWorkoutsCentersRoute, {
      method: "GET",
      headers: {
        Accept: "application/json",
      },
    })
      .then((response) => response.json())
      .then((response) => {
        markers = L.markerClusterGroup(clusterConfig);
        const geoJsonLayer = L.geoJson(response.results, {
          onEachFeature: function (feature, layer) {
            layer.bindPopup(feature.properties.details);
          },
        });
        geoJsonLayer.onEachFeature;

        markers.addLayer(geoJsonLayer);
        markers.addTo(map);

        map.fitBounds(markers.getBounds());
      });
  }
}

customElements.define("wt-heatmap", WtHeatmap);
