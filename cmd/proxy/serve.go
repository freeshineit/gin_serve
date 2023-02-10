package proxy

import (
	"fmt"
	"gin_serve/config"
	"gin_serve/helper"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var ProxyCmd = &cobra.Command{
	Use:   "proxy",
	Short: "Run proxy serve",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		mode := cmd.Flag("mode").Value.String()
		Serve(mode)
	},
}

func Serve(mode string) error {
	r := gin.Default()

	if mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	port := config.Conf.Proxy.Port

	// Create a catchall route
	r.Any("/*proxyPath", Proxy)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	log.Printf("proxy server listen: %s\n", helper.ColorBlueString("http://localhost:"+port))

	err := helper.ListenAndServe(srv)

	if err != nil {
		log.Fatal("Proxy server forced to shutdown:", err)
	}

	log.Println("Proxy server exiting")

	return err
}

func Proxy(c *gin.Context) {
	// curl http://localhost:8081
	// google 香港
	remote, err := url.Parse("https://www.google.com.hk")

	if err != nil {
		panic(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)
	// Define the director func
	// This is a good place to log, for example
	proxy.Director = func(req *http.Request) {

		header1 := map[string][]string{
			// "Accept-Encoding": {"gzip, deflate"},
			// "Accept-Language": {"en-us"},
			"User-Agent": {"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36"},
		}

		//
		req.Header = helper.MergeMap(c.Request.Header, header1)
		// req.Header = c.Request.Header
		req.Host = remote.Host
		req.URL.Scheme = remote.Scheme
		req.URL.Host = remote.Host
		req.URL.Path = c.Param("proxyPath")
		// set UserAgent
		fmt.Println(req.Header)
	}

	// fmt.Println(c.Param("proxyPath"))

	proxy.ServeHTTP(c.Writer, c.Request)
}
