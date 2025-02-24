import { browser } from "k6/browser";
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
  const page = await browser.newPage({
    colorScheme: "dark",
    viewport: {
      width: 1874,
      height: 400,
    },
  });

  try {
    await page.goto("http://localhost:8080/user/signin?lang=nl");
    await page.screenshot({ path: "docs/login-nl.png", fullPage: true });

    await page.goto("http://localhost:8080/user/signin?lang=en");
    await page.screenshot({ path: "docs/login.png", fullPage: true });

    await page.locator("#signin #username").type("jo@dwarfy.be");
    await page.locator("#signin #password").type("test");

    await page.setViewportSize({
      height: 1200,
    });
    await Promise.all([page.click("button#signin"), page.waitForNavigation()]);
    await page.screenshot({ path: "docs/dashboard.png" });

    await Promise.all([
      page.goto("http://localhost:8080/user/profile"),
      page.locator("#language").selectOption("nl"),
      page.click("button#update-profile"),
    ]);
    await page.screenshot({ path: "docs/profile-nl.png" });

    await Promise.all([
      page.goto("http://localhost:8080/user/profile"),
      page.locator("#language").selectOption("de"),
      page.click("button#update-profile"),
    ]);
    await page.screenshot({ path: "docs/profile-de.png" });

    await Promise.all([
      page.goto("http://localhost:8080/user/profile"),
      page.locator("#language").selectOption("fa"),
      page.click("button#update-profile"),
    ]);
    await page.screenshot({ path: "docs/profile-fa.png" });

    await Promise.all([
      page.goto("http://localhost:8080/user/profile"),
      page.locator("#language").selectOption("en"),
      page.click("button#update-profile"),
    ]);
    await page.screenshot({ path: "docs/profile-en.png" });

    await page.setViewportSize({
      height: 800,
    });
    await page.goto("http://localhost:8080/workouts");
    await page.screenshot({ path: "docs/workout_overview.png" });

    await page.goto("http://localhost:8080/daily?count=5");
    await page.screenshot({ path: "docs/daily_overview.png" });

    await page.setViewportSize({
      height: 400,
    });
    await page.goto("http://localhost:8080/heatmap");
    sleep(2);
    await page.screenshot({ path: "docs/heatmap.png", fullPage: true });

    await page.goto("http://localhost:8080/statistics");
    await page.screenshot({ path: "docs/statistics.png", fullPage: true });

    await page.goto("http://localhost:8080/workouts/249");
    await page.screenshot({
      path: "docs/single_workout-dark.png",
      fullPage: true,
    });

    await page.goto("http://localhost:8080/workouts/add");
    await page.screenshot({ path: "docs/upload_workouts.png", fullPage: true });

    const options = page.locator("#manual #type");
    await options.selectOption("running");
    await page.waitForSelector("#manual #location");
    await page.locator("#manual #location").fill("brussels");
    await page.waitForSelector("#manual #address-results");
    await page.screenshot({
      path: "docs/upload_workouts_manual.png",
      fullPage: true,
    });

    await page.emulateMedia({
      colorScheme: "light",
    });
    await page.goto("http://localhost:8080/workouts/249");
    await page.screenshot({
      path: "docs/single_workout-light.png",
      fullPage: true,
    });

    // Create screenshots for responsive view
    await page.setViewportSize({
      width: 600,
      height: 2000,
    });
    await page.emulateMedia({
      colorScheme: "dark",
    });

    await page.goto("http://localhost:8080/");
    await page.screenshot({ path: "docs/dashboard-responsive.png" });

    await page.goto("http://localhost:8080/workouts/249");
    await page.screenshot({ path: "docs/single_workout-responsive.png" });

    await page.goto("http://localhost:8080/statistics");
    await page.screenshot({ path: "docs/statistics-responsive.png" });
  } finally {
    await page.close();
  }
}
