package service

import (
	"context"
	"crypto/md5"
	"errors"
	"fmt"
	"os"
	"time"

	posts_pb "gateway/internal/service/posts_pb"
	stat_pb "gateway/internal/service/statistics_pb"
	"gateway/internal/user"

	"github.com/golang-jwt/jwt/v5"
)

var salt = os.Getenv("PASSWORD_SALT")
var signingKey = os.Getenv("TOKEN_SIGNING_KEY")

const (
	tokenTTL             = 12 * time.Hour
	fieldNameTimeUpdated = "time_updated"
)

type tokenClaims struct {
	jwt.RegisteredClaims
	UserId int64 `json:"user_id"`
}

func (s *Service) CreateUser(user user.User) (int64, error) {
	user.Password = generatePasswordHash(user.Password)
	user.TimeCreated = time.Now().Format(time.RFC3339)
	user.TimeUpdated = user.TimeCreated
	if user.Birthday == "" {
		user.Birthday = time.Time{}.Format(time.RFC3339)
	}

	return s.db.CreateUser(user)
}

func (s *Service) GetUsername(userId int64) (string, error) {
	if u, err := s.db.GetUserData(userId); err == nil {
		return u.Username, nil
	} else {
		return "", err
	}
}

func (s *Service) UpdateUser(userId int64, update user.UserPublic) (user.UserPublic, error) {
	userData, err := s.db.GetUserData(userId)
	if err != nil {
		return user.UserPublic{}, err
	}

	if update.Username == "" {
		update.Username = userData.Username
	}
	if update.Birthday == "" {
		update.Birthday = userData.Birthday
	}
	if update.Name == "" {
		update.Name = userData.Name
	}
	if update.Surname == "" {
		update.Surname = userData.Surname
	}
	if update.Email == "" {
		update.Email = userData.Email
	}
	if update.Phone == "" {
		update.Phone = userData.Phone
	}

	return update, s.db.UpdateUser(userId, update, time.Now().Format(time.RFC3339))
}

func (s *Service) DeleteUser(userId int64) error {
	if err := s.db.DeleteUser(userId); err != nil {
		return err
	}

	_, err := s.pClient.DeleteUser(context.Background(), &posts_pb.UserId{Id: userId})
	if err != nil {
		return err
	}

	_, err = s.sClient.DeleteUser(context.Background(), &stat_pb.UserId{Id: userId})
	return err
}

func (s *Service) GenerateToken(username, password string) (string, int64, error) {
	userId, err := s.db.GetUserId(username, generatePasswordHash(password))
	if err != nil {
		return "", 0, err
	}

	tNow := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(tNow.Add(tokenTTL)),
			IssuedAt:  jwt.NewNumericDate(tNow),
		},
		userId,
	})

	tokenS, err := token.SignedString([]byte(signingKey))
	return tokenS, userId, err
}

func (s *Service) ParseToken(accessToken string) (int64, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}

func generatePasswordHash(password string) string {
	hash := md5.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
