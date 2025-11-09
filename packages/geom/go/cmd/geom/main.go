package main

import (
	"errors"
	jsoniter "github.com/json-iterator/go"
	"github.com/mojo-lang/mojo/go/pkg/logs"
	geom2 "github.com/mojo-lang/mojo/go/pkg/mojo/geom"
	"github.com/mojo-lang/mojo/go/pkg/mojo/geom/tile"
	"github.com/urfave/cli/v2"
	"io/fs"
	"log"
	"os"
	"strings"
)

func main() {
	app := &cli.App{
		Name:  "geom",
		Usage: "geom util, transform, simplify",
		Commands: []*cli.Command{
			{
				Name:    "transform",
				Aliases: []string{"p"},
				Usage:   "transform the coordinates for geojson",

				Flags: []cli.Flag{
					&cli.StringFlag{Name: "src", Aliases: []string{"s"}},
					&cli.StringFlag{Name: "dst", Aliases: []string{"d"}},
					&cli.StringFlag{Name: "from", Aliases: []string{"f"}},
					&cli.StringFlag{Name: "to", Aliases: []string{"t"}},
				},
				Action: func(ctx *cli.Context) error {
					src := ctx.String("src")
					dst := ctx.String("dst")
					from := ctx.String("from")
					to := ctx.String("to")
					if len(src) > 0 && len(dst) > 0 && strings.HasSuffix(src, ".geojson") && strings.HasSuffix(dst, ".geojson") {
						f, err := geom2.ParseSpatialReference(from)
						if err != nil {
							logs.Warnw("invalid from SpatialReference", "from type", from)
							return err
						}

						t, err := geom2.ParseSpatialReference(to)
						if err != nil {
							logs.Warnw("invalid from SpatialReference", "to type", from)
							return err
						}

						if !geom2.CoordTransformSupported(f) {
							logs.Warnw("invalid from SpatialReference", "from type", from)
							return errors.New("not supported SpatialReference")
						}
						if !geom2.CoordTransformSupported(t) {
							logs.Warnw("invalid to SpatialReference", "to type", to)
							return errors.New("not supported SpatialReference")
						}
						return coordTransformGeojson(src, dst, f, t)
					}
					return nil
				},
			},
			{
				Name:    "simplify",
				Aliases: []string{"s"},
				Usage:   "simplify the geojson file",

				Flags: []cli.Flag{
					&cli.StringFlag{Name: "src", Aliases: []string{"s"}},
					&cli.StringFlag{Name: "dst", Aliases: []string{"d"}},
				},
				Action: func(ctx *cli.Context) error {
					return nil
				},
			},
			{
				Name:    "mvt",
				Aliases: []string{"m"},
				Usage:   "view the mvt file",

				Flags: []cli.Flag{
					&cli.StringFlag{Name: "file", Aliases: []string{"f"}},
				},
				Action: func(ctx *cli.Context) error {
					filename := ctx.String("file")
					bytes, err := os.ReadFile(filename)
					if err != nil {
						return err
					}
					collections, err := tile.UnmarshalMvt(bytes)
					if err != nil {
						return err
					}
					str, _ := jsoniter.MarshalIndent(collections, "", "    ")
					print(string(str))
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func coordTransformGeojson(fromFile, toFile string, from, to geom2.SpatialReference) error {
	ff, err := os.ReadFile(fromFile)
	if err != nil {
		return err
	}
	g, err := geom2.NewGeoJsonFrom(ff)
	if err != nil {
		return err
	}
	tg := g.CoordTransform(from, to)
	out, err := jsoniter.Marshal(tg)
	if err != nil {
		return err
	}

	return os.WriteFile(toFile, out, fs.FileMode(0666))
}
