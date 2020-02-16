package model

import (
	"crypto/sha512"
	"database/sql"
	"encoding/base64"
	"fmt"
	"time"
)

const passwordSalt = "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824"

type User struct {
	ID        int        `json:"id"`
	Email     string     `json:"email"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	LastLogin *time.Time `json:"last_login"`
}

func Login(email, password string) (*User, error) {
	result := &User{}
	hasher := sha512.New()
	hasher.Write([]byte(passwordSalt))
	hasher.Write([]byte(password))
	hasher.Write([]byte(email))

	pwd := base64.URLEncoding.EncodeToString(hasher.Sum(nil))

	row := db.QueryRow(`
		SELECT id, email, first_name, last_name FROM public.user
		WHERE email = $1 AND password = $2
	`, email, pwd)
	err := row.Scan(&result.ID, &result.Email, &result.FirstName, &result.LastName)

	switch {
	case err == sql.ErrNoRows:
		return nil, fmt.Errorf("Not Found")
	case err != nil:
		return nil, err
	}

	t := time.Now()

	_, err = db.Exec(`
		UPDATE public.user
		SET last_login = $1
		WHERE id = $2
	`, t, result.ID)
	if err != nil {
		fmt.Printf("Failed to Update time on user. %v", err)
	}
	return result, nil
}

func SignUp(email, password, firstName, lastName string) error {
	// create result as address to User
	// Query Row to make sure user doesn't exist!

	// userRow := db.QueryRow(`
	// 	SELECT email FROM public.user
	// 	WHERE email = $1
	// `, email)
	// userRef := &User{}
	// userRow.Scan(&userRef.Email)
	// if user exists, respond with error that User exists..

	// create a hasher for the password...
	hasher := sha512.New()
	hasher.Write([]byte(passwordSalt))
	hasher.Write([]byte(password))
	hasher.Write([]byte(email))

	pwd := base64.URLEncoding.EncodeToString(hasher.Sum(nil))

	t := time.Now()
	// if user doesn't write to DB with INSERT
	res, err := db.Exec(`
	INSERT INTO users (email, first_name, last_name, password, last_login)
	VALUES($1, $2, $3, $4, $5)
	`, email, firstName, lastName, pwd, t)
	// if issue with db.Exec
	fmt.Println(res)
	if err != nil {
		// print err that issue with DB writing
		return fmt.Errorf("Issue with Saving to DB for user: %v, error: %v", email, err)
	}
	// userRow = row.Scan(&userResult.ID, &userResult.Email, &userResult.FirstName, $userResult.LastName)
	// return userResult and nil for error
	return nil

}
