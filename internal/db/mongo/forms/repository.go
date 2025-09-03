package forms

import (
	ppmongo "common/infrastructure/db/ppmongo"
)

// --------------------------------------
// Ropository of specific Entity
// --------------------------------------
type FormsMongoRepository struct {
	*ppmongo.MongoRepository[FormModel, FormListModel]
}

func NewFormsMongoRepository(uri string, dbName string, collectionName string) *FormsMongoRepository {
	return &FormsMongoRepository{
		MongoRepository: ppmongo.NewMongoRepository[FormModel, FormListModel](uri, dbName, collectionName),
	}
}
