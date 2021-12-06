package cmd

import (
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var knockCmd = &cobra.Command{
	Use:   "knock",
	Short: "knock server port",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(os.Args) != 4 {
			return errors.New("Wrong number of arguments")
		}
		host := os.Args[2]
		port := os.Args[3]
		portInt, err := strconv.Atoi(port)
		if err != nil {
			return errors.New("Invalid port")
		}
		address := fmt.Sprintf("%s:%d", host, portInt)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			log.Printf("Fail to connect %s\n", address)
			return err
		}
		defer conn.Close()
		log.Printf("Connect Success")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(knockCmd)
}
