-- name: CreateOpenedPost :one
INSERT INTO opened (id, opened_at, user_id, feed_id, post_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
)
RETURNING *;