package database

import (
	"database/sql"
	"log/slog"
)

func Open() (*sql.DB, error) {
	DB, err := sql.Open("", "")

	if err != nil {
		return nil, err
	}

	err = DB.Ping()
	if err != nil {
		return nil, err
	}

	_, err = DB.Exec(`CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
	);`)

	if err != nil {
		return nil, err
	}

	slog.Info("Users table created successfully!")

	_, err = DB.Exec(`CREATE TABLE websites (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    url TEXT NOT NULL,
    interval_seconds INT DEFAULT 60,
    created_at TIMESTAMP DEFAULT NOW()
	);`)

	if err != nil {
		return nil, err
	}

	slog.Info("Websites table created successfully!")

	_, err = DB.Exec(`CREATE TABLE health_checks (
    id BIGSERIAL PRIMARY KEY,
    website_id BIGINT NOT NULL REFERENCES websites(id) ON DELETE CASCADE,
    status BOOLEAN NOT NULL,
    response_time INT,
    status_code INT,
    checked_at TIMESTAMP DEFAULT NOW()
    );`)
	if err != nil {
		return nil, err
	}

	slog.Info("health_checks table created successfully!")

	_, err = DB.Exec(`CREATE TABLE alerts (
    id BIGSERIAL PRIMARY KEY,
    website_id BIGINT NOT NULL REFERENCES websites(id) ON DELETE CASCADE,
    type VARCHAR(50) NOT NULL,
    sent_at TIMESTAMP DEFAULT NOW()
    );`)

	if err != nil {
		return nil, err
	}

	slog.Info("alerts table created successfully!")

	return DB, nil
}
