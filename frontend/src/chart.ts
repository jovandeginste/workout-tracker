export function isDark(colorMode = "browser"): boolean {
  if (colorMode === "dark") return true;
  if (colorMode === "light") return false;

  return (
    window.matchMedia &&
    window.matchMedia("(prefers-color-scheme: dark)").matches
  );
}

export function chartTheme(colorMode = "browser"): {
  dark: boolean;
  fgColor: string;
  gridColor: string;
} {
  const dark = isDark(colorMode);

  return {
    dark,
    fgColor: dark ? "#e4e4e7" : "#27272a",
    gridColor: dark ? "#3f3f46" : "#d4d4d8",
  };
}

export async function saveConfig(
  route: string,
  body: URLSearchParams,
  failureMessage: string,
): Promise<void> {
  if (!route) {
    return;
  }

  try {
    const response = await fetch(route, {
      method: "POST",
      headers: {
        "Content-Type": "application/x-www-form-urlencoded",
      },
      body,
    });

    if (response.ok) {
      return;
    }
  } catch (_err) {
    // Keep the selected chart setting even if saving the preference fails.
  }

  console.warn(failureMessage);
}
