package routes

import (
	"api/src/controllers"
	"net/http"
)

var routesVoting = []Route{
	{
		URI: "/vote",
		Method: http.MethodPost,
		Function: controllers.CreateVote,
	},
}