package utils

import "github.com/labstack/echo/v4"

func ReadRequest(ctx echo.Context, request interface{}) error {
	if err := ctx.Bind(request); err != nil {
		return err
	}
	return validate.StructCtx(ctx.Request().Context(), request)
}

func GetConfigEnv(configEnv string) string {
	if configEnv == "docker" {
		return "./App/config/config-docker.yml"
	}
	return "./App/config/config-local.yml"
}
