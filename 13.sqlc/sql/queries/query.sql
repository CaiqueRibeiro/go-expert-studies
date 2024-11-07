-- name: ListCategories :many
SELECT * FROM categories;

-- name: GetCategory :one
SELECT * FROM categories WHERE id = $1;

-- name: CreateCategory :exec
INSERT INTO categories (id, name, description) VALUES ($1, $2, $3);

-- name: UpdateCategory :exec
UPDATE categories SET name = $1, description = $2
WHERE id = $3;

-- name: DeleteCategory :exec
DELETE FROM categories WHERE id = $1;

-- name: CreateCourse :exec
INSERT INTO courses (id, name, description, category_id, price)
VALUES ($1, $2, $3, $4, $5);