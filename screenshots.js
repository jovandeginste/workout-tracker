import { browser } from "k6/experimental/browser";

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
  const page = browser.newPage();

  try {
    await page.goto("http://localhost:8080");
    page.screenshot({ path: "docs/login.png" });
  } finally {
    page.close();
  }
}
