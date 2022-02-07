const express = require("express");
const {
  createUrl,
  deleteUrl,
  getUrl,
  getAllUrls,
} = require("../controllers/urls");
const { checkLoggedIn } = require("../controllers/auths");

const urlsRouter = express.Router();

urlsRouter.get("/", checkLoggedIn, getAllUrls);
urlsRouter.get("/:id", checkLoggedIn, getUrl);
urlsRouter.post("/", checkLoggedIn, createUrl);
urlsRouter.delete("/:id", checkLoggedIn, deleteUrl);

module.exports = urlsRouter;
