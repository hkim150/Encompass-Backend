CREATE TABLE user_account (
  user_account_id bigserial PRIMARY KEY,
  username varchar NOT NULL,
  email varchar NOT NULL,
  create_time timestamptz NOT NULL DEFAULT(Now())
);

CREATE TABLE user_profile (
  user_account_id bigint PRIMARY KEY,
  zipcode varchar,
  timezone varchar
);

CREATE TABLE deal (
  deal_id bigserial PRIMARY KEY,
  author bigint NOT NULL,
  store_name varchar NOT NULL,
  description varchar NOT NULL,
  regular_price decimal(10, 2),
  sale_price decimal(10, 2),
  upvote bigint NOT NULL,
  create_time timestamptz NOT NULL DEFAULT(Now())
);

ALTER TABLE user_profile ADD FOREIGN KEY (user_account_id) REFERENCES user_account (user_account_id);

ALTER TABLE deal ADD FOREIGN KEY (author) REFERENCES user_account (user_account_id);
