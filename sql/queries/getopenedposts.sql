-- name: GetOpenedPost :many
SELECT 
    id,
    opened_at,
    user_id,
    feed_id,
    post_id
FROM opened
WHERE 
    user_id = $1
;