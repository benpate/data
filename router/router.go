package router

import (
	"github.com/benpate/data"
	"github.com/benpate/data/memory"
	"github.com/benpate/data/mongodb"
	"github.com/benpate/derp"
)

// New uses the provided database configuration information and returns a fully initialized datastore
func New(config map[string]string) (data.Datastore, *derp.Error) {

	switch config["type"] {

	case "memory":
		return memory.New(), nil

	case "mongodb":
		return mongodb.New(config["uri"], config["database"]), nil

	default:
		return nil, derp.New(404, "data.Router.New", "Unrecognized database configuration", config)
	}
}
