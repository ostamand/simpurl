const http = require("http")
const mongoose = require('mongoose');

require("dotenv").config()

const app = require("./app")
const Url = require("./models/urls.mongo")

const PORT = process.env.PORT || 3000;

mongoose.connection.on("error", (err) => {
    console.error(err)
})

async function startServer() {
    // test saving stuff to the db
    //console.log(process.env)
    await mongoose.connect(process.env.MONGO_URL);

    const newUrl = new Url({path: "https://www.google.com"})
    newUrl.save().then(() => {
        console.log("URL saved")
    })
    
    // start server
    const server = http.createServer(app)
    server.listen(PORT, () => {
        console.log(`Listening on port ${PORT}...`)
    })
}

startServer()










