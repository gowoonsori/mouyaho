package interfaces

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"likeIt/badge/infrastructure/badge"
	"likeIt/domain"
	"likeIt/domain/mocks"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
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

func Test_GetLikeBadge_Handler_Basic_Success(t *testing.T) {
	//given
	sid := "af0fds0daf1lfddfad1"
	baseUrl := "https://gowoonsori.com"
	apiUrl := "/api/like-badge"
	reqUrl := baseUrl + apiUrl
	u, _ := url.Parse(reqUrl)

	bs := initMockService()
	expectResponse := getBadge(false, "", 0, "", "", "", false)
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

func initMockService() (rr *mocks.BadgeService) {
	rr = new(mocks.BadgeService)
	return
}
