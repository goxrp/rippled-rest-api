package main

import (
	"fmt"
	"os"

	"github.com/grokify/ripple-rest-api/controllers"
	"github.com/grokify/simplego/fmt/fmtutil"
	"github.com/grokify/simplego/net/http/httpsimple"
	"github.com/grokify/simplego/strconv/strconvutil"
	"github.com/grokify/simplego/type/stringsutil"
)

func main() {
	svc := controllers.RippleApiService{
		Port:              strconvutil.AtoiOrDefault(os.Getenv("PORT"), 8080),
		Engine:            stringsutil.TrimSpaceOrDefault(os.Getenv("HTTP_ENGINE"), "nethttp"),
		DefaultJsonRpcUrl: os.Getenv("RIPPLED_SERVER_JSONRPC_URL"),
		BaseURLPath:       controllers.BaseURLPath}
	fmtutil.PrintJSON(svc)
	fmt.Printf("TRY it out: %s\n", getCmd())

	httpsimple.Serve(svc)
	fmt.Println("DONE")
}

func getCmd() string {
	return `curl -XPOST http://localhost:8080` + controllers.BaseURLPath + `account_channels  -H 'Content-Type: application/json' -d '{"account": "rN7n7otQDd6FczFgLdSqtcsAUxDkw6fzRH","destination_account": "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn","ledger_index": "validated"}'`
}
