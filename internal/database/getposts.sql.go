// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: getposts.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const getPosts = `-- name: GetPosts :many
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
WHERE 
    ff.user_id = $1
ORDER BY 
    p.published_at DESC
`

func (q *Queries) GetPosts(ctx context.Context, userID uuid.UUID) ([]Post, error) {
	rows, err := q.db.QueryContext(ctx, getPosts, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Post
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Title,
			&i.Url,
			&i.Description,
			&i.PublishedAt,
			&i.FeedID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
