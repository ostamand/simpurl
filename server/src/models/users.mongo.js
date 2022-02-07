const mongoose = require("mongoose")

const userSchema = mongoose.Schema({
    email: {
        type: String,
        required: false
    },
    username: {
        type: String,
        required: false
    },
    hashedPassword: {
        type: String,
        required: false
    }
})

module.exports = mongoose.model("User", userSchema)