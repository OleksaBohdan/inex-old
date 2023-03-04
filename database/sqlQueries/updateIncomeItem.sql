UPDATE income_items
SET name=$1
WHERE id=$2
RETURNING id;