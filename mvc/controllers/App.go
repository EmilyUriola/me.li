package controllers

import (
	"encoding/json"
	"hash/fnv"
	"html/template"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
	"github.com/me.li/mvc/models"
)

type UrlStruct struct {
	Long  string `json:"long,omitempty"`
	Short string `json:"short,omitempty"`
}

var hostName string = "http://localhost:8011/"
var notifyType int
var notifyMsg string
var t *template.Template

func Hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func CreateUrl(w http.ResponseWriter, r *http.Request) {
	var url UrlStruct
	err := json.NewDecoder(r.Body).Decode(&url)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "Bad Request."}`))
		return
	} else {
		shortUrl := models.RedisDbSave(url.Long)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message": "URL shortified.", "short_url": "` + hostName + shortUrl + `"}`))
	}
}

func CreateUrls(w http.ResponseWriter, r *http.Request) {
	urls := make([]UrlStruct, 0)
	cad := make([]string, 0)
	err := json.NewDecoder(r.Body).Decode(&urls)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "Bad Request."}`))
		return
	} else {
		for _, value := range urls {
			cad = append(cad, value.Long)
		}

		shortUrls := models.RedisDbSaveBulks(cad, hostName)
		encjson, _ := json.Marshal(shortUrls)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message": "URL shortified.", "short_url": "` + string(encjson) + `"}`))
	}
}

func GetUrl(w http.ResponseWriter, r *http.Request) {
	shortCode := mux.Vars(r)["id"]
	shortUrl, err := models.RedisDbGet(shortCode)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "URL shortified Not Found."}`))
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"long_url": "` + shortUrl + `", "short_url": "` + hostName + shortCode + `"}`))
	}
}

func DeleteUrl(w http.ResponseWriter, r *http.Request) {
	shortCode := mux.Vars(r)["id"]
	shortUrl, err := models.RedisDbDel(shortCode)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "Bad Request."}`))
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Delete URL shortified.", "short_url": "` + hostName + shortUrl + `"}`))
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	shortCode := r.URL.Path[1:]
	notifyType = 0

	if len(shortCode) != 0 {
		redirectUrl, err := models.RedisDbGet(shortCode)
		if err != nil {
			redirectUrl = hostName + "/error404/" + shortCode
		}
		RedirectTo(w, r, redirectUrl) // redirect to long url

		return
	}
}

func RedirectTo(w http.ResponseWriter, r *http.Request, urlStr string) {
	http.Redirect(w, r, urlStr, http.StatusFound)
}

func ValidateURL(longUrl string) error {
	_, err := url.ParseRequestURI(longUrl)
	return err
}
