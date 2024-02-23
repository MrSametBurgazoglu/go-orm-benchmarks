package benchmarks_test

import (
	"fmt"
	"testing"

	"github.com/FournyP/go-orm-benchmarks/gormmodels"
)

func BenchmarkGORMCreate(b *testing.B) {
	db, err := GetGORMClient()
	if err != nil {
		b.Fatalf("Error getting GORM client: %v", err)
	}

	err = TruncateGORMDatabase(db)
	if err != nil {
		b.Fatalf("Error truncating GORM database: %v", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		user := gormmodels.User{
			Name:  fmt.Sprintf("User %d", i+1),
			Email: fmt.Sprintf("user%d@example.test", i+1),
		}
		err = db.Create(&user).Error

		if err != nil {
			b.Error(err)
			continue
		}

		post := gormmodels.Post{
			Title:   fmt.Sprintf("Post %d", i+1),
			Content: fmt.Sprintf("Post %d content", i+1),
			UserID:  user.ID,
		}
		err = db.Create(&post).Error

		if err != nil {
			b.Error(err)
			continue
		}

		comment := gormmodels.Comment{
			Text:   fmt.Sprintf("Comment %d", i+1),
			PostID: post.ID,
		}
		err = db.Create(&comment).Error

		if err != nil {
			b.Error(err)
			continue
		}
	}
}

func BenchmarkGORMUpdate(b *testing.B) {
	db, err := GetGORMClient()
	if err != nil {
		b.Fatalf("Error getting GORM client: %v", err)
	}

	err = TruncateGORMDatabase(db)
	if err != nil {
		b.Fatalf("Error truncating Ent database: %v", err)
	}

	user := gormmodels.User{
		Name:  "User",
		Email: "user@example.test",
	}
	err = db.Create(&user).Error
	if err != nil {
		b.Fatalf("Error creating user entity: %v", err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := db.Model(&user).Updates(gormmodels.User{
			Name:  fmt.Sprintf("User%d", i),
			Email: fmt.Sprintf("user%d@example.test", i),
		}).Error
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkGORMDelete(b *testing.B) {
	db, err := GetGORMClient()
	if err != nil {
		b.Fatalf("Error getting GORM client: %v", err)
	}

	err = TruncateGORMDatabase(db)
	if err != nil {
		b.Fatalf("Error truncating GORM database: %v", err)
	}

	// Pre-create users to delete during the benchmark
	users := make([]gormmodels.User, b.N)
	for i := 0; i < b.N; i++ {
		user := gormmodels.User{Name: fmt.Sprintf("User%d", i), Email: fmt.Sprintf("user%d@example.test", i)}
		if err := db.Create(&user).Error; err != nil {
			b.Fatalf("Error creating user entity: %v", err)
		}
		users[i] = user
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := db.Delete(&users[i]).Error
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkGORMRead(b *testing.B) {
	db, err := GetGORMClient()
	if err != nil {
		b.Fatalf("Error getting GORM client: %v", err)
	}

	err = TruncateGORMDatabase(db)
	if err != nil {
		b.Fatalf("Error truncating GORM database: %v", err)
	}

	// Create a user for the benchmark
	if err := db.Create(&gormmodels.User{Name: "User", Email: "user@example.test"}).Error; err != nil {
		b.Fatalf("Error creating user entity: %v", err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var user gormmodels.User
		// Query for the user by name
		if err := db.Where("name = ?", "User").First(&user).Error; err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkGORMReadWithRelations(b *testing.B) {
	db, err := GetGORMClient()
	if err != nil {
		b.Fatalf("Error getting GORM client: %v", err)
	}

	err = TruncateGORMDatabase(db)
	if err != nil {
		b.Fatalf("Error truncating GORM database: %v", err)
	}

	// Assuming creation functions similar to Ent setup
	user := gormmodels.User{Name: "User", Email: "user@example.test"}
	if err := db.Create(&user).Error; err != nil {
		b.Fatalf("Error creating user entity: %v", err)
	}

	post := gormmodels.Post{Title: "Post", Content: "Body of post", UserID: user.ID}
	if err := db.Create(&post).Error; err != nil {
		b.Fatalf("Error creating post entity: %v", err)
	}

	comment := gormmodels.Comment{Text: "Comment", PostID: post.ID}
	if err := db.Create(&comment).Error; err != nil {
		b.Fatalf("Error creating comment entity: %v", err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var user gormmodels.User
		// Query for the user by name and preload posts and their comments
		if err := db.Preload("Posts.Comments").Where("name = ?", "User").First(&user).Error; err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkGormReadSingleField(b *testing.B) {
	db, err := GetGORMClient() // Assume this gets your *gorm.DB
	if err != nil {
		b.Fatalf("Error getting GORM client: %v", err)
	}

	// Truncate and setup database as before
	err = TruncateGORMDatabase(db)
	if err != nil {
		b.Fatalf("Error truncating GORM database: %v", err)
	}

	// Create a user entity
	if err := db.Create(&gormmodels.User{Name: "User", Email: "user@example.test"}).Error; err != nil {
		b.Fatalf("Error creating user entity: %v", err)
	}

	// Define a struct to hold the result of the single-field query
	type UserName struct {
		Name string
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var userName UserName
		// Query for the 'name' field of the user
		if err := db.Model(&gormmodels.User{}).Select("name").Where("email = ?", "user@example.test").First(&userName).Error; err != nil {
			b.Error(err)
		}
	}
}
