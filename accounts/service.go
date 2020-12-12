package accounts

type AccountsService interface {
	Find(string) (*User, error)
	Store(*User) error
}
