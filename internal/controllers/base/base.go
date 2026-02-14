package base

import (
	"todo-api/internal/pkg/utils"

	"github.com/go-fuego/fuego"
)

// BaseController provides common functionality for all controllers
type BaseController struct{}

func NewBaseController() *BaseController {
	return &BaseController{}
}

func (c *BaseController) SuccessResponse(message string) map[string]string {
	return map[string]string{"message": message}
}

func (c *BaseController) HandleError(err error) error {
	return utils.ToFuegoError(err)
}

func ValidateBody[T any](ctx fuego.ContextWithBody[T]) (*T, error) {
	body, err := ctx.Body()
	if err != nil {
		return nil, utils.NewBadRequestError("Invalid request body", err)
	}

	if err := utils.ValidateStruct(body); err != nil {
		return nil, utils.NewBadRequestError("Validation failed", err)
	}

	return &body, nil
}

func (c *BaseController) GetID(ctx fuego.ContextNoBody) string {
	return ctx.Request().PathValue("id")
}

func GetQuery[T any, B any](ctx fuego.Context[T, B], key string) string {
	return ctx.Request().URL.Query().Get(key)
}
