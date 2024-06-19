ifneq (,$(wildcard ./.env))
    include .env
    export
endif

align:
	fieldalignment -fix ./...

goose-init:
	export GOOSE_DRIVER=$(DATABASE_TYPE)
	mkdir -p migrations/$(DATABASE_TYPE)
	export GOOSE_DBSTRING=$(DATABASE_CONNECTION_STRING)
	export GOOSE_MIGRATION_DIR=migrations/$(DATABASE_TYPE)
	goose create create_user_table sql

goose-up :
	goose $(DATABASE_TYPE) up

