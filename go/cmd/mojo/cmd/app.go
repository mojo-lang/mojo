/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"log"
	"os"
	"sort"

	"github.com/urfave/cli/v2"
)

type App struct {
	cli.App

	Verbose bool // = flag.BoolP("verbose", "v", false, "Verbose Output")
	Help    bool // = flag.BoolP("help", "h", false, "Print usage")

	BinName   string // = filepath.Base(os.Args[0])
	Workplace string // = ""

	Version     string
	VersionDate string
}

func newApp() *App {
	app := &App{
		App: cli.App{
			UseShortOptionHandling: true,
		},
	}

	sort.Sort(Cmds(commands))
	for _, cmd := range commands {
		app.App.Commands = append(app.App.Commands, cmd.GetCmd().Command)
	}

	return app
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func (a *App) Execute() {
	err := a.App.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func Execute() {
	newApp().Execute()
}
