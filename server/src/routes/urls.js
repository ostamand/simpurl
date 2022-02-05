express = require("express");
const { createUrl } = require("../controllers/urls");

const urlsRouter = express.Router();

urlsRouter.post("/", createUrl);

module.exports = urlsRouter;
