package user_dal

import (
	"github.com/alf-grindel/dawn/conf"
	"github.com/alf-grindel/dawn/internal/dal"
	"github.com/alf-grindel/dawn/pkg/utils"
	"testing"
)

func TestUserDal(t *testing.T) {
	conf.Init()
	dal.Init()
	db := dal.Open()
	snowflake := utils.NewSnowflake(0)
	userDal := NewUserDal(db, snowflake)

	t.Run("create user", func(t *testing.T) {
		account := "dawn"
		password := "12345678"
		hashPassword := utils.HashPassword(password)
		id, err := userDal.CreateUser(account, hashPassword)
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(id)
	})

	t.Run("query user by account", func(t *testing.T) {
		account := "dawn"
		u, err := userDal.GetUserByAccount(account)
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(u)
	})

	t.Run("update user", func(t *testing.T) {
		account := "dawn"
		user, _ := userDal.GetUserByAccount(account)
		user.IsDelete = 1
		u, err := userDal.UpdateUser(user)
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(u)
	})
}
