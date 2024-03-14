package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http/httputil"
	"net/url"
	"sync"
)

var (
	backends = []string{viper.GetString("balance.url1"), viper.GetString("balance.url2"), viper.GetString("balance.url3")}
)

type LoadBalancer struct {
	mutex   sync.Mutex
	servers []string
	index   int
}

func NewLoadBalancer() *LoadBalancer {
	lb := &LoadBalancer{
		servers: backends,
	}
	return lb
}

func (lb *LoadBalancer) chooseServer() string {
	lb.mutex.Lock()
	defer lb.mutex.Unlock()

	server := lb.servers[lb.index]
	lb.index = (lb.index + 1) % len(lb.servers) //轮询使用三个地址进行分流

	return server
}
func (lb *LoadBalancer) HandleRequest(c *gin.Context) {
	server := lb.chooseServer()
	proxy := NewReverseProxy(server)
	proxy.ServeHTTP(c.Writer, c.Request)
}
func NewReverseProxy(target string) *httputil.ReverseProxy {
	url, _ := url.Parse(target)
	return httputil.NewSingleHostReverseProxy(url)
}
