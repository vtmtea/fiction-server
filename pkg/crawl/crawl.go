package crawl

import (
	"fmt"
	"github.com/gocolly/colly"
	"math/rand"
	"vtmtea.com/fiction/util"
)

func List(url string) {
	c := colly.NewCollector()
	c.UserAgent = RandomAgent()
	c.AllowURLRevisit = true

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Accept", "*/*")
		logStr := fmt.Sprintf("正在抓取列表页: %s", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Got content from: ", r.Request.URL)
	})

	c.OnXML("//*[@id='newscontent']/div[@class='l']/ul/li/span[@class='s2']/a", func(e *colly.XMLElement) {
		fmt.Println("Find book: ", e.Text)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Get error", err.Error(), string(r.Body))
		fmt.Println("Get error code: ", r.StatusCode)
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
