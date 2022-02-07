const http = require("http");
const mongoose = require("mongoose");

require("dotenv").config();

const app = require("./app");

const PORT = process.env.PORT || 3000;

mongoose.connection.on("error", (err) => {
  console.error(err);
});

async function startServer() {
  await mongoose.connect(process.env.MONGO_URL);
  const server = http.createServer(app);
  server.listen(PORT, () => {
    console.log(`Listening on port ${PORT}...`);
  });
}

startServer();
