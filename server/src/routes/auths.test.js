const request = require("supertest");

const app = require("../app");
const { mongoConnect, mongoDisconnect } = require("../services/mongo");
const { faker } = require("@faker-js/faker");
const supertest = require("supertest");

describe("Launches API", () => {
  beforeAll(async () => {
    await mongoConnect();
  });

  afterAll(async () => {
    await mongoDisconnect();
  });

  // .set('Cookie', ['myApp-token=12345667', 'myApp-other=blah'])

  describe("Test POST /signup", () => {
    test("It should respond with 201 created", async () => {
      const validUser = {
        email: faker.internet.email(),
        username: faker.word.noun(),
        password: faker.word.adjective(),
      };

      const response = await request(app)
        .post("/signup")
        .send(validUser)
        .expect(201);
      const data = response.body;
      expect("id" in data).toBeTruthy();
    });
  });

  describe("Test /login", () => {
    test("It should return session cookie", async () => {
      const validUser = {
        email: faker.internet.email(),
        username: faker.word.noun(),
        password: faker.word.adjective(),
      };

      const agent = request.agent(app);

      await agent.post("/signup").send(validUser).expect(201);

      const response = await agent
        .post("/signin")
        .send({ username: validUser.username, password: validUser.password })
        .expect(200);

      console.log(response.headers);

      await agent.get("/ping").expect(200);
    });
  });
});
