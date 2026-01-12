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
      width: 1200,
      height: 800,
    },
  });

  try {
    // 1. Sign in
    await page.goto("http://localhost:8180/user/signin?lang=en");
    await page.locator("#signin #username").type("jo@dwarfy.be");
    await page.locator("#signin #password").type("test");
    await Promise.all([page.click("button#signin"), page.waitForNavigation()]);

    // 2. Open a workout (using the same ID as in screenshots.js)
    await page.goto("http://localhost:8180/workouts/487");
    
    // Wait for the map to be rendered
    const mapElement = await page.waitForSelector("wt-map#workout-map");
    
    // Give leaflet some time to initialize tiles and layers
    sleep(2);

    // 3. Move the mouse over the track
    // We'll get the workout data to find some coordinates to hover over
    const workoutData = await page.evaluate(() => {
        const data = JSON.parse(document.getElementById("workout-data").textContent);
        return data;
    });

    const positions = workoutData.position.Data;
    const numFrames = 20;
    const step = Math.max(1, Math.floor(positions.length / numFrames));

    // Ensure the docs directory exists (managed by the test runner usually, but good for local)
    // We'll save individual frames which will later be combined into a gif using imagemagick
    for (let i = 0; i < numFrames; i++) {
        const pointIdx = Math.min(i * step, positions.length - 1);
        const point = positions[pointIdx];
        
        // We use page.evaluate to call the setMarker method on our custom element
        // This triggers the internal hover logic of the wt-map component
        await page.evaluate((idx) => {
            const wtMap = document.querySelector("wt-map#workout-map");
            // Mocking an object with data-lat, data-lng, data-title for setMarker
            const positions = JSON.parse(document.getElementById("workout-data").textContent).position.Data;
            const p = positions[idx];
            const mockObj = document.createElement('div');
            mockObj.setAttribute('data-lat', p[0]);
            mockObj.setAttribute('data-lng', p[1]);
            
            // Get tooltip content manually to provide as data-title
            const tooltipContent = wtMap.getTooltip(idx);
            mockObj.setAttribute('data-title', tooltipContent);
            
            wtMap.setMarker(mockObj);
        }, pointIdx);

        await page.screenshot({ path: `docs/track-frame-${String(i).padStart(3, '0')}.png` });
        sleep(0.1);
    }

  } finally {
    await page.close();
  }
}
