package answers

import (
	ppmongo "common/infrastructure/db/ppmongo"
)

// --------------------------------------
// Ropository of specific Entity
// --------------------------------------
type AnswersMongoRepository struct {
	*ppmongo.MongoRepository[AnswerModel, AnswerListModel]
}

func NewAnswersMongoRepository(uri string, dbName string, collectionName string) *AnswersMongoRepository {
	return &AnswersMongoRepository{
		MongoRepository: ppmongo.NewMongoRepository[AnswerModel, AnswerListModel](uri, dbName, collectionName),
	}
}
