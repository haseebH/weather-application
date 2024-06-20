package repository

type User struct {
	Name     string `json:"name" bson:"name"`
	Email    string `json:"email" bson:"email"`
	Token    string `json:"token" bson:"-"`
	Location string `json:"location" bson:"location"`
}
