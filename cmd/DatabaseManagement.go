package cmd

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/muhammadisa/go-service-boilerplate/api"
	"github.com/muhammadisa/go-service-boilerplate/api/utils/envkeyeditor"
	"github.com/spf13/cobra"
)

var (
	reinitTables = &cobra.Command{
		Use:     "reinit-tables",
		Short:   "Reinitialize all table by model",
		Long:    "Reinitialize all table by model",
		Aliases: []string{"init-tables"},
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Reinit tables")

			// Loading .env file
			err := godotenv.Load()

			// Checking error for loading .env file
			if err != nil {
				log.Fatalf("Error getting env, not coming through %v", err)
				return
			}

			// Load drop mode env
			dropEnv := os.Getenv("DROP")
			drop, err := strconv.ParseBool(dropEnv)
			if err != nil {
				log.Fatalf("Unable parsing drop env value %v", err)
				return
			}

			if drop {
				api.Seed{}.ReinitializeStructs()
				lKV, nKV, err := envkeyeditor.EnvKeyEditor("DROP", "false")
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println(fmt.Sprintf("last DROP value %s", lKV))
					fmt.Println(fmt.Sprintf("DROP switched to %s", nKV))
				}
			} else {
				fmt.Println("Can reinit tables because DROP is still false")
			}
		},
	}
)

var (
	migrate = &cobra.Command{
		Use:     "migrate",
		Short:   "Migrate",
		Long:    "Migrate",
		Aliases: []string{"migrate"},
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Silent migrate")
			api.Seed{}.Migrate()
		},
	}
)

var (
	initDummy = &cobra.Command{
		Use:     "init-dummies",
		Short:   "Initialize dummy data for development needs",
		Long:    "Initialize dummy data for development needs",
		Aliases: []string{"init-dummies"},
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Init dummies")
		},
	}
)

var (
	dropTableIfExist = &cobra.Command{
		Use:     "drop-tables",
		Short:   "Drop tables if exist indatabase",
		Long:    "Drop tables if exist indatabase",
		Aliases: []string{"drop-tables"},
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Drop tables if exist")

			// Loading .env file
			err := godotenv.Load()

			// Checking error for loading .env file
			if err != nil {
				log.Fatalf("Error getting env, not coming through %v", err)
				return
			}

			// Load drop mode env
			dropEnv := os.Getenv("DROP")
			drop, err := strconv.ParseBool(dropEnv)
			if err != nil {
				log.Fatalf("Unable parsing drop env value %v", err)
				return
			}

			if drop {
				api.Seed{}.DropTableIfExist()
				lKV, nKV, err := envkeyeditor.EnvKeyEditor("DROP", "false")
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println(fmt.Sprintf("last DROP value %s", lKV))
					fmt.Println(fmt.Sprintf("DROP switched to %s", nKV))
				}
			} else {
				fmt.Println("Can reinit tables because DROP is still false")
			}
		},
	}
)

func init() {
	cmd.AddCommand(reinitTables)
	cmd.AddCommand(migrate)
	cmd.AddCommand(initDummy)
	cmd.AddCommand(dropTableIfExist)
}
