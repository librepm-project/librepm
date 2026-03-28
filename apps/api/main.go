package main

import (
	"apps/api/app/domain"
	"apps/api/app/http"
	"apps/api/app/migration"
	"apps/api/app/seed"
	"fmt"
	"libs/jwt_utils"
	"libs/mysql_utils"
	"os"

	"github.com/google/uuid"
)

func main() {
	var command string
	var subcommand string
	if len(os.Args) > 1 {
		command = os.Args[1]
	}
	if len(os.Args) > 2 {
		subcommand = os.Args[2]
	}
	db := mysql_utils.Init()
	migration.Run(db)
	d := domain.NewDomain(db)
	switch command {
	case "server":
		http.StartHttpServer(d)
	case "seed":
		seed.NewSeedService(d).Seed(subcommand)
	case "purge":
		seed.NewSeedService(d).Purge()
	case "generate-login-link":
		if subcommand == "" {
			fmt.Fprintln(os.Stderr, "Usage: go run apps/api generate-login-link <email>")
			os.Exit(1)
		}
		user, err := d.UserRepository.FindByEmail(subcommand)
		if err != nil || user.ID == uuid.Nil {
			fmt.Fprintf(os.Stderr, "User not found: %s\n", subcommand)
			os.Exit(1)
		}
		token := jwt_utils.GenerateToken(user.ID, user.Email)
		fmt.Printf("http://localhost:8081/login?token=%s\n", token)
	default:
		http.StartHttpServer(d)
	}
}
