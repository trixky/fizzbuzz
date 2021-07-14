const { Client } = require("pg");
const redis = require("redis");
const supertest = require("supertest");

const get_cookies = (res) => {
  const header_set_cookie = res.headers["set-cookie"];
  return header_set_cookie
    ? header_set_cookie[0].split(",").map((item) => item.split(";")[0])
    : [];
};

const check_stats = (stats1, stats2) => {
  return (
    stats1.length === stats2.length &&
    stats1.requests.every(
      (request, index) =>
        request.int1 === stats2.requests[index].int1 &&
        request.int2 === stats2.requests[index].int2 &&
        request.limit === stats2.requests[index].limit &&
        request.str1 === stats2.requests[index].str1 &&
        request.str2 === stats2.requests[index].str2
    )
  );
};

const expected_stats = {
  requests: [
    {
      int1: 3,
      int2: 5,
      limit: 30,
      str1: "fizz",
      str2: "buzz",
      request_number: 70,
    },
    {
      int1: 5,
      int2: 10,
      limit: 15,
      str1: "pop",
      str2: "corn",
      request_number: 30,
    },
  ],
};

beforeAll(function (done) {
  (async () => {
    // --------------------------------------------------- INIT POSTGRES
    // https://node-postgres.com/features/connecting
    try {
      const client_postgres = new Client();
      await client_postgres.connect();
      await client_postgres.query(
        "DELETE FROM api_users WHERE pseudo <> 'admin'"
      );
      await client_postgres.query(
        "DELETE FROM fizzbuzz_request_statistics"
      );
      await client_postgres.end();
    } catch (err) {
      return console.error("INIT > POSTGRES > FAILED (" + err + ")");
    }

    // --------------------------------------------------- INIT REDIS
    // https://www.npmjs.com/package/redis
    try {
      const client_redis = redis.createClient({ host: "redis", port: 6379 });
      client_redis.on("error", (err) => {
        throw "INIT > REDIS > FAILED (" + err + ")";
      });
      client_redis.on("ready", () => {
        client_redis.flushall((err) => {
          if (err) {
            throw "INIT > REDIS > FAILED (" + err + ")";
          } else {
            client_redis.end(true);
          }
        });
      });
    } catch (err) {
      return console.error(err);
    }

    done();
  })();
});

