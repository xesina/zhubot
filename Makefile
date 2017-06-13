#@IgnoreInspection BashAddShebang
export ROOT=$(realpath $(dir $(lastword $(MAKEFILE_LIST))))
export GOOS=linux
export ENV=development
export SCHEMA=zhubot
export DEBUG=1

#########
# Migrate
#########

check-migrate:
	which migrate || go get -u -v git.vada.ir/sdp/migrate

migrate-create: check-migrate
	migrate create --name=$(NAME)

migrate-up: check-migrate
	migrate up

migrate-rollback: check-migrate
	migrate rollback

migrate-reset: check-migrate
	migrate reset

migrate-refresh: check-migrate
	migrate refresh
