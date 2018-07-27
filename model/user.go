package model

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/kdakean/kdakean/db"
	"github.com/kdakean/kdakean/errors"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	Fullname  string    `json:"fullname"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type UserSimple struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
}

type LoginForm struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

type RegisterForm struct {
	Name     string `form:"name" valid:"required,unique(users|name)"`
	Email    string `form:"email" valid:"required,email,unique(users|email)"`
	Password string `form:"password" valid:"required,minlen(8)"`
	Fullname string `form:"fullname" valid:"required"`
}

func CreateUser(f *RegisterForm) (*User, error) {
	pass, err := bcrypt.GenerateFromPassword([]byte(f.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	token, err := generateToken()
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	sum := md5.Sum(token)
	confirmToken := base64.StdEncoding.EncodeToString(sum[:])

	createdAt := time.Now().UTC()
	columns := []string{
		"name", "email", "password", "fullname",
		"confirm_token", "created_at", "updated_at",
	}
	values := []interface{}{
		f.Name, f.Email, string(pass), f.Fullname,
		confirmToken, createdAt, createdAt,
	}
	sql, args, err := db.SQ.Insert("users").Columns(columns...).
		Values(values...).Suffix("RETURNING id").ToSql()
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	user := User{
		Name:      f.Name,
		Email:     f.Email,
		Fullname:  f.Fullname,
		CreatedAt: createdAt,
		UpdatedAt: createdAt,
	}
	if err := db.DBX.QueryRow(sql, args...).Scan(&user.Id); err != nil {
		return nil, errors.ErrInternalServer
	}

	return &user, nil
}

func FindUserByLoginForm(f *LoginForm) (*User, error) {
	cond := sq.Or{
		sq.Eq{"name": f.Username},
		sq.Eq{"email": f.Username},
	}
	sql, args, err := db.SQ.Select("id, name, password").From("users").Where(cond).ToSql()
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	var user User
	if err := db.DBX.QueryRowx(sql, args...).StructScan(&user); err != nil {
		return nil, errors.ErrInvalidCredential
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(f.Password)); err != nil {
		return nil, errors.ErrInvalidCredential
	}

	return &user, nil
}

func FindUserById(userId uint) (*User, error) {
	sql, args, err := db.SQ.Select("id, name, email, fullname, created_at, updated_at").
		From("users").Where(sq.Eq{"id": userId}).ToSql()
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	var user User
	if err := db.DBX.QueryRowx(sql, args...).StructScan(&user); err != nil {
		return nil, errors.ErrRecordNotFound
	}
	return &user, nil
}

func generateToken() ([]byte, error) {
	buf := make([]byte, 32)
	if _, err := rand.Read(buf); err != nil {
		return nil, err
	}
	return buf, nil
}

// func sendConfirmEmail(user *User, token) error {
// 	confirmToken := base64.URLEncoding.EncodeToString(token)

// }
