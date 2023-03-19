// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package dependencies

import (
	"github.com/google/wire"
	"github.com/shanmukhsista/go-graphql-starter/cmd/graphql-server/graph"
	"github.com/shanmukhsista/go-graphql-starter/pkg/common/db"
	"github.com/shanmukhsista/go-graphql-starter/pkg/services/notes"
)

// Injectors from wire.go:

func NewAppResolverService() (*graph.Resolver, error) {
	pool, err := db.ProvidePgConnectionPool()
	if err != nil {
		return nil, err
	}
	database := db.ProvideNewDatabaseConnection(pool)
	repository, err := notes.ProvideNewNotesRepository(database)
	if err != nil {
		return nil, err
	}
	transactionManager, err := db.ProvideNewPostgresTransactor(database)
	if err != nil {
		return nil, err
	}
	service, err := notes.ProvideNewNotesService(repository, transactionManager)
	if err != nil {
		return nil, err
	}
	resolver := graph.ProvideNewServerResolver(service)
	return resolver, nil
}

// wire.go:

var postgresDbConnectionSet = wire.NewSet(db.ProvidePgConnectionPool, db.ProvideNewPostgresTransactor, db.ProvideNewDatabaseConnection)

var notesApiService = wire.NewSet(notes.ProvideNewNotesRepository, notes.ProvideNewNotesService)

var graphQLServerDependencySet = wire.NewSet(
	postgresDbConnectionSet,
	notesApiService, graph.ProvideNewServerResolver,
)