const mongoose = require("mongoose");

const urlSchema = new mongoose.Schema({
  userID: {
    type: String,
    required: true,
  },
  urlID: {
    type: Number,
    required: true,
  },
  symbol: {
    type: String,
  },
  tags: {
    type: [String],
  },
  url: {
    type: String,
    required: true,
  },
  description: {
    type: String,
  },
  note: {
    type: String,
  },
  createdAt: {
    type: Date,
  },
});

module.exports = mongoose.model("Url", urlSchema);
