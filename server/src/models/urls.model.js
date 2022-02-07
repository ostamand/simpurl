const Url = require("./urls.mongo");

async function getAllURLs(userID) {}

async function getLastID(userID) {
  let lastID = 0;
  let url = await Url.findOne({ userID: userID }, "urlID")
    .sort("-urlID")
    .exec();
  if (url) {
    lastID = url.urlID;
  }
  return lastID;
}

module.exports = {
  getAllURLs,
  getLastID,
};
