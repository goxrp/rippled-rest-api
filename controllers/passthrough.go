package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/grokify/simplego/encoding/jsonutil"
	"github.com/grokify/simplego/fmt/fmtutil"
	"github.com/grokify/simplego/net/anyhttp"
	"github.com/grokify/simplego/net/http/httpsimple"
	"github.com/grokify/simplego/net/httputilmore"
	"github.com/grokify/simplego/type/stringsutil"
	"github.com/rs/zerolog/log"
	"github.com/valyala/fasthttp"

	ripplenetwork "github.com/wecoinshq/ripple-network"
)

const (
	MethodLedger        = "ledger"
	MethodLedgerClosed  = "ledger_closed"
	MethodLedgerCurrent = "ledger_current"
	MethodLedgerData    = "ledger_data"
)

type RequestJsonRpc struct {
	Method string `json:"method"`
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
	httpMethod := strings.ToUpper(strings.TrimSpace(string(aReq.Method())))
	xrplMethod := strings.ToLower(strings.TrimSpace(aReq.RequestURIVar("rippled_method")))

	log.Info().
		Str("httpMethod", httpMethod).
		Str("xrplMethod", xrplMethod).
		Msg("FUNC_HandleAnyEngine__BEGIN")

	if httpMethod == http.MethodGet {
		switch xrplMethod {
		case MethodLedger:
			svc.HandleGetWithParamsAnyEngine(aRes, aReq, &LedgerRequest{})
		case MethodLedgerClosed:
			svc.HandleGetNoParamsAnyEngine(aRes, aReq)
		case MethodLedgerCurrent:
			svc.HandleGetNoParamsAnyEngine(aRes, aReq)
		case MethodLedgerData:
			svc.HandleLedgerDataAnyEngine(aRes, aReq)
		}
		return
	}

	bodyBytes, err := aReq.PostBody()
	if err == nil {
		log.Info().Msg(string(bodyBytes))
	}

	if err == nil {
		qry := aReq.QueryArgs()
		jrpcURL := strings.TrimSpace(qry.GetString("jrpcURL"))
		if len(jrpcURL) == 0 {
			jrpcURL = strings.TrimSpace(os.Getenv("JSON_RPC_URL"))
		}
		if len(jrpcURL) == 0 {
			jrpcURL = ripplenetwork.GetMainnetPublicJsonRpcUrl()
		}
		log.Info().Str("jsonRpcRemoteURL", jrpcURL)

		resp, err := ProxyApiCall(jrpcURL, xrplMethod, bodyBytes)

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

func ProxyApiCall(jrpcURL, rippledMethod string, paramsBodyBytes []byte) (*http.Response, error) {
	log.Info().Str("method", rippledMethod)
	jrpcReq := JsonRpcRequest{Method: rippledMethod}

	if len(paramsBodyBytes) > 0 {
		msi := map[string]interface{}{}
		err := json.Unmarshal(paramsBodyBytes, &msi)
		if err == nil {
			jrpcReq.Params = []map[string]interface{}{msi}
			fmtutil.PrintJSON(jrpcReq)
		}
	}

	if len(jrpcURL) == 0 {
		jrpcURL = ripplenetwork.GetMainnetPublicJsonRpcUrl()
	}
	log.Info().Str("jsonRpcRemoteURL", jrpcURL)

	sc := httpsimple.NewSimpleClient(nil, jrpcURL)
	return sc.Do(httpsimple.SimpleRequest{
		Method: http.MethodPost,
		Body:   jrpcReq,
		IsJSON: true})
}

func (svc *RippleApiService) HandleGetNoParamsAnyEngine(aRes anyhttp.Response, aReq anyhttp.Request) {
	log.Info().Msg("FUNC_HandleLedgerCurrentAnyEngine__BEGIN")

	proxyResp, err := ProxyApiCall(
		stringsutil.FirstNotEmptyTrimSpace(
			aReq.QueryArgs().GetString("jrpcURL"),
			os.Getenv("JSON_RPC_URL"),
			ripplenetwork.GetMainnetPublicJsonRpcUrl()),
		strings.ToLower(strings.TrimSpace(aReq.RequestURIVar("rippled_method"))),
		[]byte{})

	if err == nil {
		respBodyBytes, err := ioutil.ReadAll(proxyResp.Body)
		if err == nil {
			// Content-Type: text/plain; charset=utf-8
			aRes.SetHeader(httputilmore.HeaderContentType, httputilmore.ContentTypeAppJsonUtf8)
			aRes.SetBodyBytes(jsonutil.MustGetSubobjectBytes(respBodyBytes, "result"))
		}
	}
}

/*

{"error":"unknownCmd","error_code":32,"error_message":"Unknown method.","request":{"command":"account_current"},"status":"error"}

*/
