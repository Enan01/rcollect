package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Enan01/rcollect"
	"github.com/gocolly/colly/proxy"
	"github.com/spf13/cobra"
)

func main() {
	Execute()
}

var rootCmd = &cobra.Command{
	Use:   "rcollect",
	Short: "collect github star repositorys, and export to csv file.",
	Run: func(cmd *cobra.Command, args []string) {
		account, _ := cmd.Flags().GetString("account")
		output, _ := cmd.Flags().GetString("output")
		if output[len(output)-1] == '/' {
			output = output[:len(output)-1]
		}
		if len(account) == 0 {
			fmt.Println("account is nil")
			os.Exit(1)
		}
		fmt.Printf("account is %v\n", account)
		filename := fmt.Sprintf("%s-star-repos.csv", account)
		filepath := output + "/" + filename
		fmt.Printf("output file: %s\n", filepath)
		_, err := os.Stat(output)
		if err != nil && os.IsNotExist(err) {
			os.MkdirAll(output, os.ModePerm)
		}
		_, err = os.Stat(filepath)
		if err == nil {
			fmt.Printf("file %s is exists, please check.\n", filepath)
			os.Exit(1)
		}

		// TODO: proxy 调整为命令行参数
		rp, err := proxy.RoundRobinProxySwitcher("socks5://127.0.0.1:7890")
		if err != nil {
			log.Fatal(err)
		}
		collector := rcollect.NewRCollector(rcollect.WithProxy(rp), rcollect.WithAsync(true))
		githubAccount := account

		repos, err := rcollect.CollectGithubStarRepo(collector, githubAccount)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		err = rcollect.OutputToCSV(repos, githubAccount, filepath)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.Flags().StringP("account", "a", "", "github account")
	rootCmd.Flags().StringP("output", "o", "./", "file output path")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
