const express = require("express");

const authRouter = require("./auths");
const urlRouter = require("./urls");
const tagRouter = require("./tags");

const { checkLoggedIn } = require("../controllers/auths");

const apiRouter = express.Router();

apiRouter.use("/", authRouter);
apiRouter.use("/urls", urlRouter);
apiRouter.use("/tags", tagRouter);

apiRouter.get("/ping", (_, res) => {
  return res.status(200).json({ message: "pong" });
});

apiRouter.get("/ping-secure", checkLoggedIn, (_, res) => {
  return res.status(200).json({ message: "still pong but secured this time" });
});

module.exports = apiRouter;
