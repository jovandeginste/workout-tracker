import * as L from "leaflet";
import "leaflet/dist/leaflet.css";
import { formatDuration } from "../helpers.js";
import { aerialLayer, streetLayer } from "../map.js";

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
class WtMap extends HTMLElement {
  constructor() {
    super();

    this.hoverMarker = null;
    this.selectedSegment = null;
    this.selectedSegmentKey = null;
    this.selectedMarkers = [];
  }

  connectedCallback() {
    this.style.display = "block";

    const mapConfig = JSON.parse(this.getAttribute("map-config") || "{}");
    this.config = {
      center: [mapConfig.CenterLat, mapConfig.CenterLng],
      minElevation: mapConfig.MinElevation,
      maxElevation: mapConfig.MaxElevation,
      speedName: mapConfig.SpeedName,
      elevationName: mapConfig.ElevationName,
      slopeName: mapConfig.SlopeName,
      liftName: mapConfig.LiftName,
      climbsName: mapConfig.ClimbsName,
      labels: mapConfig.Labels || {},
      streetsName: mapConfig.StreetsName,
      aerialName: mapConfig.AerialName,
      showSlope: mapConfig.ShowSlope,
    };

    this.segmentLayerGroup = L.featureGroup();
    this.workout = JSON.parse(
      document.getElementById(this.getAttribute("data-el")).textContent,
    );
    this.preferredUnits = JSON.parse(
      document.getElementById(this.getAttribute("preferred-units-el"))
        .textContent,
    );
    this.trackGroup = new L.featureGroup();
    if (this.workout?.positions?.Data?.length !== 0) {
      this.makeMap();
    }
  }

  makeMap() {
    this.map = L.map(this, {
      fadeAnimation: false,
    }).setView(this.config.center, 15);

    const map = this.map;
    const layerStreet = this.getStreetLayer();
    const layerAerial = this.getAerialLayer();

    // Add features to the map
    const trackRenderer = L.canvas({ padding: 0.4 });
    const polyLineProperties = {
      renderer: trackRenderer,
      weight: 4,
      interactive: false,
    };

    map.createPane("segmentHighlight");
    map.getPane("segmentHighlight").style.zIndex = 350;
    this.segmentHighlightLayer = L.featureGroup().addTo(map);

    let prevPoint;
    const hasSpeed = !!this.workout.speed?.Data?.length;
    // Add points with tooltip to map.
    let speedLayerGroup;
    if (hasSpeed) {
      speedLayerGroup = this.getSpeedLayerGroup(polyLineProperties);
    }

    const hasSlope = !!this.workout.slope?.Data?.length;
    // Add points with tooltip to map.
    let slopeLayerGroup;
    if (hasSlope) {
      slopeLayerGroup = this.getSlopeLayerGroup(polyLineProperties);
    }

    const elevationLayerGroup = new L.featureGroup();

    this.workout.position.Data.forEach((p, i) => {
      if (p === null) {
        return;
      }

      if (prevPoint) {
        const elevation = this.workout.elevation.Data[i] || 0;
        // Add invisible point to map to allow fitBounds to work
        this.trackGroup.addLayer(
          L.circleMarker(p, {
            renderer: trackRenderer,
            opacity: 0,
            fill: false,
            radius: 4,
          })
            .addTo(map)
            .bindTooltip(() => this.getTooltip(i)),
        );

        // Elevation
        polyLineProperties["color"] = this.getColor(
          (elevation - this.config.minElevation) /
            (this.config.maxElevation - this.config.minElevation),
        );
        L.polyline([prevPoint, p], polyLineProperties).addTo(
          elevationLayerGroup,
        );
      }

      prevPoint = p;
    });

    if (slopeLayerGroup && this.config.showSlope) {
      slopeLayerGroup.addTo(map);
    } else {
      elevationLayerGroup.addTo(map);
    }

    let liftHaloGroup;
    if (this.workout.lifts?.Data?.length) {
      liftHaloGroup = this.getLiftHaloLayerGroup();
    }

    const positions = this.workout.position.Data.filter((x) => x !== null);
    const annotations = this.getAnnotationLayers();
    const liftsOverlay = liftHaloGroup
      ? L.layerGroup([liftHaloGroup, annotations.lifts])
      : annotations.lifts;

    annotations.startEnd.addTo(map);
    annotations.climbs.addTo(map);
    liftsOverlay.addTo(map);

    if (positions.length > 0) {
      if (!this.hoverMarker) {
        this.hoverMarker = L.circleMarker(positions[0], {
          color: "blue",
          radius: 8,
        });
      }

      this.hoverMarker.addTo(map); // Adding marker to the map
    }

    const overlays = {
      [this.config.elevationName]: elevationLayerGroup,
    };
    if (speedLayerGroup) {
      overlays[this.config.speedName] = speedLayerGroup;
    }
    if (slopeLayerGroup) {
      overlays[this.config.slopeName] = slopeLayerGroup;
    }

    if (this.config.climbsName && annotations.climbs.getLayers().length > 0) {
      overlays[this.config.climbsName] = annotations.climbs;
    }

    if (
      this.config.liftName &&
      (liftHaloGroup || annotations.lifts.getLayers().length > 0)
    ) {
      overlays[this.config.liftName] = liftsOverlay;
    }

    L.control.scale().addTo(map);
    L.control
      .layers(
        {
          [this.config.streetsName]: layerStreet,
          [this.config.aerialName]: layerAerial,
        },
        overlays,
      )
      .addTo(map);

    layerStreet.addTo(map);

    this.resetZoom();
  }

