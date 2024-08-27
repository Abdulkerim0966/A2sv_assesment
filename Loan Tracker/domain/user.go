package domain

import "time"

// User represents a user in the system
type User struct {
	FirstName  string    `json:"firstname" bson:"firstname"`
	LastName   string    `json:"lastname" bson:"lastname"`
	Bio        string    `json:"bio" bson:"bio"`
	Username   string    `json:"username" bson:"username" binding:"required"`
	Password   string    `json:"password" bson:"password" binding:"required"`
	Email      string    `json:"email" bson:"email" binding:"required"`
	Role       string    `json:"role" bson:"role"`
	Address    string    `json:"address" bson:"address"`
	JoinedDate time.Time `json:"joined_date" bson:"joined_date"`
}

type UserRepository interface {
	RegisterUser(user *User) error
	CheckUsernameAndEmail(username, email string) error
	GetUserByUsernameorEmail(usernameoremail string) (*User, error)
	InsertToken(token *Token) error
	GetTokenByUsername(username string) (*Token, error)
	DeleteToken(username string) error
	DeleteUser(username string) error
	Resetpassword(username, password string) error
	UpdateProfile(usernameoremail string, user *User) error
	GetUsers() ([]User, error)




}
type UserUsecase interface {
	RegisterUser(user *User) error
	CheckUsernameAndEmail(username, email string) error
	VerifyUser(token string) error
	LoginUser(usernameoremail string, password string) (string, string, error)
	RefreshToken(claims *LoginClaims) (string, error)
	ResetPassword(tokenString string) error
	UpdateProfile(user *User, claims *LoginClaims) error
	DeleteUser(username string, claim *LoginClaims) error
	
	ForgotPassword(email string, newPassword string) error
	LogoutUser(username string) error
	ChangePassword(username, oldPassword, newPassword string) error
	GetUsers(claims *LoginClaims) ([]User, error)
	GetUserByUsername(username string) (*User, error)
	
	
}

type Token struct {
	Username  string `json:"username" bson:"username"`
	ExpiresAt int64  `json:"expires_at" bson:"expires_at"`
}

type OAuthState struct {
	ID        string    `bson:"_id"`
	CreatedAt time.Time `bson:"created_at"`
	ExpiresAt time.Time `bson:"expires_at"`
}
type OAuthStateRepository interface {
	InsertState(state *OAuthState) error
	GetState(stateString string) (*OAuthState, error)
	DeleteState(state *OAuthState) error
}
