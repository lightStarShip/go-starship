package main

import (
	"fmt"
	"github.com/lightStarShip/go-server/service"
	"github.com/lightStarShip/go-server/utils"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

type SysParam struct {
	version bool
}

var (
	param = &SysParam{}
)

var rootCmd = &cobra.Command{
	Use: "starship",

	Short: "starship is a vpn proxy",

	Long: `usage description::TODO::`,

	Run: mainRun,
}

func init() {
	flags := rootCmd.Flags()

	flags.BoolVarP(&param.version, "version",
		"v", false, "starship -v")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

func mainRun(_ *cobra.Command, _ []string) {
	if param.version {
		fmt.Println(utils.VerStr())
		return
	}

	if err := service.Inst().StartUp(); err != nil {
		panic(err)
	}

	waitShutdownSignal()
}

func waitShutdownSignal() {

	pid := strconv.Itoa(os.Getpid())
	fmt.Printf("\n>>>>>>>>>>starship start at pid(%s)<<<<<<<<<<\n", pid)
	if err := ioutil.WriteFile(utils.PidFilePath(), []byte(pid), 0644); err != nil {
		fmt.Print("failed to write running pid", err)
	}
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		syscall.SIGUSR1,
		syscall.SIGUSR2)

	sig := <-sigCh
	service.Inst().ShutDown()
	fmt.Printf("\n>>>>>>>>>>process finished(%s)<<<<<<<<<<\n", sig)
}
