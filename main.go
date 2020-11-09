package main
import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
)

type Article struct {
    Title string `json:"Title"`
    Desc string `json:"desc"`
    Content string `json:"content"`
}

type Article struct {
    Id      string `json:"Id"`
    Title   string `json:"Title"`
    Desc    string `json:"desc"`
    Content string `json:"content"`
}

func returnAllArticles(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: returnAllArticles")
    json.NewEncoder(w).Encode(Articles)
    vars := mux.Vars(r)
    key := vars["id"]

    for _, article := range Articles {
        if article.Id == key {
            json.NewEncoder(w).Encode(article)
        }
    }
}
func createNewArticle(w http.ResponseWriter, r *http.Request) {
     reqBody, _ := ioutil.ReadAll(r.Body)
    var article Article 
    json.Unmarshal(reqBody, &article)
    Articles = append(Articles, article)

    json.NewEncoder(w).Encode(article)
}
func deleteArticle(w http.ResponseWriter, r *http.Request) {
    // once again, we will need to parse the path parameters
    vars := mux.Vars(r)
    // we will need to extract the `id` of the article we
    // wish to delete
    id := vars["id"]

    // we then need to loop through all our articles
    for index, article := range Articles {
        // if our id path parameter matches one of our
        // articles
        if article.Id == id {
            // updates our Articles array to remove the 
            // article
            Articles = append(Articles[:index], Articles[index+1:]...)
        }
    }

}
func handleRequests() {
    myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/", homePage)
    myRouter.HandleFunc("/articles", returnAllArticles)
    myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
    // add our new DELETE endpoint here
    myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
    myRouter.HandleFunc("/article/{id}", returnSingleArticle)
    log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

var Articles []Article
func handleRequests() {
    http.HandleFunc("/", homePage)
    log.Fatal(http.ListenAndServe(":10000", nil))
}

func handleRequests() {
    // creates a new instance of a mux router
    myRouter := mux.NewRouter().StrictSlash(true)
    // replace http.HandleFunc with myRouter.HandleFunc
    myRouter.HandleFunc("/", homePage)
    myRouter.HandleFunc("/all", returnAllArticles)
    // finally, instead of passing in nil, we want
    // to pass in our newly created router as the second
    // argument
    log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
    fmt.Println("Rest API v2.0 - Mux Routers")
    Articles = []Article{
        Articles = []Article{
        Article{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
        Article{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
    }
    handleRequests()
}
