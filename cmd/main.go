package main

import (
	"fmt"
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
		_proxy, _ := cmd.Flags().GetString("proxy")
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

		options := make([]rcollect.SetOption, 0)
		options = append(options, rcollect.WithAsync(true))
		if len(_proxy) > 0 {
			fmt.Printf("proxy: %s\n", _proxy)
			rp, err := proxy.RoundRobinProxySwitcher(_proxy)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			options = append(options, rcollect.WithProxy(rp))
		}

		collector := rcollect.NewRCollector(options...)

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
	rootCmd.Flags().StringP("proxy", "p", "", "set proxy, e.g. socks5://127.0.0.1:8888")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
