/* Create ancora DB with admin as owner */
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_database WHERE datname = 'ancora') THEN
        CREATE DATABASE ancora OWNER admin;
    END IF;
END $$;