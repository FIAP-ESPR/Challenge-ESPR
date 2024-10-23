/* Create Admin user with Super User Privileges */
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_roles WHERE rolname = 'admin') THEN
        CREATE USER admin WITH PASSWORD 'adminadmin' SUPERUSER;
    END IF;
END $$;