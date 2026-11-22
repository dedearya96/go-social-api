package services

import (
	"errors"
	"go-social-api/src/config"
	"go-social-api/src/models/auth"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct{}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (s *AuthService) LoginWithEmail(email, password string) (*auth.User, string, string, time.Time, error) {
	var user auth.User
	err := config.DBAuth.Where("email = ? AND role = ?", email, "users").First(&user).Error
	if err != nil {
		return nil, "", "", time.Time{}, errors.New("Invalid Email or Password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, "", "", time.Time{}, errors.New("Invalid Email or Password")
	}

	accessToken, err := s.GenerateJWT(user.ID)
	if err != nil {
		return nil, "", "", time.Time{}, err
	}

	refreshToken := uuid.New().String() + uuid.New().String()
	expiresAt := time.Now().Add(14 * 24 * time.Hour)

	var rt auth.RefreshToken
	res := config.DBAuth.Where("users_id = ?", user.ID).First(&rt)

	if res.Error == nil {
		config.DBAuth.Model(&rt).Updates(auth.RefreshToken{Token: refreshToken, ExpiresAt: expiresAt})
	} else {
		config.DBAuth.Create(&auth.RefreshToken{UsersID: user.ID, Token: refreshToken, ExpiresAt: expiresAt})
	}

	return &user, accessToken, refreshToken, expiresAt, nil
}

func (s *AuthService) GenerateJWT(userID string) (string, error) {
	claims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
		"iat": time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
