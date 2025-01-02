CREATE OR REPLACE STREAM messages_stream (id INT, user_id INT, recipient_id INT, message VARCHAR, created_at TIMESTAMP)
WITH (VALUE_FORMAT='JSON', KAFKA_TOPIC='messages', PARTITIONS=3);

CREATE OR REPLACE TABLE blocked_users_table (id INT PRIMARY KEY, recipient_id INT, block_user_id INT)
WITH (VALUE_FORMAT='JSON', KAFKA_TOPIC='postgres.public.blocked_users', PARTITIONS=3);

CREATE OR REPLACE TABLE blocked_users_flatten_table
WITH (PARTITIONS=3) AS
SELECT recipient_id, COLLECT_LIST(block_user_id) as block_user_list
FROM blocked_users_table
GROUP BY recipient_id;

CREATE OR REPLACE STREAM messages_filtered_block_users_stream
WITH (VALUE_FORMAT='JSON', PARTITIONS=3) AS
SELECT m.id as "id", m.user_id as "user_id", m.recipient_id as "recipient_id", m.message as "message", m.created_at as "created_at"
FROM messages_stream m
LEFT JOIN blocked_users_flatten_table buf ON (m.recipient_id = buf.recipient_id)
WHERE NOT ARRAY_CONTAINS(buf.block_user_list, m.user_id);
