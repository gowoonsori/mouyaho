package github

var (
	authorizeAPI = "https://github.com/login/oauth/authorize"
	tokenAPI     = "https://github.com/login/oauth/access_token"
)

type TokenRequest struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Code         string `json:"code"`
	RedirectUrl  string `json:"redirect_url"`
	State        string `json:"state"`
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
}
