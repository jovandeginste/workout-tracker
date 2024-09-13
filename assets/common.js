window.openModal = function (modalId) {
  document.getElementById(modalId).style.display = "block";
  document.getElementsByTagName("body")[0].classList.add("overflow-y-hidden");
};

window.closeModal = function (modalId) {
  document.getElementById(modalId).style.display = "none";
  document
    .getElementsByTagName("body")[0]
    .classList.remove("overflow-y-hidden");
};

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

const formatDuration = (s) => {
  if (s < 0) s = -s;
  const time = {
    d: Math.floor(s / 86_400),
    h: Math.floor(s / 3_600) % 24,
    m: Math.floor(s / 60) % 60,
    s: Math.floor(s) % 60,
  };
  return Object.entries(time)
    .filter((val) => val[1] !== 0)
    .map(([key, val]) => `${val}${key}`)
    .join(" ");
};

function toggleTextPassword(el, id) {
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
}

function copyToClipboard(id) {
  // Get the text field
  var copyText = document.getElementById(id);

  // Select the text field
  copyText.select();
  // Get the text field
  copyText.setSelectionRange(0, 99999); // For mobile devices

  // Copy the text inside the text field
  navigator.clipboard.writeText(copyText.value);
}

function showMessage(cls, message) {
  var al = document.getElementById("alerts");

  var msg = document.createElement("div");
  msg.classList.add(cls);
  msg.innerText = message;

  al.appendChild(msg);
}
