
import { chromium } from "playwright";


  const browser = await chromium.launch({
    executablePath: "/usr/bin/chromium",
    args: ['--no-sandbox']
  });
  const page = await browser.newPage();
  await page.goto("https://example.com/");
  await page.waitForTimeout(2000);

  const screenshot = await page.screenshot();
  await browser.close();
  console.log("Screenshot taken");
