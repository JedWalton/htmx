#!/bin/bash
PGPASSWORD=mysecretpassword psql -h localhost -p 5432 -U postgres -f outputfile.sql
