-- +goose Up
CREATE TABLE users (
	id uuid NOT NULL,
	created_at timestamp NOT NULL,
	updated_at timestamp NOT NULL,
	name varchar NOT NULL,
	CONSTRAINT newtable_pk PRIMARY KEY (id)
);

-- +goose Down
DROP TABLE users;