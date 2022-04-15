package badges

import (
	"github.com/gorilla/schema"
	"html/template"
	"mouyaho/domain"
	"net/http"
)

var tpl = template.Must(template.ParseGlob("static/badge/heart.gohtml"))

type HeartsHandler struct {
	bs domain.BadgeService
}

func NewHeartsHandler(service domain.BadgeService) *HeartsHandler {
	return &HeartsHandler{bs: service}
}

func (hh HeartsHandler) HeartsBadgeHandler(w http.ResponseWriter, r *http.Request) {
	d := domain.BadgeDto{}
	decoder := schema.NewDecoder()
	_ = decoder.Decode(&d, r.URL.Query())
	b := domain.CreateBadgeFromDto(d)

	err := tpl.ExecuteTemplate(w, "heart.gohtml", b)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
	}
}