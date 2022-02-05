const Url = require("./urls.mongo");

async function getAllURLs(userID) {}

async function getLastID(userID) {
  let lastID = await Url.findOne({ userID: userID }, "urlID")
    .sort("-urlID")
    .exec();
  if (!lastID) {
    lastID = 0;
  }
  return lastID;
}

module.exports = {
  getAllURLs,
  getLastID,
};
