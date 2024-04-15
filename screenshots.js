import { browser } from "k6/experimental/browser";
import { sleep } from "k6";

export const options = {
  scenarios: {
    browser: {
      executor: "shared-iterations",
      options: {
        browser: {
          type: "chromium",
        },
      },
    },
  },
};

export default async function () {
  const page = browser.newPage({
    colorScheme: "dark",
    viewport: {
      width: 1874,
      height: 400,
    },
  });

  try {
    await Promise.all([
      page.waitForNavigation(),
      page.goto("http://localhost:8080"),
    ]);

    page.screenshot({ path: "docs/login.png", fullPage: true });

    page.locator("#signin #username").type("jo@dwarfy.be");
    page.locator("#signin #password").type("test");

    page.setViewportSize({
      height: 1200,
    });
    await Promise.all([page.waitForNavigation(), page.click("button#signin")]);
    page.screenshot({ path: "docs/dashboard.png" });

    page.setViewportSize({
      height: 800,
    });
    await Promise.all([
      page.waitForNavigation(),
      page.goto("http://localhost:8080/workouts"),
    ]);
    page.screenshot({ path: "docs/workout_overview.png" });

    page.setViewportSize({
      height: 400,
    });
    await Promise.all([
      page.waitForNavigation(),
      page.goto("http://localhost:8080/statistics"),
    ]);
    page.screenshot({ path: "docs/statistics.png", fullPage: true });

    await Promise.all([
      page.waitForNavigation(),
      page.goto("http://localhost:8080/workouts/10"),
    ]);
    page.screenshot({ path: "docs/single_workout-dark.png", fullPage: true });

    await Promise.all([
      page.waitForNavigation(),
      page.goto("http://localhost:8080/workouts/add"),
    ]);
    page.screenshot({ path: "docs/upload_workouts.png", fullPage: true });

    page.emulateMedia({
      colorScheme: "light",
    });
    await Promise.all([
      page.waitForNavigation(),
      page.goto("http://localhost:8080/workouts/10"),
    ]);
    page.screenshot({ path: "docs/single_workout-light.png", fullPage: true });

    // Create screenshots for responsive view
    page.setViewportSize({
      width: 600,
      height: 2000,
    });
    page.emulateMedia({
      colorScheme: "dark",
    });

    await Promise.all([
      page.waitForNavigation(),
      page.goto("http://localhost:8080/"),
    ]);
    page.screenshot({ path: "docs/dashboard-responsive.png" });

    await Promise.all([
      page.waitForNavigation(),
      page.goto("http://localhost:8080/workouts/10"),
    ]);
    page.screenshot({ path: "docs/single_workout-responsive.png" });

    await Promise.all([
      page.waitForNavigation(),
      page.goto("http://localhost:8080/statistics"),
    ]);
    page.screenshot({ path: "docs/statistics-responsive.png" });
  } finally {
    page.close();
  }
}
