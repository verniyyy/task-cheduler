/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"runtime"

	"github.com/spf13/cobra"
	"github.com/verniyyy/task-cheduler/src"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "task-cheduler",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: rootCmdRun,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.task-cheduler.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// init log
	logfile, err := os.OpenFile("/var/log/task-scheduler.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic("cannnot open test.log:" + err.Error())
	}

	log.SetOutput(io.MultiWriter(logfile, os.Stdout))
	log.SetFlags(log.Ldate | log.Ltime)
}

func rootCmdRun(cmd *cobra.Command, args []string) {
	file, err := os.Open("./settings.json")
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	var settings Settings
	if err := json.NewDecoder(file).Decode(&settings); err != nil {
		log.Panic(err)
	}
	log.Printf("setting:%+v\n", settings)

	cron, err := src.CronUsecaseFactory()
	if err != nil {
		log.Panic(err)
	}

	for _, job := range settings.SendGoogleChatJobs {
		input := src.SendGoogleChatInput{
			Webhook:    job.Webhook,
			Message:    job.Message,
			EveryDayAt: job.EveryDayAt,
		}
		cron.AddSendGoocleChatJob(input)
	}

	cron.Run()
	runtime.Goexit()
}

type Settings struct {
	SendGoogleChatJobs []SendGoogleChatJob `json:"send_google_chat_jobs"`
}

type SendGoogleChatJob struct {
	Webhook    string `json:"webhook_url"`
	Message    string `json:"message"`
	EveryDayAt string `json:"every_day_at"`
}
