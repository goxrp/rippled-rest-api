package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/buaazp/fasthttprouter"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gorilla/mux"
	"github.com/grokify/gohttp/anyhttp"
	"github.com/grokify/gohttp/httpsimple"
	"github.com/grokify/mogo/net/httputilmore"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fastjson"
)

// "github.com/gorilla/schema"
// var decoder = schema.NewDecoder()

const BaseURLPath = "/api/v1/"

func (svc *RippleAPIService) HandleAPIInfoFastHTTP(ctx *fasthttp.RequestCtx) {
	svc.HandleAPIInfoAnyEngine(anyhttp.NewResReqFastHTTP(ctx))
}

func (svc *RippleAPIService) HandleAPIInfoNetHTTP(res http.ResponseWriter, req *http.Request) {
	svc.HandleAPIInfoAnyEngine(anyhttp.NewResReqNetHTTP(res, req))
}

func (svc *RippleAPIService) HandleAPIInfoAnyEngine(aRes anyhttp.Response, aReq anyhttp.Request) {
	var apiInfo = openapi3.Info{
		Title:   "GoXRP Rippled REST API Proxy",
		Version: "1.0.0",
	}
	bytes, _ := json.Marshal(apiInfo)
	// aRes.SetStatusCode(http.StatusOK)
	// aRes.SetStatusCode(400)
	aRes.SetHeader(httputilmore.HeaderContentType, httputilmore.ContentTypeAppJSONUtf8)
	aRes.SetBodyBytes(bytes)
}

type RippleAPIService struct {
	Port              int
	Engine            string
	DefaultJSONRPCURL string
	BaseURLPath       string
}

func (svc RippleAPIService) PortInt() int {
	return svc.Port
}

func (svc RippleAPIService) HTTPEngine() string {
	return svc.Engine
}

func (svc RippleAPIService) Router() http.Handler {
	mux := mux.NewRouter()
	mux.HandleFunc("/test", http.HandlerFunc(httpsimple.HandleTestNetHTTP))
	mux.HandleFunc("/test/", http.HandlerFunc(httpsimple.HandleTestNetHTTP))
	mux.HandleFunc("/api", http.HandlerFunc(svc.HandleAPIInfoNetHTTP))
	mux.HandleFunc("/api/", http.HandlerFunc(svc.HandleAPIInfoNetHTTP))
	mux.HandleFunc(BaseURLPath+"{rippled_method}", http.HandlerFunc(svc.HandleAPINetHTTP))
	mux.HandleFunc(BaseURLPath+"{rippled_method}/", http.HandlerFunc(svc.HandleAPINetHTTP))
	return mux
}

func (svc RippleAPIService) RouterFast() *fasthttprouter.Router {
	router := fasthttprouter.New()
	router.POST(BaseURLPath+":rippled_method", svc.HandleAPIFastHTTP)
	router.POST(BaseURLPath+":rippled_method/", svc.HandleAPIFastHTTP)
	return router
}

func SubobjectBytes(data []byte, key string) ([]byte, error) {
	val, err := fastjson.ParseBytes(data)
	if err != nil {
		return []byte{}, err
	}
	obj := val.GetObject(key)
	return obj.MarshalTo([]byte{}), nil
}
