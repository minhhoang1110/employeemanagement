package api

import "github.com/gorilla/mux"

func Register(r *mux.Router) {
	EmployeeApis(r)
	DepartmentApis(r)
}
