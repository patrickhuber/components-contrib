/*
Copyright 2021 The Dapr Authors
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package nethttpadaptor

import (
	"io/ioutil"
	"net"
	"net/http"
	"strconv"

	"github.com/valyala/fasthttp"

	"github.com/dapr/kit/logger"
)

// NewNetHTTPHandlerFunc wraps a fasthttp.RequestHandler in a http.HandlerFunc.
func NewNetHTTPHandlerFunc(logger logger.Logger, h fasthttp.RequestHandler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c := fasthttp.RequestCtx{}
		remoteIP := net.ParseIP(r.RemoteAddr)
		remoteAddr := net.IPAddr{remoteIP, ""} //nolint
		c.Init(&fasthttp.Request{}, &remoteAddr, nil)

		if r.Body != nil {
			reqBody, err := ioutil.ReadAll(r.Body)
			if err != nil {
				logger.Errorf("error reading request body, %+v", err)

				return
			}
			c.Request.SetBody(reqBody)
		}
		c.Request.SetRequestURI(r.URL.RequestURI())
		c.Request.URI().SetScheme(r.URL.Scheme)
		c.Request.SetHost(r.Host)
		c.Request.Header.SetMethod(r.Method)
		c.Request.Header.Set("Proto", r.Proto)
		major := strconv.Itoa(r.ProtoMajor)
		minor := strconv.Itoa(r.ProtoMinor)
		c.Request.Header.Set("Protomajor", major)
		c.Request.Header.Set("Protominor", minor)
		c.Request.Header.SetContentType(r.Header.Get("Content-Type"))
		c.Request.Header.SetContentLength(int(r.ContentLength))
		c.Request.Header.SetReferer(r.Referer())
		c.Request.Header.SetUserAgent(r.UserAgent())
		for _, cookie := range r.Cookies() {
			c.Request.Header.SetCookie(cookie.Name, cookie.Value)
		}
		for k, v := range r.Header {
			for _, i := range v {
				c.Request.Header.Add(k, i)
			}
		}

		ctx := r.Context()
		reqCtx, ok := ctx.(*fasthttp.RequestCtx)
		if ok {
			reqCtx.VisitUserValues(func(k []byte, v interface{}) {
				c.SetUserValueBytes(k, v)
			})
		}

		h(&c)

		c.Response.Header.VisitAll(func(k []byte, v []byte) {
			w.Header().Add(string(k), string(v))
		})
		c.Response.BodyWriteTo(w)
	})
}
