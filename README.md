# slackstatus

A lightweight CLI tool for setting your Slack status using reusable presets or custom messages — powered by Go and a single `SLACK_TOKEN`.

---

## Features

- ✅ Set custom Slack status text, emoji, and expiration
- 🧠 Reusable status presets (defined in YAML or JSON)
- ⏳ `--countdown` to auto-clear your status after time elapses
- 🔍 `--list` to show available presets
- 🔐 Reads token from `SLACK_TOKEN` environment variable or `--token` flag

---

## Installation

```bash
git clone https://github.com/yourusername/slackstatus.git
cd slackstatus
go mod tidy
make build
```

---

## Usage Examples

### Set status using a preset

```bash
SLACK_TOKEN=your-token ./slackstatus --preset lunch
```

### Temporary status with auto-clear

```bash
./slackstatus --preset focus --countdown
```

### Custom status with expiry

```bash
./slackstatus --emoji ":coffee:" --text "Coffee break" --expire 20m
```

### List available presets

```bash
./slackstatus --list
```

---

## Presets Configuration

Create a `presets.yaml` (or `.json`) file in your project or home config folder.

```yaml
meeting:
  text: "In a meeting"
  emoji: ":spiral_calendar_pad:"
  expiry: "0s"

lunch:
  text: "Out for lunch"
  emoji: ":fork_and_knife:"
  expiry: "1h"

focus:
  text: "Deep focus time"
  emoji: ":headphones:"
  expiry: "25m"

vacation:
  text: "On vacation"
  emoji: ":palm_tree:"
  expiry: "0s"

off:
  text: ""
  emoji: ""
  expiry: "0s"
```

Override the file path using:

```bash
./slackstatus --preset lunch --presets-file custom.yaml
```

---

## Slack Token Setup

1. Create a Slack App at [api.slack.com/apps](https://api.slack.com/apps)
2. Add the OAuth scope: `users.profile:write`
3. Install the app to your workspace
4. Copy your **User OAuth Token** (starts with `xoxp-...`)
5. Set it as an environment variable:

```bash
export SLACK_TOKEN=xoxp-...
```

---

## Development

```bash
make build         # compile
make run ARGS="--list"
make clean         # delete binary
make lint          # run go vet
```
