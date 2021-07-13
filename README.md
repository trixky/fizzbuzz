# fizzbuzz

A small fizzbuzz API implementation using [go](https://golang.org/), [postgresql](https://www.postgresql.org/), [redis](https://redis.io/), [pgadmin](https://www.pgadmin.org/), and [jest](https://jestjs.io/)/[supertest](https://www.npmjs.com/package/supertest) tester. **(docker-compose)**

## Usage

```bash
source env.sh
docker-compose up -d
```

## Api

The API is accessible on [localhost:8080/](http://localhost:8080/).
You will find the [postam](https://www.postman.com/) collection from the `fizzbuzz.postman_collection.json` file.
You will also find the API documentation [here](https://github.com/trixky/fizzbuzz/blob/main/server/README.md).

## Database

```bash
docker-compose exec postgres psql # to access the postgres cli (psql)
docker-compose exec redis sh -c 'redis-cli -h redis' # to access the redis cli
```

Pgadmin is accessible on [localhost:5050/](http://localhost:8080/).

![screenshot](https://raw.githubusercontent.com/trixky/fizzbuzz/main/demo/pgadmin_login.png)
![screenshot](https://raw.githubusercontent.com/trixky/fizzbuzz/main/demo/pgadmin_connection.png)

## Test

![screenshot](https://raw.githubusercontent.com/trixky/fizzbuzz/main/demo/test.png)

```bash
docker-compose run test
```

> The test erase all data on postgresql and redis!

## Stack

- Go
- Node.js / Jest / Supertest
- Postgresql
- Redis
