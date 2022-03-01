package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/minhhoang1110/employeemanagement/service"
)

func DepartmentApis(r *mux.Router) {
	departmentService := service.DepartmentInstance()
	r.HandleFunc("/api/department", departmentService.CreateDepartment).Methods(http.MethodPost)
	r.HandleFunc("/api/department/{id}", departmentService.GetDepartmentObject).Methods(http.MethodGet)
	r.HandleFunc("/api/department", departmentService.GetListDepartment).Methods(http.MethodGet)
}
