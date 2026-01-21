package main

import (
	"log"
	"net/http"

	"github.com/SoyebSarkar/Hiberstack/internal/config"
	"github.com/SoyebSarkar/Hiberstack/internal/engine/typesense"
	"github.com/SoyebSarkar/Hiberstack/internal/lifecycle"
	"github.com/SoyebSarkar/Hiberstack/internal/proxy"
	"github.com/SoyebSarkar/Hiberstack/internal/scheduler"
	"github.com/SoyebSarkar/Hiberstack/internal/state"
)

func main() {
	cfg := config.Load()
	ts := typesense.New(cfg.Typesense.URL, cfg.Typesense.APIKey)

	stateStore, err := state.NewSQLite("./state.db")
	if err != nil {
		log.Fatal(err)
	}
	lifecycleMgr := lifecycle.NewManager(
		ts,
		cfg.SnapshotDir,
		stateStore,
	)
	scheduler := scheduler.New(
		stateStore,
		lifecycleMgr,
		cfg.OffloadAfter,
	)

	scheduler.Start()

	proxy, err := proxy.New(cfg.Typesense.URL, lifecycleMgr, stateStore)
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	// 1️⃣ Register admin routes FIRST
	registerAdmin(mux, ts, cfg.SnapshotDir, lifecycleMgr, stateStore)

	// 2️⃣ Attach proxy as fallback
	mux.Handle("/", proxy)
	handler := loggingMiddleware(mux)

	// 3️⃣ Start server with mux
	http.ListenAndServe(":"+cfg.Port, handler)

}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf(
			"%s %s %s",
			r.Method,
			r.URL.Path,
			r.RemoteAddr,
		)
		next.ServeHTTP(w, r)
	})
}
