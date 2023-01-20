package main

import (
	"gin_server/app"
	"gin_server/proxy"
	"log"

	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"

	"gin_server/app/config"
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
		port := cmd.Flag("port").Value.String()
		mode := cmd.Flag("mode").Value.String()
		proxyPort := cmd.Flag("proxy-port").Value.String()

		log.Printf("app version %s", app.Version)

		ServerConfig := config.ServerConfig{
			Port:      port,
			Mode:      mode,
			ProxyPort: proxyPort,
		}

		g.Go(func() error {
			return app.RunServer(ServerConfig)

		})

		g.Go(func() error {
			return proxy.ProxyServer(ServerConfig)
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

	rootCmd.Flags().StringP("port", "p", "8080", "default server port 8080")
	rootCmd.Flags().StringP("proxy-port", "x", "8081", "default  proxy server port 8081")
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
