package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"kawaii_search/files"
	"kawaii_search/searcher"
	"kawaii_search/server"
	"kawaii_search/storage"
	"kawaii_search/worker"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Starts search server",
	Long:  "Start search server and build search-index for all images in the root dir",
	Run: func(cmd *cobra.Command, args []string) {
		runServer()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
	flags := serverCmd.Flags()
	flags.StringP("model", "m", "model.pt", "path to the neural network model in .pt format")
	flags.StringP("images", "i", "", "path to the images directory")
	flags.IntP("port", "p", 3000, "server port to listen to")
	err := viper.BindPFlags(flags)
	if err != nil {
		logrus.WithError(err).Fatal("Invalid command line params")
	}
}

func runServer() {
	hs := storage.NewStorage("/tmp/kawaii_search")
	fs := files.NewFilesStorage("/Users/ilya/Desktop/images")
	searcher_ := searcher.NewSearcher(hs)
	w := worker.NewHashWorker(hs, fs)
	s := server.NewServer(&server.Config{
		Address: "0.0.0.0:3000",
	}, hs, fs, searcher_)

	ctx := context.Background()

	go w.Start()
	go s.Start()

	<-ctx.Done()
}
