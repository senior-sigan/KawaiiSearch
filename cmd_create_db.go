package main

import (
	"kawaii_search/embedder"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var createDBCmd = &cobra.Command{
	Use:   "create_db [images dir path]",
	Short: "Create search index for the images in the dir",
	Long:  "For each image in the dir an Embedding is got by passing image throw the neural network. These emebeddings are vectors and stored in the binary file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		run(args[0])
	},
}

func init() {
	rootCmd.AddCommand(createDBCmd)
	flags := createDBCmd.PersistentFlags()
	flags.StringP("--model", "m", "model.pt", "path to the .pt model")
	flags.StringP("--out", "o", "db.bin", "path to the new-created database")
	flags.StringP("--config", "c", "config.json", "path to the .json model config")
	err := viper.BindPFlags(flags)
	if err != nil {
		logrus.WithError(err).Fatal("Invalid command line params")
	}
}

func run(imagesDir string) {
	modelPath := viper.GetString("model")
	if !FileExists(modelPath) {
		logrus.WithField("model", modelPath).Fatal("Model file not found")
		return
	}
	logrus.WithField("path", modelPath).Info("Loading model")
	emb := embedder.NewEmbedder(modelPath)
	logrus.WithField("model", emb.String()).WithField("path", modelPath).Info("Model is loaded. Ready to process images")

	// TODO: iterate over images and run embedder
	image := make([]float32, 224*224*3)
	out := emb.Transform(image)
	logrus.WithField("output", out).Info("Image processed")
	emb.DestroyEmbedder()

}
