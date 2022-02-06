const cors = require("cors");
const path = require('path');
const express = require("express");
const passport = require("passport");
const cookieSession = require("cookie-session");

const authRouter = require("./routes/auths");
const urlRouter = require("./routes/urls");
const redirectRouter = require("./routes/redirect");

const { checkLoggedIn } = require("./controllers/auths");

const app = express();

// middlewares
app.use(cors({"origin": "http://localhost:1234", credentials: true})) //! this is for dev
app.use(express.json());
app.use(
  cookieSession({
    secure: false,
    keys: [process.env.COOKIE_KEY_1],
    maxAge: 24 * 60 * 60 * 1000,
    httpOnly: false, //! to get access thru js for now
  })
);
app.use(passport.authenticate("session"));

// routes
app.use("/", authRouter);
app.use("/urls", urlRouter);

app.get("/ping", (_, res) => {
  return res.status(200).json({ message: "pong" });
});

app.get("/ping-secure", checkLoggedIn, (_, res) => {
  return res.status(200).json({ message: "still pong but secured this time" });
});

app.use(express.static(path.join(__dirname, "..", "public")))

app.use("/", redirectRouter);

module.exports = app;
