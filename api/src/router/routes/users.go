package routes

import (
	"devbook-api/src/controllers"
)

var usersRoutes = []Route{
	{
		URI:         "/users",
		Method:      "GET",
		Function:    controllers.GetUsers,
		RequireAuth: false,
	},
	{
		URI:         "/users/{id}",
		Method:      "GET",
		Function:    controllers.GetUser,
		RequireAuth: false,
	},
	{
		URI:         "/users",
		Method:      "POST",
		Function:    controllers.CreateUser,
		RequireAuth: false,
	},
	{
		URI:         "/users/{id}",
		Method:      "PUT",
		Function:    controllers.UpdateUser,
		RequireAuth: false,
	},
	{
		URI:         "/users/{id}",
		Method:      "DELETE",
		Function:    controllers.DeleteUser,
		RequireAuth: false,
	},
}
