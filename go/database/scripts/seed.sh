#!/bin/bash

echo "Adding all mock row entries in database"
scriptdir=`dirname "$BASH_SOURCE"`
psql -d ${PGDATABASE} -a -f $scriptdir/util/seed.sql;
