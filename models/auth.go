package models

type Auth struct {
	ID       uint   `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func GetAuth(username, password string) (auth Auth, err bool) {

	maps := make(map[string]interface{})
	maps["username"] = username
	maps["password"] = password

	db.Where(maps).First(&auth)

	if auth.ID > 0 {
		return auth, true
	}

	return auth, false
}

func (Auth) TableName() string {
	return "blog_auth"
}
