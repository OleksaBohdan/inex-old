UPDATE users
SET name=$1, surname=$2, email=$3, password_hash=$4
WHERE id=$5;