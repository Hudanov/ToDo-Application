package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func wrapJwt(
	jwt *JWTService,
	f func(http.ResponseWriter, *http.Request, *JWTService),
) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		f(rw, r, jwt)
	}
}

func main() {
	r := mux.NewRouter()

	headersOK := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	originsOK := handlers.AllowedOrigins([]string{"*"})
	methodsOK := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS", "DELETE", "PUT"})

	users := NewInMemoryUserStorage()
	userService := UserService{
		repository: users,
	}

	jwtService, err := NewJWTService("pubkey.rsa", "privkey.rsa")
	if err != nil {
		panic(err)
	}

	r.HandleFunc("/user/signup", userService.Register).Methods(http.MethodPost, "OPTIONS")
	r.HandleFunc("/user/signin", wrapJwt(jwtService, userService.JWT)).Methods(http.MethodPost)
	r.HandleFunc("/todo/lists", jwtService.jwtAuth(users, userService.createNewListHandler)).Methods(http.MethodPost, "OPTIONS")
	r.HandleFunc("/todo/lists", jwtService.jwtAuth(users, userService.getListsHandler)).Methods(http.MethodGet)
	r.HandleFunc("/todo/lists/{list_id}", jwtService.jwtAuth(users, userService.deleteListHandler)).Methods(http.MethodDelete)
	r.HandleFunc("/todo/lists/{list_id}", jwtService.jwtAuth(users, userService.updateListHandler)).Methods(http.MethodPut)
	r.HandleFunc("/todo/lists/{list_id}/tasks", jwtService.jwtAuth(users, userService.createNewTaskHandler)).Methods(http.MethodPost)
	r.HandleFunc("/todo/lists/{list_id}/tasks", jwtService.jwtAuth(users, userService.getTasksHandler)).Methods(http.MethodGet)
	r.HandleFunc("/todo/lists/{list_id}/tasks/{task_id}", jwtService.jwtAuth(users, userService.deleteTaskHandler)).Methods(http.MethodDelete)
	r.HandleFunc("/todo/lists/{list_id}/tasks/{task_id}", jwtService.jwtAuth(users, userService.updateTaskHandler)).Methods(http.MethodPut)

	srv := http.Server{
		Addr:    ":8080",
		Handler: handlers.CORS(originsOK, headersOK, methodsOK)(r),
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	go func() {
		<-interrupt
		ctx, cancel := context.WithTimeout(context.Background(),
			5*time.Second)
		defer cancel()
		srv.Shutdown(ctx)
	}()
	log.Println("Server started, hit Ctrl+C to stop")
	srv.ListenAndServe()
	if err != nil {
		log.Println("Server exited with error:", err)
	}
	log.Println("Good bye :)")

}
