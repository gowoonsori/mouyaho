package badges

import (
	"github.com/gorilla/schema"
	"html/template"
	"mouyaho/domain"
	"net/http"
)

var tpl = template.Must(template.ParseGlob("static/badge/heart.gohtml"))

func HeartsBadgeHandler(w http.ResponseWriter, r *http.Request) {
	d := domain.BadgeDto{}
	decoder := schema.NewDecoder()
	_ = decoder.Decode(&d, r.URL.Query())
	b := domain.CreateBadgeFromDto(d)

	err := tpl.ExecuteTemplate(w, "heart.gohtml", b)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
	}
}

func ReactHandler(w http.ResponseWriter, r *http.Request) {
	c, _ := r.Cookie("myh_session")
	if c == nil {
		http.Error(w, "UnAuthorization", http.StatusUnauthorized)
	}

	//token := auth.DecryptCookie(*c)

}
