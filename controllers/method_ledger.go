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
	"github.com/grokify/simplego/type/stringsutil"
	"github.com/rs/zerolog/log"

	ripplenetwork "github.com/wecoinshq/ripple-network"
)

type LedgerRequest struct {
	LedgerHash   string `json:"ledger_hash" schema:"ledger_hash"`
	LedgerIndex  string `json:"ledger_index" schema:"ledger_index"`
	Full         bool   `json:"full" schema:"full"`
	Accounts     bool   `json:"accounts" schema:"accounts"`
	Transactions bool   `json:"transactions" schema:"transactions"`
	Expand       bool   `json:"expand" schema:"expand"`
	OwnerFunds   bool   `json:"owner_funds" schema:"owner_funds"`
	Binary       bool   `json:"binary" schema:"binary"`
	Queue        bool   `json:"queue" schema:"queue"`
}

func (svc *RippleApiService) HandleGetWithParamsAnyEngine(aRes anyhttp.Response, aReq anyhttp.Request, reqParams interface{}) {
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

	//reqData := LedgerRequest{}
	qryArgs := aReq.QueryArgs()

	err = decoder.Decode(&reqParams, qryArgs.GetURLValues())
	if err != nil {
		aRes.SetStatusCode(400)
		return
	}

	rippledParamsBytes, err := json.Marshal(reqParams)
	if err != nil {
		aRes.SetStatusCode(400)
		return
	}

	proxyResp, err := ProxyApiCall(
		stringsutil.FirstNotEmptyTrimSpace(
			aReq.QueryArgs().GetString("jrpcURL"),
			os.Getenv("JSON_RPC_URL"),
			ripplenetwork.GetMainnetPublicJsonRpcUrl()),
		rippledMethod,
		rippledParamsBytes)

	if err == nil {
		respBodyBytes, err := ioutil.ReadAll(proxyResp.Body)
		if err == nil {
			// Content-Type: text/plain; charset=utf-8
			aRes.SetHeader(httputilmore.HeaderContentType, httputilmore.ContentTypeAppJsonUtf8)
			aRes.SetBodyBytes(jsonutil.MustGetSubobjectBytes(respBodyBytes, "result"))
		}
	}
}
