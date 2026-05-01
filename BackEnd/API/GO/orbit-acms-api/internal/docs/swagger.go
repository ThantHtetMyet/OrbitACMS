package docs

import "net/http"

const swaggerUIHTML = `<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1" />
  <title>Orbit ACMS API Docs</title>
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
  title: Orbit ACMS API
  version: 1.0.0
  description: API documentation for Orbit ACMS backend services.
servers:
  - url: http://localhost:8080
    description: Local development server
paths:
  /health:
    get:
      summary: Health Check
      description: Returns current API health status.
      tags:
        - System
      responses:
        "200":
          description: Service health response
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/HealthResponse"
components:
  schemas:
    HealthResponse:
      type: object
      properties:
        status:
          type: string
          example: ok
        api:
          type: string
          example: orbit-acms-api
      required:
        - status
        - api
`

func SwaggerUIHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, _ = w.Write([]byte(swaggerUIHTML))
}

func OpenAPIHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/yaml; charset=utf-8")
	_, _ = w.Write([]byte(openAPIYAML))
}

