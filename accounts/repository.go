package accounts

type Repository interface {
	Find(string) (*User, error)
	Store(*User) error
}
