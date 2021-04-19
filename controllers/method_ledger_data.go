package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/grokify/simplego/encoding/jsonutil"
	"github.com/grokify/simplego/net/anyhttp"
	"github.com/grokify/simplego/net/httputilmore"
	"github.com/rs/zerolog/log"

	ripplenetwork "github.com/wecoinshq/ripple-network"
)

type LedgerDataRequest struct {
	LedgerHash  string `json:"ledger_hash" schema:"ledger_hash"`
	LedgerIndex string `json:"ledger_index" schema:"ledger_index"`
	Binary      bool   `json:"binary" schema:"binary"`
	Limit       int    `json:"limit" schema:"limit"`
}

func (svc *RippleApiService) HandleLedgerDataAnyEngine(aRes anyhttp.Response, aReq anyhttp.Request) {
	httpMethod := strings.ToUpper(strings.TrimSpace(string(aReq.Method())))
	rippledMethod := strings.ToLower(strings.TrimSpace(aReq.RequestURIVar("rippled_method")))

	log.Info().
		Str("httpMethod", httpMethod).
		Str("rippledMethod", rippledMethod).
		Msg("FUNC_HandleLedgerDataAnyEngine__BEGIN")

	if strings.ToUpper(strings.TrimSpace(string(aReq.Method()))) == http.MethodPost {
		svc.HandleApiAnyEngine(aRes, aReq)
		return
	}
	err := aReq.ParseForm()
	if err != nil {
		aRes.SetStatusCode(400)
		return
	}

	reqData := LedgerDataRequest{}
	qryArgs := aReq.QueryArgs()

	err = decoder.Decode(&reqData, qryArgs.GetURLValues())
	if err != nil {
		aRes.SetStatusCode(400)
		return
	}

	rippledParamsBytes, err := json.Marshal(reqData)
	if err != nil {
		aRes.SetStatusCode(400)
		return
	}

	if err == nil {
		method := aReq.RequestURIVar("rippled_method")
		log.Info().Str("method", method)

		qry := aReq.QueryArgs()
		jrpcURL := strings.TrimSpace(qry.GetString("jrpcURL"))
		if len(jrpcURL) == 0 {
			jrpcURL = strings.TrimSpace(os.Getenv("JSON_RPC_URL"))
		}
		if len(jrpcURL) == 0 {
			jrpcURL = ripplenetwork.GetMainnetPublicJsonRpcUrl()
		}
		log.Info().Str("jsonRpcRemoteURL", jrpcURL)

		proxyResp, err := ProxyApiCall(jrpcURL, method, rippledParamsBytes)
		if err == nil {
			respBodyBytes, err := ioutil.ReadAll(proxyResp.Body)
			if err == nil {
				// Content-Type: text/plain; charset=utf-8
				aRes.SetHeader(httputilmore.HeaderContentType, httputilmore.ContentTypeAppJsonUtf8)
				aRes.SetBodyBytes(jsonutil.MustGetSubobjectBytes(respBodyBytes, "result"))
			}
		}
	}
}
