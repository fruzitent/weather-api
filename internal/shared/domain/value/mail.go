package value

import "net/mail"

type Mail struct{ string }

func NewMail(address string) (*Mail, error) {
	addr, err := mail.ParseAddress(address)
	if err != nil {
		return nil, err
	}
	return &Mail{addr.Address}, nil
}
