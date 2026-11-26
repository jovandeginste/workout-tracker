import L from "leaflet";
import { HeatLayer } from "leaflet.heat";
import "leaflet.markercluster";

import "leaflet/dist/leaflet.css";
import "leaflet.markercluster/dist/leaflet.markercluster";
import "leaflet.markercluster/dist/MarkerCluster.css";
import "leaflet.markercluster/dist/MarkerCluster.Default.css";

import markerIcon from "leaflet/dist/images/marker-icon.png";
import markerIcon2x from "leaflet/dist/images/marker-icon-2x.png";
import markerShadow from "leaflet/dist/images/marker-shadow.png";
import { aerialLayer, streetLayer } from "../../map.js";

class WtHeatmap extends HTMLElement {
  constructor() {
    super();
  }

  updateSize() {
    this.map?.invalidateSize(true);
  }

  async connectedCallback() {
    if (this.querySelector(".leaflet-container")) {
      return;
    }

    this.apiWorkoutsCoordinatesRoute = JSON.parse(
      this.getAttribute("api-workouts-coordinates-route"),
    );
    this.apiWorkoutsCentersRoute = JSON.parse(
      this.getAttribute("api-workouts-centers-route"),
    );
    this.i18n = JSON.parse(this.getAttribute("i18n"));
    this.types = JSON.parse(this.getAttribute("types") || "[]") || [];
    this.equipment = JSON.parse(this.getAttribute("equipment") || "[]") || [];

    this.style.display = "block";

    var map = L.map(this, {
      fadeAnimation: false,
    });
    this.map = map;
    const resizeMap = () => map.invalidateSize(true);
    requestAnimationFrame(resizeMap);
    setTimeout(resizeMap, 250);
    const layerStreet = streetLayer();
    const layerAerial = aerialLayer();

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

      map.invalidateSize(true);
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
      heatLayer = new HeatLayer(heatMapData, config);
      heatLayer.addTo(map);
      if (showMarkers) {
        markers.addTo(map);
      } else {
        markers.removeFrom(map);
      }
    };

    const translations = this.i18n || {};

    let customControl = L.Control.extend({
      options: { position: "topright" },

      onAdd: function () {
        const container = L.DomUtil.create("div", "flex flex-col p-2");

        container.style.backgroundColor = "white";
        container.innerHTML = `
        <div class="flex items-center"><label for="radius" class="w-12 text-zinc-800">Radius</label><input class="p-0" type="range" id="radius" value="10" min="1" max="30"/></div>
        <div class="flex items-center"><label for="blur" class="w-12 text-zinc-800">Blur</label><input class="p-0" type="range" id="blur" value="15" min="1" max="30"/></div>
        <div class="flex items-center"><input type="checkbox" id="showMarkers" name="showMarkers" class="mr-1" checked /><label for="showMarkers" class="text-zinc-800">${translations.markers || "Show markers"}</label></div>
        <div class="flex items-center"><input type="checkbox" id="onlyTrace" name="onlyTrace" class="mr-1" /><label for="onlyTrace" class="text-zinc-800">${translations["only-visited"] || "Only show where you've been"}</label></div>
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

    const fetchJson = async (route) => {
      const params = new URLSearchParams();
      this.types.forEach((type) => params.append("type", type));
      this.equipment.forEach((id) => params.append("equipment", id));
      const url = params.toString() ? `${route}?${params}` : route;
      const response = await fetch(url, {
        method: "GET",
        headers: { Accept: "application/json" },
      });
      return response.json();
    };

    this.loadData = async () => {
      let respJson = await fetchJson(this.apiWorkoutsCoordinatesRoute);
      heatMapData = geoJson2heat(respJson.results || { features: [] });

      if (markers !== null) {
        map.removeLayer(markers);
      }

      respJson = await fetchJson(this.apiWorkoutsCentersRoute);
      markers = L.markerClusterGroup(clusterConfig);
      const geoJsonLayer = L.geoJson(respJson.results || { features: [] }, {
        onEachFeature: function (feature, layer) {
          layer.bindPopup(feature.properties.details);
        },
      });
      markers.addLayer(geoJsonLayer);
      markers.addTo(map);

      if (markers.getLayers().length > 0) {
        map.fitBounds(markers.getBounds());
      }

      rerenderHeatMap();
    };

    await this.loadData();

    const filterForm = document.getElementById("heatmap-filter");
    filterForm?.addEventListener("change", () => {
      this.types = Array.from(
        filterForm.querySelectorAll('input[name="type"]:checked'),
      )
        .map((input) => input.value)
        .filter(Boolean);

      this.equipment = Array.from(
        filterForm.querySelectorAll('input[name="equipment"]:checked'),
      )
        .map((input) => Number(input.value))
        .filter((id) => id > 0);

      this.loadData();
    });
  }
}

customElements.define("wt-heatmap", WtHeatmap);
