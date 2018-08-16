package handler

import (
	"net/http"

	"github.com/convee/blog/model"
)

func GetArticleList(w http.ResponseWriter) {
	articles := model.GetArticlesList(1, 20)
	SendResponse(w, 0, "ok", articles)
}
