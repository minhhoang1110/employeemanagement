package service

import (
	"encoding/json"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gorilla/mux"
	"github.com/minhhoang1110/employeemanagement/database"
	"github.com/minhhoang1110/employeemanagement/models"
	"github.com/minhhoang1110/employeemanagement/types"
)

type Employee struct{}

var employeeO Employee

func EmployeeInstance() *Employee {
	return &employeeO
}

func (e *Employee) CreateEmployee(writer http.ResponseWriter, request *http.Request) {
	var newEmployee models.Employee
	var respone types.Respone
	if err := json.NewDecoder(request.Body).Decode(&newEmployee); err != nil {
		respone.ResponseWithJson(writer, http.StatusBadRequest, map[string]string{"message": "Invalid body"}, "error")
		return
	}
	if e.IsEmptyData(newEmployee) {
		respone.ResponseWithJson(writer, http.StatusBadRequest, map[string]string{"message": "Data can not empty"}, "error")
		return
	}
	if !govalidator.IsEmail(newEmployee.Email) {
		respone.ResponseWithJson(writer, http.StatusBadRequest, map[string]string{"message": "Email is invalid"}, "error")
		return
	}
	db, err := database.GetConnect()
	if err != nil {
		respone.ResponseWithJson(writer, http.StatusInternalServerError, map[string]string{"message": "Connect error"}, "error")
		return
	}
	result := db.Create(&newEmployee)
	if result.Error != nil {
		respone.ResponseWithJson(writer, http.StatusBadRequest, map[string]string{"message": result.Error.Error()}, "error")
		return
	}
	respone.ResponseWithJson(writer, http.StatusOK, newEmployee, "employee")
}

func (e *Employee) GetEmployeeObject(writer http.ResponseWriter, request *http.Request) {
	var employee models.Employee
	var respone types.Respone
	params := mux.Vars(request)
	var id string
	var ok bool
	if id, ok = params["id"]; !ok {
		respone.ResponseWithJson(writer, http.StatusBadRequest, map[string]string{"message": "Invalid ID"}, "error")
		return
	}
	db, err := database.GetConnect()
	if err != nil {
		respone.ResponseWithJson(writer, http.StatusInternalServerError, map[string]string{"message": "Connect error"}, "error")
		return
	}
	result := db.Find(&employee, "id=?", id)
	if result.Error != nil {
		respone.ResponseWithJson(writer, http.StatusBadRequest, map[string]string{"message": result.Error.Error()}, "error")
		return
	}
	if employee.ID == "" {
		respone.ResponseWithJson(writer, http.StatusBadRequest, map[string]string{"message": "Employee does not exits"}, "error")
		return
	}
	respone.ResponseWithJson(writer, http.StatusOK, employee, "employee")
}

func (e *Employee) GetListEmployee(writer http.ResponseWriter, request *http.Request) {
	var employees []models.Employee
	var respone types.Respone
	db, err := database.GetConnect()
	if err != nil {
		respone.ResponseWithJson(writer, http.StatusInternalServerError, map[string]string{"message": "Connect error"}, "error")
		return
	}
	result := db.Find(&employees)
	if result.Error != nil {
		respone.ResponseWithJson(writer, http.StatusBadRequest, map[string]string{"message": result.Error.Error()}, "error")
		return
	}
	respone.ResponseWithJson(writer, http.StatusOK, employees, "employee")
}

func (e *Employee) IsEmptyData(employee models.Employee) bool {
	return govalidator.IsNull(employee.FirstName) || govalidator.IsNull(employee.LastName) || govalidator.IsNull(employee.Address) || govalidator.IsNull(string(employee.Gender)) || govalidator.IsNull(employee.Phone)
}
