package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/buaazp/fasthttprouter"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gorilla/mux"
	"github.com/grokify/simplego/encoding/jsonutil"
	"github.com/grokify/simplego/fmt/fmtutil"
	"github.com/grokify/simplego/net/anyhttp"
	"github.com/grokify/simplego/net/http/httpsimple"
	"github.com/grokify/simplego/net/httputilmore"
	"github.com/grokify/simplego/strconv/strconvutil"
	"github.com/grokify/simplego/type/stringsutil"
	"github.com/rs/zerolog/log"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fastjson"

	ripplenetwork "github.com/wecoinshq/ripple-network"
)

func (svc *RippleApiService) HandleApiInfoFastHTTP(ctx *fasthttp.RequestCtx) {
	svc.HandleApiInfoAnyEngine(anyhttp.NewResReqFastHttp(ctx))
}

func (svc *RippleApiService) HandleApiInfoNetHTTP(res http.ResponseWriter, req *http.Request) {
	svc.HandleApiInfoAnyEngine(anyhttp.NewResReqNetHttp(res, req))
}

func (svc *RippleApiService) HandleApiInfoAnyEngine(aRes anyhttp.Response, aReq anyhttp.Request) {
	var apiInfo = openapi3.Info{
		Title:   "WeCoins Ripple API Proxy",
		Version: "1.0.0",
	}
	bytes, _ := json.Marshal(apiInfo)
	// aRes.SetStatusCode(http.StatusOK)
	// aRes.SetStatusCode(400)
	aRes.SetHeader(httputilmore.HeaderContentType, httputilmore.ContentTypeAppJsonUtf8)
	aRes.SetBodyBytes(bytes)
}

type JsonRpcRequest struct {
	Method string                   `json:"method"`
	Params []map[string]interface{} `json:"params"`
}

type RippleApiService struct {
	Port              int
	Engine            string
	DefaultJsonRpcUrl string
}

func (svc *RippleApiService) HandleApiNetHTTP(res http.ResponseWriter, req *http.Request) {
	log.Info().Msg("FUNC_HandleNetHTTP__BEGIN")
	svc.HandleApiAnyEngine(anyhttp.NewResReqNetHttp(res, req))
}

func (svc *RippleApiService) HandleApiFastHTTP(ctx *fasthttp.RequestCtx) {
	log.Info().Msg("HANDLE_FastHTTP")
	svc.HandleApiAnyEngine(anyhttp.NewResReqFastHttp(ctx))
}

func (svc *RippleApiService) HandleApiAnyEngine(aRes anyhttp.Response, aReq anyhttp.Request) {
	log.Info().Msg("FUNC_HandleAnyEngine__BEGIN")
	log.Info().Str("method", aReq.RequestURIVar("rippled_method"))

	bodyBytes, err := aReq.PostBody()
	if err == nil {
		log.Info().Msg(string(bodyBytes))
	}

	if err == nil {
		method := aReq.RequestURIVar("rippled_method")
		log.Info().Str("method", method)
		jrpcReq := JsonRpcRequest{Method: method}
		msi := map[string]interface{}{}
		err := json.Unmarshal(bodyBytes, &msi)
		if err == nil {
			jrpcReq.Params = []map[string]interface{}{msi}
			fmtutil.PrintJSON(jrpcReq)
		}

		jrpcReq = JsonRpcRequest{
			Method: method,
			Params: []map[string]interface{}{msi}}

		qry := aReq.QueryArgs()
		jrpcURL := strings.TrimSpace(qry.GetString("jrpcURL"))
		if len(jrpcURL) == 0 {
			jrpcURL = strings.TrimSpace(os.Getenv("JSON_RPC_URL"))
		}
		if len(jrpcURL) == 0 {
			jrpcURL = ripplenetwork.GetMainnetPublicJsonRpcUrl()
		}
		log.Info().Str("jsonRpcRemoteURL", jrpcURL)

		sc := httpsimple.NewSimpleClient(nil, jrpcURL)
		resp, err := sc.Do(httpsimple.SimpleRequest{
			Method: http.MethodPost,
			Body:   jrpcReq,
			IsJSON: true})
		if err == nil {
			respBodyBytes, err := ioutil.ReadAll(resp.Body)
			if err == nil {
				// Content-Type: text/plain; charset=utf-8
				aRes.SetHeader(httputilmore.HeaderContentType, httputilmore.ContentTypeAppJsonUtf8)
				aRes.SetBodyBytes(jsonutil.MustGetSubobjectBytes(respBodyBytes, "result"))
			}
		}
	}
}

func (svc RippleApiService) PortInt() int {
	return svc.Port
}

func (svc RippleApiService) HttpEngine() string {
	return svc.Engine
}

func (svc RippleApiService) Router() http.Handler {
	mux := mux.NewRouter()
	mux.HandleFunc("/test", http.HandlerFunc(httpsimple.HandleTestNetHTTP))
	mux.HandleFunc("/test/", http.HandlerFunc(httpsimple.HandleTestNetHTTP))
	mux.HandleFunc("/api", http.HandlerFunc(svc.HandleApiInfoNetHTTP))
	mux.HandleFunc("/api/", http.HandlerFunc(svc.HandleApiInfoNetHTTP))
	mux.HandleFunc("/api/ledger/v1.0/{rippled_method}", http.HandlerFunc(svc.HandleApiNetHTTP))
	mux.HandleFunc("/api/ledger/v1.0/{rippled_method}/", http.HandlerFunc(svc.HandleApiNetHTTP))
	return mux
}

func (svc RippleApiService) RouterFast() *fasthttprouter.Router {
	router := fasthttprouter.New()
	router.POST("/api/ledger/v1.0/:rippled_method", svc.HandleApiFastHTTP)
	router.POST("/api/ledger/v1.0/:rippled_method/", svc.HandleApiFastHTTP)
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

func main() {
	svc := RippleApiService{
		Port:              strconvutil.AtoiOrDefault(os.Getenv("PORT"), 8080),
		Engine:            stringsutil.TrimSpaceOrDefault(os.Getenv("HTTP_ENGINE"), "nethttp"),
		DefaultJsonRpcUrl: os.Getenv("RIPPLED_SERVER_JSONRPC_URL")}
	fmtutil.PrintJSON(svc)

	httpsimple.Serve(svc)
	fmt.Println("DONE")
}
