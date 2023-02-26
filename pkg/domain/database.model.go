package domain

import "gorm.io/gorm"

// user schema for user table to get listed all users
type User struct {
	// gorm.Model

	IdLogin      int    `json:"-" gorm:"primaryKey;autoIncrement:true;unique"`
	UserName     string `json:"username" gorm:"not null;unique"`
	Password     string `json:"password"`
	UserType     string `json:"-" postgres:"type:ENUM('admin', 'worker', 'user')" gorm:"not null"`
	Verification string `json:"-" gorm:"default:false"`
	Status       string `json:"-" gorm:"default:newuser"`
}
type Profile struct {
	IdUser        int    `json:"-" gorm:"primaryKey;autoIncrement:true;unique"`
	LoginId       int    `json:"-" gorm:"unique"`
	Login         User   `json:"-" gorm:"foreignKey:LoginId;references:IdLogin"`
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
	Code  string `json:"code"`
}

type Category struct {
	IdCategory int    `json:"-" gorm:"primaryKey;autoIncrement:true;unique"`
	Category   string `gorm:"unique" json:"category"`
}

type Job struct {
	IdJob       int      `json:"-" gorm:"primaryKey;autoIncrement:true;unique"`
	CategoryId  int      `json:"categoryid" gorm:"not null"`
	Category    Category `json:"-" gorm:"foreignKey:CategoryId;references:IdCategory"`
	IdWorker    int      `json:"-" gorm:"not null"`
	Login       User     `json:"-" gorm:"foreignKey:IdWorker;references:IdLogin"`
	Wage        int      `json:"wage" gorm:"not null"`
	Description string   `json:"desctription"`
}

type Favorite struct {
	IdFavorite int  `json:"-" gorm:"primaryKey;autoIncrement:true;unique"`
	UserId     int  `json:"-"`
	User       User `json:"-" gorm:"foreignKey:UserId;references:IdLogin"`
	JobId      int  `json:"jobid" `
	job        Job  `json:"-" gorm:"foreignKey:JobId;references:IdJob;unique"`
}

type Address struct {
	IdAddress int `gorm:"primaryKey;autoIncrement:true;unique"`
	UserId    int
	User      User `json:"-" gorm:"foreignKey:UserId;references:IdLogin"`
	HouseName string
	Place     string
	City      string
	Post      string
	Pin       string
	Phone     string
}

type Request struct {
	IdRequset int `json:"-" gorm:"primaryKey;autoIncrement:true;unique"`
	UserId    int
	User      User `json:"-" gorm:"foreignKey:UserId;references:IdLogin"`
	JobId     int
	Job       Job `json:"-" gorm:"foreignKey:JobId;references:IdJob"`
	AddressId int
	Address   Address `json:"-" gorm:"foreignKey:AddressId;references:IdAddress"`
	Status    string  `json:"-" gorm:"default:pending"`
	Date      string  `jsom:"-"`
}

type JobPayment struct {
	IdPayment int
	UserId    int
	WorkerId  int
	Amount    int
	Date      string
}
