package routes

import (
	"context"
	"gingonic-api/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(g *gin.Engine) {
	//Routes
	g.GET("/api/books", handlers.GetBooks)
	g.GET("/api/book/:id", handlers.GetBook)
	g.POST("/api/book", handlers.AddBook)
	g.PUT("/api/book/:id", handlers.UpdateBook)

	//Listening on localhost:8080/
	srv := &http.Server{
		Addr:    ":8080",
		Handler: g,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 50)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Server shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", err)
	}

	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds")
	}

	log.Println("Serving exiting")
}
