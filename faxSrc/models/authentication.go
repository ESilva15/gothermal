package models

import (
	"log"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type Authenticator interface {
	isAuthorised(string, string) bool
	Authenticate(w http.ResponseWriter, r *http.Request) bool
}
type Authentication struct{}

type LoginForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	jwt.RegisteredClaims
}

// Private functions

func openDB() *SQLite {
	var d SQLite
	err := d.Initialize("../src/data.db")
	if err != nil {
		log.Fatal("Failed to get into DB: ", err)
	}

	return &d
}

func getUserPass(user string) (string, error) {
	d := openDB()

	query := `
	SELECT password FROM users WHERE name='` + user + `';
	`

	rows, err := d.DB.Query(query)
	if err != nil {
		log.Fatal("Failed to get user data: ", err)
		return "", err
	}
	defer rows.Close()

	var p string
	for rows.Next() {
		if err := rows.Scan(&p); err != nil {
			log.Fatal("Failed to read row.", err)
		}
	}

	// No password found (user doesn't exist?)
	if len(p) == 0 {
		return "", nil
	}

	return p, nil
}

func insertToken(token string, user string, expirationDate int64) error {
	d := openDB()

	query := `
	INSERT INTO
	  sessions(user, token, expires)
		values(
		'` + user + `',
		'` + token + `',
		` + strconv.Itoa(int(expirationDate)) + `
		);
	`

	log.Println(query)

	_, err := d.DB.Exec(query)
	if err != nil {
		log.Fatal("Failed to insert data: ", err)
		return err
	}

	return nil
}

// Public functions

func ValidateCredentials(lf *LoginForm) bool {
	passwd, err := getUserPass(lf.Username)
	if err != nil {
		log.Fatal("Failed to get password from DB: ", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(passwd), []byte(lf.Password))

	return err == nil
}

func StoreSessionToken(token string, user string, expirationDate int64) error {
	err := insertToken(token, user, expirationDate)
	if err != nil {
		log.Fatal("Unable to insert token into DB: ", err)
	}

	return nil
}
