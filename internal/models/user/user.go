package user

import (
	"database/sql"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/zigapk/prpo-auth/internal/config"
	"github.com/zigapk/prpo-auth/internal/database"
	"github.com/zigapk/prpo-auth/internal/logger"
	"github.com/zigapk/prpo-auth/internal/util"
	"golang.org/x/crypto/bcrypt"
)

const bcryptCost = 12

var ErrInvalidCredentials = errors.New("invalid credentials")

type User struct {
	UID string `db:"uid"`

	Email string `db:"email"`
	Name  string `db:"name"`

	Password string `db:"password"`

	DateCreated time.Time `db:"date_created"`
}

// New user with a given username and password.
func New(email, name, password string) (*User, error) {
	// New UUID.
	uid := uuid.New().String()

	// Hash user password.
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcryptCost)
	if err != nil {
		return nil, err
	}

	// Insert user to database.
	u := &User{}
	insert := `INSERT INTO users (uid, email, name, password) VALUES ($1, $2, $3, $4) RETURNING *`
	err = database.DB.Get(u, insert, uid, email, name, string(hash))

	// Handle errors.
	if err != nil {
		return nil, err
	}

	return u, nil
}

// ForUID queries user for specified uid.
func ForUID(uid string) (*User, error) {
	u := &User{}

	query := `SELECT * FROM users WHERE uid=$1`
	err := database.DB.Get(u, query, uid)
	if err != nil {
		return nil, err
	}

	return u, nil
}

// All returns a slice of all users.
func All() ([]*User, error) {
	var users []*User

	query := `SELECT * FROM users ORDER BY name`
	err := database.DB.Select(&users, query)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u *User) Update(name string) error {
	tx, err := database.DB.Beginx()
	if err != nil {
		return err
	}

	// Set new name.
	update := `UPDATE users SET name=$2 WHERE uid=$1 RETURNING *`
	err = tx.Get(u, update, u.UID, name)
	if err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			logger.Log.Warn().Err(err).Send()
		}

		return err
	}

	// Commit.
	err = tx.Commit()
	if err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			logger.Log.Warn().Err(err).Send()
		}

		return err
	}

	return nil
}

func SetPassword(uid, password string) error {
	// Hash.
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcryptCost)
	if err != nil {
		return err
	}

	// Save new password.
	update := `UPDATE users SET password=$2 WHERE uid=$1`
	_, err = database.DB.Exec(update, uid, string(passwordHash))
	return err
}

func (u *User) PasswordMatch(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func FromToken(accessToken string) (*User, error) {
	claims, err := util.Validate(accessToken)
	if err != nil {
		return nil, err
	}

	u := &User{
		UID: claims.Subject,
	}
	return u, nil
}

func AuthorizeWithPassword(email, password string) (*User, error) {
	u := &User{}

	// Get user for specified email.
	query := `SELECT * FROM users WHERE email=$1`
	err := database.DB.Get(u, query, email)

	// Handle errors.
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrInvalidCredentials
		} else {
			return nil, err
		}
	}

	// Check that passwords match.
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	return u, nil
}

func AuthorizeWithToken(token string) (*User, error) {
	u := &User{}

	// Get user with token.
	query := `SELECT users.* FROM authenticated_devices INNER JOIN users ON authenticated_devices.user_id = users.uid
		WHERE token=$1`
	err := database.DB.Get(u, query, token)

	// Handle errors.
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrInvalidCredentials
		} else {
			return nil, err
		}
	}

	return u, nil
}

func (u *User) CreateAccessToken() (string, error) {
	now := time.Now()
	expiresAt := now.Add(time.Duration(config.Login.TokenTtl()) * time.Second)

	claims := util.Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt.Unix(),
			IssuedAt:  now.Unix(),
			Subject:   u.UID,
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	ss, err := accessToken.SignedString(config.Login.SigningPrivateKey)

	return ss, err
}

func (u *User) CreateRefreshToken() (string, error) {
	now := time.Now()
	expiresAt := now.Add(time.Duration(config.Login.RefreshTokenTtl()) * time.Second)

	claims := util.Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt.Unix(),
			IssuedAt:  now.Unix(),
			Subject:   u.UID,
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	ss, err := refreshToken.SignedString(config.Login.SigningPrivateKey)

	return ss, err
}
