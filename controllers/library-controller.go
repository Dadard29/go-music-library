package controllers

import (
	"encoding/json"
	"github.com/Dadard29/go-music-library/api"
	"github.com/Dadard29/go-music-library/managers"
	"github.com/Dadard29/go-music-library/models"
	"io/ioutil"
	"net/http"
)


// POST
// Authorization: 	token
// Params: 			None
// Body: 			models.MusicDto
func LibraryPost(w http.ResponseWriter, r *http.Request) {
	t := r.Header.Get(accessTokenKey)
	if !checkToken(t, w) {
		return
	}

	msg := "error deserializing body"
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Error(err.Error())
		api.Api.BuildErrorResponse(http.StatusInternalServerError, msg, w)
		return
	}

	var b models.MusicDto
	err = json.Unmarshal(data, &b)
	if err != nil {
		logger.Error(err.Error())
		api.Api.BuildErrorResponse(http.StatusInternalServerError, msg, w)
		return
	}


	m, msg, err := managers.LibraryManagerCreate(b, t)
	if err != nil {
		logger.Error(err.Error())
		api.Api.BuildErrorResponse(http.StatusInternalServerError, msg, w)
		return
	}

	api.Api.BuildJsonResponse(true, msg, m, w)
}


// GET
// Authorization: 	token
// Params: 			titleParam, artistParam
// Body: 			None
func LibraryGet(w http.ResponseWriter, r *http.Request) {
	t := r.Header.Get(accessTokenKey)
	if !checkToken(t, w) {
		return
	}

	mp := models.MusicParam{
		Title:  r.URL.Query().Get(titleParam),
		Artist: r.URL.Query().Get(artistParam),
		Token:  t,
	}

	if !mp.CheckSanity() {
		api.Api.BuildMissingParameter(w)
		return
	}

	m, msg, err := managers.LibraryManagerGet(mp)
	if err != nil {
		logger.Error(err.Error())
		api.Api.BuildErrorResponse(http.StatusNotFound, msg, w)
		return
	}

	api.Api.BuildJsonResponse(true, msg, m, w)
}


// DELETE
// Authorization: 	token
// Params: 			titleParam, artistParam
// Body: 			None
func LibraryDelete(w http.ResponseWriter, r *http.Request) {
	t := r.Header.Get(accessTokenKey)
	if !checkToken(t, w) {
		return
	}

	mp := models.MusicParam{
		Title:  r.URL.Query().Get(titleParam),
		Artist: r.URL.Query().Get(artistParam),
		Token:  t,
	}

	if !mp.CheckSanity() {
		api.Api.BuildMissingParameter(w)
		return
	}

	m, msg, err := managers.LibrarManagerDelete(mp)
	if err != nil {
		logger.Error(err.Error())
		api.Api.BuildErrorResponse(http.StatusInternalServerError, msg, w)
		return
	}

	api.Api.BuildJsonResponse(true, msg, m, w)
}
