const express = require("express");
const passport = require("passport");
const LocalStrategy = require("passport-local");

const { httpCreateNewUser } = require("../controllers/auths");
const { validateUser, getUserByID } = require("../models/users.model");

const router = express.Router();

passport.use(
  new LocalStrategy(async (username, password, done) => {
    const user = await validateUser(username, password);
    if (user) {
      return done(null, user);
    }
    return done(null, false);
  })
);

passport.serializeUser((user, done) => {
  done(null, { id: user._id });
});

passport.deserializeUser(async (userData, done) => {
  const user = await getUserByID(userData.id);
  if (user) {
    return done(null, user);
  }
  return done(null, false);
});

router.post("/signup", httpCreateNewUser);

router.post("/signin", passport.authenticate("local"), (req, res) => {
  return res
    .status(200)
    .json((({ username, email }) => ({ username, email }))(req.user));
});

module.exports = router;
