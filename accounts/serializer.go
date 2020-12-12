package accounts

type Serializer interface {
	Encode(*User) ([]byte, error)
	Decode([]byte) (*User, error)
}
