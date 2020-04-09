package cmd

import (
	"fmt"
	"strconv"

	"github.com/muhammadisa/go-service-boilerplate/api"
	"github.com/muhammadisa/go-service-boilerplate/api/utils/envkeyeditor"
	"github.com/sethvargo/go-password/password"
	"github.com/spf13/cobra"
)

var (
	databaseName  string
	setupDatabase = &cobra.Command{
		Use:     "database",
		Short:   "Set database which will be used",
		Long:    "Set database which will be used, for this restful api",
		Aliases: []string{"database"},
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			lastKeyValue, newKeyValue, err := envkeyeditor.EnvKeyEditor("DB_DRIVER", args[0])
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(fmt.Sprintf("last DB_DRIVER value %s", lastKeyValue))
				fmt.Println(fmt.Sprintf("DB_DRIVER switched to %s", newKeyValue))
			}

			lastKeyValue, newKeyValue, err = envkeyeditor.EnvKeyEditor("TEST_DB_DRIVER", args[0])
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(fmt.Sprintf("last TEST_DB_DRIVER value %s", lastKeyValue))
				fmt.Println(fmt.Sprintf("TEST_DB_DRIVER switched to %s", newKeyValue))
			}
		},
	}
)

var (
	switchDrop = &cobra.Command{
		Use:     "switch-drop",
		Short:   "Switch state",
		Long:    "Switch state, if false cant drop table",
		Aliases: []string{"switch-drop"},
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			// parsing args
			value, err := strconv.ParseBool(args[0])
			if err != nil {
				fmt.Println(fmt.Errorf("Only true and false are allowed, error: %v", err))
			}

			lKV, nKV, err := envkeyeditor.EnvKeyEditor("DROP", strconv.FormatBool(value))
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(fmt.Sprintf("last DROP value %s", lKV))
				fmt.Println(fmt.Sprintf("DROP switched to %s", nKV))
			}
		},
	}
)

var (
	switchDebug = &cobra.Command{
		Use:     "switch-debug",
		Short:   "Switch gorm ORM debug",
		Long:    "Switch gorm ORM debug, using debuging mode or not",
		Aliases: []string{"switch-debug"},
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			// parsing args
			value, err := strconv.ParseBool(args[0])
			if err != nil {
				fmt.Println(fmt.Errorf("Only true and false are allowed, error: %v", err))
			}

			lKV, nKV, err := envkeyeditor.EnvKeyEditor("DEBUG", strconv.FormatBool(value))
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(fmt.Sprintf("last DEBUG value %s", lKV))
				fmt.Println(fmt.Sprintf("DEBUG switched to %s", nKV))
			}
		},
	}
)

var (
	generateSecretKey = &cobra.Command{
		Use:     "generate-secret-key",
		Short:   "Generate secret key env",
		Long:    "Generate secret key inside .env file",
		Aliases: []string{"generate-secret-key"},
		Args:    cobra.MinimumNArgs(3),
		Run: func(cmd *cobra.Command, args []string) {
			// Generating secret key using password module
			var chars, nums, syms int
			c, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Println(fmt.Errorf("Char length can't be string: %v", err))
			}
			chars = c

			n, err := strconv.Atoi(args[1])
			if err != nil {
				fmt.Println(fmt.Errorf("Number Digit length can't be string: %v", err))
			}
			nums = n

			s, err := strconv.Atoi(args[2])
			if err != nil {
				fmt.Println(fmt.Errorf("Symbol length can't be string: %v", err))
			}
			syms = s

			generatedPassword, err := password.Generate(chars, nums, syms, false, false)
			if err != nil {
				fmt.Println(err)
			}

			lKV, nKV, err := envkeyeditor.EnvKeyEditor("API_SECRET", generatedPassword)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(fmt.Sprintf("combination %d chrs, %d nums, %d syms", chars, nums, syms))
				fmt.Println(fmt.Sprintf("last API_SECRET value %s", lKV))
				fmt.Println(fmt.Sprintf("API_SECRET switched to %s", nKV))
			}

			lKV, nKV, err = envkeyeditor.EnvKeyEditor("TEST_API_SECRET", generatedPassword)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(fmt.Sprintf("combination %d chrs, %d nums, %d syms", chars, nums, syms))
				fmt.Println(fmt.Sprintf("last TEST_API_SECRET value %s", lKV))
				fmt.Println(fmt.Sprintf("TEST_API_SECRET switched to %s", nKV))
			}

			lKV, nKV, err = envkeyeditor.EnvKeyEditor("SECRET", generatedPassword)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(fmt.Sprintf("combination %d chrs, %d nums, %d syms", chars, nums, syms))
				fmt.Println(fmt.Sprintf("last SECRET value %s", lKV))
				fmt.Println(fmt.Sprintf("SECRET switched to %s", nKV))
			}
		},
	}
)

var (
	webStartCmd = &cobra.Command{
		Use:     "web-start",
		Short:   "Start the service",
		Long:    "Start the service, and connecting to database credential",
		Aliases: []string{"web-start"},
		Run: func(cmd *cobra.Command, args []string) {
			api.Run()
		},
	}
)

func init() {
	cmd.AddCommand(setupDatabase)
	cmd.AddCommand(switchDrop)
	cmd.AddCommand(switchDebug)
	cmd.AddCommand(generateSecretKey)
	cmd.AddCommand(webStartCmd)
}
