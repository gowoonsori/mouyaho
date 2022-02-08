package application

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm/utils/tests"
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
	qs, err := parsingUrl(reqUrl)
	urlInfo := CreateUrlInfoFromMap(qs)
	gotSvg, err := renderLikeBadge(*urlInfo, isLike, likeCount)

	//then
	if err != nil {
		t.Error(err)
	}
	tests.AssertEqual(t, gotSvg, expectSvg)
}

func Test_Render_Like_Badge_Another_Query(t *testing.T) {
	//given
	u := "https://www.naver.com"
	bg := "#eee"
	likeColor := "red"

	reqUrl := fmt.Sprintf("https://gowoon.com/api/likeIt?url=%s&bg=%s&like_color=%s&abc=123&cd=123",
		u, url.QueryEscape(bg), url.QueryEscape(likeColor))
	isLike := true
	likeCount := 12345

	expectBadge := &badge.LikeBadge{
		IsReact:         isLike,
		LikeIconColor:   likeColor,
		CountText:       strconv.Itoa(likeCount),
		BackgroundColor: bg,
	}
	wr, _ := badge.NewLikeBadgeWriter()
	expectSvg, _ := wr.RenderBadge(*expectBadge)

	//when
	qs, err := parsingUrl(reqUrl)
	urlInfo := CreateUrlInfoFromMap(qs)
	gotSvg, err := renderLikeBadge(*urlInfo, isLike, likeCount)

	//then
	if err != nil {
		t.Error(err)
	}
	tests.AssertEqual(t, gotSvg, expectSvg)
}

func Test_GetBadge_url만있는경우_캐싱X(t *testing.T) {
	//given
	br, rr := initMockRepository()
	url := "https://www.likeIt.com/api/badge"
	badgeId := "https://gowoonsori.com"
	userId := "QD12LAD12LKAsd12DA1dls321sda"
	reqUrl := url + "?url=" + badgeId

	rr.On("FindCountByBadgeId", domain.BadgeId(badgeId)).Return(0)
	rr.On("FindByBadgeIdAndUserId", domain.BadgeId(badgeId), domain.UserId(userId)).Return(nil)
	br.On("FindById", domain.BadgeId(badgeId)).Return(nil, nil)
	br.On("Save", mock.Anything).Return(nil, nil)

	expectBadge := domain.NewBadge(domain.BadgeId(badgeId), getBadge(false, "", 0, "", "", "", false))

	//when
	ls := LikeBadgeService{
		br: br,
		rr: rr,
	}
	b := ls.GetBadge(userId, reqUrl)

	//then
	assert.Equal(t, expectBadge, b)
}

func Test_GetBadge_다른속성있는경우_캐싱X(t *testing.T) {
	//given
	br, rr := initMockRepository()
	apiUrl := "https://www.likeIt.com/api/badge"
	badgeId := "https://gowoonsori.com"
	likeColor := "#2f3f3f"
	bg := "#111111"
	transparency := true
	count := 2389
	userId := "QD12LAD12LKAsd12DA1dls321sda"
	reqQs := "url=" + url.QueryEscape(badgeId) + "&like_color=" + url.QueryEscape(likeColor) + "&bg=" + url.QueryEscape(bg) + "&transparency=" + strconv.FormatBool(transparency)
	reqUrl := apiUrl + "?" + reqQs

	rr.On("FindCountByBadgeId", domain.BadgeId(badgeId)).Return(count)
	rr.On("FindByBadgeIdAndUserId", domain.BadgeId(badgeId), domain.UserId(userId)).Return(nil)
	br.On("FindById", domain.BadgeId(badgeId)).Return(nil, nil)
	br.On("Save", mock.Anything).Return(nil, nil)

	expectBadge := domain.NewBadge(domain.BadgeId(badgeId), getBadge(false, likeColor, count, "", "", bg, transparency))

	//when
	ls := LikeBadgeService{
		br: br,
		rr: rr,
	}
	b := ls.GetBadge(userId, reqUrl)

	//then
	assert.Equal(t, expectBadge, b)
}

