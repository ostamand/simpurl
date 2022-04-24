const { getUserByID } = require("../models/users.model");

async function getAllTags(req, res) {
  const user = await getUserByID(req.user.id);
  return res.status(200).json(user.tags);
}

module.exports = {
  getAllTags,
};
