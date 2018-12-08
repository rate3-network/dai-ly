#!/bin/bash

scriptdir=`dirname "$BASH_SOURCE"`;

# Default connection parameters.
PGHOST="localhost" #safety switch
echo "\$PGHOST set to $PGHOST";
PGPORT=${PGPORT:="5432"}
echo "\$PG_PORT set to $PGPORT";


echo "Dropping database";
dropdb ${PGDATABASE};

echo "Creating database";
createdb -O ${PGUSER} ${PGDATABASE}

echo "Running migrations";
migrate -url postgres://${PGUSER}@${PGHOST}:${PGPORT}/${PGDATABASE}?sslmode=disable -path $scriptdir/migrations up;

echo "Adding all mock row entries";
psql -d ${PGDATABASE} -a -f $scriptdir/util/seed.sql;
