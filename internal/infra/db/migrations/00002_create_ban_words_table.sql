-- +goose Up
CREATE TABLE ban_words (
		id SERIAL PRIMARY KEY,
		word varchar NOT NULL,
		CONSTRAINT "UNQ__ban_words__word" UNIQUE(word),
		CHECK (word = lower(word))
);

-- +goose Down
DROP TABLE ban_words
