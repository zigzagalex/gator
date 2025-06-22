-- name: GetFeeds :many
SELECT 
    feeds.created_at,
    feeds.name,
    feeds.url,
    feeds.id,
    users.name
FROM feeds
LEFT JOIN users
    ON users.id = feeds.user_id
ORDER BY
    feeds.created_at DESC
;