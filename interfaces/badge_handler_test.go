package interfaces

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"likeIt/domain"
	"likeIt/domain/mocks"
	"likeIt/infrastructure/badge"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"
)

func Test_SetCookie(t *testing.T) {
	//given
	recorder := httptest.NewRecorder()
	http.SetCookie(recorder, &http.Cookie{Name: "test", Value: "expected"})
	request := &http.Request{Header: http.Header{"Cookie": recorder.Header()["Set-Cookie"]}}

	//when
	cookie, err := request.Cookie("test")

	//then
	require.NoError(t, err, "Failed to read 'test' Cookie: %v", err)
	require.Equal(t, cookie.Value, "expected")
}

func Test_GetLikeBadge_Handler_BasicBadge_ExistCookie_Success(t *testing.T) {
	//given
	sid := "af0fds0daf1lfddfad1"
	baseUrl := "https://gowoonsori.com"
	apiUrl := "/api/like-badge"
	reqUrl := baseUrl + apiUrl
	u, _ := url.Parse(reqUrl)

	bs := initMockService()
	expectResponse := getBadge(0, "", "", "", "", false, false)
	bs.On("GetBadgeFile", domain.UserId(sid), reqUrl).Return(expectResponse)
	lb := &LikeBadge{badgeService: bs}

	//when
	recorder := httptest.NewRecorder()
	http.SetCookie(recorder, &http.Cookie{Name: cookieId, Value: sid})
	handler := http.HandlerFunc(lb.GetLikeBadgeHandler)
	req := &http.Request{
		Header:     http.Header{"Cookie": recorder.Header()["Set-Cookie"]},
		Method:     "GET",
		URL:        u,
		RequestURI: reqUrl,
	}
	handler.ServeHTTP(recorder, req)

	//then
	status := recorder.Code
	assert.Equal(t, http.StatusOK, status, fmt.Sprintf("handler returned wrong status code: got %v want %v",
		status, http.StatusOK))

	result := recorder.Body.Bytes()
	assert.Equal(t, expectResponse, result, fmt.Sprintf("handler returned unexpected body: got %v want %v",
		recorder.Body.String(), expectResponse))
}

func Test_GetLikeBadge_Handler_BasicBadge_NotExistCookie_Success(t *testing.T) {
	//given
	baseUrl := "https://gowoonsori.com"
	apiUrl := "/api/like-badge"
	reqUrl := baseUrl + apiUrl
	u, _ := url.Parse(reqUrl)

	bs := initMockService()
	expectResponse := getBadge(0, "", "", "", "", false, false)
	bs.On("GetBadgeFile", mock.Anything, reqUrl).Return(expectResponse)
	lb := &LikeBadge{badgeService: bs}

	//when
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(lb.GetLikeBadgeHandler)
	req := &http.Request{
		Method:     "GET",
		URL:        u,
		RequestURI: reqUrl,
	}
	handler.ServeHTTP(recorder, req)

	//then
	cookies := recorder.Result().Cookies()
	var existCookie bool
	for _, c := range cookies {
		if c.Name == cookieId {
			existCookie = true
		}
	}
	assert.True(t, existCookie, "cookie가 존재하지 않습니다.")

	status := recorder.Result().StatusCode
	assert.Equal(t, http.StatusOK, status, fmt.Sprintf("handler returned wrong status code: got %v want %v",
		status, http.StatusOK))

	result := recorder.Body.Bytes()
	assert.Equal(t, expectResponse, result, fmt.Sprintf("handler returned unexpected body: got %v want %v",
		recorder.Body.String(), expectResponse))
}

func Test_GetLikeBadge_Handler_BasicBadge_ExpireCookie_Success(t *testing.T) {
	//given
	sid := "af0fds0daf1lfddfad1"
	baseUrl := "https://gowoonsori.com"
	apiUrl := "/api/like-badge"
	reqUrl := baseUrl + apiUrl
	u, _ := url.Parse(reqUrl)

	bs := initMockService()
	expectResponse := getBadge(0, "", "", "", "", false, false)
	bs.On("GetBadgeFile", domain.UserId(sid), reqUrl).Return(expectResponse)
	lb := &LikeBadge{badgeService: bs}

	//when
	recorder := httptest.NewRecorder()
	http.SetCookie(recorder, &http.Cookie{
		Name:    cookieId,
		Value:   sid,
		Expires: time.Now().Add(time.Second * -60),
	})
	handler := http.HandlerFunc(lb.GetLikeBadgeHandler)
	req := &http.Request{
		Header:     http.Header{"Cookie": recorder.Header()["Set-Cookie"]},
		Method:     "GET",
		URL:        u,
		RequestURI: reqUrl,
	}
	handler.ServeHTTP(recorder, req)

	//then
	cookies := recorder.Result().Cookies()
	var existCookie bool
	for _, c := range cookies {
		if c.Name == cookieId {
			existCookie = true
			assert.Equal(t, sid, c.Value)
		}
	}
	assert.True(t, existCookie, "cookie가 존재하지 않습니다.")

	status := recorder.Result().StatusCode
	assert.Equal(t, http.StatusOK, status, fmt.Sprintf("handler returned wrong status code: got %v want %v",
		status, http.StatusOK))

	result := recorder.Body.Bytes()
	assert.Equal(t, expectResponse, result, fmt.Sprintf("handler returned unexpected body: got %v want %v",
		recorder.Body.String(), expectResponse))
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

func initMockService() (rr *mocks.BadgeService) {
	rr = new(mocks.BadgeService)
	return
}
