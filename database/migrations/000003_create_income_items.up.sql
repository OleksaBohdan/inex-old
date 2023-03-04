CREATE TABLE IF NOT EXISTS income_items (
id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
user_id uuid NOT NULL,
name varchar(255) NOT NULL,
rang int NOT NULL DEFAULT 0,
FOREIGN KEY (user_id) REFERENCES users(id),
CONSTRAINT unique_user_income_item_name UNIQUE (user_id, name)
);
