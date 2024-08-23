package benchmarks_test

import (
	"context"
	"fmt"
	"github.com/FournyP/go-orm-benchmarks/enterprise/models"
	"testing"
)

func BenchmarkEnterpriseCreate(b *testing.B) {
	client, err := GetEnterpriseClient()
	if err != nil {
		b.Fatalf("Error getting Enterprise client: %v", err)
	}

	err = TruncateEnterpriseDatabase(client)
	if err != nil {
		b.Fatalf("Error truncating Enterprise database: %v", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		user := models.NewUsers(context.TODO(), client)
		user.SetName(fmt.Sprintf("User%d", i))
		user.SetEmail(fmt.Sprintf("user%d@example.test", i))
		err = user.Create()
		if err != nil {
			b.Error(err)
			continue
		}

		post := models.NewPosts(context.TODO(), client)
		post.SetUserID(user.GetID())
		post.SetTitle(fmt.Sprintf("Post%d", i))
		post.SetContent(fmt.Sprintf("Body of post %d", i))
		err = post.Create()
		if err != nil {
			b.Error(err)
			continue
		}

		comment := models.NewComments(context.TODO(), client)
		comment.SetPostID(post.GetID())
		comment.SetText(fmt.Sprintf("Comment on post %d", i))
		err = comment.Create()
		if err != nil {
			b.Error(err)
			continue
		}
	}
}

func BenchmarkEnterpriseUpdate(b *testing.B) {
	client, err := GetEnterpriseClient()
	if err != nil {
		b.Fatalf("Error getting Enterprise client: %v", err)
	}

	err = TruncateEnterpriseDatabase(client)
	if err != nil {
		b.Fatalf("Error truncating Enterprise database: %v", err)
	}

	user := models.NewUsers(context.Background(), client)
	user.SetName("User%d")
	user.SetEmail("user%d@example.test")
	err = user.Create()
	if err != nil {
		b.Fatalf("Error creating user entity: %v", err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		user.SetName(fmt.Sprintf("User%d", i))
		user.SetEmail(fmt.Sprintf("user%d@example.test", i))
		err = user.Update()
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkEnterpriseDelete(b *testing.B) {
	db, err := GetEnterpriseClient()
	if err != nil {
		b.Fatalf("Error getting Enterprise client: %v", err)
	}

	err = TruncateEnterpriseDatabase(db)
	if err != nil {
		b.Fatalf("Error truncating GORM database: %v", err)
	}

	users := make([]*models.Users, b.N)

	for i := 0; i < b.N; i++ {
		user := models.NewUsers(context.TODO(), db)
		user.SetName(fmt.Sprintf("User%d", i))
		user.SetEmail(fmt.Sprintf("user%d@example.test", i))
		err = user.Create()
		if err != nil {
			b.Error(err)
			continue
		}

		if err != nil {
			b.Fatalf("Error creating user entity: %v", err)
		}

		users[i] = user
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := users[i].Delete()
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkEnterpriseRead(b *testing.B) {
	client, err := GetEnterpriseClient()
	if err != nil {
		b.Fatalf("Error getting Enterprise client: %v", err)
	}

	err = TruncateEnterpriseDatabase(client)
	if err != nil {
		b.Fatalf("Error truncating Enterprise database: %v", err)
	}

	user := models.NewUsers(context.Background(), client)
	user.SetName("User")
	user.SetEmail("user%d@example.test")
	err = user.Create()
	if err != nil {
		b.Fatalf("Error creating user entity: %v", err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		newUser := models.NewUsers(context.Background(), client)
		newUser.Where(newUser.IsNameEqual("User"))
		err, _ = newUser.Get()
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkEnterpriseReadWithRelations(b *testing.B) {
	client, err := GetEnterpriseClient()
	if err != nil {
		b.Fatalf("Error getting Enterprise client: %v", err)
	}

	err = TruncateEnterpriseDatabase(client)
	if err != nil {
		b.Fatalf("Error truncating Enterprise database: %v", err)
	}

	user := models.NewUsers(context.TODO(), client)
	user.SetName("User")
	user.SetEmail("user@example.test")
	err = user.Create()
	if err != nil {
		b.Fatalf("Error creating user entity: %v", err)
	}

	post := models.NewPosts(context.TODO(), client)
	post.SetUserID(1)
	post.SetTitle("Post")
	post.SetContent("Body of post")
	err = post.Create()
	if err != nil {
		b.Fatalf("Error creating post entity: %v", err)
	}

	comment := models.NewComments(context.TODO(), client)
	comment.SetPostID(1)
	comment.SetText("Comment")
	err = comment.Create()
	if err != nil {
		b.Fatalf("Error creating comment entity: %v", err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		newUser := models.NewUsers(context.Background(), client)
		newUser.Where(newUser.IsNameEqual("User"))
		newUser.WithPostsList(func(postList *models.PostsList) {
			postList.WithCommentsList()
		})
		err, _ = newUser.Get()
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkEnterpriseReadSingleField(b *testing.B) {
	db, err := GetEnterpriseClient() // Assume this gets your *gorm.DB
	if err != nil {
		b.Fatalf("Error getting Enterprise client: %v", err)
	}

	// Truncate and setup database as before
	err = TruncateEnterpriseDatabase(db)
	if err != nil {
		b.Fatalf("Error truncating Enterprise database: %v", err)
	}

	// Create a user entity
	user := models.NewUsers(context.TODO(), db)
	user.SetName("User")
	user.SetEmail("user@example.test")
	err = user.Create()
	if err != nil {
		b.Fatalf("Error creating user entity: %v", err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// Query for the 'name' field of the user
		newUser := models.NewUsers(context.TODO(), db)
		newUser.Where(newUser.IsEmailEqual("user@example.test"))
		newUser.GetSelector().SelectName()
		if err, _ := newUser.Get(); err != nil {
			b.Error(err)
		}
	}
}
