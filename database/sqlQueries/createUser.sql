INSERT INTO users (name, surname, email, password_hash) VALUES ($1, $2, $3, $4) RETURNING id;