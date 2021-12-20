package server

import (
	"context"
	"database/sql"
	"fmt"

	// mysql driver
	_ "github.com/jackc/pgx/v4/stdlib"

	"github.com/bandajon/microservices_users/pkg/protocol/grpc"
	v1 "github.com/bandajon/microservices_users/pkg/service/v1"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "Admin"
	password = "mypass"
	dbname   = "postgres"
)

// Config is configuration for Server
type Config struct {
	// gRPC server start parameters section
	// gRPC is TCP port to listen by gRPC server
	GRPCPort string

	// DB Datastore parameters section
	// DatastoreDBHost is host of database
	DatastoreDBHost string
	// DatastoreDBUser is username to connect to database
	DatastoreDBUser string
	// DatastoreDBPassword password to connect to database
	DatastoreDBPassword string
	// DatastoreDBSchema is schema of database
	DatastoreDBSchema string
}

// RunServer runs gRPC server and HTTP gateway
func RunServer() error {
	ctx := context.Background()

	// get configuration
	/*var cfg Config
	flag.StringVar(&cfg.GRPCPort, "grpc-port", "", "gRPC port to bind")
	flag.StringVar(&cfg.DatastoreDBHost, "db-host", "", "Database host")
	flag.StringVar(&cfg.DatastoreDBUser, "db-user", "", "Database user")
	flag.StringVar(&cfg.DatastoreDBPassword, "db-password", "", "Database password")
	flag.StringVar(&cfg.DatastoreDBSchema, "db-schema", "", "Database schema")
	flag.Parse()


	if len(cfg.GRPCPort) == 0 {
		return fmt.Errorf("invalid TCP port for gRPC server: '%s'", cfg.GRPCPort)
	}

	// add MySQL driver specific parameter to parse date/time
	// Drop it for another database
	param := "parseTime=true"

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s",
		cfg.DatastoreDBUser,
		cfg.DatastoreDBPassword,
		cfg.DatastoreDBHost,
		cfg.DatastoreDBSchema,
		param)
	*/
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		return fmt.Errorf("failed to open database: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	CheckError(err)

	fmt.Println("Connected!")

	v1API := v1.NewToDoServiceServer(db)

	return grpc.RunServer(ctx, v1API, port)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
