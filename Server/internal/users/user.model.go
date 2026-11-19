package models

type User struct {
	ID           string `json:"id" bson:"_id,omitempty"`
	Name         string `json:"name" bson:"name"`
	Username     string `json:"username" bson:"username"`
	Email        string `json:"email" bson:"email"`
	Password     string `json:"password" bson:"password"`
	AccessToken  string `json:"accesss_token" bson:"access_token"`
	RefreshToken string `json:"refresh_token" bson:"refresh_token"`
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
