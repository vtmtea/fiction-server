package crawl

import (
	"fmt"
	"github.com/gocolly/colly"
	log "github.com/sirupsen/logrus"
)

func List(url string) {
	log.Infof(url)
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Got content from: ", r.Request.URL)
	})

	c.OnXML("//*[@id='newscontent']/div[@class='l']/ul/li/span[@class='s2']/a", func(e *colly.XMLElement) {
		fmt.Println("Find book: ", e.Text)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Get error", err.Error())
	})

	err := c.Visit(url)
	if err != nil {
		return
	}
}
