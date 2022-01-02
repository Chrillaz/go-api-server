package mocks

import "go-api-server/api/models"

var Todos = []models.Todo{
	{
		Name:        "Finnish api",
		Description: "Code the api server",
	},
	{
		Name:        "Finnish client GUI",
		Description: "Code the GUI with react",
	},
}
