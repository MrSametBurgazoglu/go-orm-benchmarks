package main

import (
	"github.com/FournyP/go-orm-benchmarks/enterprise/db_models"
	"github.com/MrSametBurgazoglu/enterprise/generate"
)

func main() {
	generate.Models(
		db_models.Comment(),
		db_models.Post(),
		db_models.User(),
	)
}
