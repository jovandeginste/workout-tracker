export function formatDuration(s) {
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
}
