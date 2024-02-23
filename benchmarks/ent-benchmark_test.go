package benchmarks_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/FournyP/go-orm-benchmarks/ent"
	"github.com/FournyP/go-orm-benchmarks/ent/user"
)

func BenchmarkEntCreate(b *testing.B) {
	client, err := GetEntClient()
	if err != nil {
		b.Fatalf("Error getting Ent client: %v", err)
	}

	err = TruncateEntDatabase(client)
	if err != nil {
		b.Fatalf("Error truncating Ent database: %v", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		user, err := client.User.Create().
			SetName(fmt.Sprintf("User%d", i)).
			SetEmail(fmt.Sprintf("user%d@example.test", i)).
			Save(context.Background())

		if err != nil {
			b.Error(err)
			continue
		}

		post, err := client.Post.Create().
			SetUserID(user.ID).
			SetTitle(fmt.Sprintf("Post%d", i)).
			SetContent(fmt.Sprintf("Body of post %d", i)).
			Save(context.Background())

		if err != nil {
			b.Error(err)
			continue
		}

		_, err = client.Comment.Create().
			SetText(fmt.Sprintf("Comment on post %d", i)).
			SetPost(post).
			Save(context.Background())

		if err != nil {
			b.Error(err)
			continue
		}
	}
}

func BenchmarkEntUpdate(b *testing.B) {
	client, err := GetEntClient()
	if err != nil {
		b.Fatalf("Error getting Ent client: %v", err)
	}

	err = TruncateEntDatabase(client)
	if err != nil {
		b.Fatalf("Error truncating Ent database: %v", err)
	}

	user, err := client.User.Create().
		SetName("User%d").
		SetEmail("user%d@example.test").
		Save(context.Background())

	if err != nil {
		b.Fatalf("Error creating user entity: %v", err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := client.User.UpdateOneID(user.ID).
			SetName(fmt.Sprintf("User%d", i)).
			SetEmail(fmt.Sprintf("user%d@example.test", i)).
			Save(context.Background())

		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkEntDelete(b *testing.B) {
	client, err := GetEntClient()
	if err != nil {
		b.Fatalf("Error getting Ent client: %v", err)
	}

	err = TruncateEntDatabase(client)
	if err != nil {
		b.Fatalf("Error truncating Ent database: %v", err)
	}

	users := make([]*ent.User, b.N)

	for i := 0; i < b.N; i++ {
		user, err := client.User.Create().
			SetName(fmt.Sprintf("User%d", i)).
			SetEmail(fmt.Sprintf("user%d@example.test", i)).
			Save(context.Background())

		if err != nil {
			b.Fatalf("Error creating user entity: %v", err)
		}

		users[i] = user
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := client.User.DeleteOneID(users[i].ID).Exec(context.Background())
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkEntRead(b *testing.B) {
	client, err := GetEntClient()
	if err != nil {
		b.Fatalf("Error getting Ent client: %v", err)
	}

	err = TruncateEntDatabase(client)
	if err != nil {
		b.Fatalf("Error truncating Ent database: %v", err)
	}

	_, err = client.User.Create().
		SetName("User").
		SetEmail("user%d@example.test").
		Save(context.Background())

	if err != nil {
		b.Fatalf("Error creating user entity: %v", err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := client.User.Query().Where(user.NameEQ("User")).First(context.Background())
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkEntReadWithRelations(b *testing.B) {
	client, err := GetEntClient()
	if err != nil {
		b.Fatalf("Error getting Ent client: %v", err)
	}

	err = TruncateEntDatabase(client)
	if err != nil {
		b.Fatalf("Error truncating Ent database: %v", err)
	}

	userEntity, err := client.User.Create().
		SetName("User").
		SetEmail("user@example.test").
		Save(context.Background())

	if err != nil {
		b.Fatalf("Error creating user entity: %v", err)
	}

	postEntity, err := client.Post.Create().
		SetTitle("Post").
		SetUserID(userEntity.ID).
		SetContent("Body of post").
		Save(context.Background())

	if err != nil {
		b.Fatalf("Error creating post entity: %v", err)
	}

	_, err = client.Comment.Create().
		SetText("Comment").
		SetPostID(postEntity.ID).
		Save(context.Background())

	if err != nil {
		b.Fatalf("Error creating comment entity: %v", err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := client.User.Query().
			Where(user.NameEQ("User")).
			WithPosts(func(postQuery *ent.PostQuery) {
				postQuery.WithComments()
			}).
			First(context.Background())

		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkEntReadSingleField(b *testing.B) {
	client, err := GetEntClient()
	if err != nil {
		b.Fatalf("Error getting Ent client: %v", err)
	}

	err = TruncateEntDatabase(client)
	if err != nil {
		b.Fatalf("Error truncating Ent database: %v", err)
	}

	userEntity, err := client.User.Create().
		SetName("User").
		SetEmail("user@example.test").
		Save(context.Background())

	if err != nil {
		b.Fatalf("Error creating user entity: %v", err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := client.User.Query().
			Where(user.IDEQ(userEntity.ID)).
			Select("name").
			String(context.Background())

		if err != nil {
			b.Error(err)
		}
	}
}
