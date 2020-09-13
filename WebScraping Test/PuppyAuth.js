const puppeteer = require("puppeteer");
const username = "bevy_scrape@outlook.com";
const password = "devel wer";
const USERNAME_SELECTOR = "#username";
const PASSWORD_SELECTOR = "#password";
const CTA_SELECTOR = ".btn__primary--large";

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
  console.log("we are in", url);
  await page.click(USERNAME_SELECTOR);
  await page.keyboard.type(username);
  await page.click(PASSWORD_SELECTOR);
  await page.keyboard.type(password);
  await page.click(CTA_SELECTOR);
  console.log("BUTTON CLICK ON", CTA_SELECTOR);
  //   await page.waitForNavigation();
  await page.screenshot({ path: "linkedin.png" });
  console.log("playTesting finished");

  await page.goto(
    "https://www.linkedin.com/search/results/companies/?keywords=software%20company%2C%20lebanon&origin=SWITCH_SEARCH_VERTICAL&page=5"
  );
  const x = await page.waitForSelector("search-result");
  console.log("hey", x);
  const getCompanyNames = await page.evaluate(() => {
    let companyNames = [];
    try {
      let companyElem = document.querySelectorAll("h3.search-result__title");
      companyElem.forEach(elem => {
        companyNames.push(elem.innerHTML);
      });
      return companyNames;
    } catch (err) {
      throw new Error("scarping company names failed");
    }
  });
  console.log("get to my ", getCompanyNames);
}

(async () => {
  try {
    await playTest("https://www.linkedin.com/login");
    // process.exit(1);
  } catch (err) {
    console.log(err);
    throw new Error("scraping failed");
  }
})();
