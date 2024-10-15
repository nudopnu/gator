-- name: CreateFeedFollow :one
WITH ff AS (
    INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
    VALUES (
        $1,
        $2,
        $3,
        $4,
        $5
    )
    RETURNING *
)
SELECT ff.*, f.name AS feed_name, u.name AS user_name 
FROM ff
JOIN feeds f ON ff.feed_id = f.id
JOIN users u ON ff.user_id = u.id;

-- name: GetFeedFollowsForUser :many
SELECT ff.*, f.name AS feed_name, u.name AS user_name 
FROM feed_follows ff 
JOIN feeds f ON ff.feed_id = f.id
JOIN users u ON ff.user_id = u.id
WHERE ff.user_id = $1;