package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/xshoji/chronicling-america-ui/dao"
	"github.com/zserge/lorca"
	"net"
	"net/http"
	"os"
	"os/signal"
)

func main() {

	//--------------------
	// 1. logger setting
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
	log.SetReportCaller(true)

	//--------------------
	// 2. create lorca
	ui, _ := lorca.New("", "", 1200, 840)
	deferFunc := func() {
		e := ui.Close()
		if e != nil {
			log.Error("Cannot close ui object.")
			log.Error(e)
			os.Exit(1)
		}
		log.Info("====== Exit Application ======")
	}
	// Make kill signal channel
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Kill, os.Interrupt)
	// set kill signal handling
	go func() {
		<-signals
		deferFunc()
		os.Exit(0)
	}()
	// set defer
	defer deferFunc()

	//--------------------
	// - [web applications - How do I serve CSS and JS in Go Lang - Stack Overflow](https://stackoverflow.com/questions/43601359/how-do-i-serve-css-and-js-in-go-lang)
	// 3. start server and application
	ln, err := net.Listen("tcp", "127.0.0.1:0") //監視するポートを設定します。
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	defer ln.Close()
	go http.Serve(ln, http.FileServer(FS))
	ui.Load(fmt.Sprintf("http://%s", ln.Addr()))
	ui.Bind("Search", dao.Search)
	ui.Bind("GetKeysSearchResponse", dao.GetKeysSearchResponse)
	log.Info("====== Start Application ======")
	<-ui.Done()
}
