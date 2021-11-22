package psql

// Select
const (
	getUserByEmail = `SELECT * FROM "user" WHERE email LIKE $1`
	checkEmailExist     = `SELECT "id", "email" FROM "user" WHERE "deleted_at" IS NULL AND "email" = $1`
)

// Insert
const (
	insertNewSession  = `INSERT INTO "session" (jti, user_id, client, ip) VALUES ($1, $2, $3, $4)`
	insertNewUser     = `INSERT INTO "user" (id, email, password) VALUES ($1, $2, $3)`
	insertNewProfile  = `INSERT INTO "profile" (id, user_id, fullname) VALUES ($1, $2, $3)`
	insertNewPassword = `INSERT INTO "password" (id, user_id, password) VALUES ($1, $2, $3)`
)
