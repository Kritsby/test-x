#!/bin/sh
# wait-for-postgres.sh

set -e

until PGPASSWORD=$PGPASSWORD psql -h $PGHOST -U $PGUSER -d $PGDB -c '\q'; do
  >&2 echo "Postgres is unavailable - sleeping"
  sleep 1
done

>&2 echo "Postgres is up - executing command"
exec ./api

