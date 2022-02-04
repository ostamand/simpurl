const { createUser, userToObject } = require("../models/users.model");

async function httpCreateNewUser(req, res) {
  const { user, err } = await createUser(req.body);
  if (err) {
    return res.status(500);
  }
  return res.status(201).json(userToObject(user));
}

function checkLoggedIn(req, res, next) {
  const isLoggedIn = req.isAuthenticated() && req.user;
  if (!isLoggedIn) {
    return res.status(401).json({
      error: "You must log in.",
    });
  }
  next();
}

module.exports = {
  httpCreateNewUser,
  checkLoggedIn,
};
