CREATE TABLE IF NOT EXISTS incomes(
id uuid PRIMARY KEY default gen_random_uuid(),
user_id uuid not null,
FOREIGN KEY (user_id) REFERENCES users(id),
income_id uuid not null,
FOREIGN KEY (income_id) REFERENCES income_items(id),
value float8 not null default(0),
rang int not null default(0));