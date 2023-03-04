UPDATE notes
SET text=$1
WHERE user_id=$2
RETURNING id;