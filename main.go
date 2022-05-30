package main

import (
	"fmt"
	"os"

	"github.com/grokify/gohttp/httpsimple"
	"github.com/grokify/mogo/fmt/fmtutil"
	"github.com/grokify/mogo/strconv/strconvutil"
	"github.com/grokify/mogo/type/stringsutil"

	"github.com/goxrp/rippled-rest-api/controllers"
)

const RippledJSONRPCURLEnv = "RIPPLED_JSONRPC_URL"

func main() {
	svc := controllers.RippleAPIService{
		Port:              strconvutil.AtoiOrDefault(os.Getenv("PORT"), 8080),
		Engine:            stringsutil.TrimSpaceOrDefault(os.Getenv("HTTP_ENGINE"), "nethttp"),
		DefaultJsonRpcUrl: os.Getenv(RippledJSONRPCURLEnv),
		BaseURLPath:       controllers.BaseURLPath}
	fmtutil.MustPrintJSON(svc)
	fmt.Printf("TRY it out: %s\n", getCmd())

	httpsimple.Serve(svc)
	fmt.Println("DONE")
}

func getCmd() string {
	return `curl -XPOST http://localhost:8080` + controllers.BaseURLPath + `account_channels  -H 'Content-Type: application/json' -d '{"account": "rN7n7otQDd6FczFgLdSqtcsAUxDkw6fzRH","destination_account": "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn","ledger_index": "validated"}'`
}
