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

  async connectedCallback() {
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
        return;
      }
      if (heatLayer !== null) {
        map.removeLayer(heatLayer);
      }
      const radiusEl = L.DomUtil.get("radius");
      const blurEl = L.DomUtil.get("blur");
      const showMarkers = L.DomUtil.get("showMarkers").checked;
      const onlyTrace = L.DomUtil.get("onlyTrace").checked;
      var config = {
        radius: Number(radiusEl.value),
        blur: Number(blurEl.value),
      };
      if (onlyTrace) {
        config.radius = 1;
        config.blur = 1;
        config.minOpacity = 1;
        config.gradient = { 0: "blue" };

        radiusEl.disabled = true;
        blurEl.disabled = true;
      } else {
        radiusEl.disabled = false;
        blurEl.disabled = false;
      }
      heatLayer = L.heatLayer(heatMapData, config);
      heatLayer.addTo(map);
      if (showMarkers) {
        markers.addTo(map);
      } else {
        markers.removeFrom(map);
      }
    };

    let customControl = L.Control.extend({
      options: { position: "topright" },

      onAdd: function () {
        const container = L.DomUtil.create("div", "flex flex-col p-2");

        container.style.backgroundColor = "white";
        container.innerHTML = `
        <div class="flex items-center"><label for="radius" class="w-12 text-zinc-800">Radius</label><input class="p-0" type="range" id="radius" value="10" min="1" max="30"/></div>
        <div class="flex items-center"><label for="blur" class="w-12 text-zinc-800">Blur</label><input class="p-0" type="range" id="blur" value="15" min="1" max="30"/></div>
        <div class="flex items-center"><input type="checkbox" id="showMarkers" name="showMarkers" class="mr-1" checked /><label for="showMarkers" class="text-zinc-800">Show Markers</label></div>
        <div class="flex items-center"><input type="checkbox" id="onlyTrace" name="onlyTrace" class="mr-1" /><label for="onlyTrace" class="text-zinc-800">Only show where you've been</label></div>
        `;

        // Prevent map drag when clicking control
        L.DomEvent.disableClickPropagation(container);

        container.querySelectorAll("input").forEach((element) => {
          element.oninput = L.Util.throttle(rerenderHeatMap, 50);
        });

        return container;
      },
    });

    map.addControl(new customControl());

    layerStreet.addTo(map);

    var clusterConfig = { showCoverageOnHover: false };

    let resp = await fetch(this.apiWorkoutsCoordinatesRoute, {
      method: "GET",
      headers: {
        Accept: "application/json",
      },
    });
    let respJson = await resp.json();
    heatMapData = geoJson2heat(respJson.results);

    resp = await fetch(this.apiWorkoutsCentersRoute, {
      method: "GET",
      headers: {
        Accept: "application/json",
      },
    });
    respJson = await resp.json();
    markers = L.markerClusterGroup(clusterConfig);
    const geoJsonLayer = L.geoJson(respJson.results, {
      onEachFeature: function (feature, layer) {
        layer.bindPopup(feature.properties.details);
      },
    });
    geoJsonLayer.onEachFeature;

    markers.addLayer(geoJsonLayer);
    markers.addTo(map);

    map.fitBounds(markers.getBounds());

    rerenderHeatMap();
  }
}

customElements.define("wt-heatmap", WtHeatmap);
