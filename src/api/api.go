package execapi

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/ibad69/go-devxcoaches/src/pkg/index/transport"
)

type ApiConfig struct {
	Dburi string
	Port  string
}

func Start(cfg *ApiConfig) {
	fmt.Print("testing")
	// db, err := sql.Open("pgx", cfg.Dburi)
	db, err := pgxpool.New(context.TODO(), cfg.Dburi)
	if err != nil {
		fmt.Errorf("unable to create connection pool: %w", err)
		log.Fatal("connection not made to db")
	}
	// if err != nil {
	// 	log.Println("an error occured in opening a db connection")
	// }
	// if err = db.Ping(); err != nil {
	// 	log.Fatalf("cannot ping the databse %s", err)
	// }
	log.Println("success fully connected to database")
	r := chi.NewRouter()

	// transport.Activate(r, db)
	transport.Activate(r, db)
	err = http.ListenAndServe(":5100", r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	log.Println("Server running on port: ")
}
