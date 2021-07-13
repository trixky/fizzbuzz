#!/bin/bash

# server / postgres
export POSTGRES_USER=trixky
export POSTGRES_PASSWORD=1234
export POSTGRES_DB=fizzbuzz
export POSTGRES_HOST=postgres
export POSTGRES_PORT=5432

# pgadmin
export PGADMIN_DEFAULT_EMAIL=trixky@fizz.buzz
export PGADMIN_DEFAULT_PASSWORD=1234

# test / postgres
export PGUSER=$POSTGRES_USER
export PGPASSWORD=$POSTGRES_PASSWORD
export PGDATABASE=$POSTGRES_DB

# test
export PGHOST=$POSTGRES_HOST
export PGPORT=$POSTGRES_PORT