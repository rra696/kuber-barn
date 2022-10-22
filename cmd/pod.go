/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"time"

	"github.com/rra696/kuber-barn/internal/services/pod"
	"github.com/spf13/cobra"
)

var (
	imageRegistry string
	name          string
)

// podCmd represents the pod command
var podCmd = &cobra.Command{
	Use:   "pod",
	Short: "The command line tool to run commands on pod",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("pod called")
	},
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	RunE: func(cmd *cobra.Command, args []string) error {
		pod, err := pod.NewPod(imageRegistry, name)
		if err != nil {
			return err
		}

		fmt.Printf("pod created: %s\n", pod.ID)
		fmt.Println("starting pod...")

		runningPod, err := pod.Run()
		if err != nil {
			return err
		}

		fmt.Printf("pod started: %s\n", pod.ID)

		time.Sleep(time.Second * 3)

		fmt.Println("killing pod...")

		code, err := runningPod.Kill()
		if err != nil {
			return err
		}

		fmt.Printf("pod killed: %s\n", pod.ID)

		fmt.Printf("%s exited with status %d\n", runningPod.Pod.ID, code)

		err = pod.Delete()
		if err != nil {
			return err
		}

		fmt.Printf("container deleted: %s\n", pod.ID)

		return nil
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
	},
}

func init() {
	rootCmd.AddCommand(podCmd)
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(listCmd)

	createCmd.Flags().StringVar(&imageRegistry, "registry", "", "image registry to pull (required)")
	createCmd.MarkFlagRequired("registry")
	createCmd.Flags().StringVar(&name, "name", "nameless", "the pod name")
}
