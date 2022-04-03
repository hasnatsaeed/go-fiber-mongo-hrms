package models

import (
	"github/hasnatsaeed/go-fiber-mongo-hrms/pkg/configs/database"
	"go.mongodb.org/mongo-driver/bson"
)

type Employee struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Salary float64 `json:"salary"`
	Age    int     `json:"age"`
}

func (employee *Employee) CreateEmployee() (*Employee, error) {

	collection := database.GetDb().Collection("employees")
	id, errInsert := collection.InsertOne(*database.GetCtx(), employee)
	if errInsert != nil {
		return nil, errInsert
	}

	getByIdFilter := bson.D{{Key: "_id", Value: id.InsertedID}}

	record := collection.FindOne(*database.GetCtx(), getByIdFilter)

	createdEmployee := &Employee{}

	errDecodeToEmployee := record.Decode(createdEmployee)
	if errDecodeToEmployee != nil {
		return nil, errDecodeToEmployee
	}

	return createdEmployee, nil
}

func GetEmployees() ([]Employee, error) {

	collection := database.GetDb().Collection("employees")

	employees := make([]Employee, 10)
	records, errFindRecords := collection.Find(*database.GetCtx(), bson.D{{}})
	if errFindRecords != nil {
		return nil, errFindRecords
	}

	errDecodeToEmployees := records.All(*database.GetCtx(), &employees)
	if errDecodeToEmployees != nil {
		return nil, errDecodeToEmployees
	}
	return employees, nil
}

func GetEmployee(ID string) (*Employee, error) {
	collection := database.GetDb().Collection("employees")
	record := collection.FindOne(*database.GetCtx(), bson.D{{"id", ID}})

	employee := &Employee{}
	errDecodeToEmployee := record.Decode(employee)
	if errDecodeToEmployee != nil {
		return nil, errDecodeToEmployee
	}
	return employee, nil
}

func DeleteEmployee(ID string) (int64, error) {
	collection := database.GetDb().Collection("employees")
	deleteCount, errDelete := collection.DeleteOne(*database.GetCtx(), bson.D{{"id", ID}})

	return deleteCount.DeletedCount, errDelete

}
