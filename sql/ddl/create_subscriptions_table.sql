CREATE TABLE IF NOT EXISTS subscriptions (
  user_id VARCHAR(255) NOT NULL,
  type SMALLINT NOT NULL,
  duration INT NOT NULL,
  start INT NOT NULL,
  FOREIGN KEY (user_id) REFERENCES users(id)
);