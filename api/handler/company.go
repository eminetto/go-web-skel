package handler

import (
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/eminetto/go-web-skel/pkg/company"
	"github.com/eminetto/go-web-skel/pkg/middleware"
	"github.com/gorilla/mux"
)

func companyFindAll(service company.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		return
	})
}

//MakeCompanyHandlers make url handlers
func MakeCompanyHandlers(r *mux.Router, service company.Service) {
	r.Handle("/v1/company", negroni.New(
		negroni.HandlerFunc(middleware.Cors),
		negroni.HandlerFunc(middleware.IsAuthenticated),
		negroni.Wrap(companyFindAll(service)),
	)).Methods("GET", "OPTIONS")
}
