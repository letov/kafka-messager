CREATE IF NOT EXISTS TestTableFinal (id INT PRIMARY KEY, name VARCHAR) WITH (VALUE_FORMAT='JSON', KAFKA_TOPIC='postgres.public.test_table');
