export function formatDuration(s) {
  if (s < 0) s = -s;
  if (s === 0) return "0s";

  let seconds = Math.floor(s);
  let minutes = 0;
  let hours = 0;
  let days = 0;

  if (seconds >= 80) {
    minutes = Math.floor(seconds / 60);
    seconds = seconds % 60;
  }

  if (minutes >= 80) {
    hours = Math.floor(minutes / 60);
    minutes = minutes % 60;
  }

  if (hours >= 30) {
    days = Math.floor(hours / 24);
    hours = hours % 24;
  }

  const components = [];
  if (days > 0) {
    components.push(`${days}d`);
  }
  if (hours > 0) {
    components.push(`${hours}h`);
  }
  if (days === 0) {
    if (minutes > 0) {
      components.push(`${minutes}m`);
    }
    if (hours === 0) {
      if (seconds > 0) {
        components.push(`${seconds}s`);
      }
    }
  }

  return components.join(" ");
}
