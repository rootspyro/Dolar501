package github

type GHAuth struct {
	ClientID string
	ClientSecret string
	AuthURL string
	TokenURL string
}

func NewGHAuth( cID string, cSecret string ) *GHAuth{
	return &GHAuth{
		ClientID: cID,
		ClientSecret: cSecret,
		AuthURL: "https://github.com/login/oauth/authorize",
		TokenURL: "https://github.com/login/oauth/access_token",
	}
}
