UPDATE incomes
SET value=$1
WHERE id=$2
RETURNING id;