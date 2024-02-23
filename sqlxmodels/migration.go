package sqlxmodels

const MIGRATION = `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255),
		email VARCHAR(255) UNIQUE NOT NULL
	);
	CREATE TABLE IF NOT EXISTS posts (
		id SERIAL PRIMARY KEY,
		user_id INTEGER REFERENCES users(id),
		title VARCHAR(255),
		content TEXT
	);
	CREATE TABLE IF NOT EXISTS comments (
		id SERIAL PRIMARY KEY,
		post_id INTEGER REFERENCES posts(id),
		text TEXT
	);
`
