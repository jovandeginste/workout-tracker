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
      page.goto("http://localhost:8080/user/signin?lang=nl"),
    ]);

    page.screenshot({ path: "docs/login-nl.png", fullPage: true });

    await Promise.all([
      page.waitForNavigation(),
      page.goto("http://localhost:8080/user/signin?lang=en"),
    ]);

    page.screenshot({ path: "docs/login.png", fullPage: true });

    page.locator("#signin #username").type("jo@dwarfy.be");
    page.locator("#signin #password").type("test");

    page.setViewportSize({
      height: 1200,
    });
    await Promise.all([page.waitForNavigation(), page.click("button#signin")]);
    page.screenshot({ path: "docs/dashboard.png" });

    await Promise.all([
      page.waitForNavigation(),
      page.goto("http://localhost:8090/user/profile"),
      page.locator("#language").selectOption("nl"),
      page.click("button#update-profile"),
      page.waitForNavigation(),
    ]);
    page.screenshot({ path: "docs/profile-nl.png" });

    await Promise.all([
      page.waitForNavigation(),
      page.goto("http://localhost:8090/user/profile"),
      page.locator("#language").selectOption("de"),
      page.click("button#update-profile"),
      page.waitForNavigation(),
    ]);
    page.screenshot({ path: "docs/profile-de.png" });

    await Promise.all([
      page.waitForNavigation(),
      page.goto("http://localhost:8090/user/profile"),
      page.locator("#language").selectOption("fa"),
      page.click("button#update-profile"),
      page.waitForNavigation(),
    ]);
    page.screenshot({ path: "docs/profile-fa.png" });

    await Promise.all([
      page.waitForNavigation(),
      page.goto("http://localhost:8090/user/profile"),
      page.locator("#language").selectOption("en"),
      page.click("button#update-profile"),
      page.waitForNavigation(),
    ]);
    page.screenshot({ path: "docs/profile-en.png" });

    page.setViewportSize({
      height: 800,
    });
    await Promise.all([
      page.waitForNavigation(),
      page.goto("http://localhost:8080/workouts"),
    ]);
    page.screenshot({ path: "docs/workout_overview.png" });

    await Promise.all([
      page.waitForNavigation(),
      page.goto("http://localhost:8080/daily?count=5"),
    ]);
    page.screenshot({ path: "docs/daily_overview.png" });

    page.setViewportSize({
      height: 400,
    });
    await Promise.all([
      page.waitForNavigation(),
      page.goto("http://localhost:8080/heatmap"),
    ]);
    sleep(2);
    page.screenshot({ path: "docs/heatmap.png", fullPage: true });

    await Promise.all([
      page.waitForNavigation(),
      page.goto("http://localhost:8080/statistics"),
    ]);
    page.screenshot({ path: "docs/statistics.png", fullPage: true });

    await Promise.all([
      page.waitForNavigation(),
      page.goto("http://localhost:8080/workouts/249"),
    ]);
    page.screenshot({ path: "docs/single_workout-dark.png", fullPage: true });

    await Promise.all([
      page.waitForNavigation(),
      page.goto("http://localhost:8080/workouts/add"),
    ]);
    page.screenshot({ path: "docs/upload_workouts.png", fullPage: true });
    const options = page.locator("#manual #type");
    await options.selectOption("running");
    await page.waitForSelector("#manual #location");
    page.locator("#manual #location").fill("brussels");
    await page.waitForSelector("#manual #address-results");
    page.screenshot({
      path: "docs/upload_workouts_manual.png",
      fullPage: true,
    });

    page.emulateMedia({
      colorScheme: "light",
    });
    await Promise.all([
      page.waitForNavigation(),
      page.goto("http://localhost:8080/workouts/249"),
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
      page.goto("http://localhost:8080/workouts/249"),
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
