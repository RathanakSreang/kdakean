package model

type User struct {
	Name     string `valid:"required~err_name_required"`
	Fullname string `valid:"-"`
	Password string `valid:"required~err_password_required"`
	Email    string `valid:"required~err_email_required,email~err_email_ivalid"`
}
