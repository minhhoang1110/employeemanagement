package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/minhhoang1110/employeemanagement/service"
)

func EmployeeApis(r *mux.Router) {
	employeeService := service.EmployeeInstance()
	r.HandleFunc("/api/employee", employeeService.CreateEmployee).Methods(http.MethodPost)
}
