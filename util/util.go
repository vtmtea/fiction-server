package util

import (
	"github.com/gin-gonic/gin"
	"github.com/teris-io/shortid"
	"net/url"
	"strings"
)

func GenShortId() (string, error) {
	return shortid.Generate()
}

func GetReqID(c *gin.Context) string {
	v, ok := c.Get("X-Request-Id")
	if !ok {
		return ""
	}
	if requestId, ok := v.(string); ok {
		return requestId
	}
	return ""
}

func GetUrlHost(urlString string) (string, error) {
	urlParse, err := url.Parse(urlString)
	if err != nil {
		return "", err
	}

	return urlParse.Host, nil
}

func FormatCategoryName(name string) string {
	categories := []string{
		"玄幻", "奇幻", "武侠", "仙侠", "都市", "现实", "军事", "历史", "游戏", "体育", "科幻", "诸天无限", "悬疑", "轻小说", "言情",
	}

	for _, category := range categories {
		if strings.Contains(name, category) {
			return category
		}
	}

	return "三体"
}
