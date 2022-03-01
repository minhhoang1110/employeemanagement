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

type Department struct{}

var departmentO Department

func DepartmentInstance() *Department {
	return &departmentO
}

func (d *Department) CreateDepartment(writer http.ResponseWriter, request *http.Request) {
	var newDepartment models.Department
	var respone types.Respone
	if err := json.NewDecoder(request.Body).Decode(&newDepartment); err != nil {
		respone.ResponseWithJson(writer, http.StatusBadRequest, map[string]string{"message": "Invalid body"}, "error")
		return
	}
	if d.IsEmptyData(newDepartment) {
		respone.ResponseWithJson(writer, http.StatusBadRequest, map[string]string{"message": "Data can not empty"}, "error")
		return
	}
	db, err := database.GetConnect()
	if err != nil {
		respone.ResponseWithJson(writer, http.StatusInternalServerError, map[string]string{"message": "Connect error"}, "error")
		return
	}
	result := db.Create(&newDepartment)
	if result.Error != nil {
		respone.ResponseWithJson(writer, http.StatusBadRequest, map[string]string{"message": result.Error.Error()}, "error")
		return
	}
	respone.ResponseWithJson(writer, http.StatusOK, newDepartment, "employee")
}

func (d *Department) GetDepartmentObject(writer http.ResponseWriter, request *http.Request) {
	var department models.Department
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
	result := db.Find(&department, "id=?", id)
	if result.Error != nil {
		respone.ResponseWithJson(writer, http.StatusBadRequest, map[string]string{"message": result.Error.Error()}, "error")
		return
	}
	if department.ID == "" {
		respone.ResponseWithJson(writer, http.StatusBadRequest, map[string]string{"message": "Department does not exits"}, "error")
		return
	}
	respone.ResponseWithJson(writer, http.StatusOK, department, "employee")
}

func (d *Department) GetListDepartment(writer http.ResponseWriter, request *http.Request) {
	var departments []models.Department
	var respone types.Respone
	db, err := database.GetConnect()
	if err != nil {
		respone.ResponseWithJson(writer, http.StatusInternalServerError, map[string]string{"message": "Connect error"}, "error")
		return
	}
	result := db.Find(&departments)
	if result.Error != nil {
		respone.ResponseWithJson(writer, http.StatusBadRequest, map[string]string{"message": result.Error.Error()}, "error")
		return
	}
	respone.ResponseWithJson(writer, http.StatusOK, departments, "employee")
}

func (d *Department) IsEmptyData(department models.Department) bool {
	return govalidator.IsNull(department.Name) || govalidator.IsNull(department.Phone) || govalidator.IsNull(department.Address)
}
