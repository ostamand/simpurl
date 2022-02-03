const express = require("express");
const passport = require("passport");
const cookieSession = require("cookie-session");

const authRouter = require("./routes/auths");
const { checkLoggedIn } = require("./controllers/auths");

const app = express();

// middlewares
app.use(express.json());

app.use(
  cookieSession({
    keys: [process.env.COOKIE_KEY_1],
    maxAge: 24 * 60 * 60 * 1000,
  })
);
app.use(passport.authenticate("session"));

// routes
app.use("/", authRouter);

app.get("/ping", checkLoggedIn, (_, res) => {
  return res.status(200).json({ message: "pong" });
});

module.exports = app;
