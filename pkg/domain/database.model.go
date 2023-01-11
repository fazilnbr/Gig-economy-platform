package domain

import "gorm.io/gorm"

// user schema for user table to get listed all users
type Login struct {
	// gorm.Model

	IdLogin      int    `json:"id_login" gorm:"primaryKey;autoIncrement:true;unique"`
	UserName     string `json:"username" gorm:"not null;unique"`
	Password     string `json:"password"`
	UserType     string `json:"usertype" postgres:"type:ENUM('admin', 'worker', 'user')" gorm:"not null"`
	Verification string `json:"verification" gorm:"default:false"`
	Status       string `json:"status" gorm:"default:newuser"`
}
type Profile struct {
	IdUser        int    `gorm:"primaryKey;autoIncrement:true;unique"`
	IdLogin       int    `gorm:"unique"`
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

//to store mail verification details

type Verification struct {
	gorm.Model
	Email string `json:"email"`
	Code  int    `json:"code"`
}

type Category struct {
	IdCategory int    `gorm:"primaryKey;autoIncrement:true;unique"`
	Category   string `gorm:"unique" json:"category"`
}

type Job struct {
	IdJob      int    `gorm:"primaryKey;autoIncrement:true;unique"`
	IdCategory string `json:"categoryid"`
	IdWorker   string `json:"workerid"`
}
