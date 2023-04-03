package usecase_test

import (
	"ormapi/app/features/user"
	"ormapi/app/features/user/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	ul := usecase.New(&MockSuccess{})

	t.Run("Sukses login", func(t *testing.T) {
		result, err := ul.Login("12345", "tonohaha577")

		assert.Nil(t, err)
		assert.Equal(t, "12345", result.HP)
		assert.Equal(t, "jerry", result.Nama)
	})
}

type MockSuccess struct{}

func (ms *MockSuccess) Login(hp string, password string) (user.Core, error) { // asumsi kembalian SUKSES dari repository
	return user.Core{Nama: "jerry", HP: "12345"}, nil
}

func (ms *MockSuccess) Insert(newUser user.Core) (user.Core, error) {
	return user.Core{}, nil
}
