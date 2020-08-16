package handler

import (
	"net/http"
	"net/http/httputil"
)

func Handler(w http.ResponseWriter, r *http.Request) {
    httputil.NewSingleHostReverseProxy("http://pmiappserverprod-env.eba-i3hnze8t.us-east-2.elasticbeanstalk.com/").ServeHTTP(res, req)
}