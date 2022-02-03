const bcrypt = require("bcrypt");

const User = require("../models/users.mongo");

function hashPassword(password) {
  return new Promise((resolve, _) => {
    bcrypt.genSalt(10, (_, salt) => {
      bcrypt.hash(password, salt, (_, hash) => {
        resolve(hash);
      });
    });
  });
}

function validatePassword(password, hash) {
  return new Promise((resolve, _) => {
    bcrypt.compare(password, hash, (err, result) => {
      if (!err && result) {
        resolve(true);
      }
      resolve(false);
    });
  });
}

/**
 * Save new user to database
 *
 * @param {Object} data
 * @param {String} data.email
 * @param {String} data.username
 * @param {String} data.password
 */
async function createUser(data) {
  const hash = await hashPassword(data.password);
  data.hashedPassword = hash;
  let user = new User(data);
  delete data.hashedPassword; // to not modify the data provided
  try {
    user = await user.save();
    return { user, err: null };
  } catch (err) {
    return { user: null, err };
  }
}

async function getUserByID(id) {
  return await User.findById(id);
}

async function validateUser(username, password) {
  const user = await User.findOne({ username }).exec();
  if (user) {
    const ok = await validatePassword(password, user.hashedPassword);
    if (ok) {
      return user;
    }
  }
  return null;
}

module.exports = {
  createUser,
  validateUser,
  getUserByID,
};
