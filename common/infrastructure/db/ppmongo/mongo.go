package mongo

import (
	"common/domain"
	"common/domain/logger"
	"common/utils"
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"common/utils/cerrs"
)

// --------------------------------------
// Ropository of specific Infra
// --------------------------------------
// MongoRepository es una implementación genérica de Repository usando generics.
type MongoRepository[T domain.IEntity] struct {
	Client     *mongo.Client
	Database   *mongo.Database
	Collection *mongo.Collection
}

// NewMongoRepository crea una nueva instancia de MongoRepository.
func NewMongoRepository[T domain.IEntity](uri string, dbName string, collectionName string) *MongoRepository[T] {
	// Context con timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Configurar las opciones del cliente
	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	// Crear referencia a la base de datos y la colección
	database := client.Database(dbName)
	collection := database.Collection(collectionName)

	return &MongoRepository[T]{
		Client:     client,
		Database:   database,
		Collection: collection,
	}
}

func (m *MongoRepository[T]) Save(ctx context.Context, document T) utils.Result[string] {
	result, err := m.Collection.InsertOne(ctx, document)
	if err != nil {
		return utils.Result[string]{Err: cerrs.NewCustomError(http.StatusInternalServerError, err.Error(), "mongo.save")}
	}

	// Convertir el InsertedID a primitive.ObjectID
	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return utils.Result[string]{Err: cerrs.NewCustomError(http.StatusInternalServerError, "el InsertedID no es un ObjectID válido", "mongo.save")}
	}

	// Retornar el ID como string
	return utils.Result[string]{Data: insertedID.Hex()}
}

// InsertDocumentWithID inserta un documento en MongoDB usando un UUID como _id
func (m *MongoRepository[T]) SaveWithID(ctx context.Context, id string, document T) utils.Result[string] {
	// Convertir el documento a BSON
	data, err := bson.Marshal(document)
	if err != nil {
		return utils.Result[string]{Err: cerrs.NewCustomError(http.StatusInternalServerError, err.Error(), "mongo.save_with_id")}
	}

	// Deserializar el BSON a un mapa para poder modificarlo
	var docMap bson.M
	err = bson.Unmarshal(data, &docMap)
	if err != nil {
		return utils.Result[string]{Err: cerrs.NewCustomError(http.StatusInternalServerError, err.Error(), "mongo.save_with_id")}
	}

	// Establecer el _id con el valor proporcionado
	docMap["_id"] = id

	// Insertar el documento modificado en la colección
	_, err = m.Collection.InsertOne(ctx, docMap)
	if err != nil {
		return utils.Result[string]{Err: cerrs.NewCustomError(http.StatusInternalServerError, err.Error(), "mongo.save_with_id")}
	}

	// Retornar el id como confirmación
	return utils.Result[string]{Data: id}
}

// Update actualiza un documento existente en la colección usando el id obtenido del entity.
func (m *MongoRepository[T]) Update(ctx context.Context, entity T) error {
	// Obtener el id desde el entity
	id := entity.GetID()

	// Convertir la entidad a BSON para poder actualizarla
	data, err := bson.Marshal(entity)
	if err != nil {
		return cerrs.NewCustomError(http.StatusInternalServerError, err.Error(), "mongo.update")
	}

	// Deserializar el BSON a un mapa para poder utilizarlo en la actualización
	var docMap bson.M
	err = bson.Unmarshal(data, &docMap)
	if err != nil {
		return cerrs.NewCustomError(http.StatusInternalServerError, err.Error(), "mongo.update")
	}

	// Crear el filtro basado en el _id
	filter := bson.M{"_id": id}

	// Realizar la actualización usando $set para reemplazar los campos existentes
	update := bson.M{"$set": docMap}
	result, err := m.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return cerrs.NewCustomError(http.StatusInternalServerError, err.Error(), "mongo.update")
	}

	// Verificar que se haya encontrado y actualizado el documento
	if result.MatchedCount == 0 {
		return cerrs.NewCustomError(http.StatusInternalServerError, fmt.Errorf("no se encontró documento con id %s", id).Error(), "mongo.update")
	}

	return nil
}

