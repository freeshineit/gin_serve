package proxy

import (
	"fmt"
	"gin_serve/helper"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

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
