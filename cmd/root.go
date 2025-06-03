package cmd

import (
	"fmt"
	"time"
	"slackstatus/internal/slack"
	"github.com/urfave/cli/v2"
)

func Run(args []string) error {
	app := &cli.App{
		Name:  "slackstatus",
		Usage: "Set your Slack status using presets or custom input",
		Commands: []*cli.Command{
		},
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "token", EnvVars: []string{"SLACK_TOKEN"}, Usage: "Slack OAuth token"},
			&cli.StringFlag{Name: "preset", Usage: "Name of status preset"},
			&cli.StringFlag{Name: "text", Usage: "Custom status text"},
			&cli.StringFlag{Name: "emoji", Usage: "Custom status emoji"},
			&cli.StringFlag{Name: "expire", Usage: "Duration before expiry (e.g., 25m, 1h)"},
			&cli.StringFlag{Name: "presets-file", Value: "presets.yaml", Usage: "Path to presets YAML or JSON"},
			&cli.BoolFlag{Name: "countdown", Usage: "Temporarily set status and auto-clear after duration"},
			&cli.BoolFlag{Name: "list", Usage: "List all available status presets"},
		},
		Action: func(c *cli.Context) error {
			// List presets
			if c.Bool("list") {
				return slack.ListPresets(c.String("presets-file"))
			}

			// Load token from env or fallback to saved file
			token := c.String("token")
			if token == "" {
			    return cli.Exit("Missing Slack token. Set with --token or SLACK_TOKEN env var.", 1)
			}

			var text, emoji, expireStr string

			// Use preset if specified
			if preset := c.String("preset"); preset != "" {
				presets, err := slack.LoadPresets(c.String("presets-file"))
				if err != nil {
					return cli.Exit(fmt.Sprintf("Error loading presets: %v", err), 1)
				}

				def, ok := presets[preset]
				if !ok {
					return cli.Exit("Unknown preset: "+preset, 1)
				}

				text = def.Text
				emoji = def.Emoji
				expireStr = def.Expiry
			} else {
				text = c.String("text")
				emoji = c.String("emoji")
				expireStr = c.String("expire")
			}

			// Skip validation if preset used (even if empty like 'off')
			if c.String("preset") == "" && text == "" && emoji == "" {
				return cli.Exit("Must provide either --preset or both --text and --emoji", 1)
			}

			expireUnix := slack.ParseDuration(expireStr)

			fmt.Printf("Setting status: %s %s (expires in: %s)\n", emoji, text, expireStr)
			if err := slack.SetStatus(token, text, emoji, expireUnix); err != nil {
				return cli.Exit(fmt.Sprintf("Failed to set status: %v", err), 1)
			}

			if c.Bool("countdown") && expireUnix > 0 {
				d, err := time.ParseDuration(expireStr)
				if err != nil {
					return cli.Exit(fmt.Sprintf("Invalid duration: %v", err), 1)
				}

				fmt.Printf("⏳ Waiting %s before clearing...\n", d)
				time.Sleep(d)
				fmt.Println("Clearing status...")
				return slack.ClearStatus(token)
			}

			return nil
		},
	}

	return app.Run(args)
}
