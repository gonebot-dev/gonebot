package gonebot

import (
	"log"
	"net/http"
	"net/http/pprof"
	"testing"

	"github.com/gonebot-dev/gonebot/plugins/builtinplugins"
)

func TestMain(m *testing.M) {
	mux := http.NewServeMux()
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	go func() { log.Fatal(http.ListenAndServe(":8080", mux)) }()

	LoadPlugin(builtinplugins.Echo)
	StartBackend("onebot11")
}
