SET search_path TO pulse, public;

DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1
        FROM pg_type
        WHERE typname = 'monitor_status'
    ) THEN
        CREATE TYPE monitor_status AS ENUM ('up', 'down');
    END IF;
END
$$;
