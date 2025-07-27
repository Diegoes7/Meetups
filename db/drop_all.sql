-- db/drop_all.sql

DO $$
DECLARE
    stmt TEXT;
BEGIN
    -- Drop foreign keys
    FOR stmt IN
        SELECT 'ALTER TABLE "' || tab.relname || '" DROP CONSTRAINT "' || con.conname || '";'
        FROM pg_constraint con
        JOIN pg_class tab ON con.conrelid = tab.oid
        JOIN pg_namespace ns ON ns.oid = tab.relnamespace
        WHERE ns.nspname = 'public' AND con.contype = 'f'
    LOOP
        EXECUTE stmt;
    END LOOP;

    -- Drop tables
    FOR stmt IN
        SELECT 'DROP TABLE IF EXISTS "' || tablename || '" CASCADE;'
        FROM pg_tables
        WHERE schemaname = 'public'
    LOOP
        EXECUTE stmt;
    END LOOP;

    -- Drop sequences
    FOR stmt IN
        SELECT 'DROP SEQUENCE IF EXISTS "' || sequence_name || '" CASCADE;'
        FROM information_schema.sequences
        WHERE sequence_schema = 'public'
    LOOP
        EXECUTE stmt;
    END LOOP;
END$$;
