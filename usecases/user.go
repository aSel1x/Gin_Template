package usecases

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/aSel1x/Gin_Template/adapters"
	"github.com/aSel1x/Gin_Template/models"
	"github.com/aSel1x/Gin_Template/usecases/security"
)

type UserUsecase struct {
	*adapters.Adapters
	appSecret string
}

func NewUserUsecase(adapters adapters.Adapters, appSecret string) *UserUsecase {
	return &UserUsecase{&adapters, appSecret}
}

func (us *UserUsecase) Create(user models.UserCreate) (*models.User, error) {
	storedUser, err := us.UserRepo.RetrieveByUsername(user.Username)
	if storedUser != nil {
		return nil, fmt.Errorf("user already exists")
	}

	hashedPassword, _ := security.HashPwd(user.Password)
	newUser := models.User{
		UserBase: models.UserBase{
			Username: user.Username,
			IsActive: true,
		},
		Password: hashedPassword,
	}

	err = us.UserRepo.Create(&newUser)

	return nil, err
}

func (us *UserUsecase) Authenticate(user models.UserCreate) (*models.User, error) {
	storedUser, err := us.UserRepo.RetrieveByUsername(user.Username)
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}
	if !security.CheckPwd(user.Password, storedUser.Password) {
		return nil, fmt.Errorf("invalid password")
	}
	return storedUser, nil
}

func (us *UserUsecase) OAuth2(user *models.User) (*models.UserAuth, error) {
	expiresIn := 3600 // 1 hour
	expirationTime := time.Now().Add(time.Duration(expiresIn) * time.Second)

	accessToken, err := security.Encode(
		us.appSecret,
		jwt.MapClaims{"id": user.ID},
		&expirationTime,
	)
	if err != nil {
		return nil, fmt.Errorf("could not generate access token")
	}

	refreshToken, err := security.Encode(
		us.appSecret,
		jwt.MapClaims{"id": user.ID},
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("could not generate refresh token")
	}

	return &models.UserAuth{
		AccessToken:  accessToken,
		TokenType:    "Bearer",
		ExpiresIn:    expiresIn,
		RefreshToken: refreshToken,
	}, nil
}

func (us *UserUsecase) RetrieveByToken(token string) (*models.User, error) {
	payload, err := security.Decode(us.appSecret, token)
	if err != nil || payload == nil {
		return nil, fmt.Errorf("invalid token")
	}
	id := int(payload["id"].(float64))

	user, err := us.UserRepo.RetrieveOne(id)

	return user, err
}

func (us *UserUsecase) RefreshOAuth2(refreshToken string) (*models.UserAuth, error) {
	user, err := us.RetrieveByToken(refreshToken)
	if err != nil {
		return nil, err
	} else {
		return us.OAuth2(user)
	}
}
