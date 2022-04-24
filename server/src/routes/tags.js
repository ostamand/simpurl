const express = require("express");
const { checkLoggedIn } = require("../controllers//auths");
const { getAllTags } = require("../controllers/tags");

const tagsRouter = express.Router();

tagsRouter.get("", checkLoggedIn, getAllTags);

module.exports = tagsRouter;