  getStreetLayer() {
    return streetLayer();
  }

  getAerialLayer() {
    return aerialLayer();
  }

  buildColorLayer(positions, polyLineProperties, colorFor) {
    const group = new L.featureGroup();

    let prevPoint;
    positions.forEach((p, i) => {
      if (p === null) {
        return;
      }

      if (prevPoint) {
        polyLineProperties.color = colorFor(i);
        group.addLayer(L.polyline([prevPoint, p], polyLineProperties));
      }

      prevPoint = p;
    });

    return group;
  }

  getLiftHaloLayerGroup() {
    const group = new L.featureGroup();

    for (const positions of this.workout.lifts.Data) {
      if (!positions || positions.length < 2) {
        continue;
      }

      L.polyline(positions, {
        pane: "segmentHighlight",
        color: "#f59e0b",
        weight: 10,
        opacity: 0.85,
        interactive: false,
      }).addTo(group);
    }

    return group;
  }

  getAnnotationLayers() {
    const startEnd = new L.featureGroup();
    const climbs = new L.featureGroup();
    const lifts = new L.featureGroup();

    const allPositions = this.workout.position.Data;
    const positions = allPositions.filter((x) => x !== null);

    if (positions.length > 0) {
      const first = positions[0];
      const last = positions[positions.length - 1];

      startEnd.addLayer(
        this.getFlagAnnotation(
          first,
          this.label("start", "Start"),
          "start",
          this.getTooltip(allPositions.indexOf(first)),
        ),
      );
      startEnd.addLayer(
        this.getFlagAnnotation(
          last,
          this.label("end", "End"),
          "end",
          this.getTooltip(allPositions.lastIndexOf(last)),
        ),
      );
    }

    for (const segment of this.workout.segments?.Data || []) {
      const group = segment.kind === "lift" ? lifts : climbs;
      const markers = [
        this.getSegmentAnnotation(segment, "start"),
        this.getSegmentAnnotation(segment, "end"),
      ];

      for (const marker of markers) {
        marker.on("mouseover", () => this.highlightSegment(segment, markers));
        marker.on("mouseout", () => {
          this.clearSegmentHighlight();
          this.restoreSegmentHighlight();
        });
        marker.on("click", () => this.toggleSegmentSelection(segment, markers));
        group.addLayer(marker);
      }
    }

    return { startEnd, climbs, lifts };
  }

  getAnnotationIcon(html) {
    return L.divIcon({
      className: "wt-map-annotation",
      html,
      iconSize: [0, 0],
      iconAnchor: [0, 0],
    });
  }

  getFlagAnnotation(point, text, kind, tooltip) {
    const html = `<span class="wt-map-flag wt-map-flag-${kind}"><span class="wt-map-flag-dot"></span><span class="wt-map-flag-label">${text}</span></span>`;

    return L.marker(point, {
      icon: this.getAnnotationIcon(html),
      keyboard: false,
    }).bindTooltip(tooltip, {
      direction: "top",
      offset: [0, -6],
      opacity: 0.95,
    });
  }

  label(key, fallback) {
    return this.config.labels?.[key] || fallback;
  }

