# хотел сделать фильтрацию через KSQLDB blocked_users_table и ban_words_table, но на уровне запросов невозможно(
# поэтому репликация по факту не используется
CREATE OR REPLACE TABLE blocked_users_table (id INT PRIMARY KEY, user_id INT, block_user_id INT) WITH (VALUE_FORMAT='JSON', KAFKA_TOPIC='postgres.public.blocked_users');
CREATE OR REPLACE TABLE ban_words_table (id INT PRIMARY KEY, word VARCHAR) WITH (VALUE_FORMAT='JSON', KAFKA_TOPIC='postgres.public.ban_words');

