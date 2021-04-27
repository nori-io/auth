package totp

type TotpHelper interface {
	Generate(email string) (string, error)
}
