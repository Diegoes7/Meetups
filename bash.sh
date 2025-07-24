#!/bin/bash
"/c/Program Files/PostgreSQL/17/bin/createdb" -U postgres meetup_dev

# set for default user
export PGUSER=postgres

#set the password not to prompt each time
export PGPASSWORD="victoria7"

# set the right path so type just the command not the whole path
export PATH="$PATH:/c/Program Files/PostgreSQL/17/bin"

# gqlgen generate regenerate the schema

# prompt to migrate the database, if the table not exist if gave error relation already have
psql -U postgres -d meetup_dev -f db/migrations/000003_meetup_invitation_table.up.sql

#$ create image foem which will make container
# docker build -t meetup-app .

#& start the application
# docker compose up -d


#! Run the docker container command
# docker run -it -p 8080:8080 \
#   -e POSTGRES_HOST=host.docker.internal \
#   -e POSTGRES_PORT=5432 \
#   -e POSTGRES_USER=postgres \
#   -e POSTGRES_PASSWORD=postgres \
#   -e POSTGRES_DB=meetup_dev \
#   -e REDIS_HOST=host.docker.internal \
#   -e REDIS_PORT=6379 \
#   --name meetup-app \
#   meetup-app

# export POSTGRESQL_URL=postgresql://postgres:victoria7@localhost:5432/meetup_dev?sslmode=disable
# export JWT_TOKEN=asecret