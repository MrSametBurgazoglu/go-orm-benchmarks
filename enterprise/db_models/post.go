package db_models

import (
	"github.com/MrSametBurgazoglu/enterprise/models"
)

func Post() *models.Table {
	idField := models.UintField("ID").AddSerial()

	tb := &models.Table{
		Fields: []models.FieldI{
			idField,
			models.StringField("Title"),
			models.StringField("Content"),
			models.UintField("UserID"),
		},
		Relations: []*models.Relation{
			models.ManyToOne("Users", idField.DBName, "user_id"),
			models.OneToMany("Comments", "id", "post_id"),
		},
	}

	tb.SetTableName("Posts")
	tb.SetIDField(idField)

	return tb
}
