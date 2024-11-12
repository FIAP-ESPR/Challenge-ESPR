CREATE USER admin WITH PASSWORD 'adminadmin' SUPERUSER;
CREATE DATABASE ancora OWNER admin;
/* Users Table */

-- DROP TABLE public.users;

CREATE TABLE public.users (
	id serial4 NOT NULL,
	username varchar NOT NULL,
	"password" bytea NOT NULL,
	CONSTRAINT users_pk PRIMARY KEY (id)
);

/* TODO */