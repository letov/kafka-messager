-- +goose Up
CREATE TABLE ban_words (
		id SERIAL PRIMARY KEY,
		word varchar NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		CONSTRAINT "UNQ__ban_words__word" UNIQUE(word),
		CHECK (word = lower(word))
);

-- +goose Down
DROP TABLE ban_words
