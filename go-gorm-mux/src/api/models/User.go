package models

import (
	"errors"
	"html"
	"log"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Name  	  string    `gorm:"size:255;not null;unique" json:"name"`
	Email     string    `gorm:"size:100;not null;unique" json:"email"`
	Password  string    `gorm:"size:100;not null;" json:"password"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *User) BeforeSave() error {
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) Prepare() {
	u.ID = 0
	u.Name = html.EscapeString(strings.TrimSpace(u.Name))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.Name == "" {
			return errors.New("Required Name")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}

		return nil
	case "login":
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil

	default:
		if u.Name == "" {
			return errors.New("Required Name")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil
	}
}

func (u *User) SaveUser(db *gorm.DB) (*User, error) {
	err := db.Debug().Create(&u).Error
	
	if err != nil {
		return &User{}, err
	}

	return u, nil
}

func (u *User) FindAllUsers(db *gorm.DB) (*[]User, error) {
	users := []User{}
	err := db.Debug().Model(&User{}).Limit(100).Find(&users).Error
	
	if err != nil {
		return &[]User{}, err
	}
	
	return &users, err
}

func (u *User) FindUserByID(db *gorm.DB, uid uint32) (*User, error) {
	err := db.Debug().Model(User{}).Where("id = ?", uid).Take(&u).Error
	
	if err != nil {
		return &User{}, err
	}
	
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &User{}, errors.New("User Not Found")
	}
	
	return u, err
}

func (u *User) UpdateUser(db *gorm.DB, uid uint32) (*User, error) {
	// To hash the password
	err := u.BeforeSave()
	if err != nil {
		log.Fatal(err)
	}

	var user User

	if err := db.Where("id =?", uid).First(&user).Error; err != nil {
		log.Fatal("User not found")
		return nil, err
	}

	if err := db.Model(&user).Updates(User{Name: u.Name, Password: u.Password, Email: u.Email}).Error; err != nil {
		log.Fatal("Can't update user")
		return nil, err
	}

	// db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).UpdateColumns(
	// 	map[string]interface{}{
	// 		"password":  	u.Password,
	// 		"name":  	 	u.Name,
	// 		"email":     	u.Email,
	// 		"updated_at": 	time.Now(),
	// 	},
	// )

	// if db.Error != nil {
	// 	return &User{}, db.Error
	// }

	// err = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&u).Error
	// if err != nil {
	// 	return &User{}, err
	// }

	return &user, nil
}

func (u *User) DeleteUser(db *gorm.DB, uid uint32) (int64, error) {
	var user User

	if res := db.Exec("PRAGMA foreign_keys = OFF", nil); res.Error != nil {
		log.Fatal(res.Error)
		return 0, res.Error
	}

	if err := db.Where("id = ?", uid).First(&user).Error; err != nil {
		return 0, err
	}

	if err := db.Delete(&user).Error; err != nil {
		return 0, err
	}
	
	return db.RowsAffected, nil
}
