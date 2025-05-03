psql -U postgres -d postgres -c "DROP DATABASE meetup_dev;"
psql -U postgres -d postgres -c "CREATE DATABASE meetup_dev;"
psql -U postgres -d meetup_dev -f db/migrations/000001_create_users_table.up.sql
