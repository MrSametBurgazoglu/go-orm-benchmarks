package db_models

import (
	"github.com/MrSametBurgazoglu/enterprise/models"
)

func User() *models.Table {
	idField := models.UintField("ID").AddSerial()

	tb := &models.Table{
		Fields: []models.FieldI{
			idField,
			models.StringField("Name"),
			models.StringField("Email"),
		},
		Relations: []*models.Relation{
			models.OneToMany("Posts", "id", "user_id"),
		},
	}

	tb.SetTableName("Users")
	tb.SetIDField(idField)

	return tb
}
