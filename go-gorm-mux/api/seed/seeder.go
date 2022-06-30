package seed

import (
	"go-gorm-mux/api/models"
	"log"

	"gorm.io/gorm"
)

// Represents the users sample data.
var users = []models.User{
	models.User{
		Name: 	  "User 1",
		Email:    "user1@gmail.com",
		Password: "password",
	},
	models.User{
		Name: 	  "User 2",
		Email:    "user2@gmail.com",
		Password: "password",
	},
}

// Represents the posts sample data.
var posts = []models.Post{
	models.Post{
		Title:   "Title 1",
		Content: "Hello world 1",
	},
	models.Post{
		Title:   "Title 2",
		Content: "Hello world 2",
	},
}

// Load represents seeding of the database.
func Load(db *gorm.DB) {
	err := db.Debug().Migrator().DropTable(&models.Post{}, &models.User{})

	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}

	err = db.Debug().AutoMigrate(&models.User{}, &models.Post{})
	
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	for i := range users {
		if r := db.Create(&users[i]); r.Error != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		posts[i].AuthorID = users[i].ID

		if r := db.Create(&posts[i]); r.Error != nil {
			log.Fatalf("cannot seed posts table: %v", err)
		}
	}
}
