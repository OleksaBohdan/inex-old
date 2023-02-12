CREATE TABLE IF NOT EXISTS cost_items(
id uuid PRIMARY KEY default gen_random_uuid(),
user_id uuid not null,
FOREIGN KEY (user_id) REFERENCES users(id),
name varchar(255),
rang int not null default(0));
