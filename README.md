# fizzbuzz

A small fizzbuzz API implementation using [go](https://golang.org/), [postgresql](https://www.postgresql.org/), [redis](https://redis.io/), [pgadmin](https://www.pgadmin.org/), and [jest](https://jestjs.io/)/[supertest](https://www.npmjs.com/package/supertest) tester. **(docker-compose)**

## Usage

```bash
source env.sh
docker-compose up -d
```

## Api

The API is accessible in localhost on port 8080.
You will find the [postam](https://www.postman.com/) collection from the `fizzbuzz.postman_collection.json` file.
You will also find the API documentation [here](https://github.com/trixky/fizzbuzz/blob/main/server/README.md).

## Database

```bash
docker-compose exec postgres psql # to access the postgres cli (psql)
docker-compose exec redis sh -c 'redis-cli -h redis' # to access the redis cli
```

## Test

![screenshot](https://raw.githubusercontent.com/trixky/volte_face/master/demo/screenshot.png)

```bash
docker-compose run test
```

> The test erase all data on postgresql and redis!

## Stack

- Go
- Node.js / Jest / Supertest
- Postgresql
- Redis
