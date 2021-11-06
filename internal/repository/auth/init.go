package auth

type Repository struct{}

// TODO: ini untuk dependancy injection yang diperlukan db
func New() *Repository {
	return &Repository{}
}
