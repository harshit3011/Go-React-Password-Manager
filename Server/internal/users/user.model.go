package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type User struct {
	ID            bson.ObjectID `json:"id" bson:"_id,omitempty"`
	Name          string        `json:"name,omitempty" bson:"name,omitempty"`
	Username      string        `json:"username" bson:"username"`
	Email         string        `json:"email" bson:"email"`
	EmailVerified bool          `json:"emailVerified" bson:"email_verified"`

	AuthHash   string     `json:"-" bson:"auth_hash"`
	KDF        KDFConfig  `json:"-" bson:"kdf"`
	WrappedVEK WrappedKey `json:"-" bson:"wrapped_vek"`

	Revision    int64      `json:"-" bson:"revision"`
	LastLoginAt *time.Time `json:"lastLoginAt,omitempty" bson:"last_login_at,omitempty"`
	CreatedAt   time.Time  `json:"createdAt" bson:"created_at"`
	UpdatedAt   time.Time  `json:"updatedAt" bson:"updated_at"`
}

type KDFConfig struct {
	Algorithm string    `json:"algorithm" bson:"algorithm" validate:"required,oneof=argon2id"`
	Salt      string    `json:"salt" bson:"salt" validate:"required,base64"`
	Params    KDFParams `json:"params" bson:"params"`
}

type KDFParams struct {
	MemoryKiB   uint32 `json:"memoryKiB" bson:"memory_kib" validate:"required,min=19456,max=1048576"`
	Iterations  uint32 `json:"iterations" bson:"iterations" validate:"required,min=2,max=10"`
	Parallelism uint8  `json:"parallelism" bson:"parallelism" validate:"required,min=1,max=16"`
	KeyLength   uint32 `json:"keyLength" bson:"key_length" validate:"required,eq=32"`
}

type WrappedKey struct {
	Version    uint8  `json:"version" bson:"version" validate:"required,min=1"`
	Nonce      string `json:"nonce" bson:"nonce" validate:"required,base64"`
	Ciphertext string `json:"ciphertext" bson:"ciphertext" validate:"required,base64"`
}

type RegisterRequest struct {
	Name       string     `json:"name" validate:"omitempty,max=100"`
	Username   string     `json:"username" validate:"required,min=3,max=32,alphanum"`
	Email      string     `json:"email" validate:"required,email,max=254"`
	AuthKey    string     `json:"authKey" validate:"required,base64,len=44"`
	KDF        KDFConfig  `json:"kdf"`
	WrappedVEK WrappedKey `json:"wrappedVek"`
}

type PreloginRequest struct {
	Identifier string `json:"identifier" validate:"required,max=254"`
}

type LoginRequest struct {
	Identifier string `json:"identifier" validate:"required,max=254"`
	AuthKey    string `json:"authKey" validate:"required,base64,len=44"`
}

type PreloginResponse struct {
	KDF KDFConfig `json:"kdf"`
}

type LoginResponse struct {
	AccessToken string       `json:"accessToken"`
	WrappedVEK  WrappedKey   `json:"wrappedVek"`
	User        UserResponse `json:"user"`
}

type UserResponse struct {
	ID            string `json:"id"`
	Name          string `json:"name,omitempty"`
	Username      string `json:"username"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"emailVerified"`
}
