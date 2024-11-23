package auth

import (
	"errors"
	"time"

	"aidanwoods.dev/go-paseto"
	"github.com/revandpratama/reflect/auth-service/internal/entity"
)

var secretKey = paseto.NewV4SymmetricKey()

func CreateToken(user *entity.User) (string, error) {
	token := paseto.NewToken()

	//set rule
	token.SetIssuedAt(time.Now()) // Set the `iat` claim
	token.SetNotBefore(time.Now())
	token.SetExpiration(time.Now().Add(2 * time.Minute))

	//insert paylaod
	token.SetString("user_id", user.ID.String())
	token.SetString("name", user.Name)
	token.SetString("username", user.Username)
	token.SetString("email", user.Email)

	// Private key (DO NOT SHARE)

	// signed := token.V4Sign(secretKey, nil)
	encrypted := token.V4Encrypt(secretKey, nil)

	var err error
	if encrypted == "" {
		return "", errors.New("failed creating token")
	}

	return encrypted, err
	// return &Token{
	// 	publicKeyHex: publicKey.ExportHex(),
	// 	signedToken:  signed,
	// }, err
}

func VerifyToken(encryptedToken string) (*paseto.Token, error) {

	parser := paseto.NewParser()
	parser.AddRule(paseto.NotExpired())
	parser.AddRule(paseto.ValidAt(time.Now()))

	// publicKey, err := paseto.NewV4AsymmetricPublicKeyFromHex(publicKeyHex)
	// if err != nil {
	// 	return nil, err
	// }

	parsedToken, err := parser.ParseV4Local(secretKey, encryptedToken, nil)
	if err != nil {
		return nil, err
	}
	// Parse and validate the token
	// parsedToken, err := parser.ParseV4Public(publicKey, signedToken, nil)
	// if err != nil {
	// 	return nil, err
	// }

	return parsedToken, err
}

// func AuthTest() {

// 	user := entity.User{
// 		ID:       uuid.New(),
// 		Name:     "Black Star",
// 		Username: "black.star",
// 		Email:    "black@gmail.com",
// 	}

// 	signedToken, err := CreateToken(&user)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	fmt.Println("signedToken :", signedToken)

// 	parsedToken, err := VerifyToken(signedToken)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}

// 	fmt.Println("parsedToken : ", parsedToken)

// 	userID, err := parsedToken.GetString("user_id")
// 	name, err := parsedToken.GetString("name")
// 	username, err := parsedToken.GetString("username")

// 	fmt.Println(userID)
// 	fmt.Println(name)
// 	fmt.Println(username)

// }
