package models

type User struct {
    ID       uint   `json:"id"`
    Username string `json:"username"`
    Password string `json:"password"`
}

// Exported Users map to simulate a database
var Users = make(map[string]User)

var DB *gorm.DB