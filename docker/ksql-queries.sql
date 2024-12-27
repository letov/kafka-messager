CREATE OR REPLACE TABLE blocked_users (id INT PRIMARY KEY, user_id INT, block_user_id INT) WITH (VALUE_FORMAT='JSON', KAFKA_TOPIC='postgres.public.blocked_users');
CREATE OR REPLACE TABLE ban_words (id INT PRIMARY KEY, word VARCHAR) WITH (VALUE_FORMAT='JSON', KAFKA_TOPIC='postgres.public.ban_words');
