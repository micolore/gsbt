package main

import (
	"bytes"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"strings"
)

func main() {
	cmd := newCommand()
	cmd.AddCommand(newNestedCommand())
	rootCmd := &cobra.Command{}
	rootCmd.AddCommand(cmd)

	if err := rootCmd.Execute(); err != nil {
		println(err.Error())
	}
}
func newCommand() *cobra.Command {
	cmd := &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			originCmd := exec.Command("/bin/bash", "-c", "docker ps -a")
			var out bytes.Buffer
			originCmd.Stdout = &out
			err := originCmd.Run()
			if err != nil {
				panic(err)
			}
			getDockerPsList(out.String())
		},
		Use:   `pd`,
		Short: "ps docker list",
		Long:  "This is  ps docker command",
	}
	return cmd
}
func newNestedCommand() *cobra.Command {
	cmd := &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			println(`Bar`)
		},
		Use:   `bar`,
		Short: "Command bar",
		Long:  "This is a nested command",
	}
	return cmd
}

var CompletionCmd = &cobra.Command{
	Use:                   "completion [bash|zsh|fish|powershell]",
	Short:                 "Generate completion script",
	Long:                  "To load completions",
	DisableFlagsInUseLine: true,
	ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
	Args:                  cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "bash":
			cmd.Root().GenBashCompletion(os.Stdout)
		case "zsh":
			cmd.Root().GenZshCompletion(os.Stdout)
		case "fish":
			cmd.Root().GenFishCompletion(os.Stdout, true)
		case "powershell":
			cmd.Root().GenPowerShellCompletionWithDesc(os.Stdout)
		}
	},
}

// 获取docker ps -a的结果
func getDockerPsList(resultString string) {
	fmt.Printf("getDockerPsList:\n%s\n", resultString)
	//strDestFileName := "./output.txt"
	//fWrite, err := os.Create(strDestFileName)
	//if err != nil {
	//	fmt.Println(err.Error())
	//	panic(err)
	//}
	//fWrite.WriteString(resultString)
	//file ,err :=ioutil.ReadFile("./output.txt")
	//if err != nil{
	//	panic(err)
	//}
	lines := strings.Split(resultString, "\n")
	for i := 0; i < len(lines); i++ {
		fmt.Println("l:" + string(i) + "  " + lines[i])
		ls := lines[i]
		lis := strings.Split(ls, "        ")
		for i := 0; i < len(lis); i++ {
			ss := lis[i]
			fmt.Println(ss)
		}
	}
}

type PsDockerListResult struct {
	ContainerId string
	Image       string
	Command     string
	Created     string
	Status      string
	Ports       string
	Names       string
}
