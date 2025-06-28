-- name: GetUnreadCount :one
SELECT COUNT(*) FROM posts p
JOIN feed_follows ff ON p.feed_id = ff.feed_id
LEFT JOIN opened o
  ON o.post_id = p.id AND o.user_id = ff.user_id
WHERE
  ff.user_id = $1
  AND ff.feed_id = $2
  AND o.id IS NULL
  AND p.published_at > ff.created_at;
