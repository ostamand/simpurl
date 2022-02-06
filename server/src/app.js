const cors = require("cors");
const path = require('path');
const express = require("express");
const passport = require("passport");
const cookieSession = require("cookie-session");

const redirectRouter = require("./routes/redirect");
const apiRouter = require("./routes/api")

const { checkLoggedIn } = require("./controllers/auths");

const app = express();

// middlewares
app.use(cors({"origin": "http://localhost:1234, https://shorturl-w723ubjq4a-uk.a.run.app", credentials: true})) //! this is for dev
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
app.use("/api", apiRouter)

app.use(express.static(path.join(__dirname, "..", "public")))

app.use("/", redirectRouter);

module.exports = app;
