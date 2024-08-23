package benchmarks_test

import (
	"context"
	"fmt"
	"github.com/FournyP/go-orm-benchmarks/ent"
	"github.com/FournyP/go-orm-benchmarks/enterprise/db_models"
	"github.com/FournyP/go-orm-benchmarks/enterprise/models"
	"github.com/FournyP/go-orm-benchmarks/gormmodels"
	"github.com/FournyP/go-orm-benchmarks/sqlxmodels"
	"github.com/MrSametBurgazoglu/enterprise/migrate"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	DATABASE_URL_FORMAT = "host=127.0.0.1 port=%d user=root dbname=main password=root sslmode=disable"
)

const (
	ENT_DATABASE_PORT  = 5432
	GORM_DATABASE_PORT = 5433
	SQLX_DATABASE_PORT = 5434
)

func GetEntClient() (*ent.Client, error) {
	entDatabaseUrl := fmt.Sprintf(DATABASE_URL_FORMAT, ENT_DATABASE_PORT)
	client, err := ent.Open("postgres", entDatabaseUrl)

	if err != nil {
		return nil, err
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		return nil, fmt.Errorf("failed creating schema resources: %v", err)
	}

	return client, nil
}

func GetGORMClient() (*gorm.DB, error) {
	gormDatabaseUrl := fmt.Sprintf(DATABASE_URL_FORMAT, GORM_DATABASE_PORT)
	dialector := postgres.Open(gormDatabaseUrl)

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(gormmodels.User{}, gormmodels.Post{}, gormmodels.Comment{})
	if err != nil {
		return nil, fmt.Errorf("failed creating schema resources: %v", err)
	}

	return db, nil
}

func GetSqlxClient() (*sqlx.DB, error) {
	// Replace with your actual database DSN
	sqlxDatabaseUrl := fmt.Sprintf(DATABASE_URL_FORMAT, SQLX_DATABASE_PORT)
	db, err := sqlx.Connect("postgres", sqlxDatabaseUrl)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(sqlxmodels.MIGRATION)
	if err != nil {
		return nil, fmt.Errorf("failed creating schema resources: %v", err)
	}

	return db, nil
}

func GetEnterpriseClient() (models.IDatabase, error) {
	dbUrl := "postgresql://root:root@127.0.0.1:5435/main?search_path=public&sslmode=disable"
	db, err := models.NewDB(dbUrl)
	if err != nil {
		panic(err)
	}
	err = DeleteEnterpriseDatabase(db)
	if err != nil {
		panic(err)
	}
	migrate.AutoApplyMigration(context.TODO(), dbUrl, "migration",
		db_models.Post(),
		db_models.Comment(),
		db_models.User())

	return db, nil
}

func TruncateEntDatabase(client *ent.Client) error {
	_, err := client.Comment.Delete().Exec(context.Background())
	if err != nil {
		return err
	}

	_, err = client.Post.Delete().Exec(context.Background())
	if err != nil {
		return err
	}

	_, err = client.User.Delete().Exec(context.Background())
	if err != nil {
		return err
	}

	return nil
}

func TruncateGORMDatabase(db *gorm.DB) error {
	err := db.Exec("TRUNCATE TABLE comments, posts, users").Error
	if err != nil {
		return err
	}

	return nil
}

func TruncateSqlxDatabase(db *sqlx.DB) error {
	_, err := db.Exec("TRUNCATE TABLE comments, posts, users")
	if err != nil {
		return err
	}

	return nil
}

func TruncateEnterpriseDatabase(db models.IDatabase) error {
	_, err := db.Exec(context.TODO(), "TRUNCATE TABLE \"comments\", \"posts\", \"users\"")
	if err != nil {
		return err
	}

	return nil
}

func DeleteEnterpriseDatabase(db models.IDatabase) error {
	_, err := db.Exec(context.TODO(), "DROP TABLE \"comments\", \"posts\", \"users\"")
	if err != nil {
		return err
	}

	return nil
}
