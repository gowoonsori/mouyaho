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
	"strings"
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
	qs, err := ParsingUrl(reqUrl)
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
	qs, err := ParsingUrl(reqUrl)
	urlInfo := CreateUrlInfoFromMap(qs)
	gotSvg, err := renderLikeBadge(*urlInfo, isLike, likeCount)

	//then
	if err != nil {
		t.Error(err)
	}
	tests.AssertEqual(t, gotSvg, expectSvg)
}

func Test_Get_Badge_url만있는경우(t *testing.T) {
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

func Test_Get_Badge_다른속성있는경우(t *testing.T) {
	//given
	br, rr := initMockRepository()
	url := "https://www.likeIt.com/api/badge"
	badgeId := "https://gowoonsori.com"
	likeColor := "#2f3f3f"
	bg := "#111111"
	transparency := true
	userId := "QD12LAD12LKAsd12DA1dls321sda"
	reqUrl := url + "?url=" + badgeId + "&like_color=" + likeColor + "&bg=" + bg + "&transparency=" + strconv.FormatBool(transparency)

	rr.On("FindCountByBadgeId", domain.BadgeId(badgeId)).Return(0)
	rr.On("FindByBadgeIdAndUserId", domain.BadgeId(badgeId), domain.UserId(userId)).Return(nil)
	br.On("FindById", domain.BadgeId(badgeId)).Return(nil, nil)
	br.On("Save", mock.Anything).Return(nil, nil)

	expectBadge := domain.NewBadge(domain.BadgeId(badgeId), getBadge(false, likeColor, 0, "", "", bg, transparency))

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

func getDefaultBadge() []byte {
	return []byte(strings.TrimSpace(`<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" width="99" height="30">
      <linearGradient id="smooth" x2="0" y2="100%">
        <stop offset="0" stop-color="#bbb" stop-opacity=".1"/>
        <stop offset="1" stop-opacity=".1"/>
      </linearGradient>
      <mask id="round">
        <rect width="99" height="30" rx="15" ry="15" fill="#fff"/>
      </mask>
      <g mask="url(#round)">
        <rect width="35" height="30" fill="#eee" fill-opacity="1"/>
        <rect x="35" width="29" height="30" fill="#eee" fill-opacity="1"/>
        <rect x="64" width="35" height="30" fill="#eee" fill-opacity="1"/>
        <rect width="99" height="30" fill="url(#smooth)" fill-opacity="1"/>
      </g>
      <g fill="$fff" text-anchor="middle" font-family="Arial,Sans,Verdana,Helvetica,sans-serif" font-size="11">
        <text x="44.5" y="18" fill="black" fill-opacity=".3">0</text>
        <text x="44.5" y="18" fill="#000">0</text>
      </g>
      <g class="react_icon" transform="translate(15,7)">
        <path transform="scale(0.035,0.035)" d="M 433.601 67.001 C 408.901 42.301 376.201 28.801 341.301 28.801 C 306.401 28.801 273.601 42.401 248.901 67.101 L 236.001 80.001 L 222.901 66.901 C 198.201 42.201 165.301 28.501 130.401 28.501 C 95.601 28.501 62.801 42.101 38.201 66.701 C 13.501 91.401 -0.099 124.201 0.001 159.101 C 0.001 194.001 13.701 226.701 38.401 251.401 L 226.201 439.201 C 228.801 441.801 232.301 443.201 235.701 443.201 C 239.101 443.201 242.601 441.901 245.201 439.301 L 433.401 251.801 C 458.101 227.101 471.701 194.301 471.701 159.401 C 471.801 124.501 458.301 91.701 433.601 67.001 Z"
        class="react_off"
        onclick="react()"
        fill-opacity="0" fill="red" stroke-width="20" stroke="red"/>
      </g>
      <g class="share_icon" transform="translate(9, -11)">
        <path transform="scale(0.57,0.57)" d="M 125.01 47.955 C 123.524 47.955 122.193 48.633 121.31 49.696 L 115.565 46.434 C 115.744 45.932 115.842 45.391 115.842 44.828 C 115.842 44.265 115.744 43.724 115.565 43.221 L 121.309 39.959 C 122.192 41.022 123.523 41.701 125.01 41.701 C 127.662 41.701 129.82 39.542 129.82 36.889 C 129.82 34.237 127.663 32.079 125.01 32.079 C 122.358 32.079 120.2 34.237 120.2 36.889 C 120.2 37.452 120.298 37.993 120.477 38.496 L 114.732 41.758 C 113.849 40.696 112.518 40.018 111.032 40.018 C 108.379 40.018 106.221 42.175 106.221 44.828 C 106.221 47.48 108.379 49.638 111.032 49.638 C 112.518 49.638 113.849 48.96 114.732 47.897 L 120.477 51.159 C 120.298 51.662 120.2 52.202 120.2 52.766 C 120.2 55.418 122.358 57.576 125.01 57.576 C 127.662 57.576 129.82 55.418 129.82 52.766 C 129.82 50.113 127.663 47.955 125.01 47.955 Z M 125.01 33.762 C 126.734 33.762 128.137 35.165 128.137 36.889 C 128.137 38.614 126.734 40.018 125.01 40.018 C 123.286 40.018 121.883 38.614 121.883 36.889 C 121.883 35.165 123.286 33.762 125.01 33.762 Z M 111.032 47.955 C 109.307 47.955 107.904 46.552 107.904 44.828 C 107.904 43.104 109.307 41.701 111.032 41.701 C 112.756 41.701 114.159 43.104 114.159 44.828 C 114.159 46.552 112.756 47.955 111.032 47.955 Z M 125.01 55.893 C 123.286 55.893 121.883 54.49 121.883 52.766 C 121.883 51.041 123.286 49.638 125.01 49.638 C 126.734 49.638 128.137 51.041 128.137 52.766 C 128.137 54.49 126.734 55.893 125.01 55.893 Z"
        fill="black"/>
      </g>
      <style>
        .react_icon:hover{
          cursor:pointer;
        }
        .react_icon .react_on{
          fill-opacity:1;
        }
        .share_icon:hover{
          cursor:pointer;
        }
      </style>
    </svg>`))
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
