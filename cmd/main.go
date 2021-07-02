package main

import (
	"fmt"
	"log"

	"github.com/Enan01/rcollect"
	"github.com/gocolly/colly/proxy"
)

func main() {
	rp, err := proxy.RoundRobinProxySwitcher("socks5://127.0.0.1:7890")
	if err != nil {
		log.Fatal(err)
	}
	collector := rcollect.NewRCollector(rcollect.WithProxy(rp), rcollect.WithAsync(true))
	githubAccount := "Enan01"

	repos, err := rcollect.CollectGithubStarRepo(collector, githubAccount)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("star count:", len(repos))
	// for _, v := range repos {
	// 	fmt.Printf("repository=%+v\n", v)
	// }

	err = rcollect.OutputToCSV(repos, githubAccount)
	if err != nil {
		log.Fatal(err)
	}
}
