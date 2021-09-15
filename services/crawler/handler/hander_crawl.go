package handler

import (
	"fmt"
	"micro_backend_film/services/biz/pb/pb_crawl"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

type HandlerCrawl struct {
	Colly *colly.Collector
}

func (hc *HandlerCrawl) CrawlFilm() *pb_crawl.CrawledFilms {
	crawledFilms := &pb_crawl.CrawledFilms{}

	hc.Colly.OnHTML("tbody.lister-list", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(i int, h *colly.HTMLElement) {
			var crawledFilm pb_crawl.CrawledFilm
			crawledFilm.Name = h.ChildText("td.titleColumn > a")

			year := h.ChildText("td.titleColumn > span.secondaryInfo")
			year = strings.Replace(year, "(", "", -1)
			year = strings.Replace(year, ")", "", -1)
			crawledFilm.Year = year

			rating := h.ChildText("td.ratingColumn > strong")
			rating = strings.Replace(rating, ",", ".", -1)
			ratingInt, _ := strconv.ParseFloat(rating, 64)
			crawledFilm.Rating = ratingInt
			crawledFilms.Films = append(crawledFilms.Films, &crawledFilm)
		})
	})

	hc.Colly.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	hc.Colly.Visit("https://www.imdb.com/chart/moviemeter/?ref_=nv_mv_mpm")
	return crawledFilms
}
