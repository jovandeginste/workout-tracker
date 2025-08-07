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
    const clean = new URLSearchParams(window.location.search).get("clean");

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

    layerStreet.addTo(map);

    var heatConfig = { radius: 10 };
    if (clean == "1") {
      heatConfig = {
        radius: 1,
        minOpacity: 1,
        blur: 1,
        gradient: { 0: "blue" },
      };
    }
    var clusterConfig = { showCoverageOnHover: false };

    fetch(this.apiWorkoutsCoordinatesRoute, {
      method: "GET",
      headers: {
        Accept: "application/json",
      },
    })
      .then((response) => response.json())
      .then((response) => {
        var data = geoJson2heat(response.results);
        L.heatLayer(data, heatConfig).addTo(map);
      });

    fetch(this.apiWorkoutsCentersRoute, {
      method: "GET",
      headers: {
        Accept: "application/json",
      },
    })
      .then((response) => response.json())
      .then((response) => {
        var markers = L.markerClusterGroup(clusterConfig);
        var geoJsonLayer = L.geoJson(response.results, {
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
