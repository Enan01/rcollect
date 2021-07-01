package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/proxy"
)

func main() {
	initCollect()
	visitStarPage("Enan01")
	// time.Sleep(20 * time.Second)
}

const (
	GithubUrl         = "https://github.com"
	GithubStarPageUrl = "https://github.com/%s?tab=stars"
)

var c *colly.Collector

func initCollect() {
	c = colly.NewCollector(colly.MaxDepth(0))

	// set proxy
	rp, err := proxy.RoundRobinProxySwitcher("socks5://127.0.0.1:7890")
	if err != nil {
		log.Fatal(err)
	}
	c.SetProxyFunc(rp)
}

func visitStarPage(name string) {
	target := fmt.Sprintf(GithubStarPageUrl, name)

	c.OnHTML(".col-12.d-block.width-full.py-4.border-bottom.color-border-secondary", func(e *colly.HTMLElement) {
		linktag := e.DOM.Find(".d-inline-block.mb-1").Children().Children()
		link, _ := linktag.Attr("href")

		desctag := e.DOM.Find(".d-inline-block.col-9.color-text-secondary.pr-4")
		desc := desctag.Text()

		startag := e.DOM.Find(".Link--muted.mr-3").First()
		star := startag.Text()

		fmt.Printf("repository=%s, desc=%s, star=%s\n", link, strings.TrimSpace(desc), strings.TrimSpace(star))
	})
	c.OnHTML(".paginate-container", func(e *colly.HTMLElement) {
		nexttag := e.DOM.Find("a[href]").Last()

		nextlink, _ := nexttag.Attr("href")
		if !strings.Contains(nextlink, "after") {
			return
		}
		fmt.Println("-----------------------------")
		fmt.Printf("nextlink=%s\n", nextlink)
		fmt.Println("-----------------------------")
		err := c.Visit(nextlink)
		if err != nil {
			fmt.Println(err)
			return
		}
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>", r.URL.String())
		fmt.Println()
	})

	err := c.Visit(target)
	if err != nil {
		fmt.Println(err)
		return
	}
}
