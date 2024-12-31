-- +goose Up
CREATE TABLE blocked_users (
		id SERIAL PRIMARY KEY,
		recipient_id integer NOT NULL,
		block_user_id integer NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		CONSTRAINT "UNQ__blocked_users__recipient_id__block_user_id" UNIQUE(recipient_id, block_user_id),
		CHECK (recipient_id <> block_user_id)
);

-- +goose Down
DROP TABLE blocked_users
