package controllers

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "encoding/json"
    "github.com/gorilla/mux"

	"aza-crochet-web-services/api/models"
)

var Articles []models.Article
var Article models.Article

func HomePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
}

func ReturnAllArticles(w http.ResponseWriter, r *http.Request){
    json.NewEncoder(w).Encode(Articles)
}

func ReturnSingleArticle(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
    key := vars["id"]

    for _, article := range Articles {
        if article.ID == key {
            json.NewEncoder(w).Encode(article)
        }
    }
}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    for index, article := range Articles {
        if article.ID == id {
            Articles = append(Articles[:index], Articles[index+1:]...)
        }
    }
}

func CreateNewArticle(w http.ResponseWriter, r *http.Request) {
    reqBody, _ := ioutil.ReadAll(r.Body)
    json.Unmarshal(reqBody, &Article)
    Articles = append(Articles, Article)

    json.NewEncoder(w).Encode(Article)
}