const express = require("express");
const { checkLoggedIn } = require("../controllers/auths");
const { httpRedirect } = require("../controllers/redirect");

const redirectRouter = express.Router();

redirectRouter.get("/:symbol", checkLoggedIn, httpRedirect);

module.exports = redirectRouter;
