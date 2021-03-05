
package auth

// User - struct
type User struct {
	ID string `json:"id"  bson:"_id"`
	UserName string `json:"user_name"  bson:"user_name"`
	Email string `json:"email"  bson:"email" binding:"required"`
	Password string `json:"password"  bson:"password" binding:"required"`
}

// Login struct
type Login struct {
	Email string `json:"email"  bson:"email" binding:"required"`
	Password string `json:"password"  bson:"password" binding:"required"`
}
