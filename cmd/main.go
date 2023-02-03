package main

import (
	"gin_serve/cmd/app"
	"gin_serve/cmd/proxy"
	"log"

	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"

	"gin_serve/config"
)

var cfgFile string

var (
	g errgroup.Group
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "server",
	Short: "server is a simple restful api server",
	Long: `server is a simple restful api server
    use help get more ifo`,
	Run: func(cmd *cobra.Command, args []string) {
		mode := cmd.Flag("mode").Value.String()

		// log.Printf("server version %s", version)

		g.Go(func() error {
			return app.Serve(mode)
		})

		g.Go(func() error {
			return proxy.Serve(mode)
		})

		if err := g.Wait(); err != nil {
			log.Fatal(err)
		}
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
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default: ./app/config/config.toml)")

	rootCmd.Flags().StringP("mode", "m", "debug", "default  server running in 'debug' mode")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {

	if cfgFile != "" {

	} else {
		if err := config.InitConfig(cfgFile); err != nil {
			panic(err)
		}
		log.Printf("Load config success ...")
		// c.WatchConfig(configChange)
	}
}

// 执行
func main() {
	execute()
}
