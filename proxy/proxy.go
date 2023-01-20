package proxy

import (
	"fmt"
	"go_python_serve/app/config"
	"go_python_serve/utils"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func Proxy(c *gin.Context) {
	remote, err := url.Parse("http://www.baidu.com")
	if err != nil {
		panic(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)
	// Define the director func
	// This is a good place to log, for example
	proxy.Director = func(req *http.Request) {
		req.Header = c.Request.Header
		req.Host = remote.Host
		req.URL.Scheme = remote.Scheme
		req.URL.Host = remote.Host
		req.URL.Path = c.Param("proxyPath")
		// set UserAgent
	}

	fmt.Println(c.Param("proxyPath"))

	proxy.ServeHTTP(c.Writer, c.Request)
}

func ProxyServer(conf config.ServerConfig) error {
	r := gin.Default()

	if conf.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	//Create a catchall route
	r.Any("/*proxyPath", Proxy)

	// err := r.Run(":" + conf.ProxyPort)

	srv := &http.Server{
		Addr:    ":" + conf.ProxyPort,
		Handler: r,
	}

	log.Printf("proxy server listen: http://localhost:%s\n", conf.ProxyPort)

	err := utils.ListenAndServe(srv)

	if err != nil {
		log.Fatal("Proxy server forced to shutdown:", err)
	}

	log.Println("Proxy server exiting")

	return err
}
