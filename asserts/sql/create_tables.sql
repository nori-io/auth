CREATE TABLE users
(id bigserial PRIMARY KEY,
 email  VARCHAR (32) NOT NULL,
 password VARCHAR (32) NOT NULL,
 status   SMALLINT NOT NULL,
 created_at TIMESTAMP,
 updated_at TIMESTAMP);

