package net

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/go-ping/ping"
	"github.com/spf13/cobra"
)

var (
	urlPath string
)

func pingSite(domain string) {
	pinger, err := ping.NewPinger(domain)
	if err != nil {
		panic(err)
	}

	// listen for ctrl-C signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			pinger.Stop()
		}
	}()

	pinger.OnRecv = func(pkt *ping.Packet) {
		fmt.Printf("%d bytes from %s: icmp_seq=%d time=%v\n",
			pkt.Nbytes, pkt.IPAddr, pkt.Seq, pkt.Rtt)
	}
	pinger.OnFinish = func(stats *ping.Statistics) {
		fmt.Printf("\n--- %s ping statistics ---\n", stats.Addr)
		fmt.Printf("%d packets transmitted, %d packets received, %v%% packet loss\n",
			stats.PacketsSent, stats.PacketsRecv, stats.PacketLoss)
		fmt.Printf("round-trip min/avg/max/stddev = %v/%v/%v/%v\n",
			stats.MinRtt, stats.AvgRtt, stats.MaxRtt, stats.StdDevRtt)
	}

	fmt.Printf("PING %s (%s):\n", pinger.Addr(), pinger.IPAddr())
	pinger.Run()
}

// pingCmd represents the ping command
var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "This pings a remote URL and returns the response time.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		// Logic
		pingSite(urlPath)
	},
}

func init() {
	pingCmd.Flags().StringVarP(&urlPath, "url", "u", "", "URL to ping")

	if err := pingCmd.MarkFlagRequired("url"); err != nil {
		fmt.Println("Error: ", err)
	}

	NetCmd.AddCommand(pingCmd)
}
