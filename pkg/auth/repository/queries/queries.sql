-- name: GetUserByEmail :one
SELECT * FROM TBL_USERS
WHERE EMAIL = $1 LIMIT 1;
