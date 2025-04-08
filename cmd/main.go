package main

import (
	"log"
	"os"

	"shortlink/internal/i18n"
	"shortlink/internal/repository"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  i18n.AppName.Locale(),
		Usage: i18n.AppUsage.Locale(),
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "lang,l",
				Usage: i18n.FlagLangUsage.Locale(),
				Value: "en",
			},
		},
		Before: func(c *cli.Context) error {
			if lang := c.String("lang"); lang != "" {
				i18n.SetLocale(lang)
			}
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:  "start",
				Usage: i18n.StartCommandUsage.Locale(),
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "host,H",
						Usage:   i18n.FlagHostUsage.Locale(),
						Value:   "localhost",
						EnvVars: []string{"DB_HOST"},
					},
					&cli.StringFlag{
						Name:    "port,P",
						Usage:   i18n.FlagPortUsage.Locale(),
						Value:   "5432",
						EnvVars: []string{"DB_PORT"},
					},
					&cli.StringFlag{
						Name:    "user,u",
						Usage:   i18n.FlagUserUsage.Locale(),
						Value:   "postgres",
						EnvVars: []string{"DB_USER"},
					},
					&cli.StringFlag{
						Name:    "password,p",
						Usage:   i18n.FlagPasswordUsage.Locale(),
						Value:   "postgres",
						EnvVars: []string{"DB_PASSWORD"},
					},
					&cli.StringFlag{
						Name:    "dbname,d",
						Usage:   i18n.FlagDBNameUsage.Locale(),
						Value:   "shortlink_db",
						EnvVars: []string{"DB_NAME"},
					},
					&cli.StringFlag{
						Name:    "sslmode,s",
						Usage:   i18n.FlagSSLModeUsage.Locale(),
						Value:   "disable",
						EnvVars: []string{"DB_SSLMODE"},
					},
				},
				Before: func(c *cli.Context) error {
					c.Command.Usage = i18n.StartCommandUsage.Locale()
					for _, flag := range c.Command.Flags {
						stringFlag, ok := flag.(*cli.StringFlag)
						if !ok {
							continue
						}
						switch stringFlag.Name {
						case "host,H":
							stringFlag.Usage = i18n.FlagHostUsage.Locale()
						case "port,P":
							stringFlag.Usage = i18n.FlagPortUsage.Locale()
						case "user,u":
							stringFlag.Usage = i18n.FlagUserUsage.Locale()
						case "password,p":
							stringFlag.Usage = i18n.FlagPasswordUsage.Locale()
						case "dbname,d":
							stringFlag.Usage = i18n.FlagDBNameUsage.Locale()
						case "sslmode,s":
							stringFlag.Usage = i18n.FlagSSLModeUsage.Locale()
						}
					}
					return nil
				},
				Action: func(c *cli.Context) error {
					cfg := &repository.Config{
						Host:     c.String("host"),
						Port:     c.String("port"),
						User:     c.String("user"),
						Password: c.String("password"),
						DBName:   c.String("dbname"),
						SSLMode:  c.String("sslmode"),
					}

					repo, err := repository.NewRepository(cfg)
					if err != nil {
						return err
					}
					defer repo.Close()

					log.Println(i18n.ConnectedToDatabase.Locale())

					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
