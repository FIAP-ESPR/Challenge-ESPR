DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_roles WHERE rolname = 'admin') THEN
        CREATE USER admin WITH PASSWORD 'adminadmin' SUPERUSER;
    END IF;
    IF NOT EXISTS (SELECT 1 FROM pg_database WHERE datname = 'ancora') THEN
        CREATE DATABASE ancora OWNER admin;
    END IF;
END $$;

/* TODO */
