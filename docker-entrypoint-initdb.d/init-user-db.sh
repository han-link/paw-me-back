#!/usr/bin/env bash
set -e

echo "Init supertokens database"

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" <<-EOSQL
	CREATE DATABASE supertokens;
	GRANT ALL PRIVILEGES ON DATABASE supertokens TO $POSTGRES_USER;
EOSQL
