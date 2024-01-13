/**
 * @license
 * Copyright 2017 Google Inc.
 * SPDX-License-Identifier: Apache-2.0
 */

'use strict';

import puppeteer from 'puppeteer';

(async () => {
  const browser = await puppeteer.launch({
    executablePath: "/usr/bin/chromium",
    args: ['--no-sandbox']
  });
  const page = await browser.newPage();
  await page.goto('http://example.com');
  await page.screenshot({path: 'example.png'});
  await browser.close();
})();