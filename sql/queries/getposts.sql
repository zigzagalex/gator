-- name: GetPosts :many
SELECT
    p.id,
    p.created_at,
    p.updated_at,
    p.title,
    p.url,
    p.description,
    p.published_at,
    p.feed_id 
FROM posts as p
LEFT JOIN feed_follows as ff
    ON ff.feed_id = p.feed_id
WHERE ff.user_id = $1
ORDER BY 
    p.published_at DESC
;