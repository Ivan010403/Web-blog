package add

import (
	"fmt"
	"net/http"
	resp "webblog/internal/lib/api/response"

	"github.com/go-chi/render"
)

type Request struct {
	Author  string `json:"author"`
	Topic   string `json:"topic"`
	Content string `json:"content"`
}

type Response struct {
	resp.Response
	Date string `json:"date"`
}

type SaverArticle interface {
	SaveArticle(author string, topic string, content string) error
}

func New(articleSaver SaverArticle) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "http-server.handlers.url.Add.New"

		var req *Request

		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			render.JSON(w, r, resp.Error("failed to decode request"))

			return
		}

		fmt.Println("body of request is decoded")

		err = articleSaver.SaveArticle(req.Author, req.Topic, req.Content)
	}
}
