package forms

import (
	ppmongo "common/infrastructure/db/ppmongo"
)

// --------------------------------------
// Ropository of specific Entity
// --------------------------------------
type FormsMongoRepository struct {
	*ppmongo.MongoRepository[FormModel]
}

func NewFormsMongoRepository(uri string, dbName string, collectionName string) *FormsMongoRepository {
	return &FormsMongoRepository{
		MongoRepository: ppmongo.NewMongoRepository[FormModel](uri, dbName, collectionName),
	}
}
