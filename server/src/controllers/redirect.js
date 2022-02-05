const Url = require("../models/urls.mongo");

async function httpRedirect(req, res) {
  const url = await Url.findOne({
    userID: req.user.id,
    symbol: req.params.symbol,
  });
  if (url) {
    return res.redirect(url.url);
  }
  return res.redirect("/"); //TODO: where to redirect?
}

module.exports = {
  httpRedirect,
};
