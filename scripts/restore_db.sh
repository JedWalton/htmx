#!/bin/bash
# PGPASSWORD=mysecretpassword psql -h localhost -p 5432 -U postgres -f outputfile.sql
# Setting environment variables
export PGHOST=localhost
export PGPORT=5432
export PGUSER=postgres
export PGPASSWORD=mysecretpassword

# Executing psql with the SQL file
psql -f outputfile.sql