func Test_GetBadge_다른속성있고_Encoding안된경우_캐싱X(t *testing.T) {
	//given
	br, rr := initMockRepository()
	apiUrl := "https://www.likeIt.com/api/badge"
	badgeId := "https://gowoonsori.com"
	likeColor := "#2f3f3f"
	bg := "#111111"
	transparency := true
	userId := "QD12LAD12LKAsd12DA1dls321sda"
	reqQs := "url=" + badgeId + "&like_color=" + likeColor + "&bg=" + bg + "&transparency=" + strconv.FormatBool(transparency)
	reqUrl := apiUrl + "?" + reqQs

	rr.On("FindCountByBadgeId", domain.BadgeId(badgeId)).Return(0)
	rr.On("FindByBadgeIdAndUserId", domain.BadgeId(badgeId), domain.UserId(userId)).Return(nil)
	br.On("FindById", domain.BadgeId(badgeId)).Return(nil, nil)
	br.On("Save", mock.Anything).Return(nil, nil)

	//속성코드의 #이 앞에서 짤려 transparency의 값을 읽지 못한다.
	expectBadge := domain.NewBadge(domain.BadgeId(badgeId), getBadge(false, "", 0, "", "", "", false))

	//when
	ls := LikeBadgeService{
		br: br,
		rr: rr,
	}
	b := ls.GetBadge(userId, reqUrl)

	//then
	assert.Equal(t, expectBadge, b)
}

func Test_GetBadge_좋아요상태인경우_캐싱X(t *testing.T) {
	//given
	br, rr := initMockRepository()
	apiUrl := "https://www.likeIt.com/api/badge"
	badgeId := "https://gowoonsori.com"
	likeColor := "#2f3f3f"
	bg := "#111111"
	transparency := true
	count := 2389
	userId := "QD12LAD12LKAsd12DA1dls321sda"
	reqQs := "url=" + url.QueryEscape(badgeId) + "&like_color=" + url.QueryEscape(likeColor) + "&bg=" + url.QueryEscape(bg) + "&transparency=" + strconv.FormatBool(transparency)
	reqUrl := apiUrl + "?" + reqQs

	rr.On("FindCountByBadgeId", domain.BadgeId(badgeId)).Return(count)
	rr.On("FindByBadgeIdAndUserId", domain.BadgeId(badgeId), domain.UserId(userId)).Return(domain.NewReact(domain.ReactId(1), domain.BadgeId(badgeId), domain.UserId(userId)))
	br.On("FindById", domain.BadgeId(badgeId)).Return(nil, nil)
	br.On("Save", mock.Anything).Return(nil, nil)

	expectBadge := domain.NewBadge(domain.BadgeId(badgeId), getBadge(true, likeColor, count, "", "", bg, transparency))

	//when
	ls := LikeBadgeService{
		br: br,
		rr: rr,
	}
	b := ls.GetBadge(userId, reqUrl)

	//then
	assert.Equal(t, expectBadge, b)
}

func Test_GetBadge_Badge_캐싱O(t *testing.T) {
	//given
	br, rr := initMockRepository()
	apiUrl := "https://www.likeIt.com/api/badge"
	badgeId := "https://gowoonsori.com"
	likeColor := "#2f3f3f"
	bg := "#111111"
	transparency := true
	count := 2389
	userId := "QD12LAD12LKAsd12DA1dls321sda"
	reqQs := "url=" + url.QueryEscape(badgeId) + "&like_color=" + url.QueryEscape(likeColor) + "&bg=" + url.QueryEscape(bg) + "&transparency=" + strconv.FormatBool(transparency)
	reqUrl := apiUrl + "?" + reqQs

	expectBadge := domain.NewBadge(domain.BadgeId(badgeId), getBadge(true, likeColor, count, "", "", bg, transparency))

	rr.On("FindCountByBadgeId", domain.BadgeId(badgeId)).Return(count)
	rr.On("FindByBadgeIdAndUserId", domain.BadgeId(badgeId), domain.UserId(userId)).Return(domain.NewReact(domain.ReactId(1), domain.BadgeId(badgeId), domain.UserId(userId)))
	br.On("FindById", domain.BadgeId(badgeId)).Return(expectBadge, nil)

	//when
	ls := LikeBadgeService{
		br: br,
		rr: rr,
	}
	b := ls.GetBadge(userId, reqUrl)

	//then
	assert.Equal(t, expectBadge, b)
}

func initMockRepository() (br *mocks.BadgeRepository, rr *mocks.ReactRepository) {
	br = new(mocks.BadgeRepository)
	rr = new(mocks.ReactRepository)
	return
}

func getBadge(isReact bool, likeColor string, countText int, textColor string, shareColor string, bg string, transparency bool) []byte {
	b := &badge.LikeBadge{
		IsReact:         isReact,
		LikeIconColor:   likeColor,
		CountText:       strconv.Itoa(countText),
		CountTextColor:  textColor,
		ShareIconColor:  shareColor,
		BackgroundColor: bg,
		IsTransparency:  transparency,
	}
	wr, err := badge.NewLikeBadgeWriter()
	if err != nil {
		panic(err)
	}

	//when
	svg, err := wr.RenderBadge(*b)
	if err != nil {
		panic(err)
	}

	return svg
}
