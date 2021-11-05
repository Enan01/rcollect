package rcollect

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

const (
	GithubUrl         = "https://github.com"
	GithubStarPageUrl = "https://github.com/%s?tab=stars"
)

type RepoInfo struct {
	Link       string `json:"link"`
	Desc       string `json:"desc"`
	Star       int    `json:"star"`
	UpdateTime string `json:"updateTime"`
}

func CollectGithubStarRepo(c *RCollector, accountName string) ([]RepoInfo, error) {
	var (
		err   error
		repos []RepoInfo
	)

	target := fmt.Sprintf(GithubStarPageUrl, accountName)
	c.c.OnHTML(".col-12.d-block.width-full.py-4.border-bottom.color-border-muted", func(e *colly.HTMLElement) {
		linktag := e.DOM.Find(".d-inline-block.mb-1").Find("div>h3>a")
		link, _ := linktag.Attr("href")

		desctag := e.DOM.Find("div[class=py-1]").Find("div>p")
		desc := strings.TrimSpace(desctag.Text())

		infotag := e.DOM.Find(".f6.color-fg-muted.mt-2")
		startag := infotag.Find("div>a").First()
		star, _ := strconv.Atoi(strings.ReplaceAll(strings.TrimSpace(startag.Text()), ",", ""))

		lastupdatetag := infotag.Find("div>relative-time")
		// TODO: 时间格式化
		lastUpdate, _ := lastupdatetag.Attr("datetime")

		// fmt.Printf("repository=%s, desc=%s, star=%d, lastUpdate=%s\n", link, desc, star, lastUpdate)

		repos = append(repos, RepoInfo{
			Link:       link,
			Desc:       desc,
			Star:       star,
			UpdateTime: lastUpdate,
		})
	})

	c.c.OnHTML(".paginate-container", func(e *colly.HTMLElement) {
		nexttag := e.DOM.Find("a[href]").Last()

		nextlink, _ := nexttag.Attr("href")
		if !strings.Contains(nextlink, "after") {
			return
		}
		err := c.c.Visit(nextlink)
		if err != nil {
			fmt.Println(err)
			return
		}
	})
	c.c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})
	c.c.OnError(func(r *colly.Response, e error) {
		fmt.Println("err=", e)
		os.Exit(1)
	})

	err = c.c.Visit(target)
	if err != nil {
		fmt.Println("error=", err)
		return nil, err
	}

	if c.opt.Async {
		c.c.Wait()
	}
	return repos, nil
}

func OutputToCSV(repos []RepoInfo, account string, filepath string) error {
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	bufwriter := bufio.NewWriterSize(file, 1024*2)
	defer func() {
		bufwriter.Flush()
	}()

	bufwriter.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM，避免使用Microsoft Excel打开乱码

	for _, v := range repos {
		_, err := bufwriter.WriteString(fmt.Sprintf("%s,\"%s\",\"%d\",%s\n", GithubUrl+v.Link, v.Desc, v.Star, v.UpdateTime))
		if err != nil {
			return err
		}
	}

	return nil
}