  segmentKey(segment) {
    return `${segment.kind}-${segment.index}`;
  }

  toggleSegmentSelection(segment, markers) {
    if (this.selectedSegmentKey === this.segmentKey(segment)) {
      this.selectedSegment = null;
      this.selectedSegmentKey = null;
      this.selectedMarkers = [];
    } else {
      this.selectedSegment = segment;
      this.selectedSegmentKey = this.segmentKey(segment);
      this.selectedMarkers = markers;
    }

    this.clearSegmentHighlight();
    this.restoreSegmentHighlight();
  }

  restoreSegmentHighlight() {
    if (!this.selectedSegment) {
      return;
    }

    this.drawSegmentHighlight(this.selectedSegment, this.selectedMarkers);
  }

  highlightSegment(segment, markers) {
    this.clearSegmentHighlight();
    this.drawSegmentHighlight(segment, markers);
  }

  drawSegmentHighlight(segment, markers) {
    const positions = this.workout.position.Data.slice(
      segment.startIdx,
      segment.endIdx + 1,
    ).filter((p) => p !== null);

    if (positions.length > 1) {
      const isLift = segment.kind === "lift";

      L.polyline(positions, {
        pane: "segmentHighlight",
        color: isLift ? "#fbbf24" : "#fff",
        weight: isLift ? 10 : 9,
        opacity: isLift ? 1 : 0.9,
        interactive: false,
      }).addTo(this.segmentHighlightLayer);
    }

    for (const marker of markers) {
      marker.getElement()?.classList.add("wt-map-annotation-active");
    }
  }

  clearSegmentHighlight() {
    this.segmentHighlightLayer?.clearLayers();

    for (const element of this.map
      .getContainer()
      .querySelectorAll(".wt-map-annotation-active")) {
      element.classList.remove("wt-map-annotation-active");
    }
  }

  getSegmentAnnotation(segment, which) {
    const glyphs = {
      climb: "▲",
      descent: "▼",
      lift: '<span class="icon-[fa-solid--tram]"></span>',
    };

    const glyph = glyphs[segment.kind] || "?";
    const name = this.label(segment.kind, "Segment");
    const html = `<span class="wt-map-badge wt-map-badge-${segment.kind} wt-map-badge-${which}">${glyph}<b>${segment.index}</b></span>`;
    const direction = segment.kind === "descent" ? "↓" : "↑";
    const distance =
      segment.length >= 1000
        ? `${(segment.length / 1000).toFixed(1)} km`
        : `${Math.round(segment.length)} m`;
    const when = this.label(which, which);
    const slope = `${Math.round(100 * segment.avgSlope)}%`;
    const tooltip = `<b>${name} ${
      segment.index
    }</b> — ${when}<br>${direction} ${Math.round(
      segment.gain,
    )} m · ${distance} · ${slope} · ${formatDuration(segment.duration)}`;

    return L.marker(segment[which], {
      icon: this.getAnnotationIcon(html),
      keyboard: false,
    }).bindTooltip(tooltip, {
      direction: "top",
      offset: [0, -10],
      opacity: 0.95,
    });
  }

  getSlopeLayerGroup(polyLineProperties = {}) {
    return this.buildColorLayer(
      this.workout.position.Data,
      polyLineProperties,
      (i) => this.colorForSlope(this.workout.slope.Data[i] || 0),
    );
  }

  getSpeedLayerGroup(polyLineProperties = {}) {
    const MOVING_AVERAGE_LENGTH = 15;
    const movingSpeeds = [];
    const speeds = this.workout.speed?.Data.filter((x) => x !== null) || [];

    const averageSpeed = speeds.reduce((a, x) => a + x, 0) / speeds.length;
    const stdevSpeed = Math.sqrt(
      speeds.reduce((a, x) => a + Math.pow(x - averageSpeed, 2), 0) /
        (speeds.length - 1),
    );

    return this.buildColorLayer(
      this.workout.position.Data,
      polyLineProperties,
      (i) => {
        const speed = this.workout.speed.Data[i] || null;
        if (speed === null || speed < 0.1) {
          return "rgb(0,0,0)"; // Pausing
        }

        if (movingSpeeds.length > MOVING_AVERAGE_LENGTH) {
          movingSpeeds.shift();
        }
        movingSpeeds.push(speed);
        const movingAverageSpeed =
          movingSpeeds.reduce((a, x) => a + x) / movingSpeeds.length;

        const zScore =
          ((movingAverageSpeed || averageSpeed) - averageSpeed) / stdevSpeed;

        return this.getColor(0.5 + zScore / 2);
      },
    );
  }

