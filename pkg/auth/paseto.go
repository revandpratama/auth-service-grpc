package auth

import (
	"errors"
	"strconv"
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
	token.SetString("user_id", strconv.Itoa(user.ID))
	token.SetString("role_id", strconv.Itoa(user.RoleID))
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

func VerifyToken(encryptedToken string) (*entity.User, error) {

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

	user, err := getPayloadFromParsedToken(parsedToken)

	return user, err
}

func getPayloadFromParsedToken(parsedToken *paseto.Token) (*entity.User, error) {
	userIDStr, err := parsedToken.GetString("user_id")
	if err != nil {
		return nil, err
	}
	roleIDStr, err := parsedToken.GetString("role_id")
	if err != nil {
		return nil, err
	}
	name, err := parsedToken.GetString("name")
	if err != nil {
		return nil, err
	}
	username, err := parsedToken.GetString("username")
	if err != nil {
		return nil, err
	}
	email, err := parsedToken.GetString("email")
	if err != nil {
		return nil, err
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		return nil, err
	}
	roleID, err := strconv.Atoi(roleIDStr)
	if err != nil {
		return nil, err
	}

	user := entity.User{
		ID:       userID,
		RoleID:   roleID,
		Name:     name,
		Email:    email,
		Username: username,
	}

	return &user, nil
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
