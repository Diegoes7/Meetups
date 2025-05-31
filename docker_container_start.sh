#!/bin/bash

# Load environment variables from .env if it exists
if [ -f .env ]; then
  set -o allexport
  source <(grep -v '^#' .env | xargs -d '\n')
  set +o allexport
else
  echo ".env file not found. Using default values."
fi

# Default values (fallbacks if not set in .env)
POSTGRES_HOST=${POSTGRES_HOST:-host.docker.internal}
POSTGRES_PORT=${POSTGRES_PORT:-5432}
POSTGRES_USER=${POSTGRES_USER:-postgres}
POSTGRES_PASSWORD=${POSTGRES_PASSWORD:-postgres}
POSTGRES_DB=${POSTGRES_DB:-meetup_dev}
REDIS_HOST=${REDIS_HOST:-host.docker.internal}
REDIS_PORT=${REDIS_PORT:-6379}
CONTAINER_NAME=${CONTAINER_NAME:-meetup-app}
IMAGE_NAME=${IMAGE_NAME:-meetup-app}

echo "Starting container: $CONTAINER_NAME"

docker run -it --rm \
  -p 8080:8080 \
  -e POSTGRES_HOST=$POSTGRES_HOST \
  -e POSTGRES_PORT=$POSTGRES_PORT \
  -e POSTGRES_USER=$POSTGRES_USER \
  -e POSTGRES_PASSWORD=$POSTGRES_PASSWORD \
  -e POSTGRES_DB=$POSTGRES_DB \
  -e REDIS_HOST=$REDIS_HOST \
  -e REDIS_PORT=$REDIS_PORT \
  --name $CONTAINER_NAME \
  $IMAGE_NAME
echo "Container started successfully."