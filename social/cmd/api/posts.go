package main

import (
	"net/http"

	"github.com/johnwr-response/Backend-Engineering-with-Go/social/internal/store"
)

type CreatePostPayload struct {
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
}

func (app *application) createPostHandler(w http.ResponseWriter, r *http.Request) {
	var payload CreatePostPayload
	if err := readJSON(w, r, &payload); err != nil {
		err := writeJSONError(w, http.StatusBadRequest, err.Error())
		if err != nil {
			return
		}
		return
	}
	//fmt.Println(payload)

	post := &store.Post{
		Title:   payload.Title,
		Content: payload.Content,
		Tags:    payload.Tags,
		// TODO: Change after auth
		UserID: 1,
	}

	ctx := r.Context()

	if err := app.store.Posts.Create(ctx, post); err != nil {
		err := writeJSONError(w, http.StatusInternalServerError, err.Error())
		if err != nil {
			return
		}
		return
	}

	if err := writeJSON(w, http.StatusCreated, post); err != nil {
		err := writeJSONError(w, http.StatusInternalServerError, err.Error())
		if err != nil {
			return
		}
		return
	}

}
