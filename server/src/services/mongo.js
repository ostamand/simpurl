const mongoose = require("mongoose");
require("dotenv").config();

mongoose.connection.once("connected", () => {
  console.log("Connected to MongoDB");
});

mongoose.connection.on("error", (err) => {
  console.error(err);
});

/**
 * Rename _id to id and remove __v
 * @param {mongoose.Model} model
 */
function modelToObject(model) {
  const obj = model.toObject();
  obj.id = obj._id;
  delete obj._id;
  delete obj.__v;
  return obj;
}

async function mongoConnect() {
  await mongoose.connect(process.env.MONGO_URL);
}

async function mongoDisconnect() {
  await mongoose.disconnect();
}

module.exports = {
  mongoConnect,
  mongoDisconnect,
  modelToObject,
};
