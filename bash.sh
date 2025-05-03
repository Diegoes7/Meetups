#!/bin/bash
"/c/Program Files/PostgreSQL/17/bin/createdb" -U postgres meetup_dev

# set for default user
export PGUSER=postgres

#set the password not to prompt each time
export PGPASSWORD="victoria7"

# set the right path so type just the command not the whole path
export PATH="$PATH:/c/Program Files/PostgreSQL/17/bin"

# gqlgen generate regenerate the schema

#prompt to migrate the database, if the table not exist if gave error relation already have
psql -U postgres -d meetup_dev -f db/migrations/000001_create_users_table.up.sql