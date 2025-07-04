package def

type User struct {
	Name  string
	Email string
	Token string
}

func (u *User) String() string {
	return "Name: " + u.Name + " Email: " + u.Email + " Token: " + u.Token
}

type Users []User

func (u Users) String() string {
	s := ""
	for _, u := range u {
		s += u.String() + "\n"
	}
	return s
}

// type UserCreationIngredient struct {
// 	Name     string `json:"name" bson:"name"`
// 	Email    string `json:"email"`
// 	Password string `json:"password"`
// }

type UserCreationIngredient struct {
	Name     string
	Email    string
	Password string
}

type UserLogin struct {
	Name     string
	Password string
}
