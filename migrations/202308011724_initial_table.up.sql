CREATE TABLE users (
	id VARCHAR (255) PRIMARY KEY,
	username VARCHAR ( 255 ) UNIQUE NOT NULL,
	image_id VARCHAR ( 255 ) NULL,
	password TEXT NOT NULL,
	created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE images (
	id VARCHAR (255) PRIMARY KEY,
	binary_image BYTEA NOT NULL,
	created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP
);