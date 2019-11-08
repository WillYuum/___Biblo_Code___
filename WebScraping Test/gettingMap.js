const puppeteer = require("puppeteer");
const mapLink = "";
const mapImage = "lu-fs";

async function startBrowser() {
  const browser = await puppeteer.launch({ headless: false });
  const page = await browser.newPage();
  console.log("starting browser finished");
  return { browser, page };
}

async function closeBrowser(browser) {
  return browser.close();
}

async function playTest(url) {
  const { browser, page } = await startBrowser();
  page.setViewport({ width: 1366, height: 768 });
  await page.goto(url);
  //   await page.waitForNavigation();
  await page.screenshot({ path: "googleSearch.png" });
  console.log("playTesting finished");

  const getData = await page.evaluate(() => {
    mapImage = document.querySelector("img.lu-fs").getAttribute("src");
    return mapImage;
  });
  console.log("second", getData);
}

(async () => {
  try {
    await playTest("https://www.google.com/search?client=firefox-b-d&q=Myki");
    // process.exit(1);
  } catch (err) {
    console.log(err);
    throw new Error("scraping failed");
  }
})();
