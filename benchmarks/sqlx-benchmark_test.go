package benchmarks_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/FournyP/go-orm-benchmarks/sqlxmodels"
)

const (
	INSERT_USER_QUERY               = "INSERT INTO users (name, email) VALUES ($1, $2)"
	DELETE_USER_QUERY               = "DELETE FROM users WHERE id = $1"
	SELECT_USER_QUERY_BY_EMAIL      = "SELECT * FROM users WHERE email = $1"
	SELECT_USER_ID_QUERY_BY_EMAIL   = "SELECT id FROM users WHERE email = $1"
	INSERT_POST_QUERY               = "INSERT INTO posts (user_id, title, content) VALUES ($1, $2, $3)"
	SELECT_POST_QUERY_BY_USER_ID    = "SELECT * FROM posts WHERE user_id = $1"
	SELECT_POST_ID_QUERY_BY_USER_ID = "SELECT id FROM posts WHERE user_id = $1"
	INSERT_COMMENT_QUERY            = "INSERT INTO comments (post_id, text) VALUES ($1, $2)"
	SELECT_COMMENT_QUERY_BY_USER_ID = "SELECT * FROM comments WHERE post_id = $1"
)

func BenchmarkSqlxCreate(b *testing.B) {
	db, err := GetSqlxClient()
	if err != nil {
		log.Fatalf("Error getting SQLx DB: %v", err)
	}

	err = TruncateSqlxDatabase(db)
	if err != nil {
		log.Fatalf("Error truncating SQLx database: %v", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Create a user
		_, err := db.Exec(INSERT_USER_QUERY, fmt.Sprintf("User%d", i), fmt.Sprintf("user%d@example.test", i))
		if err != nil {
			b.Error("Failed to insert user:", err)
			continue
		}
		var userId int64
		err = db.Get(&userId, SELECT_USER_ID_QUERY_BY_EMAIL, fmt.Sprintf("user%d@example.test", i))
		if err != nil {
			b.Error("Failed to get user ID:", err)
			continue
		}

		// Create a post
		_, err = db.Exec(INSERT_POST_QUERY, userId, fmt.Sprintf("Post%d", i), fmt.Sprintf("Body of post %d", i))
		if err != nil {
			b.Error("Failed to insert post:", err)
			continue
		}
		var postId int64
		err = db.Get(&postId, SELECT_POST_ID_QUERY_BY_USER_ID, userId)
		if err != nil {
			b.Error("Failed to get post ID:", err)
			continue
		}

		// Create a comment
		_, err = db.Exec(INSERT_COMMENT_QUERY, postId, fmt.Sprintf("Comment on post %d", i))
		if err != nil {
			b.Error("Failed to insert comment:", err)
		}
	}
}

func BenchmarkSqlxUpdate(b *testing.B) {
	db, err := GetSqlxClient()
	if err != nil {
		log.Fatalf("Error getting SQLx DB: %v", err)
	}

	err = TruncateSqlxDatabase(db)
	if err != nil {
		log.Fatalf("Error truncating SQLx database: %v", err)
	}

	// Create a user to update
	_, err = db.Exec(INSERT_USER_QUERY, "User", "user@example.test")
	if err != nil {
		log.Fatalf("Error creating user entity: %v", err)
	}
	var userID uint
	err = db.Get(&userID, SELECT_USER_ID_QUERY_BY_EMAIL, "user@example.test")
	if err != nil {
		log.Fatalf("Error getting user ID: %v", err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := db.Exec(
			"UPDATE users SET name = $1, email = $2 WHERE id = $3",
			fmt.Sprintf("User%d", i),
			fmt.Sprintf("user%d@example.test", i),
			userID)

		if err != nil {
			b.Error("Failed to update user:", err)
		}
	}
}

func BenchmarkSqlxDelete(b *testing.B) {
	db, err := GetSqlxClient()
	if err != nil {
		log.Fatalf("Error getting SQLx DB: %v", err)
	}

	err = TruncateSqlxDatabase(db)
	if err != nil {
		log.Fatalf("Error truncating SQLx database: %v", err)
	}

	userIDs := make([]int64, b.N)
	for i := 0; i < b.N; i++ {
		// Create a user
		_, err := db.Exec(
			INSERT_USER_QUERY,
			fmt.Sprintf("User%d", i),
			fmt.Sprintf("user%d@example.test", i))
		if err != nil {
			log.Fatalf("Error creating user entity: %v", err)
		}

		var userId int64
		err = db.Get(&userId, SELECT_USER_ID_QUERY_BY_EMAIL, fmt.Sprintf("user%d@example.test", i))
		if err != nil {
			b.Error("Failed to get user ID:", err)
			continue
		}
		userIDs[i] = userId
	}

	b.ResetTimer()

	for _, id := range userIDs {
		// Delete the user
		_, err := db.Exec(DELETE_USER_QUERY, id)
		if err != nil {
			b.Error("Failed to delete user:", err)
		}
	}
}

func BenchmarkSqlxRead(b *testing.B) {
	db, err := GetSqlxClient()
	if err != nil {
		log.Fatalf("Error getting SQLx DB: %v", err)
	}

	err = TruncateSqlxDatabase(db)
	if err != nil {
		log.Fatalf("Error truncating SQLx database: %v", err)
	}

	// Create a user for the benchmark
	_, err = db.Exec(INSERT_USER_QUERY, "User", "user@example.test")
	if err != nil {
		log.Fatalf("Error creating user entity: %v", err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var user sqlxmodels.User
		// Query for the user by name
		err := db.Get(&user, SELECT_USER_QUERY_BY_EMAIL, "user@example.test")
		if err != nil {
			b.Error("Failed to query user:", err)
		}
	}
}

func BenchmarkSqlxReadWithRelations(b *testing.B) {
	db, err := GetSqlxClient()
	if err != nil {
		b.Fatalf("Error getting SQLx DB: %v", err)
	}

	err = TruncateSqlxDatabase(db)
	if err != nil {
		b.Fatalf("Error truncating SQLx database: %v", err)
	}

	// INSERT USER
	_, err = db.Exec(INSERT_USER_QUERY, "User", "user@example.test")
	if err != nil {
		b.Fatalf("Failed to insert user: %v", err)
	}
	var userId int
	err = db.Get(&userId, SELECT_USER_ID_QUERY_BY_EMAIL, "user@example.test")
	if err != nil {
		b.Fatalf("Failed to get user ID: %v", err)
	}

	// INSERT POST
	_, err = db.Exec(INSERT_POST_QUERY, userId, "Post", "Body of post")
	if err != nil {
		b.Fatalf("Failed to insert post: %v", err)
	}
	var postId int64
	err = db.Get(&postId, SELECT_POST_ID_QUERY_BY_USER_ID, userId)
	if err != nil {
		b.Fatalf("Failed to get post ID: %v", err)
	}

	// INSERT COMMENT
	_, err = db.Exec(INSERT_COMMENT_QUERY, postId, "Comment")
	if err != nil {
		b.Fatalf("Failed to insert comment: %v", err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var user sqlxmodels.User
		err := db.Get(&user, SELECT_USER_QUERY_BY_EMAIL, "user@example.test")
		if err != nil {
			b.Error("Failed to fetch user:", err)
		}

		// Fetch user's posts
		err = db.Select(&user.Posts, SELECT_POST_QUERY_BY_USER_ID, user.ID)
		if err != nil {
			b.Error("Failed to fetch user:", err)
		}

		// For each post, fetch its comments
		for i := range user.Posts {
			post := &user.Posts[i]
			err = db.Select(&post.Comments, SELECT_COMMENT_QUERY_BY_USER_ID, post.ID)
			if err != nil {
				b.Error("Failed to fetch comments:", err)
			}
		}
	}
}

func BenchmarkSqlxReadSingleField(b *testing.B) {
	db, err := GetSqlxClient()
	if err != nil {
		b.Fatalf("Error getting SQLx DB: %v", err)
	}

	err = TruncateSqlxDatabase(db)
	if err != nil {
		b.Fatalf("Error truncating SQLx database: %v", err)
	}

	// Create a user entity
	_, err = db.Exec(INSERT_USER_QUERY, "User", "user@example.test")
	if err != nil {
		b.Fatalf("Error creating user entity: %v", err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var userId int64
		// Query for the 'name' field of the user
		err := db.Get(&userId, SELECT_USER_ID_QUERY_BY_EMAIL, "user@example.test")
		if err != nil {
			b.Error("Failed to query user name:", err)
		}
		// At this point, 'name' holds the name of the user
	}
}
