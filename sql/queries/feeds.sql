-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: GetFeeds :many
SELECT F.*, U.name AS user_name FROM (
    users U JOIN feeds F ON U.id = F.user_id
);

-- name: GetFeedByUrl :one
SELECT * FROM feeds WHERE feeds.url = $1;

-- name: MarkFeedFetched :exec
UPDATE feeds SET updated_at=CURRENT_TIMESTAMP, last_fetched_at=CURRENT_TIMESTAMP WHERE id=$1;

-- name: GetNextFeedToFetch :one
SELECT * FROM feeds 
ORDER BY last_fetched_at ASC NULLS FIRST;