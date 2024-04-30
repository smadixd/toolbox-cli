package net

import (
	"fmt"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

var (
	urlPath string
	client  = http.Client{
		Timeout: 2 * time.Second,
	}
)

func ping(domain string) (int, error) {
	url := "https://" + domain
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	resp.Body.Close()
	return resp.StatusCode, nil
}

// pingCmd represents the ping command
var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "This pings a remote URL and returns the response time.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		// Logic

		if resp, err := ping(urlPath); err != nil {
			fmt.Println("Error: ", err)
		} else {
			fmt.Println("Response: ", resp)
		}
	},
}

func init() {
	pingCmd.Flags().StringVarP(&urlPath, "url", "u", "", "URL to ping")

	if err := pingCmd.MarkFlagRequired("url"); err != nil {
		fmt.Println("Error: ", err)
	}

	NetCmd.AddCommand(pingCmd)
}
