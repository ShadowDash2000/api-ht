package handler

import "net/http"

type DepartmentHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Patch(w http.ResponseWriter, r *http.Request)
}

type EmployeeHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
}

func RegisterRoutes(router *http.ServeMux, departmentHandler DepartmentHandler, employeeHandler EmployeeHandler) {
	router.HandleFunc("POST /departments", departmentHandler.Create)
	router.HandleFunc("GET /departments/{id}", departmentHandler.Get)
	router.HandleFunc("DELETE /departments/{id}", departmentHandler.Delete)
	router.HandleFunc("PATCH /departments/{id}", departmentHandler.Patch)

	router.HandleFunc("POST /departments/{id}/employees", employeeHandler.Create)
}
