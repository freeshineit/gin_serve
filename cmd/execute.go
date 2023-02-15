package cmd

import (
	"gin_serve/cmd/app"
	"gin_serve/cmd/proxy"
	"gin_serve/cmd/socket"
	"gin_serve/cmd/version"
	"log"
	"os"

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
		var mode string // = gin.DebugMode
		if mode = cmd.Flag("mode").Value.String(); mode == "" {
			mode = os.Getenv("GIN_MODE")
		}

		g.Go(func() error {
			return app.Serve(mode)
		})

		g.Go(func() error {
			return proxy.Serve(mode)
		})

		g.Go(func() error {
			return socket.Serve(mode)
		})

		if err := g.Wait(); err != nil {
			log.Fatal(err)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {

	rootCmd.AddCommand(version.VersionCmd)
	// rootCmd.AddCommand(all.AllCmd)
	rootCmd.AddCommand(app.AppCmd)
	rootCmd.AddCommand(proxy.ProxyCmd)
	rootCmd.AddCommand(socket.SocketCmd)

	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default: ./config/config.toml)")

	rootCmd.PersistentFlags().StringP("mode", "m", "debug", "default server running in 'debug' mode")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if err := config.InitConfig(cfgFile); err != nil {
		panic(err)
	}
	// log.Printf("Load config success ...")
}
