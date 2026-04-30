package dto

// HealthResponse is the API response shape for health checks.
type HealthResponse struct {
	Status string `json:"status"`
	API    string `json:"api"`
}
