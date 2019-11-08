const puppeteer = require('puppeteer');

(async () => {
  // Set up browser and page.
  const browser = await puppeteer.launch({headless:false});
  const page = await browser.newPage();
  page.setViewport({ width: 1280, height: 926 });

  // Navigate to this blog post and wait a bit.
  await page.goto('https://intoli.com/blog/saving-images/');
  await page.waitForSelector('#png');

  // Select the #svg img element and save the screenshot.
  const svgImage = await page.$('#png');
  const a = await svgImage.screenshot({
    path: 'dir1/dir2/bob.png',
    type: "png",
    omitBackground: true,
    encoding: "base64"
  });
  console.log("hi", a.toJson())

  await browser.close();
})();