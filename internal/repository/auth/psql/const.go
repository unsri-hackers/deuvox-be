package psql

// Select
const (
	getUserByEmailQuery = `SELECT * FROM "user" WHERE email LIKE $1`
)

// Insert
const (
	insertNewSession = `INSERT INTO "session" (jti, user_id, client, ip) VALUES ($1, $2, $3, $4)`
)
