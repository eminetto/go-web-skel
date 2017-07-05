package middlewares

import (
	"bytes"
	"net/http"

	"gitlab.com/thecodenation/thecodenation/errors"
	"gitlab.com/thecodenation/thecodenation/session"
	"gitlab.com/thecodenation/thecodenation/storage"
)

//FileUpload faz o tratamento do upload
func FileUpload(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	resume, _, err := r.FormFile("Resume")
	if resume == nil {
		next(w, r)
		return
	}
	defer resume.Close()
	location, err := storage.Dial()
	if err != nil {
		errors.HandleError(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer location.Close()

	session, err := session.Store.Get(r, "auth-session")
	if err != nil {
		errors.HandleError(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	profile, _ := session.Values["profile"].(string)

	container, err := storage.Container(location, profile)
	if err != nil {
		errors.HandleError(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var buff bytes.Buffer
	fileSize, err := buff.ReadFrom(resume)
	_, _ = resume.Seek(0, 0)
	name := profile + ".pdf"
	item, err := container.Put(name, resume, fileSize, nil)
	if err != nil {
		errors.HandleError(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session.Values["resume"] = item.URL().String()
	err = session.Save(r, w)
	if err != nil {
		errors.HandleError(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	next(w, r)
}
