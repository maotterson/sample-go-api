package controllers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	. "github.com/maotterson/sample-go-api/models"
	"github.com/maotterson/sample-go-api/store"
)

func GetArticles(w http.ResponseWriter, r *http.Request){
	//set response header
	w.Header().Set("Content-Type","application/json")

	//get the articles (from the data store)
	articles := store.Articles

	//send the response with all articles in the body
	json.NewEncoder(w).Encode(articles)
}
func GetArticle(w http.ResponseWriter, r *http.Request){
	//set response header
	w.Header().Set("Content-Type","application/json")

	//find the article by passed in id
	params := mux.Vars(r)
	articles := store.Articles
	for _, item := range articles {
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	//if no article with the id is found
	json.NewEncoder(w).Encode(&Article{})
}
func CreateArticle(w http.ResponseWriter, r *http.Request){
	//set response header
	w.Header().Set("Content-Type","application/json")

	//create new article from the request body & assign id
	var article Article
	_ = json.NewDecoder(r.Body).Decode(&article)
	article.ID = strconv.Itoa(rand.Intn(100000))

	//add article to data store
	store.Articles = append(store.Articles, article)

	//send response with the article in the body
	json.NewEncoder(w).Encode(article)
}
func EditArticle(w http.ResponseWriter, r *http.Request){
	//set response header
	w.Header().Set("Content-Type","application/json")

	//find the article by passed in id
	params := mux.Vars(r)
	articles := store.Articles
	for index, item := range articles {
		if item.ID == params["id"]{
			//delete existing entry
			articles = append(articles[:index], articles[index+1:]...)
			
			//create new entry to replace it
			var article Article
			_ = json.NewDecoder(r.Body).Decode(&article)
			article.ID = params["id"]
		
			//add article to data store
			store.Articles = append(store.Articles, article)
		
			//send response with the article in the body
			json.NewEncoder(w).Encode(article)
			return

		}
	}

	//if no article with the id is found
	json.NewEncoder(w).Encode(&Article{})
}
func DeleteArticle(w http.ResponseWriter, r *http.Request){
	//set response header
	w.Header().Set("Content-Type","application/json")

	//find the article by passed in id and delete it
	params := mux.Vars(r)
	articles := store.Articles
	p := &store.Articles
	for index, item := range articles {
		if item.ID == params["id"]{
			articles = append(articles[:index], articles[index+1:]...)
			break
		}
	}
	*p = articles
	json.NewEncoder(w).Encode(*p)
}