/*
Copyright © 2022 Jimmy Ungerman jimmy.ungerman@gmail.com

*/
package cmd

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
	"github.com/spf13/cobra"
)

// A Response struct to map the Entire Response
type Response struct {
    Pokemon []Pokemon `json:"results"`
	
}

// A Pokemon Struct to map every pokemon to.
type Pokemon struct {
	Name	string	`json:"name"`
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pokeCLI",
	Short: "CLI for Pokemon API",
	Long: `CLI to interact with the Pokemon API and get details about any pokemon.
	
	`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) { 
		const baseURL = "https://pokeapi.co/api/v2/pokemon-species?limit=100000&offset=0"
		response, err := http.Get(baseURL)
		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}
	
		responseData, err := ioutil.ReadAll(response.Body)

		if err != nil {
			log.Fatal(err)
		}
		var responseObject Response
		json.Unmarshal(responseData, &responseObject)

		for i := 0; i < len(responseObject.Pokemon); i++ {
			fmt.Println(responseObject.Pokemon[i].Name)
		}
	
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.pokeCLI.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

