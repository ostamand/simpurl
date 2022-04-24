const { getUserByID } = require("../models/users.model");

//TODO add error management
async function addTagIfNeeded(userID, tags) {
  const user = await getUserByID(userID);
  let added = false;
  tags.forEach((tag) => {
    if (!user.tags.includes(tag)) {
      user.tags.push(tag);
      added = true;
    }
  });
  if (added) {
    await user.save();
  }
}

module.exports = {
  addTagIfNeeded,
};
