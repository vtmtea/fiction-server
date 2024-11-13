package crawl

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/sirupsen/logrus"
	"math/rand"
	"strings"
	"vtmtea.com/fiction/handler/log"
	"vtmtea.com/fiction/model"
	"vtmtea.com/fiction/util"
)

func List(url string) {
	c := colly.NewCollector()
	c.UserAgent = RandomAgent()
	c.AllowURLRevisit = true
	var logs []model.Log

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Accept", "*/*")
		logStr := fmt.Sprintf("正在抓取列表页: %s", r.URL)
		log.Create(logStr, 1)
	})

	c.OnResponse(func(r *colly.Response) {
		logStr := fmt.Sprintf("成功抓取列表页: %s", r.Request.URL)
		log.Create(logStr, 1)
	})

	c.OnXML("//*[@id='newscontent']/div[@class='l']/ul/li", func(e *colly.XMLElement) {
		categoryReplaceChars := []string{"[", "", "]", ""}
		categoryReplacer := strings.NewReplacer(categoryReplaceChars...)

		bookUrl := e.ChildAttr("/span[@class='s2']/a", "href")
		bookName := e.ChildText("/span[@class='s2']/a")
		category := e.ChildText("/span[@class='s1']")
		lastChapterTitle := e.ChildText("/span[@class='s3']/a")
		lastChapterUrl := e.ChildAttr("/span[@class='s3']/a", "href")
		logrus.Infof("Get book: %s, url: %s, category: %s", bookName, bookUrl, categoryReplacer.Replace(category))
		logrus.Infof("Last chapter: %s, url: %s", lastChapterTitle, lastChapterUrl)
		//logStr := fmt.Sprintf("发现小说: %s", e.Text)
		//logs = append(logs, model.Log{
		//	Type:    1,
		//	Message: logStr,
		//})

	})

	c.OnScraped(func(r *colly.Response) {
		logrus.Printf("%v\n", logs)
		//log.CreateMultiple(logs)
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
