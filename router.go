package mux

import (
	"github.com/gorilla/mux"
	// "path"
)

const (
	RepoGetRoute  = "repo"
	RepoListRoute = "repo.list"
)

func NewAPIRouter() *mux.Router {
	// Define the routes but don't attach handlers.
	m := mux.NewRouter()
	m.Path("/api/repos/{Name:.*}").Name(RepoGetRoute)
	m.Path("/api/repos").Name(RepoListRoute)
	return m
}
