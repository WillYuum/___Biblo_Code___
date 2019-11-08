const puppeteer = require("puppeteer");

let url = "https://www.linkedin.com/company/mykisecurity/";

(async () => {
  const browser = await puppeteer.launch({ headless: true });
  const page = await browser.newPage();
  await page.setViewport({ width: 1920, height: 926 });
  await page.goto(url);
  console.log("here", page);

  // get hotel details
  let hotelData = await page.evaluate(() => {
    let hotels = [];
    // get the hotel elements

    try {
      let hotelsElms = document.querySelector(
        "div.org-top-card-summary__info-item"
      ).innerHTML;

      console.log("here", hotelsElms);
    } catch (exception) {}

    return hotels;
  });

  console.log(hotelData);
})();
