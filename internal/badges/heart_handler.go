package badges

import (
	"github.com/gorilla/schema"
	"html/template"
	"mouyaho/domain"
	"net/http"
)

var tpl = template.Must(template.ParseGlob("static/badge/heart.gohtml"))

func HeartBadgeHandler(w http.ResponseWriter, r *http.Request) {
	d := domain.BadgeDto{}
	decoder := schema.NewDecoder()
	_ = decoder.Decode(&d, r.URL.Query())
	b := domain.CreateBadgeFromDto(d)

	err := tpl.ExecuteTemplate(w, "heart.gohtml", b)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
	}
}
