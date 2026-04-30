package router

import (
	"net/http"

	"orbit-teleport-api/internal/controller"
	"orbit-teleport-api/internal/docs"
)

func New() *http.ServeMux {
	mux := http.NewServeMux()

	healthController := controller.NewHealthController()
	userController := controller.NewUserController()
	userImageController := controller.NewUserImageController()
	imageTypeController := controller.NewImageTypeController()
	userRoleController := controller.NewUserRoleController()
	userPermissionController := controller.NewUserPermissionController()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/swagger", http.StatusFound)
	})
	mux.HandleFunc("GET /swagger", docs.SwaggerUIHandler)
	mux.HandleFunc("GET /swagger/", docs.SwaggerUIHandler)
	mux.HandleFunc("GET /swagger/openapi.yaml", docs.OpenAPIHandler)
	mux.HandleFunc("GET /health", healthController.GetHealth)
	mux.HandleFunc("GET /users", userController.List)
	mux.HandleFunc("GET /users/{id}", userController.GetByID)
	mux.HandleFunc("POST /users", userController.Create)
	mux.HandleFunc("PUT /users/{id}", userController.Update)
	mux.HandleFunc("DELETE /users/{id}", userController.Delete)
	mux.HandleFunc("GET /user-images", userImageController.List)
	mux.HandleFunc("GET /user-images/{id}", userImageController.GetByID)
	mux.HandleFunc("POST /user-images", userImageController.Create)
	mux.HandleFunc("PUT /user-images/{id}", userImageController.Update)
	mux.HandleFunc("DELETE /user-images/{id}", userImageController.Delete)
	mux.HandleFunc("GET /image-types", imageTypeController.List)
	mux.HandleFunc("GET /image-types/{id}", imageTypeController.GetByID)
	mux.HandleFunc("POST /image-types", imageTypeController.Create)
	mux.HandleFunc("PUT /image-types/{id}", imageTypeController.Update)
	mux.HandleFunc("DELETE /image-types/{id}", imageTypeController.Delete)
	mux.HandleFunc("GET /user-roles", userRoleController.List)
	mux.HandleFunc("GET /user-roles/{id}", userRoleController.GetByID)
	mux.HandleFunc("POST /user-roles", userRoleController.Create)
	mux.HandleFunc("PUT /user-roles/{id}", userRoleController.Update)
	mux.HandleFunc("DELETE /user-roles/{id}", userRoleController.Delete)
	mux.HandleFunc("GET /user-permissions", userPermissionController.List)
	mux.HandleFunc("GET /user-permissions/{id}", userPermissionController.GetByID)
	mux.HandleFunc("POST /user-permissions", userPermissionController.Create)
	mux.HandleFunc("PUT /user-permissions/{id}", userPermissionController.Update)
	mux.HandleFunc("DELETE /user-permissions/{id}", userPermissionController.Delete)

	return mux
}
