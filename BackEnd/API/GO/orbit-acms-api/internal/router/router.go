package router

import (
	"net/http"

	"orbit-acms-api/internal/controller"
	"orbit-acms-api/internal/docs"
)

// New wires all API routes and returns the root mux.
func New() *http.ServeMux {
	mux := http.NewServeMux()

	health := controller.NewHealthController()
	imageType := controller.NewImageTypeController()
	userImage := controller.NewUserImageController()
	userPermission := controller.NewUserPermissionController()
	userRole := controller.NewUserRoleController()

	// Health
	mux.HandleFunc("GET /health", health.GetHealth)

	// Swagger docs
	mux.HandleFunc("GET /swagger", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/swagger/index.html", http.StatusMovedPermanently)
	})
	mux.HandleFunc("GET /swagger/index.html", docs.SwaggerUIHandler)
	mux.HandleFunc("GET /swagger/openapi.yaml", docs.OpenAPIHandler)

	// Image types
	mux.HandleFunc("GET /api/image-types", imageType.List)
	mux.HandleFunc("POST /api/image-types", imageType.Create)
	mux.HandleFunc("GET /api/image-types/{id}", imageType.GetByID)
	mux.HandleFunc("PUT /api/image-types/{id}", imageType.Update)
	mux.HandleFunc("DELETE /api/image-types/{id}", imageType.Delete)

	// User images
	mux.HandleFunc("GET /api/user-images", userImage.List)
	mux.HandleFunc("POST /api/user-images", userImage.Create)
	mux.HandleFunc("GET /api/user-images/{id}", userImage.GetByID)
	mux.HandleFunc("PUT /api/user-images/{id}", userImage.Update)
	mux.HandleFunc("DELETE /api/user-images/{id}", userImage.Delete)

	// User permissions
	mux.HandleFunc("GET /api/user-permissions", userPermission.List)
	mux.HandleFunc("POST /api/user-permissions", userPermission.Create)
	mux.HandleFunc("GET /api/user-permissions/{id}", userPermission.GetByID)
	mux.HandleFunc("PUT /api/user-permissions/{id}", userPermission.Update)
	mux.HandleFunc("DELETE /api/user-permissions/{id}", userPermission.Delete)

	// User roles
	mux.HandleFunc("GET /api/user-roles", userRole.List)
	mux.HandleFunc("POST /api/user-roles", userRole.Create)
	mux.HandleFunc("GET /api/user-roles/{id}", userRole.GetByID)
	mux.HandleFunc("PUT /api/user-roles/{id}", userRole.Update)
	mux.HandleFunc("DELETE /api/user-roles/{id}", userRole.Delete)

	return mux
}
