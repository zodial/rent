package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/zodial/rent/backend/internal/db"
	"github.com/zodial/rent/backend/internal/models"
)

func main() {
	migrate := flag.Bool("migrate", false, "run migrations")
	flag.Parse()

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "./data/dev.sqlite"
	}

	gdb, err := db.Open(dsn)
	if err != nil {
		log.Fatalf("failed to open db: %v", err)
	}

	if *migrate {
		if err := models.Migrate(gdb); err != nil {
			log.Fatalf("migrate failed: %v", err)
		}
		fmt.Println("migrations complete")
		return
	}

	// Minimal server placeholder. Full API handlers to be added in internal/api.
	fmt.Println("Server would start here. Run with --migrate to perform migrations.")
}
