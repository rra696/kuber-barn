/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
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

		_, err = pod.Run()
		if err != nil {
			return err
		}

		fmt.Printf("pod started: %s\n", pod.ID)

		return nil
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List existing pods",
	RunE: func(cmd *cobra.Command, args []string) error {
		runningPods, err := pod.ListRunningPods()
		if err != nil {
			return err
		}

		for _, pod := range runningPods {
			fmt.Println(pod)
		}

		return nil
	},
}

var killCmd = &cobra.Command{
	Use: "kill",
	Short: "kill existing pod",
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := pod.KillPod(name)
		if err != nil {
			return err
		}

		fmt.Println(id)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(podCmd)

	// Create
	podCmd.AddCommand(createCmd)
	createCmd.Flags().StringVar(&imageRegistry, "registry", "", "image registry to pull (required)")
	createCmd.MarkFlagRequired("registry")
	createCmd.Flags().StringVar(&name, "name", "nameless", "the pod name")

	// List
	podCmd.AddCommand(listCmd)

	// Kill
	podCmd.AddCommand(killCmd)
	killCmd.Flags().StringVar(&name, "id", "", "the pod id (required)")
	killCmd.MarkFlagRequired("id")
}
