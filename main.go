package main

import (
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/buaazp/fasthttprouter"
	"github.com/gorilla/mux"
	"github.com/grokify/ripple-api/network"
	"github.com/grokify/simplego/fmt/fmtutil"
	"github.com/grokify/simplego/net/anyhttp"
	"github.com/grokify/simplego/net/http/httpsimple"
	log "github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
)

type JsonRpcRequest struct {
	Method string                   `json:"method"`
	Params []map[string]interface{} `json:"params"`
}

type RippleApiService struct {
	Port   int
	Engine string
}

func (svc *RippleApiService) HandleTestNetHTTP(res http.ResponseWriter, req *http.Request) {
	log.Info("FUNC_HandleNetHTTP__BEGIN")
	fmt.Fprintf(res, `{"foo":"%q"}`, html.EscapeString(req.URL.Path))
}

func (svc *RippleApiService) HandleApiNetHTTP(res http.ResponseWriter, req *http.Request) {
	log.Info("FUNC_HandleNetHTTP__BEGIN")
	svc.HandleApiAnyEngine(anyhttp.NewResReqNetHttp(res, req))
}

func (svc *RippleApiService) HandleApiFastHTTP(ctx *fasthttp.RequestCtx) {
	log.Info("HANDLE_FastHTTP")
	svc.HandleApiAnyEngine(anyhttp.NewResReqFastHttp(ctx))
}

func (svc *RippleApiService) HandleApiAnyEngine(aRes anyhttp.Response, aReq anyhttp.Request) {
	log.Info("FUNC_HandleAnyEngine__BEGIN")

	log.Info(fmt.Sprintf("PATH_VAR [%s]", aReq.RequestURIVar("rippled_method")))

	bodyBytes, err := aReq.PostBody()
	if err == nil {
		log.Info(string(bodyBytes))
	}

	if err == nil {
		method := aReq.RequestURIVar("rippled_method")
		log.Info(fmt.Sprintf("METHOD [%s]", method))
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
			jrpcURL = network.RipplePublicServer1JsonRpcURL
		}
		log.Info(fmt.Sprintf("FUNC_HandleAnyEngine__RemoteUrl [%s]", jrpcURL))

		sc := httpsimple.NewSimpleClient(nil, jrpcURL)
		resp, err := sc.Do(httpsimple.SimpleRequest{
			Method: http.MethodPost,
			URL:    "",
			Body:   jrpcReq,
			IsJSON: true})
		if err == nil {
			respBodyBytes, err := ioutil.ReadAll(resp.Body)
			if err == nil {
				aRes.SetBodyBytes(respBodyBytes)
			}
		}
	}
}

/*
curl -H 'Content-Type: application/json' -d '{"method":"account_info","params":[{"account":"r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59","strict":true,"ledger_index":"current","queue":true}]}' https://s1.ripple.com:51234/
*/

func (svc RippleApiService) PortInt() int {
	return svc.Port
}

func (svc RippleApiService) HttpEngine() string {
	return svc.Engine
}

func (svc RippleApiService) Router() http.Handler {
	mux := mux.NewRouter()
	mux.HandleFunc("/", http.HandlerFunc(svc.HandleTestNetHTTP))
	mux.HandleFunc("/api/v1/{rippled_method}", http.HandlerFunc(svc.HandleApiNetHTTP))
	mux.HandleFunc("/api/v1/{rippled_method}/", http.HandlerFunc(svc.HandleApiNetHTTP))
	return mux
}

func (svc RippleApiService) RouterFast() *fasthttprouter.Router {
	router := fasthttprouter.New()
	router.POST("/api/v1/:rippled_method", svc.HandleApiFastHTTP)
	router.POST("/api/v1/:rippled_method/", svc.HandleApiFastHTTP)
	return router
}

func main() {
	svc := RippleApiService{
		Port:   8080,
		Engine: os.Getenv("HTTP_ENGINE")}
	fmtutil.PrintJSON(svc)

	httpsimple.Serve(svc)
	fmt.Println("DONE")
}

/*

curl -XPOST 'http://localhost:8080/api/account_channels/?abc-def&ghi-jkl#anchor' -d '{"account":"rN7n7otQDd6FczFgLdSqtcsAUxDkw6fzRH","destination_account":"rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn","ledger_index":"validated"}'


{
        "account": "rN7n7otQDd6FczFgLdSqtcsAUxDkw6fzRH",
        "destination_account": "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
        "ledger_index": "validated"
    }

{
    "method": "account_channels",
    "params": [{
        "account": "rN7n7otQDd6FczFgLdSqtcsAUxDkw6fzRH",
        "destination_account": "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
        "ledger_index": "validated"
    }]
}


*/
