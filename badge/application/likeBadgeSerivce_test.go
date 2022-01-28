package application

import (
	"fmt"
	"gorm.io/gorm/utils/tests"
	"likeIt/badge/infrastructure/badge"
	"net/url"
	"strconv"
	"testing"
)

func Test_PARSE_QUERY_URL(t *testing.T) {
	//given
	u := "https://www.naver.com"
	bg := "#eee"
	likeColor := "red"
	textColor := "#aaa"
	shareColor := "black"
	transparency := "true"
	qs := fmt.Sprintf("url=%s&bg=%s&like_color=%s&text_color=%s&share_color=%s&transparency=%s", u, bg, likeColor, textColor, shareColor, transparency)
	reqUrl := "https://gowoon.com/api/likeIt?" + url.QueryEscape(qs)

	//when
	p, _ := url.Parse(reqUrl)
	rq, _ := url.QueryUnescape(p.RawQuery)
	m, err := url.ParseQuery(rq)
	if err != nil {
		t.Error("파싱 에러 발생 : " + fmt.Sprintf("%v", err))
	}

	//then
	tests.AssertEqual(t, m["url"][0], u)
	tests.AssertEqual(t, m["bg"][0], bg)
	tests.AssertEqual(t, m["like_color"][0], likeColor)
	tests.AssertEqual(t, m["text_color"][0], textColor)
	tests.AssertEqual(t, m["share_color"][0], shareColor)
	tests.AssertEqual(t, m["transparency"][0], transparency)
}

func Test_Render_Like_Badge_Success(t *testing.T) {
	//given
	u := "https://www.naver.com"
	bg := "#eee"
	likeColor := "red"
	textColor := "#aaa"
	shareColor := "black"
	transparency := "true"

	reqUrl := fmt.Sprintf("https://gowoon.com/api/likeIt?url=%s&bg=%s&like_color=%s&text_color=%s&share_color=%s&transparency=%s",
		u, url.QueryEscape(bg), url.QueryEscape(likeColor), url.QueryEscape(textColor), url.QueryEscape(shareColor), transparency)
	isLike := true
	likeCount := 12345

	expectBadge := &badge.LikeBadge{
		IsReact:         isLike,
		LikeIconColor:   likeColor,
		CountText:       strconv.Itoa(likeCount),
		CountTextColor:  textColor,
		ShareIconColor:  shareColor,
		BackgroundColor: bg,
		IsTransparency:  true,
	}
	wr, _ := badge.NewLikeBadgeWriter()
	expectSvg, _ := wr.RenderBadge(*expectBadge)

	//when
	lb := LikeBadgeService{}
	gotSvg, err := lb.renderLikeBadge(reqUrl, isLike, likeCount)

	//then
	if err != nil {
		t.Error(err)
	}
	tests.AssertEqual(t, gotSvg, expectSvg)
}
