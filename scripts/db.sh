#!/bin/sh
echo "*** TEMPORARY SCRIPT BEFORE IMPLEMENTING MIGRATIONS ***"

. ./scripts/values.sh

# values.sh should contain these values (fill in passwords)

# PGUSER=postgres
# PGPASSWORD=
# PGDATABASE=postgres
# DEV_USERNAME=facadeuser
# DEV_PASSWORD=facadepass
# DEV_DATABASE=facade

setup() {
    psql postgres://$PGUSER:$PGPASSWORD@localhost/postgres -c "CREATE USER $DEV_USERNAME WITH PASSWORD '$DEV_PASSWORD'"
    psql postgres://$PGUSER:$PGPASSWORD@localhost/postgres -c "CREATE DATABASE $DEV_DATABASE"
    psql postgres://$PGUSER:$PGPASSWORD@localhost/postgres -c "GRANT ALL PRIVILEGES ON DATABASE $DEV_DATABASE TO $DEV_USERNAME"

    psql postgres://$DEV_USERNAME:$DEV_PASSWORD@localhost/$DEV_DATABASE -f "scripts/setup.sql"
}

teardown() {
    psql postgres://$PGUSER:$PGPASSWORD@localhost/postgres -c "SELECT pg_terminate_backend(pg_stat_activity.pid) FROM pg_stat_activity WHERE pg_stat_activity.datname = '$DEV_DATABASE' AND pid <> pg_backend_pid()"
    psql postgres://$PGUSER:$PGPASSWORD@localhost/postgres -c "DROP DATABASE $DEV_DATABASE"
    psql postgres://$PGUSER:$PGPASSWORD@localhost/postgres -c "DROP USER $DEV_USERNAME"
}

case "$1" in
setup)
    setup
;;
teardown)
    teardown
;;
*)
echo "Usage: $0 (setup|teardown)"
;;
esac

# docker run -d \
#   --name dev-postgres \
#   -e POSTGRES_USER='ADD values.sh PGUSER HERE' \
#   -e POSTGRES_PASSWORD='ADD values.sh PGPASSWORD' \
#   -e POSTGRES_DB=postgres \
#   -p 5432:5432 postgres

# psql postgres://user:pass@localhost/facade