// UpdateFields recibe el ID en hex string, aplica los updates y devuelve el objeto actualizado.
func (m *MongoRepository[T]) UpdateFields(ctx context.Context, id string, updates map[string]interface{}) utils.Result[T] {

	entry, done := logger.FromContextWithExit(ctx)
	defer done()
	entry.Info("Updating fields for document")

	var updated T

	// 1. Convertir hex string a ObjectID
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		entry.Error("Error al convertir el ID a ObjectID", err)
		return utils.Result[T]{Err: cerrs.NewCustomError(http.StatusInternalServerError, fmt.Errorf("id inválido (%s): %w", id, err).Error(), "mongo.update_fields")}
	}

	// 2. Validar campos a actualizar
	if len(updates) == 0 {
		entry.Error("No se proporcionaron campos para actualizar")
		return utils.Result[T]{Err: cerrs.NewCustomError(http.StatusInternalServerError, fmt.Errorf("no se proporcionaron campos para actualizar").Error(), "mongo.update_fields")}
	}

	// 3. Preparar filtro y documento de actualización
	filter := bson.M{"_id": oid}
	update := bson.M{"$set": updates}

	// 4. Opciones: devolver el documento *después* de aplicar el update
	opts := options.FindOneAndUpdate().
		SetReturnDocument(options.After).
		SetUpsert(false) // evita crear uno nuevo si no existe

	// 5. Ejecutar FindOneAndUpdate y decodificar en `updated`
	err = m.Collection.FindOneAndUpdate(ctx, filter, update, opts).Decode(&updated)
	if err != nil {
		entry.Error("Error al actualizar el documento", err)
		if err == mongo.ErrNoDocuments {
			entry.Errorf("no se encontró ningún documento con id %s", id)
			return utils.Result[T]{Err: cerrs.NewCustomError(http.StatusInternalServerError, fmt.Errorf("no se encontró ningún documento con id %s", id).Error(), "mongo.update_fields")}
		}
		return utils.Result[T]{Err: cerrs.NewCustomError(http.StatusInternalServerError, fmt.Errorf("error al devolver el documento actualizado: %w", err).Error(), "mongo.update_fields")}
	}

	return utils.Result[T]{Data: updated}
}

// Delete elimina un documento de la colección usando el id proporcionado.
func (m *MongoRepository[T]) Delete(ctx context.Context, id string) error {
	// Crear el filtro usando _id
	filter := bson.M{"_id": id}

	// Ejecutar la eliminación
	result, err := m.Collection.DeleteOne(ctx, filter)
	if err != nil {
		return cerrs.NewCustomError(http.StatusInternalServerError, fmt.Errorf("error al eliminar el documento: %w", err).Error(), "mongo.delete")
	}

	// Verificar que se haya eliminado algún documento
	if result.DeletedCount == 0 {
		return cerrs.NewCustomError(http.StatusInternalServerError, fmt.Errorf("no se encontró documento con id %s", id).Error(), "mongo.delete")
	}

	return nil
}

func (m *MongoRepository[T]) Find(ctx context.Context, id string) utils.Result[T] {
	var result T

	// Crear el filtro basado en el _id
	filter := bson.M{"_id": id}

	// Ejecutar la búsqueda en la colección
	err := m.Collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return utils.Result[T]{Err: cerrs.NewCustomError(http.StatusInternalServerError, fmt.Errorf("no se encontró documento con id %s", id).Error(), "mongo.find")}
		}
		return utils.Result[T]{Err: cerrs.NewCustomError(http.StatusInternalServerError, fmt.Errorf("error al buscar el documento: %w", err).Error(), "mongo.find")}
	}

	return utils.Result[T]{Data: result}
}
