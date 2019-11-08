// const axios = require("axios");

// const url = "https://myki.com/blog/";
// let cheerio = require("cheerio");
// axios.get(url).then(res => {
//   getHtml(res.data, ".post-card-title");
// });

// const getHtml = async (html, selector) => {
//   content = [];
//   try {
//     const $ = cheerio.load(html);
//     $(selector).each((i, elem) => {
//       content.push({
//         meta: $(elem).text()
//       });
//       console.log(content);
//     });
//     debugger;
//   } catch (err) {
//     console.log(err);
//     throw new Error("scraping failed");
//   }
// };

const axios = require("axios");
const puppeteer = require("puppeteer");

const url = "https://www.google.com/search?client=firefox-b-d&q=Myki";
let cheerio = require("cheerio");

(async () => {
  const browser = await puppeteer.launch({ headless: true });
  const page = await browser.newPage();
  await page.setViewport({ width: 1920, height: 926 });
  await page.goto(url);

  const $ = cheerio.load(page.evaluate());

  let data = document.querySelectorAll(
    "img.lu-fs"
  )
  try {
    $(selector).each((i, elem) => {
      console.log(elem);
      content.push({
        meta: $(elem).text()
      });
      console.log(content);
    });
    debugger;
  } catch (err) {
    console.log(err);
    throw new Error("scraping failed");
  }
})();


const getHtml = async (html, selector) => {
  content = [];
  try {
    const $ = cheerio.load(html);
    $(selector).each((i, elem) => {
      console.log(elem);
      content.push({
        meta: $(elem).text()
      });
      console.log(content);
    });
    debugger;
  } catch (err) {
    console.log(err);
    throw new Error("scraping failed");
  }
};
