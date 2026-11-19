package models

import "time"

type User struct {
	ID            string    `json:"id" bson:"_id,omitempty"`
	Name          string    `json:"name" bson:"name"`
	Username      string    `json:"username" bson:"username"`
	Email         string    `json:"email" bson:"email"`
	PasswordHash  string    `json:"-" bson:"password_hash"`
	Salt          string    `json:"-" bson:"kdf_salt"`
	KDF           string    `json:"-" bson:"kdf_algorithm"`
	KDFParameters KDFParams `json:"-" bson:"kdf_params"`
	WrappedVEK    string    `json:"-" bson:"wrapped_vek"`
	CreatedAt     time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt" bson:"updatedAt"`
}

type RegisterUserDetails struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required,min=6"`
}

type LoginUserDetails struct {
	Username string `json:"username" validate:"required_without=Email"`
	Email    string `json:"email" validate:"required_without=Username"`
	Password string `json:"password" validate:"required,min=6"`
}

type KDFParams struct {
	Memory      uint32 `json:"-" bson:"memory"`
	Iterations  uint32 `json:"-" bson:"iterations"`
	Parallelism uint8  `json:"-" bson:"parallelism"`
}
