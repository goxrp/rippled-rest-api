package controllers

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/grokify/gohttp/anyhttp"
	"github.com/grokify/mogo/encoding/jsonutil"
	"github.com/grokify/mogo/net/httputilmore"
	"github.com/grokify/mogo/type/stringsutil"
	"github.com/rs/zerolog/log"
	"github.com/valyala/fasthttp"

	gorippled "github.com/goxrp/go-rippled"
	ripplenetwork "github.com/goxrp/ripple-network"
)

const (
	MethodLedger        = "ledger"
	MethodLedgerClosed  = "ledger_closed"
	MethodLedgerCurrent = "ledger_current"
	MethodLedgerData    = "ledger_data"
)

type RequestJSONRPC struct {
	Method string `json:"method"`
}

func (svc *RippleAPIService) HandleApiNetHTTP(res http.ResponseWriter, req *http.Request) {
	log.Info().Msg("FUNC_HandleNetHTTP__BEGIN")
	svc.HandleApiAnyEngine(anyhttp.NewResReqNetHTTP(res, req))
}

func (svc *RippleAPIService) HandleApiFastHTTP(ctx *fasthttp.RequestCtx) {
	log.Info().Msg("HANDLE_FastHTTP")
	svc.HandleApiAnyEngine(anyhttp.NewResReqFastHTTP(ctx))
}

func (svc *RippleAPIService) HandleApiAnyEngine(aRes anyhttp.Response, aReq anyhttp.Request) {
	httpMethod := strings.ToUpper(strings.TrimSpace(string(aReq.Method())))

	acHeaders := strings.TrimSpace(aReq.HeaderString(httputilmore.HeaderAccessControlRequestHeaders))
	if len(acHeaders) > 0 {
		aRes.SetHeader(httputilmore.HeaderAccessControlAllowHeaders, acHeaders)
	}
	aRes.SetHeader(httputilmore.HeaderAccessControlAllowMethods, http.MethodPost)
	aRes.SetHeader(httputilmore.HeaderAccessControlAllowOrigin, "*")
	if httpMethod == http.MethodOptions {
		return
	}

	xrplMethod := strings.ToLower(strings.TrimSpace(aReq.RequestURIVar("rippled_method")))

	log.Info().
		Str("httpMethod", httpMethod).
		Str("xrplMethod", xrplMethod).
		Msg("FUNC_HandleAnyEngine__BEGIN")

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
			jrpcURL = ripplenetwork.GetMainnetPublicJSONRPCURL()
		}
		log.Info().Str("jsonRpcRemoteURL", jrpcURL)

		resp, err := gorippled.DoApiJsonRpcSplit(jrpcURL, xrplMethod, bodyBytes)

		if err == nil {
			respBodyBytes, err := ioutil.ReadAll(resp.Body)
			if err == nil {
				aRes.SetHeader(httputilmore.HeaderAccessControlAllowOrigin, "*")
				aRes.SetHeader(httputilmore.HeaderContentType, httputilmore.ContentTypeAppJSONUtf8)
				aRes.SetBodyBytes(jsonutil.MustGetSubobjectBytes(respBodyBytes, "result"))
			}
		}
	}
}

func (svc *RippleAPIService) HandleGetNoParamsAnyEngine(aRes anyhttp.Response, aReq anyhttp.Request) {
	log.Info().Msg("FUNC_HandleLedgerCurrentAnyEngine__BEGIN")

	proxyResp, err := gorippled.DoApiJsonRpcSplit(
		stringsutil.FirstNotEmptyTrimSpace(
			aReq.QueryArgs().GetString("jrpcURL"),
			os.Getenv("JSON_RPC_URL"),
			ripplenetwork.GetMainnetPublicJSONRPCURL()),
		strings.ToLower(strings.TrimSpace(aReq.RequestURIVar("rippled_method"))),
		[]byte{})

	if err == nil {
		respBodyBytes, err := ioutil.ReadAll(proxyResp.Body)
		if err == nil {
			// Content-Type: text/plain; charset=utf-8
			aRes.SetHeader(httputilmore.HeaderContentType, httputilmore.ContentTypeAppJSONUtf8)
			aRes.SetBodyBytes(jsonutil.MustGetSubobjectBytes(respBodyBytes, "result"))
		}
	}
}

/*

{"error":"unknownCmd","error_code":32,"error_message":"Unknown method.","request":{"command":"account_current"},"status":"error"}

*/
