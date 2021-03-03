package main

import (
	"Next_fit_algorithm/domain"
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	infoLog    *log.Logger
	CacheTable *domain.CacheTable
}

func main() {
	addr := flag.String("addr", ":4001", "HTTP network address")
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app := &application{
		infoLog: infoLog,
		CacheTable: &domain.CacheTable{
			MemoryBlocks: []*domain.MemoryBlock{
				{
					Id:             1,
					FreeMemoryLeft: 0,
					Description:    "Too Busy to take any task",
				},
				{
					Id:             2,
					FreeMemoryLeft: 305,
					Description:    "Doing some operation",
				},
				{
					Id:             3,
					FreeMemoryLeft: 0,
					Description:    "Too Busy to take any task",
				},
				{
					Id:             4,
					FreeMemoryLeft: 150,
					Description:    "Doing some operation",
				},
				{
					Id:             5,
					FreeMemoryLeft: 70,
					Description:    "Doing some operation",
				},
				{
					Id:             6,
					FreeMemoryLeft: 0,
					Description:    "Too Busy to take any task",
				},
				{
					Id:             7,
					FreeMemoryLeft: 90,
					Description:    "Doing some operation",
				},
			},
		},
	}

	srv := &http.Server{
		Addr:    *addr,
		Handler: app.routes(),
	}
	infoLog.Printf("Starting  server on %v", *addr)
	_ = srv.ListenAndServe()
}
