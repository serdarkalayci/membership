package postgres

import (
	"context"
	"fmt"
	"time"

	pgxpool "github.com/jackc/pgx/v5/pgxpool"
	"github.com/nicholasjackson/env"
	"github.com/rs/zerolog/log"
	"github.com/serdarkalayci/membership/api/application"
)

var databaseName = env.String("DB_NAME", false, "membership", "The database name")
var dbHost = env.String("DB_HOST", false, "localhost", "Database connection string")
var dbPort = env.String("DB_PORT", false, "5432", "Database port")
var username = env.String("DB_USER", false, "membershipuser", "Database username")
var password = env.String("DB_PASSWORD", false, "mysecretpassword", "Database password")

// NewDataContext returns a new mongoDB backed DataContext
func NewDataContext() (*application.DataContext, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// We try to get connectionstring value from the environment variables, if not found it falls back to local database
	env.Parse()
	if *username == "" {
		*username = "defaultuser"
	}
	if *password == "" {
		*password = "defaultpassword"
	}
	if *dbHost == "" {
		*dbHost = "localhost"
	}
	if *dbPort == "" {
		*dbPort = "5432"
	}
	if *databaseName == "" {
		*databaseName = "membership"
	}
	connectionString := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s", *username, *password, *dbHost, *dbPort, *databaseName)
	
	dbPool, err := pgxpool.New(ctx, connectionString)
	// execute the select query and get result rows
	_, err = dbPool.Query(ctx, "select 1")

	if err != nil {
		log.Error().Err(err).Msgf("An error occured while connecting to tha database")
	}
	dataContext := application.DataContext{}
	uRepo := newUserRepository(dbPool, *databaseName)
	mRepo := newMemberRepository(dbPool, *databaseName)
	lRepo := newLookupRepository(dbPool, *databaseName)
	dataContext.SetRepositories(uRepo, mRepo, lRepo)
	return &dataContext, nil
}
