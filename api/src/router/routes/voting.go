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
	{
		URI: "/results",
		Method: http.MethodGet,
		Function: controllers.GetResults,
	},
	{
		URI: "/votes",
		Method: http.MethodGet,
		Function: controllers.GetVotes,
	},

}