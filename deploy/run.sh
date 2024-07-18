#!/bin/bash

DB_CHANGELOG_PATH="../migrations/liquibase/migrations.xml"

run() {
  echo "Run application"
  go run ../cmd/main.go
}

compose(){
  docker compose --file docker-compose.yml --env-file .env up -d
}

update_db() {
  liquibase \
    --driver="org.postgresql.Driver" \
    --url=jdbc:postgresql://localhost:5555/postgres \
    --username=postgres \
    --password=postgres \
    --changeLogFile="$DB_CHANGELOG_PATH" \
    --liquibaseSchemaName=public \
    --defaultSchemaName=public \
    --logLevel=debug \
    update
}

rollback_db() {
  liquibase \
    --driver="org.postgresql.Driver" \
    --url=jdbc:postgresql://localhost:5555/postgres \
    --username=postgres \
    --password=postgres \
    --changeLogFile="$DB_CHANGELOG_PATH" \
    --liquibaseSchemaName=public \
    --defaultSchemaName=public \
    --logLevel=debug \
    rollbackToDate 2022-07-01
}

# Main script logic
case "$1" in
run)
  run
  ;;
compose)
  compose
  ;;
update_db)
  update_db
  ;;
rollback_db)
  rollback_db
  ;;
*)
  echo "Invalid command."
  exit 1
  ;;
esac