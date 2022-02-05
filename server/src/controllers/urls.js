const Url = require("../models/urls.mongo");
const { getLastID } = require("../models/urls.model");

async function createUrl(req, res) {
  const data = req.body;
  // check if exists already
  let url = await Url.findOne({ url: data.url });
  if (url) {
    return res.status(409).json(url);
  }

  data.userID = req.user.id;
  data.urlID = (await getLastID(data.userID)) + 1;
  data.createdAt = Date.now();

  try {
    url = await new Url(data).save();
    return res.status(200).json(url);
  } catch (err) {
    console.error(err);
    return res.status(500).json({ message: "Internal error." });
  }
}

async function deleteUrl(req, res) {
  const id = Number(req.params.id);
  const result = await Url.deleteOne({ userID: req.user.id, urlID: id });
  if (result.deletedCount > 0) {
    return res.status(200).json({ message: `URL with id ${id} deleted.` });
  }
  return res.status(404).json({ message: `URL with id ${id} not found.` });
}

async function getUrl(req, res) {
  const id = Number(req.params.id);
  const url = await Url.findOne({ userID: req.user.id, urlID: id });
  if (url) {
    return res.status(200).json(url);
  }
  return res.status(404).json({ message: `URL with id ${id} not found.` });
}

module.exports = {
  getUrl,
  createUrl,
  deleteUrl,
};
