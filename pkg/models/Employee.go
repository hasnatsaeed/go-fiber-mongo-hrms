package models

import (
	"context"
	"github/hasnatsaeed/go-fiber-mongo-hrms/pkg/configs/database"
	"go.mongodb.org/mongo-driver/bson"
)

const EMPLOYEE_COLLECTION string = "employees"

type Employee struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Salary float64 `json:"salary"`
	Age    int     `json:"age"`
}

func (employee *Employee) CreateEmployee() (*Employee, error) {

	collection := database.GetCollection(EMPLOYEE_COLLECTION)
	id, errInsert := collection.InsertOne(context.TODO(), employee)
	if errInsert != nil {
		return nil, errInsert
	}

	insertedEmployee := &Employee{}
	errorFetch := collection.FindOne(context.TODO(), bson.D{{Key: "_id", Value: id.InsertedID}}).Decode(insertedEmployee)

	if errorFetch != nil {
		return nil, errorFetch
	}

	return insertedEmployee, nil
}

func GetEmployees() ([]Employee, error) {

	collection := database.GetCollection(EMPLOYEE_COLLECTION)

	cursor, errFindRecords := collection.Find(context.TODO(), bson.D{{}})
	if errFindRecords != nil {
		return nil, errFindRecords
	}

	employees := make([]Employee, 10)

	errDecodeToEmployees := cursor.All(context.TODO(), &employees)
	if errDecodeToEmployees != nil {
		return nil, errDecodeToEmployees
	}
	return employees, nil
}

func GetEmployee(ID string) (*Employee, error) {
	collection := database.GetCollection(EMPLOYEE_COLLECTION)

	employee := &Employee{}

	errorFetch := collection.FindOne(context.TODO(), bson.D{{"id", ID}}).Decode(employee)

	if errorFetch != nil {
		return nil, errorFetch
	}
	return employee, nil
}

func DeleteEmployee(ID string) (int64, error) {
	collection := database.GetCollection(EMPLOYEE_COLLECTION)
	deleteResult, errDelete := collection.DeleteOne(context.TODO(), bson.D{{"id", ID}})
	return deleteResult.DeletedCount, errDelete
}

func UpdateEmployee(ID string, employee Employee) (int64, error) {
	collection := database.GetCollection(EMPLOYEE_COLLECTION)

	update := bson.D{{"$set", bson.D{
		{"name", employee.Name},
		{"salary", employee.Salary},
		{"age", employee.Age}}}}

	record, errUpdate := collection.UpdateOne(context.TODO(), bson.D{{"id", ID}}, update)
	if errUpdate != nil {
		return 0, errUpdate
	}
	return record.ModifiedCount, errUpdate
}
