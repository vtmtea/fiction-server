package crawl

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/sirupsen/logrus"
	"math/rand"
	"time"
	"vtmtea.com/fiction/handler/author"
	"vtmtea.com/fiction/handler/log"
	"vtmtea.com/fiction/model"
	book2 "vtmtea.com/fiction/service/book"
	"vtmtea.com/fiction/service/category"
	"vtmtea.com/fiction/service/chapter"
	"vtmtea.com/fiction/util"
)

func List(link string) {
	host, err := util.GetUrlHost(link)

	if err != nil {
		logrus.Fatal(err)
		return
	}

	var (
		logs      []model.Log
		bookNames []string
		bookUrls  []string
	)

	c := colly.NewCollector(
		colly.AllowedDomains(host),
		colly.UserAgent(RandomAgent()),
		colly.AllowURLRevisit(),
	)

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
		bookNames = append(bookNames, e.Text)
		bookUrls = append(bookUrls, e.Request.AbsoluteURL(e.Attr("href")))
	})

	c.OnScraped(func(r *colly.Response) {
		for _, name := range bookNames {
			logs = append(logs, model.Log{
				Type:    1,
				Message: fmt.Sprintf("发现小说: %s", name),
			})
		}
		log.CreateMultiple(logs)

		for _, url := range bookUrls {
			go Single(url)
		}
	})

	c.OnError(func(r *colly.Response, err error) {
		logStr := fmt.Sprintf("抓取错误，错误码: %d，错误原因：%s", r.StatusCode, err.Error())
		log.Create(logStr, 2)
	})

	err = c.Visit(link)
	if err != nil {
		return
	}
}

func Single(bookUrl string) {
	host, err := util.GetUrlHost(bookUrl)

	if err != nil {
		logrus.Fatal(err)
		return
	}

	book := model.Book{
		SourceID:  1,
		SourceURL: bookUrl,
	}
	var chapters []*model.Chapter

	c := colly.NewCollector(
		colly.AllowedDomains(host),
		colly.UserAgent(RandomAgent()),
		colly.AllowURLRevisit(),
	)

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Accept", "*/*")
		logStr := fmt.Sprintf("正在抓取小说: %s", r.URL)
		log.Create(logStr, 1)
	})

	c.OnResponse(func(r *colly.Response) {
		logStr := fmt.Sprintf("成功抓取小说: %s", r.Request.URL)
		log.Create(logStr, 1)
	})

	c.OnXML("/html/head/meta[@property='og:title']", func(e *colly.XMLElement) {
		book.Name = e.Attr("content")
	})

	c.OnXML("/html/head/meta[@property='og:description']", func(e *colly.XMLElement) {
		book.Description = e.Attr("content")
	})

	c.OnXML("/html/head/meta[@property='og:image']", func(e *colly.XMLElement) {
		book.Cover = e.Attr("content")
	})

	c.OnXML("/html/head/meta[@property='og:novel:category']", func(e *colly.XMLElement) {
		categoryEntity := category.Get(e.Attr("content"))
		book.CategoryID = categoryEntity.ID
	})

	c.OnXML("/html/head/meta[@property='og:novel:author']", func(e *colly.XMLElement) {
		authorEntity := author.Get(e.Attr("content"))
		book.AuthorID = authorEntity.ID
	})

	c.OnXML("/html/head/meta[@property='og:novel:status']", func(e *colly.XMLElement) {
		book.UpdateStatus = e.Attr("content")
	})

	c.OnXML("/html/head/meta[@property='og:novel:latest_chapter_name']", func(e *colly.XMLElement) {
		book.LastChapterName = e.Attr("content")
	})

	c.OnXML("/html/head/meta[@property='og:novel:latest_chapter_url']", func(e *colly.XMLElement) {
		book.LastChapterURL = e.Attr("content")
	})

	c.OnXML("/html/head/meta[@property='og:novel:update_time']", func(e *colly.XMLElement) {
		updateTime := e.Attr("content")
		parseTime, err := time.Parse("2006-01-02 15:04:05", updateTime)
		if err != nil {
			return
		}
		book.LastUpdateTime = parseTime
	})

	c.OnXML("//*[@id='list']/dl/dd/a", func(e *colly.XMLElement) {
		chapters = append(chapters, &model.Chapter{
			SourceID:  1,
			SourceURL: e.Attr("href"),
			Name:      e.Text,
			Order:     int32(len(chapters) + 1),
		})
	})

	c.OnScraped(func(r *colly.Response) {
		bookModel := book2.GetBySourceUrl(bookUrl)

		if bookModel.ID == 0 {
			bookModel = book2.Create(book)
		}

		// Add book chapters
		bookChapterCount := chapter.GetBookChapterCount(bookModel.ID)
		for _, m := range chapters {
			m.BookID = bookModel.ID
		}
		chapter.CreateMultiple(chapters[bookChapterCount:])
	})

	c.OnError(func(r *colly.Response, err error) {
		logStr := fmt.Sprintf("抓取%s错误，错误码: %d，错误原因：%s", r.Request.URL, r.StatusCode, err.Error())
		log.Create(logStr, 2)
	})

	err = c.Visit(bookUrl)
	if err != nil {
		return
	}
}

func RandomAgent() string {
	var randomNumber = rand.Intn(len(util.AgentList))

	return util.AgentList[randomNumber]
}
