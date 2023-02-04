CREATE TABLE IF NOT EXISTS note(
id uuid PRIMARY KEY default gen_random_uuid(),
text varchar(280)
);
