package tools

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

const (
	TOKEN_ISSUER = "StarkSim"
	SECRETKEY    = "42wqTE23123wffLU94342wgadgFs"

	AccessTokenExp  = time.Hour * 2
	RefreshTokenExp = time.Hour * 12
)

type CustomClaims struct {
	UserId int64
	jwt.RegisteredClaims
}

// GetToken 生成 JWT
func GetToken(CreateAt time.Time, UserID int64) (string, error) {
	customClaims := &CustomClaims{
		UserId: UserID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    TOKEN_ISSUER,
			ExpiresAt: &jwt.NumericDate{Time: CreateAt.Add(AccessTokenExp)},
			IssuedAt:  &jwt.NumericDate{Time: CreateAt},
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims)
	signedToken, err := token.SignedString([]byte(SECRETKEY))
	//refreshToken := base64.URLEncoding.EncodeToString(utils.NewSHA1(utils.Must(utils.NewRandom()), []byte(access)).Bytes())
	//refreshToken = strings.ToUpper(strings.TrimRight(refreshToken, "="))
	return signedToken, err
}

// ParseToken 解析token
func ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SECRETKEY), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
