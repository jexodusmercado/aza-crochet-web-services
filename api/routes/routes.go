package routes

import(
    "github.com/gorilla/mux"
	"aza-crochet-web-services/api/controllers"
)

func HandleRequests(r *mux.Router) {

    // match as '/
    r.HandleFunc("/", controllers.HomePage).Methods("GET")

    // match as '/article
    article := r.PathPrefix("/article").Subrouter()
    article.HandleFunc("/", controllers.ReturnAllArticles).Methods("GET")
    article.HandleFunc("/", controllers.CreateNewArticle).Methods("POST")
    article.HandleFunc("/{id}", controllers.DeleteArticle).Methods("DELETE")
    article.HandleFunc("/{id}", controllers.ReturnSingleArticle).Methods("GET")
}
