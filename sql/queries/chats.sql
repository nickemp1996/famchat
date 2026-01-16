-- name: CreateChat :one
INSERT INTO chats (id, created_at, updated_at, name, user_id)
VALUES (
    $1,
    NOW(),
    NOW(),
    $2,
    $3
)
RETURNING *;

-- name: GetChat :one
SELECT * FROM chats
WHERE name = $1 LIMIT 1;

-- name: DeleteChats :exec
DELETE FROM chats;