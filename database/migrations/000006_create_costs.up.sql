CREATE TABLE IF NOT EXISTS costs(
id uuid PRIMARY KEY default gen_random_uuid(),
user_id uuid not null,
FOREIGN KEY (user_id) REFERENCES users(id),
cost_id uuid not null,
FOREIGN KEY (cost_id) REFERENCES cost_items(id),
value float8 not null default(0),
rang int not null default(0));