  getTooltip(i) {
    const tooltipDisplay = {
      time: "",
      distance: this.preferredUnits.distance,
      duration: "",
      speed: this.preferredUnits.speed,
      elevation: this.preferredUnits.elevation,
      "heart-rate": this.preferredUnits.heartRate,
      cadence: this.preferredUnits.cadence,
      calories: this.preferredUnits.calories,
      power: this.preferredUnits.power,
      temperature: this.preferredUnits.temperature,
      slope: "%",
    };

    let tooltip = `<ul>`;
    for (const [field, unit] of Object.entries(tooltipDisplay)) {
      if (this.workout[field]?.Data[i] !== undefined) {
        const label = this.workout[field].Label;
        const val = this.workout[field].Data[i];
        let formattedVal =
          typeof val === "number" && val % 1 !== 0 ? val.toFixed(2) : val;
        if (field === "duration") {
          formattedVal = formatDuration(val);
        } else if (field === "slope" && val !== null) {
          formattedVal = Math.round(100 * val);
        } else if (field === "time" && val !== null) {
          formattedVal = new Date(val).toTimeString().substr(0, 5);
        }

        if (val !== null) {
          tooltip += `<li><b>${label}</b>: ${formattedVal} ${unit}</li>`;
        }
      }
    }
    tooltip += `</ul>`;
    return tooltip;
  }

  // Determine color for a value; value from 0 to 1
  // Linearly interpolate between blue and green
  getColor(value) {
    value = Math.max(0, Math.min(1, value)); // Clamp to 0...1

    const lowColor = [50, 50, 255];
    const highColor = [50, 255, 50];
    const color = [0, 1, 2].map((i) =>
      Math.floor(value * (highColor[i] - lowColor[i]) + lowColor[i]),
    );
    return `rgb(${color.join(",")})`;
  }

  setMarker(obj) {
    const lat = obj.getAttribute("data-lat");
    const lng = obj.getAttribute("data-lng");
    const title = obj.getAttribute("data-title");

    if (!this.hoverMarker) return;

    if (title != null) {
      this.hoverMarker.bindTooltip(title);
    }

    this.hoverMarker.setLatLng([lat, lng]);

    // Adding popup to the marker
    this.hoverMarker.openTooltip();
  }

  clearMarker() {
    if (!this.hoverMarker) return;
    this.hoverMarker.closeTooltip();
  }

  setSegment(_title, data) {
    this.segmentLayerGroup.clearLayers();

    const positions = data.position.filter((p) => p !== null);
    if (positions.length < 2) {
      return;
    }

    for (let i = 1; i < positions.length; i++) {
      L.polyline([positions[i - 1], positions[i]], {
        color: "red",
      }).addTo(this.segmentLayerGroup);
    }

    this.segmentLayerGroup.addTo(this.map);
  }

  fitSegmentBounds() {
    this.map.fitBounds(this.segmentLayerGroup.getBounds(), {
      animate: false,
      padding: [20, 20],
    });
  }

  resetZoom() {
    this.map.fitBounds(this.trackGroup.getBounds(), { animate: false });
  }

  clearSegment() {
    this.segmentLayerGroup.clearLayers();
    this.resetZoom();
  }

  updateSize() {
    this.map.invalidateSize(true);
  }

  colorForSlope(slope) {
    // See https://github.com/alexgasconn/GPX-Analyzer/blob/bcf16608dda6c748d0fe6d87cd358ca8af4291e5/components/core/utils.py#L4
    switch (true) {
      case slope >= 0.18:
        return "#8B0000"; // Dark Red
      case slope >= 0.1:
        return "#FF8C00"; // Dark Orange
      case slope >= 0.02:
        return "#FFFF00"; // Yellow
      case slope >= 0.0:
        return "#ADFF2F"; // GreenYellow
      case slope >= -0.02:
        return "#ADD8E6"; // LightBlue
      case slope >= -0.1:
        return "#0000FF"; // Blue
      default:
        return "#00008B"; // Dark Blue
    }
  }
}

customElements.define("wt-map", WtMap);
