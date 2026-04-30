package docs

import "net/http"

const swaggerUIHTML = `<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1" />
  <title>Orbit Teleport API Docs</title>
  <link rel="stylesheet" href="https://unpkg.com/swagger-ui-dist@5/swagger-ui.css" />
</head>
<body>
  <div id="swagger-ui"></div>
  <script src="https://unpkg.com/swagger-ui-dist@5/swagger-ui-bundle.js"></script>
  <script>
    window.ui = SwaggerUIBundle({
      url: "/swagger/openapi.yaml",
      dom_id: "#swagger-ui",
      deepLinking: true,
      docExpansion: "list"
    });
  </script>
</body>
</html>
`

const openAPIYAML = `openapi: 3.0.3
info:
  title: Orbit Teleport API
  version: 1.0.0
  description: API documentation for Orbit Teleport backend services.
servers:
  - url: http://localhost:8080
    description: Local development server
paths:
  /health:
    get:
      tags: [System]
      summary: Health Check
      responses:
        "200":
          description: Service health response
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/HealthResponse"
  /users:
    get:
      tags: [User]
      summary: List users
      responses:
        "200":
          description: List of users
          content:
            application/json:
              schema:
                type: array
                items:
                  $ref: "#/components/schemas/UserDTO"
    post:
      tags: [User]
      summary: Create user
      requestBody:
        required: true
        content:
          application/json:
            schema:
              $ref: "#/components/schemas/UserDTO"
      responses:
        "201":
          description: User created
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/UserDTO"
  /users/{id}:
    parameters:
      - $ref: "#/components/parameters/ID"
    get:
      tags: [User]
      summary: Get user by ID
      responses:
        "200":
          description: User detail
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/UserDTO"
    put:
      tags: [User]
      summary: Update user
      requestBody:
        required: true
        content:
          application/json:
            schema:
              $ref: "#/components/schemas/UserDTO"
      responses:
        "200":
          description: Updated user
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/UserDTO"
    delete:
      tags: [User]
      summary: Delete user
      responses:
        "204":
          description: User deleted
  /user-images:
    get:
      tags: [UserImage]
      summary: List user images
      responses:
        "200":
          description: List of user images
          content:
            application/json:
              schema:
                type: array
                items:
                  $ref: "#/components/schemas/UserImageDTO"
    post:
      tags: [UserImage]
      summary: Create user image
      requestBody:
        required: true
        content:
          application/json:
            schema:
              $ref: "#/components/schemas/UserImageDTO"
      responses:
        "201":
          description: User image created
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/UserImageDTO"
  /user-images/{id}:
    parameters:
      - $ref: "#/components/parameters/ID"
    get:
      tags: [UserImage]
      summary: Get user image by ID
      responses:
        "200":
          description: User image detail
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/UserImageDTO"
    put:
      tags: [UserImage]
      summary: Update user image
      requestBody:
        required: true
        content:
          application/json:
            schema:
              $ref: "#/components/schemas/UserImageDTO"
      responses:
        "200":
          description: Updated user image
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/UserImageDTO"
    delete:
      tags: [UserImage]
      summary: Delete user image
      responses:
        "204":
          description: User image deleted
  /image-types:
    get:
      tags: [ImageType]
      summary: List image types
      responses:
        "200":
          description: List of image types
          content:
            application/json:
              schema:
                type: array
                items:
                  $ref: "#/components/schemas/ImageTypeDTO"
    post:
      tags: [ImageType]
      summary: Create image type
      requestBody:
        required: true
        content:
          application/json:
            schema:
              $ref: "#/components/schemas/ImageTypeDTO"
      responses:
        "201":
          description: Image type created
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/ImageTypeDTO"
  /image-types/{id}:
    parameters:
      - $ref: "#/components/parameters/ID"
    get:
      tags: [ImageType]
      summary: Get image type by ID
      responses:
        "200":
          description: Image type detail
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/ImageTypeDTO"
    put:
      tags: [ImageType]
      summary: Update image type
      requestBody:
        required: true
        content:
          application/json:
            schema:
              $ref: "#/components/schemas/ImageTypeDTO"
      responses:
        "200":
          description: Updated image type
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/ImageTypeDTO"
    delete:
      tags: [ImageType]
      summary: Delete image type
      responses:
        "204":
          description: Image type deleted
  /user-roles:
    get:
      tags: [UserRole]
      summary: List user roles
      responses:
        "200":
          description: List of user roles
          content:
            application/json:
              schema:
                type: array
                items:
                  $ref: "#/components/schemas/UserRoleDTO"
    post:
      tags: [UserRole]
      summary: Create user role
      requestBody:
        required: true
        content:
          application/json:
            schema:
              $ref: "#/components/schemas/UserRoleDTO"
      responses:
        "201":
          description: User role created
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/UserRoleDTO"
  /user-roles/{id}:
    parameters:
      - $ref: "#/components/parameters/ID"
    get:
      tags: [UserRole]
      summary: Get user role by ID
      responses:
        "200":
          description: User role detail
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/UserRoleDTO"
    put:
      tags: [UserRole]
      summary: Update user role
      requestBody:
        required: true
        content:
          application/json:
            schema:
              $ref: "#/components/schemas/UserRoleDTO"
      responses:
        "200":
          description: Updated user role
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/UserRoleDTO"
    delete:
      tags: [UserRole]
      summary: Delete user role
      responses:
        "204":
          description: User role deleted
  /user-permissions:
    get:
      tags: [UserPermission]
      summary: List user permissions
      responses:
        "200":
          description: List of user permissions
          content:
            application/json:
              schema:
                type: array
                items:
                  $ref: "#/components/schemas/UserPermissionDTO"
    post:
      tags: [UserPermission]
      summary: Create user permission
      requestBody:
        required: true
        content:
          application/json:
            schema:
              $ref: "#/components/schemas/UserPermissionDTO"
      responses:
        "201":
          description: User permission created
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/UserPermissionDTO"
  /user-permissions/{id}:
    parameters:
      - $ref: "#/components/parameters/ID"
    get:
      tags: [UserPermission]
      summary: Get user permission by ID
      responses:
        "200":
          description: User permission detail
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/UserPermissionDTO"
    put:
      tags: [UserPermission]
      summary: Update user permission
      requestBody:
        required: true
        content:
          application/json:
            schema:
              $ref: "#/components/schemas/UserPermissionDTO"
      responses:
        "200":
          description: Updated user permission
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/UserPermissionDTO"
    delete:
      tags: [UserPermission]
      summary: Delete user permission
      responses:
        "204":
          description: User permission deleted
components:
  parameters:
    ID:
      name: id
      in: path
      required: true
      schema:
        type: string
  schemas:
    HealthResponse:
      type: object
      properties:
        status:
          type: string
          example: ok
        api:
          type: string
          example: orbit-teleport-api
      required: [status, api]
    UserDTO:
      type: object
      properties:
        id: { type: string }
        userRoleId: { type: string }
        firstName: { type: string }
        lastName: { type: string }
        email: { type: string }
        mobileNo: { type: string }
        remark: { type: string, nullable: true }
        lastLogin: { type: string, format: date-time, nullable: true }
        isActive: { type: boolean }
        isDeleted: { type: boolean }
        createdDate: { type: string, format: date-time }
        updatedDate: { type: string, format: date-time }
        createdBy: { type: string, nullable: true }
        updatedBy: { type: string, nullable: true }
    UserImageDTO:
      type: object
      properties:
        id: { type: string }
        userId: { type: string }
        imageTypeId: { type: string }
        imageName: { type: string }
        storedDirectory: { type: string }
        isDeleted: { type: boolean }
        uploadedDate: { type: string, format: date-time }
        uploadedBy: { type: string, nullable: true }
    ImageTypeDTO:
      type: object
      properties:
        id: { type: string }
        imageTypeName: { type: string }
        isDeleted: { type: boolean }
        createdDate: { type: string, format: date-time }
        updatedDate: { type: string, format: date-time }
        createdBy: { type: string, nullable: true }
        updatedBy: { type: string, nullable: true }
    UserRoleDTO:
      type: object
      properties:
        id: { type: string }
        roleName: { type: string }
        description: { type: string, nullable: true }
        isDeleted: { type: boolean }
        createdDate: { type: string, format: date-time, nullable: true }
        updatedDate: { type: string, format: date-time, nullable: true }
        createdBy: { type: string, nullable: true }
        updatedBy: { type: string, nullable: true }
    UserPermissionDTO:
      type: object
      properties:
        id: { type: string }
        userRoleId: { type: string }
        moduleName: { type: string }
        canCreate: { type: boolean }
        canRead: { type: boolean }
        canUpdate: { type: boolean }
        canDelete: { type: boolean }
        isDeleted: { type: boolean }
        createdDate: { type: string, format: date-time, nullable: true }
        updatedDate: { type: string, format: date-time, nullable: true }
        createdBy: { type: string, nullable: true }
        updatedBy: { type: string, nullable: true }
`

func SwaggerUIHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, _ = w.Write([]byte(swaggerUIHTML))
}

func OpenAPIHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/yaml; charset=utf-8")
	_, _ = w.Write([]byte(openAPIYAML))
}
