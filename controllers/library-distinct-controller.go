package controllers

import (
	"github.com/Dadard29/go-music-library/api"
	"github.com/Dadard29/go-music-library/managers"
	"net/http"
)

// GET
// Authorization: 	token
// Params: 			None
// Body: 			None

// get the list of albums in library
func LibraryAlbumsGet(w http.ResponseWriter, r *http.Request) {
	t := r.Header.Get(accessTokenKey)
	if !checkToken(t, w) {
		return
	}

	albumList, msg, err := managers.LibraryAlbumManagerGet(t)
	if err != nil {
		logger.Error(err.Error())
		api.Api.BuildErrorResponse(http.StatusInternalServerError, msg, w)
		return
	}

	api.Api.BuildJsonResponse(true, msg, albumList, w)
}

// GET
// Authorization: 	token
// Params: 			None
// Body: 			None

// get the list of artist in library
func LibraryArtistGet(w http.ResponseWriter, r *http.Request) {
	t := r.Header.Get(accessTokenKey)
	if !checkToken(t, w) {
		return
	}

	artistList, msg, err := managers.LibraryArtistManagerGet(t)
	if err != nil {
		logger.Error(err.Error())
		api.Api.BuildErrorResponse(http.StatusInternalServerError, msg, w)
		return
	}

	api.Api.BuildJsonResponse(true, msg, artistList, w)
}
