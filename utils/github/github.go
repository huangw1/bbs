/**
 * @Author: huangw1
 * @Date: 2019/7/25 15:31
 */

package github

import (
	"context"
	"encoding/json"
	"github.com/huangw1/bbs/utils/config"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	"gopkg.in/resty.v1"
)

var OauthConfig *oauth2.Config

func InitGithub() {
	OauthConfig = &oauth2.Config{
		ClientID:     config.Conf.Github.ClientID,
		ClientSecret: config.Conf.Github.ClientSecret,
		RedirectURL:  config.Conf.BaseUrl + "/user/github/callback",
		Scopes:       []string{"public_repo", "user"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://github.com/login/oauth/authorize",
			TokenURL: "https://github.com/login/oauth/access_token",
		},
	}
}

type UserInfo struct {
	Id        int64  `json:"id"`
	Login     string `json:"login"`
	NodeId    string `json:"node_id"`
	AvatarUrl string `json:"avatar_url"`
	Url       string `json:"url"`
	HtmlUrl   string `json:"html_url"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	Bio       string `json:"bio"`
	Company   string `json:"company"`
	Blog      string `json:"blog"`
	Location  string `json:"location"`
}

const APIBaseURL = "https://api.github.com/user"

func GetToken(code string) (string, error) {
	token, err := OauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return "", err
	}
	return token.AccessToken, nil
}

func GetUserInfo(token string) (*UserInfo, error) {
	res, err := resty.R().SetQueryParam("access_token", token).Get(APIBaseURL)
	if err != nil {
		logrus.Errorf("github access_token failed: %s", err.Error())
		return nil, err
	}
	userInfo := &UserInfo{}
	if err := json.Unmarshal(res.Body(), userInfo); err != nil {
		logrus.Errorf("github access_token unmarshal failed: %s", err.Error())
		return nil, err
	}
	return userInfo, nil
}