// ===================================== API TEST
describe("API test", function () {
  const request = supertest("server:8080");
  let users_session_cookies = {
    admin: [],
    chat: [],
    chien: [],
  };

  // ===================================== REGISTER
  describe("/register", function () {
    it("register @chat:abcdABCD1234@", (done) => {
      request
        .post("/register")
        .set("Content-type", "application/json")
        .send({
          pseudo: "chat",
          password: "abcdABCD1234@",
        })
        .expect(200)
        .end((err) => {
          try {
            expect(err).toBe(null);
            done();
          } catch (err) {
            done(err);
          }
        });
    });
    it("register @chien:abcdABCD1234@", (done) => {
      request
        .post("/register")
        .set("Content-type", "application/json")
        .send({
          pseudo: "chien",
          password: "abcdABCD1234@",
        })
        .expect(200)
        .end((err) => {
          try {
            expect(err).toBe(null);
            done();
          } catch (err) {
            done(err);
          }
        });
    });
    it("register @chien:abcdABCD1234@ #must_fail", (done) => {
      request
        .post("/register")
        .set("Content-type", "application/json")
        .send({
          pseudo: "chien",
          password: "abcdABCD1234@",
        })
        .expect(409)
        .end((err) => {
          try {
            expect(err).toBe(null);
            done();
          } catch (err) {
            done(err);
          }
        });
    });
  });
  // ===================================== LOGIN
  describe("/login", function () {
    it("login @?:? #must_fail", (done) => {
      request
        .post("/login")
        .set("Content-type", "application/json")
        .send({})
        .expect(401)
        .end((err, res) => {
          try {
            expect(err).toBe(null);
            users_session_cookies.admin = get_cookies(res);
            done();
          } catch (err) {
            done(err);
          }
        });
    });
    it("login @?:abcdABCD1234@ #must_fail", (done) => {
      request
        .post("/login")
        .set("Content-type", "application/json")
        .send({
          password: "abcdABCD1234@",
        })
        .expect(401)
        .end((err, res) => {
          try {
            expect(err).toBe(null);
            users_session_cookies.admin = get_cookies(res);
            done();
          } catch (err) {
            done(err);
          }
        });
    });
    it("login @admin:? #must_fail", (done) => {
      request
        .post("/login")
        .set("Content-type", "application/json")
        .send({
          pseudo: "admin",
        })
        .expect(401)
        .end((err, res) => {
          try {
            expect(err).toBe(null);
            users_session_cookies.admin = get_cookies(res);
            done();
          } catch (err) {
            done(err);
          }
        });
    });
    it("login @admin:bad_password #must_fail", (done) => {
      request
        .post("/login")
        .set("Content-type", "application/json")
        .send({
          pseudo: "admin",
          password: "bad_passwordz",
        })
        .expect(401)
        .end((err, res) => {
          try {
            expect(err).toBe(null);
            users_session_cookies.admin = get_cookies(res);
            done();
          } catch (err) {
            done(err);
          }
        });
    });
    it("login @admin:abcdABCD1234@", (done) => {
      request
        .post("/login")
        .set("Content-type", "application/json")
        .send({
          pseudo: "admin",
          password: "abcdABCD1234@",
        })
        .expect(200)
        .end((err, res) => {
          try {
            expect(err).toBe(null);
            users_session_cookies.admin = get_cookies(res);
            done();
          } catch (err) {
            done(err);
          }
        });
    });
    it("login @chat:abcdABCD1234@", (done) => {
      request
        .post("/login")
        .set("Content-type", "application/json")
        .send({
          pseudo: "chat",
          password: "abcdABCD1234@",
        })
        .expect(200)
        .end((err, res) => {
          try {
            expect(err).toBe(null);
            users_session_cookies.chat = get_cookies(res);
            done();
          } catch (err) {
            done(err);
          }
        });
    });
    it("login @chien:abcdABCD1234@", (done) => {
      request
        .post("/login")
        .set("Content-type", "application/json")
        .send({
          pseudo: "chien",
          password: "abcdABCD1234@",
        })
        .expect(200)
        .end((err, res) => {
          try {
            expect(err).toBe(null);
            users_session_cookies.chien = get_cookies(res);
            done();
          } catch (err) {
            done(err);
          }
        });
    });
  });
  // ===================================== FIZZBUZZ
  describe("/fizzbuzz", function () {
    it("fizzbuzz @? #must_fail", (done) => {
      request
        .get("/fizzbuzz")
        .query({
          int1: "3",
          int2: "5",
          limit: "30",
          str1: "fizz",
          str2: "buzz",
        })
        .expect(401)
        .end((err) => {
          try {
            expect(err).toBe(null);
            done();
          } catch (err) {
            done(err);
          }
        });
    });
    it("fizzbuzz @admin", (done) => {
      request
        .get("/fizzbuzz")
        .query({
          int1: "3",
          int2: "5",
          limit: "30",
          str1: "fizz",
          str2: "buzz",
        })
        .set("Cookie", users_session_cookies.admin)
        .expect(200)
        .end((err) => {
          try {
            expect(err).toBe(null);
            done();
          } catch (err) {
            done(err);
          }
        });
    });
    it("fizzbuzz @chat", (done) => {
      request
        .get("/fizzbuzz")
        .query({
          int1: "3",
          int2: "5",
          limit: "30",
          str1: "fizz",
          str2: "buzz",
        })
        .set("Cookie", users_session_cookies.chat)
        .expect(200)
        .end((err) => {
          try {
            expect(err).toBe(null);
            done();
          } catch (err) {
            done(err);
          }
        });
    });
    it("fizzbuzz @chien", (done) => {
      request
        .get("/fizzbuzz")
        .query({
          int1: "5",
          int2: "10",
          limit: "15",
          str1: "pop",
          str2: "corn",
        })
        .set("Cookie", users_session_cookies.chien)
        .expect(200)
        .end((err) => {
          try {
            expect(err).toBe(null);
            done();
          } catch (err) {
            done(err);
          }
        });
    });
  });
  // ===================================== STATS
  describe("/stats", function () {
    it("stats @? #must_fail", (done) => {
      request
        .get("/stats")
        .expect(401)
        .end((err) => {
          try {
            expect(err).toBe(null);
            done();
          } catch (err) {
            done(err);
          }
        });
    });
    it("stats @admin", (done) => {
      request
        .get("/stats")
        .set("Cookie", users_session_cookies.admin)
        .expect(200)
        .end((err, res) => {
          try {
            expect(check_stats(expected_stats, res.body)).toBe(true);
            expect(err).toBe(null);
            done();
          } catch (err) {
            done(err);
          }
        });
    });
    it("stats @chat", (done) => {
      request
        .get("/stats")
        .set("Cookie", users_session_cookies.chat)
        .expect(200)
        .end((err, res) => {
          try {
            expect(check_stats(expected_stats, res.body)).toBe(true);
            expect(err).toBe(null);
            done();
          } catch (err) {
            done(err);
          }
        });
    });
    it("stats @chien", (done) => {
      request
        .get("/stats")
        .set("Cookie", users_session_cookies.chien)
        .expect(200)
        .end((err, res) => {
          try {
            expect(check_stats(expected_stats, res.body)).toBe(true);
            expect(err).toBe(null);
            done();
          } catch (err) {
            done(err);
          }
        });
    });
  });
  // ===================================== BLOCK
  describe("/block", function () {
    it("block @? > @chien [true] #must_fail", (done) => {
      request
        .patch("/block")
        .set("Content-type", "application/json")
        .send({
          pseudo: "chien",
          block_status: "true",
        })
        .expect(401)
        .end((err) => {
          try {
            expect(err).toBe(null);
            done();
          } catch (err) {
            done(err);
          }
        });
    });
    it("block @chat > @chien [true] #must_fail", (done) => {
      request
        .patch("/block")
        .set("Cookie", users_session_cookies.chat)
        .set("Content-type", "application/json")
        .send({
          pseudo: "chien",
          block_status: "true",
        })
        .expect(401)
        .end((err) => {
          try {
            expect(err).toBe(null);
            done();
          } catch (err) {
            done(err);
          }
        });
    });
    it("block @admin > @chien [true]", (done) => {
      request
        .patch("/block")
        .set("Cookie", users_session_cookies.admin)
        .set("Content-type", "application/json")
        .send({
          pseudo: "chien",
          block_status: "true",
        })
        .expect(200)
        .end((err) => {
          try {
            expect(err).toBe(null);
            done();
          } catch (err) {
            done(err);
          }
        });
    });
    it("fizzbuzz @chien #must_fail", (done) => {
      request
        .get("/fizzbuzz")
        .query({
          int1: "5",
          int2: "10",
          limit: "15",
          str1: "pop",
          str2: "corn",
        })
        .set("Cookie", users_session_cookies.chien)
        .expect(401)
        .end((err) => {
          try {
            expect(err).toBe(null);
            done();
          } catch (err) {
            done(err);
          }
        });
    });
    it("login @chien:abcdABCD1234@ #must_fail", (done) => {
      request
        .post("/login")
        .set("Content-type", "application/json")
        .send({
          pseudo: "chien",
          password: "abcdABCD1234@",
        })
        .expect(401)
        .end((err, res) => {
          try {
            expect(err).toBe(null);
            users_session_cookies.chien = get_cookies(res);
            done();
          } catch (err) {
            done(err);
          }
        });
    });
    it("block @admin > @chien [false]", (done) => {
      request
        .patch("/block")
        .set("Cookie", users_session_cookies.admin)
        .set("Content-type", "application/json")
        .send({
          pseudo: "chien",
          block_status: "false",
        })
        .expect(200)
        .end((err) => {
          try {
            expect(err).toBe(null);
            done();
          } catch (err) {
            done(err);
          }
        });
    });
    it("login @chien:abcdABCD1234@", (done) => {
      request
        .post("/login")
        .set("Content-type", "application/json")
        .send({
          pseudo: "chien",
          password: "abcdABCD1234@",
        })
        .expect(200)
        .end((err, res) => {
          try {
            expect(err).toBe(null);
            users_session_cookies.chien = get_cookies(res);
            done();
          } catch (err) {
            done(err);
          }
        });
    });
  });
});
