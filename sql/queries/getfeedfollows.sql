-- name: GetFeedFollowsForUser :many
SELECT 
    feed_follows.id,
    feed_follows.created_at,
    users.name AS user_name,
    users.id AS user_id,
    feeds.name AS feed_name,
    feeds.id AS feed_id
FROM feed_follows
LEFT JOIN users 
    ON users.id = feed_follows.user_id
LEFT JOIN feeds
    ON feeds.id = feed_follows.feed_id
WHERE users.name = $1
;