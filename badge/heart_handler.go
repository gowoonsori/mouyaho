package badge

import (
	"github.com/gorilla/schema"
	"html/template"
	"mouyaho/domain"
	"net/http"
)

var tpl = template.Must(template.ParseGlob("badge/heart.gohtml"))

func GetBadge(w http.ResponseWriter, r *http.Request) {
	d := domain.BadgeDto{}
	decoder := schema.NewDecoder()
	_ = decoder.Decode(&d, r.URL.Query())
	b := domain.CreateBadgeFromDto(d)

	err := tpl.ExecuteTemplate(w, "heart.gohtml", b)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
	}
}
