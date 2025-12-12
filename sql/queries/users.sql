-- name: GetUserById :one

SELECT user_id, age, country, device, language, interests
FROM users
WHERE user_id = $1;
