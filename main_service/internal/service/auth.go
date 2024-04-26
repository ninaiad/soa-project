package service

import (
	"crypto/md5"
	"errors"
	"fmt"
	"os"
	"soa-main/internal/user"
	"soa-main/internal/database"
	"time"

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
	UserId int `json:"user_id"`
}

type AuthService struct {
	db database.Authorization
}

func CreateAuthService(db database.Authorization) *AuthService {
	return &AuthService{db: db}
}

func (a *AuthService) CreateUser(user user.User) error {
	user.Password = generatePasswordHash(user.Password)
	user.TimeCreated = time.Now().Format(time.RFC3339)
	user.TimeUpdated = user.TimeCreated
	if user.Birthday == "" {
		user.Birthday = time.Time{}.Format(time.RFC3339)
	}

	return a.db.CreateUser(user)
}

func (a *AuthService) UpdateUser(userId int, update user.UserPublic) (user.UserPublic, error) {
	userData, err := a.db.GetUserData(userId)
	if err != nil {
		return user.UserPublic{}, err
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

	return update, a.db.UpdateUser(userId, update, time.Now().Format(time.RFC3339))
}

func (a *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := a.db.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	tNow := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(tNow.Add(tokenTTL)),
			IssuedAt:  jwt.NewNumericDate(tNow),
		},
		user.Id,
	})

	return token.SignedString([]byte(signingKey))
}

func (a *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
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
