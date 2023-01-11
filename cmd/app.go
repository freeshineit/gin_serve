package main

import (
	"go_python_serve/app"
	"log"

	"github.com/spf13/cobra"

	"go_python_serve/app/config"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "server",
	Short: "server is a simple restful api server",
	Long: `server is a simple restful api server
    use help get more ifo`,
	Run: func(cmd *cobra.Command, args []string) {
		port := cmd.Flag("port").Value.String()
		mode := cmd.Flag("mode").Value.String()
		app.RunServer(app.ServerConfig{
			Port: port,
			Mode: mode,
		})
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default: ./app/config/config.yaml)")

	rootCmd.Flags().StringP("port", "p", "8080", "default server port 8080")
	rootCmd.Flags().StringP("mode", "m", "debug", "default  server running in 'debug' mode")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {

	if cfgFile != "" {

	} else {
		c := config.Config{
			Name: cfgFile,
		}

		if err := c.InitConfig(); err != nil {
			panic(err)
		}
		log.Printf("载入配置成功")
		// c.WatchConfig(configChange)
	}
}

// 执行
func main() {
	execute()
}
