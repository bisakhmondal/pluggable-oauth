package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"golang.org/x/oauth2/github"
	"golang.org/x/oauth2/google"
)

var (
	googleOauthConfig *oauth2.Config
	fbOauthConfig     *oauth2.Config
	ghOauthConfig     *oauth2.Config

	conf *Config
)

func init() {
	conf = &Config{}
	conf.Parse()

	fmt.Println(conf)

	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:3000/google/callback",
		ClientID:     conf.Credentials.Google.Id,
		ClientSecret: conf.Credentials.Google.Secret,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
	fbOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:3000/facebook/callback",
		ClientID:     conf.Credentials.Facebook.Id,
		ClientSecret: conf.Credentials.Facebook.Secret,
		Scopes:       []string{"email"},
		Endpoint:     facebook.Endpoint,
	}
	ghOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:3000/github/callback",
		ClientID:     conf.Credentials.Github.Id,
		ClientSecret: conf.Credentials.Github.Secret,
		Scopes:       []string{"email"},
		Endpoint:     github.Endpoint,
	}

}

func main() {
	mx := mux.NewRouter()

	mx.HandleFunc("/", handleMain)
	mx.HandleFunc("/google/login", googleLogin)
	mx.HandleFunc("/facebook/login", facebookLogin)
	mx.HandleFunc("/github/login", githubLogin)
	mx.HandleFunc("/google/login/callback", googleCallback)
	mx.HandleFunc("/facebook/login/callback", facebookCallback)
	mx.HandleFunc("/github/login/callback", githubCallback)

	corsH := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "http://127.0.0.1:3000"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"X-Requested-With", "Content-Type", "Authorization", "Content-Length"},
	})

	server := &http.Server{

		Addr:         ":4000",
		Handler:      corsH.Handler(mx),
		ReadTimeout:  8 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  300 * time.Second,
	}
	server.ListenAndServe()
}

func handleMain(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("hello world"))
}

func googleLogin(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(googleOauthConfig.AuthCodeURL("random string")))
}
func facebookLogin(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fbOauthConfig.AuthCodeURL("random string")))
}
func githubLogin(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(ghOauthConfig.AuthCodeURL("random string")))
}

func googleCallback(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("code")
	token, err := googleOauthConfig.Exchange(context.TODO(), code)

	if Giveup(&w, err) {
		return
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if Giveup(&w, err) {
		return
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if Giveup(&w, err) {
		return
	}

	w.WriteHeader(200)
	w.Write(content)
}

func facebookCallback(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("code")
	token, err := fbOauthConfig.Exchange(context.TODO(), code)

	if Giveup(&w, err) {
		return
	}

	resp, err := http.Get("https://graph.facebook.com/me?access_token=" +
		url.QueryEscape(token.AccessToken))
	if Giveup(&w, err) {
		return
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if Giveup(&w, err) {
		return
	}

	w.WriteHeader(200)
	w.Write(content)
}

func githubCallback(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("code")
	token, err := ghOauthConfig.Exchange(context.TODO(), code)

	if Giveup(&w, err) {
		return
	}

	req, _ := http.NewRequest("GET", "https://api.github.com/user", nil)
	header := fmt.Sprintf("token %s", token.AccessToken)

	req.Header.Set("Authorization", header)

	resp, err := http.DefaultClient.Do(req)

	if Giveup(&w, err) {
		return
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if Giveup(&w, err) {
		return
	}

	w.WriteHeader(200)
	w.Write(content)
}

func Giveup(w *http.ResponseWriter, err error) bool {
	if err != nil {
		(*w).WriteHeader(http.StatusInternalServerError)
		(*w).Write([]byte("error occured: " + err.Error()))
		return true
	}
	return false
}
