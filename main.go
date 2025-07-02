package main

import (
	"github.com/spf13/cobra"
)

var (
	udpListenAddrStr  string = "0.0.0.0:8081"
	httpListenAddrStr string = "0.0.0.0:8080"
)

var rootCmd = &cobra.Command{
	Use:   "mesh",
	Short: "Start the mesh network daemon",
	Run: func(cmd *cobra.Command, args []string) {
		// UDP API listener
		// HTTP API listener
		// Heartbeat loop
		udpAPIlistener(udpListenAddrStr)
	},
}

func main() {
	rootCmd.PersistentFlags().StringVar(&udpListenAddrStr, "udp-addr", "", "address to listen for udp packets on (default is 0.0.0.0:8081)")
	rootCmd.PersistentFlags().StringVar(&udpListenAddrStr, "http-addr", "", "address to listen for http packets on (default is 0.0.0.0:8080)")

	rootCmd.Execute()

}
