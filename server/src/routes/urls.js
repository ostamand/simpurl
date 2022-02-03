express = require("express")
const { addURL } = require("../controllers/urls");

const urlsRouter = express.Router();

urlsRouter.post("/", addURL)

module.exports = urlsRouter