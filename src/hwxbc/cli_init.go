package main

import (
	"blog"
	"cli"
	"cli/color/colorable"
	"errors"
	"fmt"
	"os"
)

func colorError(msg string) {
	stdout := colorable.NewColorable(os.Stdout)
	fmt.Fprintf(stdout, "\x1b[101;97m%s\x1b[0m", msg)
}

func colorTip(msg string) {
	stdout := colorable.NewColorable(os.Stdout)
	fmt.Fprintf(stdout, "\x1b[42;97m%s\x1b[0m", msg)
}

// root command
type rootT struct {
	cli.Helper
	Version bool `cli:"!v,version" usage:"display version"`
}

// child command: build
type childBuildT struct {
	cli.Helper
	Input  string `cli:"*i,input" usage:"specify the post directory,eg:./_post"`
	Output string `cli:"o,output" usage:"specify the post directory, default is ../source/_data/"`
	Prefix string `cli:"p,prefix" usage:"specify the output prefix"`
	Suffix string `cli:"s,suffix" usage:"specify the output suffix"`
	Type   string `cli:"t,type" usage:"<time_top_page | time_tag_page | time_category_page>"`
	All    bool   `cli:"a,all" usage:"build all types"`
}

// child command: create
type childCreateT struct {
	cli.Helper
	Categories string `cli:"c,categories" usage:"set blog categories,eg: 技术,生活"`
	Output     string `cli:"o,output" usage:"specify the post directory, default is ../source/_posts/"`
	Tags       string `cli:"t,tags" usage:"set blog tags,eg: Git,Python,C/C++"`
	Author     string `cli:"a,author" usage:"specify the author,eg: Ethan"`
	Title      string `cli:"i,title" usage:"set title"`
}

// child command: clean
type childCleanT struct {
}

var root = &cli.Command{
	Desc: "Hexo Post Compiler(hexopc)\n\nhexopc is a hexo post compiler." +
		"\npowered by Ethan(https://github.com/yicm).",
	// Argv is a factory function of argument object
	// ctx.Argv() is if Command.Argv == nil or Command.Argv() is nil
	Argv: func() interface{} { return new(rootT) },
	Fn: func(ctx *cli.Context) error {
		argv := ctx.Argv().(*rootT)
		if argv.Version {
			ctx.String("v0.1.0\n")
		} else {
			// windows
			colorError(blog.ERR_Params)
			return errors.New(blog.ERR_Params)
		}
		return nil
	},
}

var childBuild = &cli.Command{
	Name: "build",
	Desc: "this is a build command",
	Argv: func() interface{} { return new(childBuildT) },
	Fn: func(ctx *cli.Context) error {
		argv := ctx.Argv().(*childBuildT)
		// build:input
		if len(argv.Input) <= 0 {
			colorError(blog.ERR_Params_Input)
			return errors.New(blog.ERR_Params_Input)
		}
		// build:prefix
		// build:suffix
		// build:Type
		// build:All
		if argv.All {
			buildAll(argv.Input, argv.Output, argv.Prefix, argv.Suffix, callBuildAll)
		} else if !argv.All && len(argv.Type) > 0 {
			for i := 0; i < len(blog.BUILD_TYPES); i++ {
				if argv.Type == blog.BUILD_TYPES[i] {
					buildType(argv.Type, argv.Input, argv.Output, argv.Prefix, argv.Suffix, callBuildType)
					return nil
				}
			}
			colorError(blog.ERR_Type_Not_Support + argv.Type)
			return errors.New(blog.ERR_Type_Not_Support)
		} else {
			colorError(blog.ERR_Params_Output)
			return errors.New(blog.ERR_Params_Output)
		}

		return nil
	},
}

var childCreate = &cli.Command{
	Name: "create",
	Desc: "this is a create command",
	Argv: func() interface{} { return new(childCreateT) },
	Fn: func(ctx *cli.Context) error {
		argv := ctx.Argv().(*childCreateT)
		createArticle(argv.Categories, argv.Output, argv.Tags, argv.Author, argv.Title, callCreateArticle)
		return nil
	},
}

var childClean = &cli.Command{
	Name: "clean",
	Desc: "this is a clean command",
	Argv: func() interface{} { return new(childCleanT) },
	Fn: func(ctx *cli.Context) error {
		cleanBuild(callCleanBuild)

		return nil
	},
}

func cliInit() {
	if err := cli.Root(root,
		cli.Tree(childBuild),
		cli.Tree(childCreate),
		cli.Tree(childClean),
	).Run(os.Args[1:]); err != nil {
		fmt.Println("")
		fmt.Fprintln(os.Stderr, err)
		fmt.Println("")
		os.Exit(1)
	}
}

// Function callback
type CallbackBuildAll func(input string, output string, prefix string, suffix string)
type CallbackBuildType func(build_type string, input string, output string, prefix string, suffix string)
type CallbackCreateArticle func(categories string, output string, tags string, author string, title string)
type CallbackClean func()

func buildAll(input string, output string, prefix string, suffix string, callback CallbackBuildAll) {
	callback(input, output, prefix, suffix)
}

func buildType(build_type string, input string, output string, prefix string, suffix string, callback CallbackBuildType) {
	callback(build_type, input, output, prefix, suffix)
}

func createArticle(categories string, output string, tags string, author string, title string, callback CallbackCreateArticle) {
	callback(categories, output, tags, author, title)
}

func cleanBuild(callback CallbackClean) {
	callback()
}
