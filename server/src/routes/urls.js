const express = require("express");
const { createUrl, deleteUrl, getUrl } = require("../controllers/urls");
const { checkLoggedIn } = require("../controllers/auths");

const urlsRouter = express.Router();

urlsRouter.get("/:id", checkLoggedIn, getUrl);
urlsRouter.post("/", checkLoggedIn, createUrl);
urlsRouter.delete("/:id", checkLoggedIn, deleteUrl);

module.exports = urlsRouter;
