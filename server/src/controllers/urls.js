const Url = require("../models/urls.mongo");

async function createUrl(req, res) {
  const data = req.body;
  // check if exists already
  let url = await Url.findOne({ url: data.url }).exec();
  if (url) {
    return res.status(409).json(url);
  }
  data.userID = req.user.id;
  url = new Url(data);
  try {
    url = await url.save();
    return res.status(200).json(url);
  } catch (err) {
    return res.status(500).json({ message: "Internal error." });
  }
}

module.exports = {
  createUrl,
};
