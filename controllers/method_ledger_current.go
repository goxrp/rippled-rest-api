package controllers

import (
	"io/ioutil"
	"os"

	"github.com/grokify/simplego/encoding/jsonutil"
	"github.com/grokify/simplego/net/anyhttp"
	"github.com/grokify/simplego/net/httputilmore"
	"github.com/grokify/simplego/type/stringsutil"
	"github.com/rs/zerolog/log"

	ripplenetwork "github.com/wecoinshq/ripple-network"
)

type RequestJsonRpc struct {
	Method string `json:"method"`
}

func (svc *RippleApiService) HandleLedgerCurrentAnyEngine(aRes anyhttp.Response, aReq anyhttp.Request) {
	log.Info().Msg("FUNC_HandleLedgerCurrentAnyEngine__BEGIN")

	qry := aReq.QueryArgs()
	jrpcURL := stringsutil.FirstNotEmptyTrimSpace(
		qry.GetString("jrpcURL"),
		os.Getenv("JSON_RPC_URL"),
		ripplenetwork.GetMainnetPublicJsonRpcUrl())
	log.Info().Str("jsonRpcRemoteURL", jrpcURL)

	proxyResp, err := ProxyApiCall(jrpcURL, MethodLedgerCurrent, []byte{})
	if err == nil {
		respBodyBytes, err := ioutil.ReadAll(proxyResp.Body)
		if err == nil {
			// Content-Type: text/plain; charset=utf-8
			aRes.SetHeader(httputilmore.HeaderContentType, httputilmore.ContentTypeAppJsonUtf8)
			aRes.SetBodyBytes(jsonutil.MustGetSubobjectBytes(respBodyBytes, "result"))
		}
	}
}
