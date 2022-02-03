const request = require("supertest");

const app = require("../app");
const { mongoConnect, mongoDisconnect } = require("../services/mongo");
const { faker } = require("@faker-js/faker");

describe("Launches API", () => {
  beforeAll(async () => {
    await mongoConnect();
  });

  afterAll(async () => {
    await mongoDisconnect();
  });

  const validUser = {
    email: faker.internet.email(),
    username: faker.word.noun(),
    password: faker.word.adjective(),
  };

  // .set('Cookie', ['myApp-token=12345667', 'myApp-other=blah'])

  describe("Test POST /signup", () => {
    test("It should respond with 201 created", async () => {
      const response = await request(app)
        .post("/signup")
        .send(validUser)
        .expect(201);
      const data = response.body;
      expect("_id" in data).toBeTruthy();
      expect(data.hashedPassword).toBeTruthy();
    });
  });

  describe("Test /login", () => {
    test("It should return session cookie", async () => {
      const response = await request(app)
        .post("/login")
        .send({ username: validUser.username, password: validUser.password })
        .expect(200);
    });
  });
});
