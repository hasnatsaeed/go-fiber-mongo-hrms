package repositories

import (
	"context"
	"github/hasnatsaeed/go-fiber-mongo-hrms/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewEmployeeMongoDbRepository(col *mongo.Collection) Repository {
	return &repositoryMongoDb{
		col: col,
	}
}

type repositoryMongoDb struct {
	col *mongo.Collection
}

const employeeCollectionName = "employees"

func (repository *repositoryMongoDb) SaveEmployee(ctx context.Context, employee *models.Employee) (*models.Employee, error) {

	id, errInsert := repository.col.InsertOne(ctx, employee)
	if errInsert != nil {
		return nil, errInsert
	}

	insertedEmployee := &models.Employee{}
	errorFetch := repository.col.FindOne(ctx, bson.D{{Key: "_id", Value: id.InsertedID}}).Decode(insertedEmployee)

	if errorFetch != nil {
		return nil, errorFetch
	}

	return insertedEmployee, nil
}

func (repository *repositoryMongoDb) FetchEmployeeById(ctx context.Context, id string) (*models.Employee, error) {

	employee := &models.Employee{}

	errorFetch := repository.col.FindOne(ctx, bson.D{{"id", id}}).Decode(employee)

	if errorFetch != nil {
		return nil, errorFetch
	}
	return employee, nil
}

func (repository *repositoryMongoDb) FetchAllEmployees(ctx context.Context) ([]models.Employee, error) {

	cursor, errFindRecords := repository.col.Find(ctx, bson.D{{}})
	if errFindRecords != nil {
		return nil, errFindRecords
	}

	employees := make([]models.Employee, 10)

	errDecodeToEmployees := cursor.All(ctx, &employees)
	if errDecodeToEmployees != nil {
		return nil, errDecodeToEmployees
	}
	return employees, nil
}
func (repository *repositoryMongoDb) UpdateEmployee(ctx context.Context, id string, employeeToUpdate *models.Employee) (*models.Employee, error) {

	update := bson.D{{"$set", bson.D{
		{"name", employeeToUpdate.Name},
		{"salary", employeeToUpdate.Salary},
		{"age", employeeToUpdate.Age},
	}}}

	_, errUpdate := repository.col.UpdateOne(ctx, bson.D{{"id", id}}, update)
	if errUpdate != nil {
		return nil, errUpdate
	}
	return employeeToUpdate, errUpdate
}
func (repository *repositoryMongoDb) DeleteEmployee(ctx context.Context, id string) (*models.Employee, error) {

	employee, errFetch := repository.FetchEmployeeById(ctx, id)
	if errFetch != nil {
		return nil, errFetch
	}

	_, errDelete := repository.col.DeleteOne(ctx, employee)
	return employee, errDelete
}
