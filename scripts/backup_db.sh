#!/bin/bash
PGPASSWORD=mysecretpassword pg_dumpall -h localhost -p 5432 -U postgres > outputfile.sql
