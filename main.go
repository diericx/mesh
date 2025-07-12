package main

import (
	"fmt"
	"os"

	"github.com/google/uuid"
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
		}
	},
}

var joinCmd = &cobra.Command{
	Use:   "networks",
	Short: "Lists the networks on the target node",
	Run: func(cmd *cobra.Command, args []string) {
		reqId, err := uuid.NewV6()
		if err != nil {
			fmt.Println("Error creating new request id: %v", err)
			os.Exit(1)
		}
		sendUDPRequest("0.0.0.0:8081", UDPRequest{
			Endpoint: "/api/v1/networks",
			ReqId:    reqId.String(),
		})
	},
}

func main() {
	rootCmd.PersistentFlags().BoolVar(&isController, "controller", false, "whether or not this node is a controller")
	rootCmd.PersistentFlags().StringVar(&udpListenAddrStr, "udp-addr", "0.0.0.0:8081", "address to listen for udp packets on (default is 0.0.0.0:8081)")
	rootCmd.PersistentFlags().StringVar(&httpListenAddrStr, "http-addr", "0.0.0.0:8080", "address to listen for http packets on (default is 0.0.0.0:8080)")

	rootCmd.AddCommand(joinCmd)

	rootCmd.Execute()

}
