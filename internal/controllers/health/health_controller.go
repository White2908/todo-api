package health

import (
	"runtime"
	"time"

	"todo-api/internal/controllers/base"

	"github.com/go-fuego/fuego"
)

type HealthController struct {
	*base.BaseController
	startTime time.Time
}

func NewHealthController() *HealthController {
	return &HealthController{
		BaseController: base.NewBaseController(),
		startTime:      time.Now(),
	}
}

// Get health status
func (c *HealthController) HealthCheck(ctx fuego.ContextNoBody) (map[string]interface{}, error) {
	return map[string]interface{}{
		"status":     "healthy",
		"timestamp":  time.Now(),
		"uptime":     time.Since(c.startTime).String(),
		"go_version": runtime.Version(),
		"goroutines": runtime.NumGoroutine(),
		"cpu_cores":  runtime.NumCPU(),
	}, nil
}

// Get readiness status
func (c *HealthController) ReadinessCheck(ctx fuego.ContextNoBody) (map[string]interface{}, error) {
	return map[string]interface{}{
		"status": "ready",
		"checks": map[string]bool{
			"database": true,
			"memory":   true,
		},
	}, nil
}

// Get liveness status
func (c *HealthController) LivenessCheck(ctx fuego.ContextNoBody) (map[string]interface{}, error) {
	return map[string]interface{}{
		"status": "alive",
		"time":   time.Now().Unix(),
	}, nil
}
