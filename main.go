package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
)
func main(){
	port := "8080"
	if fromEnv := os.Getenv("PORT");fromEnv != ""{
		port= fromEnv
	}
	log.Printf("starting up on http://localhost:%s",port)

router:=chi.NewRouter()
router.Use(middleware.Logger)
router.Get("/",func(w http.ResponseWriter,r *http.Request){
_,err:=w.Write([]byte("a simple api"))
if err!=nil{
	log.Println(err)
}
})
log.Fatal(http.ListenAndServe(":"+port,router))

}