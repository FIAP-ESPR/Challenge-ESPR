CREATE USER admin WITH PASSWORD 'adminadmin' SUPERUSER;
CREATE DATABASE ancora OWNER admin;
/* Users Table */

-- public.mecanica_empresa definição

-- Drop table

-- DROP TABLE public.mecanica_empresa;

CREATE TABLE public.mecanica_empresa (
	id serial4 NOT NULL,
	nome varchar NOT NULL,
	cnpj varchar NOT NULL,
	CONSTRAINT mecanica_empresa_pk PRIMARY KEY (id),
	CONSTRAINT mecanica_empresa_unique UNIQUE (cnpj)
);


-- public.users definição

-- Drop table

-- DROP TABLE public.users;

CREATE TABLE public.users (
	id serial4 NOT NULL,
	username varchar NOT NULL,
	"password" bytea NOT NULL,
	CONSTRAINT users_pk PRIMARY KEY (id)
);


-- public.cliente definição

-- Drop table

-- DROP TABLE public.cliente;

CREATE TABLE public.cliente (
	id serial4 NOT NULL,
	nome varchar NOT NULL,
	documento varchar NOT NULL,
	telefone varchar NULL,
	email varchar NULL,
	CONSTRAINT cliente_pk PRIMARY KEY (id),
	CONSTRAINT cliente_users_fk FOREIGN KEY (id) REFERENCES public.users(id)
);


-- public.mecanico definição

-- Drop table

-- DROP TABLE public.mecanico;

CREATE TABLE public.mecanico (
	id serial4 NOT NULL,
	nome varchar NULL,
	documento varchar NULL,
	mecanica_empresa serial4 NOT NULL,
	CONSTRAINT mecanico_pk PRIMARY KEY (id),
	CONSTRAINT mecanico_mecanica_empresa_fk FOREIGN KEY (mecanica_empresa) REFERENCES public.mecanica_empresa(id),
	CONSTRAINT users_fk FOREIGN KEY (id) REFERENCES public.users(id)
);


-- public.carro definição

-- Drop table

-- DROP TABLE public.carro;

CREATE TABLE public.carro (
	dono serial4 NOT NULL,
	mecanico serial4 NOT NULL,
	chassi varchar NULL,
	modelo varchar NULL,
	placa varchar NULL,
	CONSTRAINT carro_pk PRIMARY KEY (dono),
	CONSTRAINT carro_cliente_fk FOREIGN KEY (dono) REFERENCES public.cliente(id),
	CONSTRAINT carro_mecanico_fk FOREIGN KEY (mecanico) REFERENCES public.mecanico(id)
);

/* TODO */