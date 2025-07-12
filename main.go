package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	udpListenAddrStr  string = ""
	httpListenAddrStr string = ""
	isController      bool   = false
)

var rootCmd = &cobra.Command{
	Use:   "mesh",
	Short: "Start the mesh network daemon",
	Run: func(cmd *cobra.Command, args []string) {
		if isController {
			udpAPIlistener(udpListenAddrStr)
		} else {
			fmt.Println("Sending a message...")
			send()
		}
	},
}

func main() {
	rootCmd.PersistentFlags().BoolVar(&isController, "controller", false, "whether or not this node is a controller")
	rootCmd.PersistentFlags().StringVar(&udpListenAddrStr, "udp-addr", "0.0.0.0:8081", "address to listen for udp packets on (default is 0.0.0.0:8081)")
	rootCmd.PersistentFlags().StringVar(&httpListenAddrStr, "http-addr", "0.0.0.0:8080", "address to listen for http packets on (default is 0.0.0.0:8080)")

	rootCmd.Execute()

}
