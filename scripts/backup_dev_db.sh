#!/bin/bash

# Your PostgreSQL URL
# POSTGRESQL_URL="postgres://postgres:mysecretpassword@localhost:5432/devdb?sslmode=disable"
#
export PGHOST=localhost
export PGPORT=5432
export PGUSER=postgres
export PGPASSWORD=mysecretpassword


# Running pg_dumpall
pg_dumpall > outputfile.sql

