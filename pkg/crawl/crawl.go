package crawl

import (
	"fmt"
	"github.com/gocolly/colly"
	"math/rand"
	"vtmtea.com/fiction/handler/log"
	"vtmtea.com/fiction/util"
)

func List(url string) {
	c := colly.NewCollector()
	c.UserAgent = RandomAgent()
	c.AllowURLRevisit = true

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Accept", "*/*")
		logStr := fmt.Sprintf("正在抓取列表页: %s", r.URL)
		log.Create(logStr, 1)
	})

	c.OnResponse(func(r *colly.Response) {
		logStr := fmt.Sprintf("成功抓取列表页: %s", r.Request.URL)
		log.Create(logStr, 1)
	})

	c.OnXML("//*[@id='newscontent']/div[@class='l']/ul/li/span[@class='s2']/a", func(e *colly.XMLElement) {
		logStr := fmt.Sprintf("发现小说: %s", e.Text)
		log.Create(logStr, 1)
	})

	c.OnError(func(r *colly.Response, err error) {
		logStr := fmt.Sprintf("抓取错误，错误码: %d，错误原因：%s", r.StatusCode, err.Error())
		log.Create(logStr, 2)
	})

	err := c.Visit(url)
	if err != nil {
		return
	}
}

func RandomAgent() string {
	var randomNumber = rand.Intn(len(util.AgentList))

	return util.AgentList[randomNumber]
}
