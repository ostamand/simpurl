const mongoose = require("mongoose")

const urlSchema = new mongoose.Schema({
    path: {
        type: String,
        required: true,
    }
})

module.exports = mongoose.model("Url", urlSchema)