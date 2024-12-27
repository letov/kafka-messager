-- +goose Up
CREATE TABLE blocked_users (
		id SERIAL PRIMARY KEY,
		user_id integer NOT NULL,
		block_user_id integer NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		CONSTRAINT "UNQ__blocked_users__user_id__block_user_id" UNIQUE(user_id, block_user_id),
		CHECK (user_id <> block_user_id)
);

-- +goose Down
DROP TABLE blocked_users
