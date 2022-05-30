-- name: CreateCustomer :one
INSERT INTO customers (
  full_name,
  phone_number
) VALUES (
  $1, $2
) RETURNING *;

-- name: GetCustomer :one
SELECT * FROM customers
WHERE id = $1 LIMIT 1;

-- name: ListCustomers :many
SELECT * FROM customers
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateCustomer :one
UPDATE customers
SET full_name = $2, phone_number = $3
WHERE id = $1
RETURNING *;


-- name: DeleteCustomer :exec
DELETE FROM customers
WHERE id = $1;