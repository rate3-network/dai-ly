#!/bin/bash

echo "Deleting all row entries in database"
scriptdir=`dirname "$BASH_SOURCE"`
psql -d ${PGDATABASE} -a -f $scriptdir/util/truncate.sql;
