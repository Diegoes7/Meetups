# psql -U postgres -d postgres -c "DROP DATABASE meetup_dev;"
# psql -U postgres -d postgres -c "CREATE DATABASE meetup_dev;"
# psql -U postgres -d meetup_dev -f db/migrations/000001_create_users_table.up.sql

#!/bin/bash

# Connection variables
# PGUSER="meetup_postgresql_user"
# PGPASSWORD="FCEvGwGOZdwDvY6C9xpmAL4bk4ca7RK5"
# PGHOST="dpg-d22h8bu3jp1c738u4uo0-a.oregon-postgres.render.com"
# PGDATABASE="meetup_postgresql"
# PGPORT=5432
# PGSSLMODE="require"

# export PGPASSWORD

# psql "postgresql://${PGUSER}@${PGHOST}:${PGPORT}/${PGDATABASE}?sslmode=${PGSSLMODE}" <<EOF
# DO \$\$
# DECLARE
#     rec RECORD;
# BEGIN
#     FOR rec IN
#         SELECT tablename FROM pg_tables WHERE schemaname = 'public'
#     LOOP
#         EXECUTE format('TRUNCATE TABLE public.%I CASCADE;', rec.tablename);
#     END LOOP;
# END;
# \$\$;
# EOF

# unset PGPASSWORD

#!/bin/bash

CONN="postgresql://meetup_postgresql_user:FCEvGwGOZdwDvY6C9xpmAL4bk4ca7RK5@dpg-d22h8bu3jp1c738u4uo0-a.oregon-postgres.render.com/meetup_postgresql?sslmode=require"

psql "$CONN" -f ./db/drop_all.sql
