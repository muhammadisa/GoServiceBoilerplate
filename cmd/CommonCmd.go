package cmd

import (
	"fmt"
	"strconv"

	"github.com/muhammadisa/restful-api-boilerplate/api/utils/envkeyeditor"
	"github.com/sethvargo/go-password/password"
	"github.com/spf13/cobra"
)

var (
	debug       bool
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

			lastKeyValue, newKeyValue, err := envkeyeditor.EnvKeyEditor("DEBUG", strconv.FormatBool(value))
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(fmt.Sprintf("last DEBUG value %s", lastKeyValue))
				fmt.Println(fmt.Sprintf("DEBUG switched to %s", newKeyValue))
			}
		},
	}
)

var (
	generateSecretKey = &cobra.Command{
		Use:     "generate-secret-key",
		Short:   "Generate secret key env",
		Long:    "Generate secret key env",
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

			lastKeyValue, newKeyValue, err := envkeyeditor.EnvKeyEditor("SECRET", generatedPassword)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(fmt.Sprintf("SECRET generated with combination %d chars, %d nums, %d syms", chars, nums, syms))
				fmt.Println(fmt.Sprintf("last SECRET value %s", lastKeyValue))
				fmt.Println(fmt.Sprintf("SECRET switched to %s", newKeyValue))
			}
		},
	}
)

func init() {
	switchDebug.Flags().BoolVarP(&debug, "debug", "d", false, "debug true/false")
	cmd.AddCommand(switchDebug)
	cmd.AddCommand(generateSecretKey)
}
