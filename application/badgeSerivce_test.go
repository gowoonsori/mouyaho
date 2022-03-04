package application

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"likeIt/domain"
	"likeIt/domain/mocks"
	"likeIt/infrastructure/badge"
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
	clear := true
	isLike := true
	likeCount := 12345
	reqUrl := fmt.Sprintf("https://gowoon.com/api/likeIt?url=%s&bg=%s&like_color=%s&text_color=%s&share_color=%s&clear=%t",
		u, url.QueryEscape(bg), url.QueryEscape(likeColor), url.QueryEscape(textColor), url.QueryEscape(shareColor), clear)

	expectBadge := getBadge(likeCount, likeColor, textColor, shareColor, bg, isLike, clear)

	//when
	urlInfo := badge.CreateUrlInfoFromUrl(reqUrl)
	gotSvg, err := badge.GenerateLikeBadge(*urlInfo, isLike, likeCount)

	//then
	if err != nil {
		t.Errorf("error가 발생하여 file 이 nil 입니다. err : %#v", err)
	}
	assert.Equal(t, gotSvg, expectBadge)
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

	expectBadge := getBadge(likeCount, likeColor, "", "", bg, isLike, false)

	//when
	urlInfo := badge.CreateUrlInfoFromUrl(reqUrl)
	gotSvg, err := badge.GenerateLikeBadge(*urlInfo, isLike, likeCount)

	//then
	if err != nil {
		t.Errorf("error가 발생하여 file 이 nil 입니다. err : %#v", err)
	}
	assert.Equal(t, gotSvg, expectBadge)
}

func Test_GetBadgeFile_url만있는경우_캐싱X(t *testing.T) {
	//given
	rr := initMockRepository()
	url := "https://www.likeIt.com/api/badge"
	badgeId := "https://gowoonsori.com"
	userId := "QD12LAD12LKAsd12DA1dls321sda"
	reqUrl := url + "?url=" + badgeId

	expectBadge := getBadge(0, "", "", "", "", false, false)

	rr.On("FindCountByBadgeId", domain.BadgeId(badgeId)).Return(0)
	rr.On("FindByBadgeIdAndUserId", domain.BadgeId(badgeId), domain.UserId(userId)).Return(nil)

	//when
	ls := LikeBadgeService{rr: rr}
	b := ls.GetBadgeFile(domain.UserId(userId), reqUrl)

	//then
	assert.Equal(t, expectBadge, b)
}

func Test_GetBadgeFile_다른속성있는경우_캐싱X(t *testing.T) {
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

	expectBadge := getBadge(count, likeColor, "", "", bg, false, clear)

	rr.On("FindCountByBadgeId", domain.BadgeId(badgeId)).Return(count)
	rr.On("FindByBadgeIdAndUserId", domain.BadgeId(badgeId), domain.UserId(userId)).Return(nil)

	//when
	ls := LikeBadgeService{rr: rr}
	b := ls.GetBadgeFile(domain.UserId(userId), reqUrl)

	//then
	assert.Equal(t, expectBadge, b)
}

func Test_GetBadgeFile_다른속성있고_Encoding안된경우_캐싱X(t *testing.T) {
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
	expectBadge := getBadge(0, "", "", "", "", false, false)

	rr.On("FindCountByBadgeId", domain.BadgeId(badgeId)).Return(0)
	rr.On("FindByBadgeIdAndUserId", domain.BadgeId(badgeId), domain.UserId(userId)).Return(nil)

	//when
	ls := LikeBadgeService{rr: rr}
	b := ls.GetBadgeFile(domain.UserId(userId), reqUrl)

	//then
	assert.Equal(t, expectBadge, b)
}

func Test_GetBadgeFile_좋아요상태인경우_캐싱X(t *testing.T) {
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

	expectBadge := getBadge(count, likeColor, "", "", bg, true, clear)

	rr.On("FindCountByBadgeId", domain.BadgeId(badgeId)).Return(count)
	rr.On("FindByBadgeIdAndUserId", domain.BadgeId(badgeId), domain.UserId(userId)).Return(domain.NewReact(domain.ReactId(1), domain.BadgeId(badgeId), *domain.NewReader(domain.UserId(userId)), true))

	//when
	ls := LikeBadgeService{rr: rr}
	b := ls.GetBadgeFile(domain.UserId(userId), reqUrl)

	//then
	assert.Equal(t, expectBadge, b)
}

func Test_GetBadgeFile_Badge_캐싱O(t *testing.T) {
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

	expectBadge := getBadge(count, likeColor, "", "", bg, true, clear)

	rr.On("FindCountByBadgeId", domain.BadgeId(badgeId)).Return(count)
	rr.On("FindByBadgeIdAndUserId", domain.BadgeId(badgeId), domain.UserId(userId)).Return(domain.NewReact(domain.ReactId(1), domain.BadgeId(badgeId), *domain.NewReader(domain.UserId(userId)), true))

	//when
	ls := LikeBadgeService{rr: rr}
	b := ls.GetBadgeFile(domain.UserId(userId), reqUrl)

	//then
	assert.Equal(t, expectBadge, b)
}

func Test_LikeBadge_처음좋아요(t *testing.T) {
	//given
	rr := initMockRepository()
	badgeId := "https://gowoonsori.com"
	userId := "QD12LAD12LKAsd12DA1dls321sda"
	expectedCount := 12

	rr.On("FindByBadgeIdAndUserId", domain.BadgeId(badgeId), domain.UserId(userId)).Return(nil)
	rr.On("Save", domain.ByOn(domain.UserId(userId), domain.BadgeId(badgeId))).Return(domain.ByOn(domain.UserId(userId), domain.BadgeId(badgeId)), nil)
	rr.On("FindCountByBadgeId", domain.BadgeId(badgeId)).Return(expectedCount)

	//when
	ls := LikeBadgeService{rr: rr}
	b := ls.LikeBadge(domain.UserId(userId), domain.BadgeId(badgeId))

	//then
	assert.Equal(t, &ReactDto{
		Count:  expectedCount,
		IsLike: true,
	}, b)
}

func Test_LikeBadge_좋아요상태변경(t *testing.T) {
	//given
	rr := initMockRepository()
	badgeId := "https://gowoonsori.com"
	userId := "QD12LAD12LKAsd12DA1dls321sda"
	reactId := domain.ReactId(12345)
	react := domain.NewReact(reactId, domain.BadgeId(badgeId), *domain.NewReader(domain.UserId(userId)), false)
	expectedCount := 12

	rr.On("FindByBadgeIdAndUserId", domain.BadgeId(badgeId), domain.UserId(userId)).Return(react)
	rr.On("UpdateLikeStatusById", react.Id(), true).Return(nil)
	rr.On("FindCountByBadgeId", domain.BadgeId(badgeId)).Return(expectedCount)

	//when
	ls := LikeBadgeService{rr: rr}
	b := ls.LikeBadge(domain.UserId(userId), domain.BadgeId(badgeId))

	//then
	assert.Equal(t, &ReactDto{
		Count:  expectedCount,
		IsLike: true,
	}, b)
}

func initMockRepository() (rr *mocks.ReactRepository) {
	rr = new(mocks.ReactRepository)
	return
}

func getBadge(countText int, likeColor, textColor, shareColor, bg string, isReact, isClear bool) []byte {
	u := badge.UrlInfo{
		Url:             "",
		LikeIconColor:   likeColor,
		CountTextColor:  textColor,
		ShareIconColor:  shareColor,
		BackgroundColor: bg,
		IsClear:         isClear,
	}

	//when
	f, _ := badge.GenerateLikeBadge(u, isReact, countText)
	return f
}
