package mfa_totp

import (
	"testing"

	config2 "github.com/nori-plugins/authentication/internal/config"
)

func Test_totpHelper_Generate(t *testing.T) {
	conf := config2.Config{
		Issuer: func() string {
			return "nori"
		},
	}

	h := New(Params{Config: conf})
	image, key, err := h.Generate("test@mail.ru")

	t.Log(image)
	t.Log(key)
	t.Log(err)
}
