package application

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"likeIt/badge/infrastructure/badge"
	"likeIt/domain"
	"likeIt/domain/mocks"
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
	clear := "true"
	qs := fmt.Sprintf("url=%s&bg=%s&like_color=%s&text_color=%s&share_color=%s&clear=%s", u, bg, likeColor, textColor, shareColor, clear)
	reqUrl := "https://gowoon.com/api/likeIt?" + url.QueryEscape(qs)

	//when
	p, _ := url.Parse(reqUrl)
	rq, _ := url.QueryUnescape(p.RawQuery)
	m, err := url.ParseQuery(rq)
	if err != nil {
		t.Error("파싱 에러 발생 : " + fmt.Sprintf("%v", err))
	}

	//then
	assert.Equal(t, m["url"][0], u)
	assert.Equal(t, m["bg"][0], bg)
	assert.Equal(t, m["like_color"][0], likeColor)
	assert.Equal(t, m["text_color"][0], textColor)
	assert.Equal(t, m["share_color"][0], shareColor)
	assert.Equal(t, m["clear"][0], clear)
}

func Test_Render_Like_Badge_Success(t *testing.T) {
	//given
	u := "https://www.naver.com"
	bg := "#eee"
	likeColor := "red"
	textColor := "#aaa"
	shareColor := "black"
	clear := "true"
	isLike := true
	likeCount := 12345
	reqUrl := fmt.Sprintf("https://gowoon.com/api/likeIt?url=%s&bg=%s&like_color=%s&text_color=%s&share_color=%s&clear=%s",
		u, url.QueryEscape(bg), url.QueryEscape(likeColor), url.QueryEscape(textColor), url.QueryEscape(shareColor), clear)

	expectBadge := badge.NewLikeBadge(likeColor, textColor, shareColor, bg, likeCount, isLike, true)
	wr, _ := badge.NewLikeBadgeWriter()
	expectSvg, _ := wr.RenderBadge(expectBadge)

	//when
	urlInfo := CreateUrlInfoFromUrl(reqUrl)
	bs := LikeBadgeService{rr: initMockRepository()}
	gotSvg, err := bs.renderBadge(*urlInfo, isLike, likeCount)

	//then
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, gotSvg, expectSvg)
}

func Test_Render_Like_Badge_Another_Query(t *testing.T) {
	//given
	u := "https://www.naver.com"
	bg := "#eee"
	likeColor := "red"
	isLike := true
	likeCount := 12345
	reqUrl := fmt.Sprintf("https://gowoon.com/api/likeIt?url=%s&bg=%s&like_color=%s&abc=123&cd=123",
		u, url.QueryEscape(bg), url.QueryEscape(likeColor))

	expectBadge := badge.NewLikeBadge(likeColor, "", "", bg, likeCount, isLike, false)
	wr, _ := badge.NewLikeBadgeWriter()
	expectSvg, _ := wr.RenderBadge(expectBadge)

	//when
	urlInfo := CreateUrlInfoFromUrl(reqUrl)
	bs := LikeBadgeService{rr: initMockRepository()}
	gotSvg, err := bs.renderBadge(*urlInfo, isLike, likeCount)

	//then
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, gotSvg, expectSvg)
}

func Test_GetBadge_url만있는경우_캐싱X(t *testing.T) {
	//given
	rr := initMockRepository()
	url := "https://www.likeIt.com/api/badge"
	badgeId := "https://gowoonsori.com"
	userId := "QD12LAD12LKAsd12DA1dls321sda"
	reqUrl := url + "?url=" + badgeId

	expectBadge := getBadge(false, "", 0, "", "", "", false)

	rr.On("FindCountByBadgeId", domain.BadgeId(badgeId)).Return(0)
	rr.On("FindByBadgeIdAndUserId", domain.BadgeId(badgeId), domain.UserId(userId)).Return(nil)

	//when
	ls := LikeBadgeService{rr: rr}
	b := ls.GetBadge(userId, reqUrl)

	//then
	assert.Equal(t, expectBadge, b)
}

func Test_GetBadge_다른속성있는경우_캐싱X(t *testing.T) {
	//given
	rr := initMockRepository()
	apiUrl := "https://www.likeIt.com/api/badge"
	badgeId := "https://gowoonsori.com"
	likeColor := "#2f3f3f"
	bg := "#111111"
	clear := true
	count := 2389
	userId := "QD12LAD12LKAsd12DA1dls321sda"
	reqQs := "url=" + url.QueryEscape(badgeId) + "&like_color=" + url.QueryEscape(likeColor) + "&bg=" + url.QueryEscape(bg) + "&clear=" + strconv.FormatBool(clear)
	reqUrl := apiUrl + "?" + reqQs

	expectBadge := getBadge(false, likeColor, count, "", "", bg, clear)

	rr.On("FindCountByBadgeId", domain.BadgeId(badgeId)).Return(count)
	rr.On("FindByBadgeIdAndUserId", domain.BadgeId(badgeId), domain.UserId(userId)).Return(nil)

	//when
	ls := LikeBadgeService{rr: rr}
	b := ls.GetBadge(userId, reqUrl)

	//then
	assert.Equal(t, expectBadge, b)
}

