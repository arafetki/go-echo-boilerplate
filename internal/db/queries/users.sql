-- name: InsertUser :exec
INSERT INTO users (id,name,email,avatar_url,is_email_verified,created_at) VALUES ($1,$2,$3,$4,$5,$6);