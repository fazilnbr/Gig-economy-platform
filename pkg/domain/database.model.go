package domain

import "gorm.io/gorm"

// user schema for user table to get listed all users
type Login struct {
	// gorm.Model

	ID_Login     int    `json:"id_login" gorm:"primaryKey;autoIncrement:true;unique"`
	UserName     string `json:"username" gorm:"not null;unique"`
	Password     string `json:"password"`
	UserType     string `json:"usertype" postgres:"type:ENUM('admin', 'worker', 'user')" gorm:"not null"`
	Verification string `json:"verification" gorm:"default:false"`
	Status       string `json:"status" gorm:"default:newuser"`
}
type Profile struct {
	ID_User       int    `gorm:"primaryKey;autoIncrement:true;unique"`
	ID_Login      int    `gorm:"unique"`
	Name          string `json:"name"`
	Gender        string `json:"gender"`
	DateOfBirth   string `json:"dateofbirth"`
	HouseName     string `json:"housename"`
	Place         string `json:"place"`
	Post          string `json:"post"`
	Pin           string `json:"pin"`
	ContactNumber string `gorm:"unique" json:"contactnumber"`
	EmailID       string `gorm:"unique" json:"emailid"`
	Photo         string `json:"photo"`
}

// type Worker struct {
// 	ID_User       int    `gorm:"primaryKey;autoIncrement:true;unique"`
// 	ID_Login      int    `gorm:"unique"`
// 	Name          string `json:"name"`
// 	Gender        string `json:"gender"`
// 	DateOfBirth   string `json:"dateofbirth"`
// 	HouseName     string `json:"housename"`
// 	Place         string `json:"place"`
// 	Post          string `json:"post"`
// 	Pin           string `json:"pin"`
// 	ContactNumber string `gorm:"unique" json:"contactnumber"`
// 	EmailID       string `gorm:"unique" json:"emailid"`
// 	Photo         string `json:"photo"`
// }

//to store mail verification details

type Verification struct {
	gorm.Model
	Email string `json:"email"`
	Code  int    `json:"code"`
}

type Category struct {
	ID_category int    `gorm:"primaryKey;autoIncrement:true;unique"`
	Category    string `gorm:"unique" json:"contactnumber"`
}
