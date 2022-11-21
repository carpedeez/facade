#!/bin/sh
echo "*** TEMPORARY SCRIPT BEFORE IMPLEMENTING MIGRATIONS ***"

. ./scripts/values.sh

# values.sh should contain these values

# PGUSER=postgres
# PGPASSWORD=p4ssword
# PGDATABASE=postgres
# DEV_USERNAME=facadeuser
# DEV_PASSWORD=facadepass
# DEV_DATABASE=facade

up() {
    psql postgres://$PGUSER:$PGPASSWORD@localhost/postgres -c "CREATE USER $DEV_USERNAME WITH PASSWORD '$DEV_PASSWORD'"
    psql postgres://$PGUSER:$PGPASSWORD@localhost/postgres -c "CREATE DATABASE $DEV_DATABASE"
    psql postgres://$PGUSER:$PGPASSWORD@localhost/postgres -c "GRANT ALL PRIVILEGES ON DATABASE $DEV_DATABASE TO $DEV_USERNAME"

    psql postgres://$DEV_USERNAME:$DEV_PASSWORD@localhost/$DEV_DATABASE -f "scripts/setup.sql"
}

down() {
    psql postgres://$PGUSER:$PGPASSWORD@localhost/postgres -c "SELECT pg_terminate_backend(pg_stat_activity.pid) FROM pg_stat_activity WHERE pg_stat_activity.datname = '$DEV_DATABASE' AND pid <> pg_backend_pid()"
    psql postgres://$PGUSER:$PGPASSWORD@localhost/postgres -c "DROP DATABASE $DEV_DATABASE"
    psql postgres://$PGUSER:$PGPASSWORD@localhost/postgres -c "DROP USER $DEV_USERNAME"
}

case "$1" in
up)
    up
;;
down)
    down
;;
*)
echo "Usage: $0 (up|down)"
;;
esac

# psql postgres://user:pass@localhost/facade