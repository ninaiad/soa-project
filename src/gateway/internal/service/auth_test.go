package service_test

import (
	"fmt"
	"testing"

	"gateway/internal/service"
	"gateway/internal/user"

	"github.com/stretchr/testify/assert"
)

type MockDB struct {
	users []user.User
}

func TestCreateUser(t *testing.T) {
	mockDB := MockDB{users: []user.User{}}
	a := service.CreateAuthService(&mockDB)

	for i, tc := range []user.User{
		{
			Username: "username",
			Password: "pAsSwOrD",
			Name:     "NotNina",
			Surname:  "Knot",
			Email:    "alsdkfja@asldflsdkfj",
			Phone:    "2342oi3u423",
		},
		{
			Username: "sdkfjs",
			Password: "sfsdfksjdlf",
			Name:     "sdkfjsldkfn",
			Surname:  "xc,mvnxcv",
			Email:    "alsdkfja@",
			Phone:    "932r",
		},
	} {
		t.Run(tc.Name, func(t *testing.T) {
			id, err := a.CreateUser(tc)
			assert.NoError(t, err)
			assert.Equal(t, id, mockDB.users[i].Id)
			username, err := a.GetUsername(id)
			assert.NoError(t, err)
			assert.Equal(t, username, tc.Username)
			contains := false
			for _, u := range mockDB.users {
				if u.Username == username {
					contains = true
					break
				}
			}
			assert.True(t, contains)
			pub, err := a.UpdateUser(id, user.UserPublic{Name: "new " + tc.Name})
			assert.NoError(t, err)
			assert.Equal(t, pub.Name, "new "+tc.Name)
			assert.Equal(t, pub.Surname, tc.Surname)
			assert.Equal(t, pub.Email, tc.Email)
			assert.Equal(t, pub.Phone, tc.Phone)
		})
	}
}

func TestToken(t *testing.T) {
	mockDB := MockDB{users: []user.User{}}
	a := service.CreateAuthService(&mockDB)
	users := []user.User{{
		Username: "username1",
		Password: "pAsSwOrD1",
	},
		{
			Username: "username2",
			Password: "pAsSwOrD2",
		}}

	t.Run("token", func(t *testing.T) {
		id1, err := a.CreateUser(users[0])
		assert.NoError(t, err)
		assert.Equal(t, id1, mockDB.users[0].Id)
		token1, uid1, err := a.GenerateToken(users[0].Username, users[0].Password)
		assert.NoError(t, err)
		assert.Equal(t, uid1, id1)

		id2, err := a.CreateUser(users[1])
		assert.NoError(t, err)
		assert.Equal(t, id2, mockDB.users[1].Id)
		token2, uid2, err := a.GenerateToken(users[1].Username, users[1].Password)
		assert.NoError(t, err)
		assert.Equal(t, uid2, id2)

		assert.NotEqual(t, uid1, uid2)
		assert.NotEqual(t, id1, id2)
		assert.NotEqual(t, token1, token2)

		uid1, err = a.ParseToken(token1)
		assert.NoError(t, err)
		uid2, err = a.ParseToken(token2)
		assert.NoError(t, err)
		assert.Equal(t, uid1, id1)
		assert.Equal(t, uid2, id2)
	})
}

func (db *MockDB) CreateUser(user user.User) (int64, error) {
	user.Id = int64(len(db.users) + 1)
	db.users = append(db.users, user)
	return int64(len(db.users)), nil
}

func (db *MockDB) GetUserId(username, password string) (int64, error) {
	for _, user := range db.users {
		if user.Username == username && user.Password == password {
			return user.Id, nil
		}
	}
	return -1, fmt.Errorf("Not found")
}

func (db *MockDB) GetUserData(userId int64) (user.UserPublic, error) {
	for _, u := range db.users {
		if u.Id == userId {
			return user.UserPublic{
				Username: u.Username,
				Name:     u.Name,
				Surname:  u.Surname,
				Birthday: u.Birthday,
				Email:    u.Email,
				Phone:    u.Phone,
			}, nil
		}
	}

	return user.UserPublic{}, fmt.Errorf("Not found")
}

func (db *MockDB) UpdateUser(userId int64, update user.UserPublic, timeUpdated string) error {
	for i, u := range db.users {
		if u.Id == userId {
			db.users[i] = user.User{
				Id:       userId,
				Name:     update.Name,
				Surname:  update.Surname,
				Birthday: update.Birthday,
				Email:    update.Email,
				Phone:    update.Phone,
			}
			return nil
		}
	}

	return fmt.Errorf("Not found")
}

func (db *MockDB) DeleteUser(userId int64) error {
	toDelete := -1
	for i, u := range db.users {
		if u.Id == userId {
			toDelete = i
			break
		}
	}
	if toDelete == -1 {
		return fmt.Errorf("Not found")
	}

	if len(db.users) == 1 {
		db.users = []user.User{}
		return nil
	}

	db.users[toDelete] = db.users[len(db.users)-1]
	db.users = db.users[:1]
	return nil
}
