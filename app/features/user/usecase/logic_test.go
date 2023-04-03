package usecase_test

import (
	"errors"
	"ormapi/app/features/user"
	"ormapi/app/features/user/mocks"
	"ormapi/app/features/user/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	// fixture
	repo := mocks.NewRepository(t)
	ul := usecase.New(repo)
	successCaseData := user.Core{Nama: "jerry", HP: "12345", Password: "tonohaha577"}

	t.Run("Sukses login", func(t *testing.T) {
		// panggil repository palsu dengan parameter dan return sesuai yg kita inginkan
		// jalankan logic yang akan test sesuai dengan parameter yg digunakan dalam repository
		// cek hasil logic menggunakan assert secukup
		// jalankan .assertexpectation untuk cek apakah expektasi repository berjalan dengan baik

		repo.On("Login", successCaseData.HP, successCaseData.Password).Return(user.Core{Nama: "jerry", HP: "12345"}, nil).Once()
		result, err := ul.Login("12345", "tonohaha577")

		assert.Nil(t, err)
		assert.Equal(t, "12345", result.HP)
		assert.Equal(t, "jerry", result.Nama)
		repo.AssertExpectations(t)
	})

	t.Run("Password salah", func(t *testing.T) {
		repo.On("Login", successCaseData.HP, "tonohaha").Return(user.Core{}, errors.New("password salah")).Once()
		result, err := ul.Login("12345", "tonohaha")

		assert.Error(t, err)
		assert.ErrorContains(t, err, "password salah")
		assert.Empty(t, result.Nama)
		repo.AssertExpectations(t)
	})

	t.Run("Data tidak ditemukan", func(t *testing.T) {
		repo.On("Login", "6789", "tonohaha").Return(user.Core{}, errors.New("data tidak ditemukan")).Once()
		result, err := ul.Login("6789", "tonohaha")

		assert.Error(t, err)
		assert.ErrorContains(t, err, "data tidak ditemukan")
		assert.Empty(t, result.Nama)
		repo.AssertExpectations(t)
	})

	t.Run("Kesalahan pada server", func(t *testing.T) {
		repo.On("Login", successCaseData.HP, "tonohaha").Return(user.Core{}, errors.New("column not exist")).Once()
		result, err := ul.Login("12345", "tonohaha")

		assert.Error(t, err)
		assert.ErrorContains(t, err, "server")
		assert.Empty(t, "", result.Nama)
		repo.AssertExpectations(t)
	})
}

// func TestLogin(t *testing.T) {
// 	ul := usecase.New(&MockSuccess{}) // fixture

// 	t.Run("Sukses login", func(t *testing.T) {
// 		result, err := ul.Login("12345", "tonohaha577")

// 		assert.Nil(t, err)
// 		assert.Equal(t, "12345", result.HP)
// 		assert.Equal(t, "jerry", result.Nama)
// 	})

// 	t.Run("Password salah", func(t *testing.T) {
// 		ul := usecase.New(&MockGagal{})

// 		result, err := ul.Login("12345", "tonohaha")

// 		assert.Error(t, err)
// 		assert.ErrorContains(t, err, "password salah")
// 		assert.Empty(t, "", result.Nama)
// 	})
// }

// fixture
type MockSuccess struct{}

func (ms *MockSuccess) Login(hp string, password string) (user.Core, error) { // asumsi kembalian SUKSES dari repository
	return user.Core{Nama: "jerry", HP: "12345"}, nil
}

func (ms *MockSuccess) Insert(newUser user.Core) (user.Core, error) {
	return user.Core{}, nil
}

type MockGagal struct{}

func (mg *MockGagal) Login(hp string, password string) (user.Core, error) { // asumsi kembalian SUKSES dari repository
	return user.Core{}, errors.New("password salah")
}

func (mg *MockGagal) Insert(newUser user.Core) (user.Core, error) {
	return user.Core{}, nil
}
