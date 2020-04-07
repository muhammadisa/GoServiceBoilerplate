package cmd

import (
	"fmt"

	"github.com/muhammadisa/restful-api-boilerplate/api"
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
			api.Seed{}.ReinitializeStructs()
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
			api.Seed{}.DropTableIfExist()
		},
	}
)

func init() {
	cmd.AddCommand(reinitTables)
	cmd.AddCommand(migrate)
	cmd.AddCommand(initDummy)
	cmd.AddCommand(dropTableIfExist)
}
