package postgres

import (
	"fmt"
	"log"
	"time"

	"github.com/nicholasjackson/env"
	"github.com/serdarkalayci/membership/api/application"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var databaseName = env.String("DatabaseName", false, "membership", "The database name for arangodb")
var connectionString = env.String("ConnectionString", false, "localhost:5432", "Database connection string")
var username = env.String("DbUserName", false, "membershipuser", "Database username")
var password = env.String("DbPassword", false, "mysecretpassword", "Database password")

// NewDataContext returns a new mongoDB backed DataContext
func NewDataContext() (*application.DataContext, error) {
	env.Parse()
	// ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Second)
	// defer cancel()
	// We try to get connectionstring value from the environment variables, if not found it falls back to local database
	dsn := fmt.Sprintf("postgres://%s:%s@%s/%s", *username, *password, *connectionString, *databaseName)
	fmt.Printf("Connecting to %s\n", dsn)
	// Client object
	// db, err := pgx.Connect(ctx, dsn)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	  }), &gorm.Config{})
    if err != nil {
		log.Fatalln(err)
        return &application.DataContext{}, fmt.Errorf("cannot create the database client on %s", *connectionString)
    }
	// Open a database. In case the database is not ready yet, we retry a few times
	count := 0
	for count < 5 {
		var result int
		if db.Raw("select 1;").Scan(&result); result == 1 {
			break
		}
		count++
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		return &application.DataContext{}, fmt.Errorf("cannot connect to the database on %s", *connectionString)
	}
	dataContext := application.DataContext{}
	uRepo := newUserRepository(db)
	mRepo := newMemberRepository(db)
	dataContext.SetRepositories(uRepo, mRepo)
	return &dataContext, nil
}
