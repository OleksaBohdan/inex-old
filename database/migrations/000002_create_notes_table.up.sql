CREATE TABLE IF NOT EXISTS notes(
id uuid PRIMARY KEY default gen_random_uuid(),
user_id uuid not null,
FOREIGN KEY (user_id) REFERENCES users(id),
text varchar(280)
);