func Test_GetBadge_다른속성있고_Encoding안된경우_캐싱X(t *testing.T) {
	//given
	rr := initMockRepository()
	apiUrl := "https://www.likeIt.com/api/badge"
	badgeId := "https://gowoonsori.com"
	likeColor := "#2f3f3f"
	bg := "#111111"
	clear := true
	userId := "QD12LAD12LKAsd12DA1dls321sda"
	reqQs := "url=" + badgeId + "&like_color=" + likeColor + "&bg=" + bg + "&clear=" + strconv.FormatBool(clear)
	reqUrl := apiUrl + "?" + reqQs

	//속성코드의 #이 앞에서 짤려 clear의 값을 읽지 못한다.
	expectBadge := getBadge(false, "", 0, "", "", "", false)

	rr.On("FindCountByBadgeId", domain.BadgeId(badgeId)).Return(0)
	rr.On("FindByBadgeIdAndUserId", domain.BadgeId(badgeId), domain.UserId(userId)).Return(nil)

	//when
	ls := LikeBadgeService{rr: rr}
	b := ls.GetBadge(userId, reqUrl)

	//then
	assert.Equal(t, expectBadge, b)
}

func Test_GetBadge_좋아요상태인경우_캐싱X(t *testing.T) {
	//given
	rr := initMockRepository()
	apiUrl := "https://www.likeIt.com/api/badge"
	badgeId := "https://gowoonsori.com"
	likeColor := "#2f3f3f"
	bg := "#111111"
	clear := true
	count := 2389
	userId := "QD12LAD12LKAsd12DA1dls321sda"
	reqQs := "url=" + url.QueryEscape(badgeId) + "&like_color=" + url.QueryEscape(likeColor) + "&bg=" + url.QueryEscape(bg) + "&clear=" + strconv.FormatBool(clear)
	reqUrl := apiUrl + "?" + reqQs

	expectBadge := getBadge(true, likeColor, count, "", "", bg, clear)

	rr.On("FindCountByBadgeId", domain.BadgeId(badgeId)).Return(count)
	rr.On("FindByBadgeIdAndUserId", domain.BadgeId(badgeId), domain.UserId(userId)).Return(domain.NewReact(domain.ReactId(1), domain.BadgeId(badgeId), *domain.NewReader(domain.UserId(userId))))

	//when
	ls := LikeBadgeService{rr: rr}
	b := ls.GetBadge(userId, reqUrl)

	//then
	assert.Equal(t, expectBadge, b)
}

func Test_GetBadge_Badge_캐싱O(t *testing.T) {
	//given
	rr := initMockRepository()
	apiUrl := "https://www.likeIt.com/api/badge"
	badgeId := "https://gowoonsori.com"
	likeColor := "#2f3f3f"
	bg := "#111111"
	clear := true
	count := 2389
	userId := "QD12LAD12LKAsd12DA1dls321sda"
	reqQs := "url=" + url.QueryEscape(badgeId) + "&like_color=" + url.QueryEscape(likeColor) + "&bg=" + url.QueryEscape(bg) + "&clear=" + strconv.FormatBool(clear)
	reqUrl := apiUrl + "?" + reqQs

	expectBadge := getBadge(true, likeColor, count, "", "", bg, clear)

	rr.On("FindCountByBadgeId", domain.BadgeId(badgeId)).Return(count)
	rr.On("FindByBadgeIdAndUserId", domain.BadgeId(badgeId), domain.UserId(userId)).Return(domain.NewReact(domain.ReactId(1), domain.BadgeId(badgeId), *domain.NewReader(domain.UserId(userId))))

	//when
	ls := LikeBadgeService{rr: rr}
	b := ls.GetBadge(userId, reqUrl)

	//then
	assert.Equal(t, expectBadge, b)
}

func initMockRepository() (rr *mocks.ReactRepository) {
	rr = new(mocks.ReactRepository)
	return
}

func getBadge(isReact bool, likeColor string, countText int, textColor string, shareColor string, bg string, isClear bool) []byte {
	b := badge.NewLikeBadge(likeColor, textColor, shareColor, bg, countText, isReact, isClear)

	wr, err := badge.NewLikeBadgeWriter()
	if err != nil {
		panic(err)
	}

	//when
	svg, err := wr.RenderBadge(b)
	if err != nil {
		panic(err)
	}

	return svg
}
