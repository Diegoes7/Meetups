1. Stop and remove all containers and volumes, if want to start fresh:

    docker compose down --volumes --remove-orphans

2. Build images:

    docker compose build

3. Start containers and trigger seed.sql: from the root directory of the project

    docker compose up -d

    This app is live on this url: https://meetups-y7ke.onrender.com/
