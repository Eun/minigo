package root

import (
	"fmt"
	"os"
	"strings"
	"unicode"

	"io"

	"github.com/Eun/minigo/minigo"
	"github.com/hjson/hjson-go"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "minigo [flags] <.go files to run>",
		Short: "A mini golang interpreter",
		Long: `A mini golang interpreter based on yaegi-template and yaegi.
https://github.com/Eun/minigo`,
		Run:  run,
		Args: cobra.MinimumNArgs(1),
	}

	templatingFlag bool
	writeToFlag    string
	contextFlag    string
)

func Version(version string) {
	rootCmd.Version = version
}

func init() {
	rootCmd.Flags().BoolVarP(&templatingFlag, "template", "t", false, "enable templating")
	rootCmd.Flags().StringVarP(&writeToFlag, "out", "o", "", "write output to file")
	rootCmd.Flags().StringVarP(&contextFlag, "context", "c", "", "set context to the specified json object")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run(cmd *cobra.Command, args []string) {
	var w io.Writer = os.Stdout
	if writeToFlag != "" {
		f, err := os.Create(writeToFlag)
		if err != nil {
			fmt.Fprintf(os.Stderr, "unable to write file `%s': %v\n", writeToFlag, err)
			os.Exit(1)
		}
		defer f.Close()
		w = f
	}

	context := makeContext()

	amountOfGoFiles := len(args)
	if amountOfGoFiles <= 0 {
		return
	}

	g, err := minigo.New(minigo.Config{
		TemplateMode: templatingFlag,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to create minigo instance: %v\n", err)
		os.Exit(1)
	}

	for _, s := range args {
		f, err := os.Open(s)
		if err != nil {
			fmt.Fprintf(os.Stderr, "unable to open file `%s': %v\n", s, err)
			os.Exit(1)
		}
		err = g.Run(f, context, w)
		_ = f.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "unable to exec file `%s': %v\n", s, err)
			os.Exit(1)
		}
	}
}

func makeContext() interface{} {
	if contextFlag == "" {
		return nil
	}

	contextFlag = strings.TrimFunc(contextFlag, func(r rune) bool {
		return unicode.IsSpace(r) || r == '\'' || r == '"'
	})
	if contextFlag == "" {
		return nil
	}
	var dat map[string]interface{}
	if err := hjson.Unmarshal([]byte(contextFlag), &dat); err != nil {
		fmt.Fprintf(os.Stderr, "unable to decode context `%s': %v\n", contextFlag, err)
		os.Exit(1)
	}
	context, err := minigo.ConvertMapToStruct(dat)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to convert context map to struct `%s': %v\n", contextFlag, err)
		os.Exit(1)
	}
	return context
}
