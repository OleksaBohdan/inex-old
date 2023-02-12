INSERT INTO notes (user_id, text) VALUES ($1, $2) RETURNING id;

