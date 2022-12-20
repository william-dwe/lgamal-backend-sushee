-- =============================================
-- Author:      William Wibowo Ciptono
-- Create date: 14 Des 2022
-- Description: Initiate restaurant DB tables
-- =============================================

CREATE DATABASE DB_RESTAURANT;

CREATE TABLE IF NOT EXISTS roles (
	id INT PRIMARY KEY,
    role_name VARCHAR,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP NULL
);

CREATE TABLE IF NOT EXISTS users (
	id SERIAL PRIMARY KEY,
	full_name VARCHAR NOT NULL,
	phone VARCHAR UNIQUE check (phone ~ '^[0-9]*$'),
	email VARCHAR UNIQUE NOT NULL,
	username VARCHAR UNIQUE NOT NULL check (username ~ '^[a-z0-9]+'),
	password VARCHAR NOT NULL,
	register_date date NOT NULL,
	profile_picture BYTEA,
	play_attempt INT,
	role_id INT NOT NULL,
	foreign key (role_id) references roles(id),
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP NULL
);

CREATE TABLE IF NOT EXISTS session (
	id SERIAL PRIMARY KEY,
	refresh_token VARCHAR,
	user_id INT,
	foreign key (user_id) references users(id),
	expired_at TIMESTAMP,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP NULL
);

-- CREATE SEQUENCE wallet_id_sequence
--   INCREMENT 1
--   MINVALUE 1
--   MAXVALUE 999000
--   START 100000
--   CACHE 1;

-- CREATE TABLE IF NOT EXISTS wallets(
-- 	id INTEGER PRIMARY KEY DEFAULT NEXTVAL('666666666wallet_id_sequence'),
-- 	user_id INT,
-- 	FOREIGN KEY (user_id) REFERENCES users(id),
-- 	balance NUMERIC,
-- 	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
-- 	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
-- 	deleted_at TIMESTAMP NULL
-- );

-- ALTER SEQUENCE wallet_id_sequence
-- OWNED BY wallets.id;