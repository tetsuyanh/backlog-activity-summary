package main

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/urfave/cli"

	"github.com/tetsuyanh/backlog-activity-summary/nulab/backlog"
	"github.com/tetsuyanh/backlog-activity-summary/renderer"
	"github.com/tetsuyanh/backlog-activity-summary/service"
)

const (
	EXIT_OK    = 0
	EXIT_ERROR = 1

	ENV_BACKLOG_TEAM  = "BACKLOG_TEAM"
	ENV_BACKLOG_TOKEN = "BACKLOG_TOKEN"
)

func main() {
	os.Exit(Run(os.Args))
}

// Run is to run application
func Run(args []string) int {
	team := os.Getenv(ENV_BACKLOG_TEAM)
	token := os.Getenv(ENV_BACKLOG_TOKEN)
	if team == "" || token == "" {
		log.Printf("require env %s, %s\n", ENV_BACKLOG_TEAM, ENV_BACKLOG_TOKEN)
		return EXIT_ERROR
	}
	backlog.Setup(team, token)

	app := cli.NewApp()
	app.Name = "backlog-activity-summary"
	app.Usage = "export summary of backlog-activity"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "begin, b",
			Usage: "'begin' of period [e.g.'20180601']",
		},
		cli.StringFlag{
			Name:  "end, e",
			Usage: "'end' of period [e.g.'20180630']",
		},
	}
	app.Action = func(c *cli.Context) error {
		bStr := c.String("b")
		if bStr == "" {
			return cli.NewExitError(errors.New("require 'begin' of period: -b 20180601"), 1)
		}
		eStr := c.String("e")
		if eStr == "" {
			return cli.NewExitError(errors.New("require 'end'' of period: -e 20180630"), 1)
		}

		bTime, _ := time.Parse("20060102", bStr)
		begin := time.Date(bTime.Year(), bTime.Month(), bTime.Day(), 0, 0, 0, 0, time.UTC)
		eTime, _ := time.Parse("20060102", eStr)
		end := time.Date(eTime.Year(), eTime.Month(), eTime.Day(), 23, 59, 59, 999, time.UTC)

		us := service.NewUserService()
		me, errMyself := us.Myself()
		if errMyself != nil {
			log.Printf("errMyself: %s\n", errMyself)
			return cli.NewExitError(errMyself, 1)
		}
		sm, errSummary := us.SummaryPeriod(me.ID, me.Name, begin, end)
		if errSummary != nil {
			log.Printf("errSummary: %s\n", errSummary)
			return cli.NewExitError(errSummary, 1)
		}

		if errRender := renderer.NewMdRenderer().RenderSummary(sm); errRender != nil {
			log.Printf("errRender: %s\n", errRender)
			return cli.NewExitError(errRender, 1)
		}
		return nil
	}

	app.Run(os.Args)
	return EXIT_OK
}
