package cmd

import (
	"github.com/bykof/go-plantuml/astParser"
	"github.com/bykof/go-plantuml/domain"
	"github.com/bykof/go-plantuml/formatter"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
)

var (
	outPath     string
	directories []string
	files       []string
	recursive   bool
	generateCmd = &cobra.Command{
		Use:   "generate [OPTIONS]",
		Short: "Generate a plantuml diagram from given paths",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			var classes domain.Classes
			for _, file := range files {
				classes = append(classes, astParser.ParseFile(file)...)
			}

			for _, directory := range directories {
				classes = append(classes, astParser.ParseDirectory(directory, recursive)...)
			}

			formattedClasses := formatter.FormatClasses(classes)
			formattedRelations := formatter.FormatRelations(classes)

			formattedPlantUML := formatter.FormatPlantUMLWrapper(formattedClasses, formattedRelations)
			err := ioutil.WriteFile(outPath, []byte(formattedPlantUML), 0644)
			if err != nil {
				log.Fatal(err)
			}
		},
	}
)

func init() {
	generateCmd.Flags().StringSliceVarP(
		&directories,
		"directories",
		"d",
		[]string{"."},
		"the go source directories",
	)
	generateCmd.Flags().StringSliceVarP(
		&files,
		"files",
		"f",
		[]string{},
		"the go source files",
	)
	generateCmd.Flags().StringVarP(
		&outPath,
		"out",
		"o",
		"graph.puml",
		"the graphfile",
	)
	generateCmd.Flags().BoolVarP(
		&recursive,
		"recursive",
		"r",
		false,
		"traverse the given directories recursively",
	)
	rootCmd.AddCommand(generateCmd)
}
