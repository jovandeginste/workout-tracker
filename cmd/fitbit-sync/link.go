package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"path"
	"time"

	"github.com/skratchdot/open-golang/open"
)

func (fs *fitbitSync) CheckRedirectURI() error {
	u, err := url.Parse(fs.bind)
	if err != nil {
		return err
	}

	u.Path = path.Join(u.Path, "/link")

	fs.redirectURI = u

	return nil
}

func (fs *fitbitSync) getToken() error {
	if err := fs.CheckRedirectURI(); err != nil {
		return err
	}

	go fs.configureServer()

	fs.authCodeURL, fs.state, fs.codeVerifier = fs.fitbitClient.AuthCodeURL(fs.redirectURI.String())

	fmt.Println("Please open the following link in your browser, if it does not open automatically:")
	fmt.Println(fs.authCodeURL.String())

	if err := open.Run(fs.authCodeURL.String()); err != nil {
		log.Println("Could not open browser:", err)
	}

	fmt.Println("Allow this application access, then continue here.")

	<-fs.waitForAuth

	return nil
}

func (fs *fitbitSync) configureServer() {
	http.HandleFunc("/link", fs.linkFunc)

	fs.server = &http.Server{
		Addr:              fs.redirectURI.Host,
		ReadHeaderTimeout: 3 * time.Second,
	}

	log.Fatal(fs.server.ListenAndServe())
}

func (fs *fitbitSync) linkFunc(w http.ResponseWriter, req *http.Request) {
	defer func() { fs.waitForAuth <- true }()

	requestQuery := req.URL.Query()
	if requestQuery.Has("error") {
		http.Error(w, requestQuery.Get("error"), http.StatusInternalServerError)
		return
	}

	if requestQuery.Get("state") != fs.state {
		http.Error(w, "state mismatched.", http.StatusBadRequest)
		return
	}

	linkResp, err := fs.fitbitClient.Link(req.Context(), requestQuery.Get("code"), fs.codeVerifier, fs.redirectURI.String())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fs.cfg.FitbitConfig.UserID = linkResp.UserID
	fs.cfg.Token = linkResp.Token

	if err := fs.cfg.saveConfig(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if _, err := io.WriteString(w, "You may now close this page and check the application's output."); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
