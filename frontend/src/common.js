// Close all modals when press ESC
document.onkeydown = function (event) {
  event = event || window.event;
  if (event.keyCode === 27) {
    document
      .getElementsByTagName("body")[0]
      .classList.remove("overflow-y-hidden");
    let modals = document.getElementsByClassName("modal");
    Array.prototype.slice.call(modals).forEach((i) => {
      i.style.display = "none";
    });
  }
};

globalThis.toggleTextPassword = function toggleTextPassword(el, id) {
  var x = document.getElementById(id);

  if (x.type === "password") {
    // Remove class from span
    el.classList.remove("icon-eye");
    el.classList.add("icon-eye-slash");
    x.type = "text";
  } else {
    el.classList.remove("icon-eye-slash");
    el.classList.add("icon-eye");
    x.type = "password";
  }
};

globalThis.copyToClipboard = function copyToClipboard(id) {
  // Get the text field
  var copyText = document.getElementById(id);

  // Select the text field
  copyText.select();
  // Get the text field
  copyText.setSelectionRange(0, 99999); // For mobile devices

  // Copy the text inside the text field
  navigator.clipboard.writeText(copyText.value);
  var noticeText = copyText.getAttribute("data-notice");
  if (noticeText != null) {
    showMessage("notice", noticeText);
  }
};

globalThis.showMessage = function showMessage(cls, message) {
  var al = document.getElementById("alerts");

  var msg = document.createElement("div");
  msg.classList.add(cls);
  msg.innerText = message;

  al.appendChild(msg);
};

globalThis.geoJson2heat = function geoJson2heat(geojson, intensity) {
  return geojson.features.map(function (feature) {
    return [
      parseFloat(feature.geometry.coordinates[1]),
      parseFloat(feature.geometry.coordinates[0]),
      intensity,
    ];
  });
};

globalThis.editDaily = function editDaily(obj) {
  var date = obj.getAttribute("data-date");
  var height = obj.getAttribute("data-height");
  var weight = obj.getAttribute("data-weight");
  var steps = obj.getAttribute("data-steps");

  document.getElementById("date").value = date;
  document.getElementById("height").value = height;
  document.getElementById("weight").value = weight;
  document.getElementById("steps").value = steps;

  readDailyHeight();
};

globalThis.updateDailyHeight = function updateDailyHeight() {
  var ft = document.getElementById("ft");
  var inch = document.getElementById("in");
  var height = document.getElementById("height");

  height.value = parseInt(ft.value) * 12 + parseInt(inch.value);
};

globalThis.readDailyHeight = function readDailyHeight() {
  var ft = document.getElementById("ft");
  var inch = document.getElementById("in");
  var height = document.getElementById("height");
  ft.value = Math.floor(height.value / 12);
  inch.value = height.value % 12;
};

globalThis.fullMap = function fullMap(map) {
  const d = document.getElementById("map-container");
  const mapEl = document.getElementById(map);

  d.classList.toggle("small-size");
  d.classList.toggle("full-size");

  if (mapEl) {
    mapEl.updateSize();
  }

  return false;
};

globalThis.showTab = function showTab(parentId, elemId) {
  var parent = document.getElementById(parentId);
  if (!parent) return;

  var origSize = parent.offsetHeight;
  if (parent.style.minHeight === "") {
    parent.style.minHeight = origSize + "px";
  }

  parent.querySelectorAll(".tabbed-content .tab").forEach((element) => {
    if (element.id === "tab-" + elemId) {
      element.classList.add("selected");
    } else if (element.id === "content-" + elemId) {
      element.classList.add("selected");
    } else {
      element.classList.remove("selected");
    }
  });

  // We only let the height increase, this avoids jarring transitions and
  // scrolling
  var newSize = parent.offsetHeight;
  parent.style.minHeight = Math.max(origSize, newSize) + "px";
};
