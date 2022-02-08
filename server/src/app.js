const path = require("path");
const express = require("express");
const passport = require("passport");
const cookieSession = require("cookie-session");

const redirectRouter = require("./routes/redirect");
const apiRouter = require("./routes/api");

const { checkLoggedIn } = require("./controllers/auths");

require("dotenv").config();

const app = express();

//middlewares
app.use((req, res, next) => {
  const allowedOrigins = process.env.CORS.split(",").map((s) => s.trim());
  const origin = req.headers.origin;
  if (allowedOrigins.includes(origin)) {
    res.setHeader("Access-Control-Allow-Origin", origin);
  }
  res.header("Access-Control-Allow-Headers", "Content-Type,Authorization");
  res.header("Access-Control-Allow-Credentials", true);
  res.header(
    "Access-Control-Allow-Methods",
    "POST, GET, OPTIONS, PATCH, DELETE"
  );
  return next();
});

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
app.use("/api", apiRouter);

app.use(express.static(path.join(__dirname, "..", "public")));

app.use("/", redirectRouter);

module.exports = app;
