CREATE TABLE IF NOT EXISTS oauth (
  user_id VARCHAR(255) NOT NULL,
  token VARCHAR(255) NOT NULL,
  refresh_token VARCHAR(255) NOT NULL,
  start_time INT NOT NULL,
  expires_in INT NOT NULL,
  FOREIGN KEY (user_id) REFERENCES users(id)
);