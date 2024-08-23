package db_models

import "github.com/MrSametBurgazoglu/enterprise/models"

func Comment() *models.Table {
	idField := models.UintField("ID").AddSerial()

	tb := &models.Table{
		Fields: []models.FieldI{
			idField,
			models.StringField("Text"),
			models.UintField("PostID"),
		},
		Relations: []*models.Relation{
			models.OneToMany("Posts", "id", "post_id"),
		},
	}

	tb.SetTableName("Comments")
	tb.SetIDField(idField)

	return tb
}
