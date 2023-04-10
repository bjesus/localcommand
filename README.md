# Local Command

Local Command is an interface for running local commands directly from the web. It lets you use links with the `cmd://` scheme, like `cmd://ls -l` or any other command you want to run. It can be useful for triggering actions directly from internal dashboards.

## Usage

The `localcommand` program expects a URL-encoded command as an argument, for example, `ls%20%2Fhome%2F` (for `ls /home/`). After properly registering it as scheme handler for `command://`, it means you can use links like `command://ls /home`. Local Command asks for confirmation before running a command, unless you define this command as pre-approved. You can configure almost everything about the way it runs: whether to run commands in a terminal (and which), how to approve running commands, and what to do with their output.

## Configuration

Local Command requires a configuration file to run. On Linux, this file will be placed under `~/.config/localcommand/config.toml`. If you're unsure about the path in your OS, check out [XDG](https://github.com/adrg/xdg), or just run Local Command and it will tell you where it expects the file.

See [example.toml](example.toml) for configuration reference.

## Installation

### Linux

1. Place the binary in your path
2. Place the [desktop file](localcommand.desktop) in /usr/share/applications

For Arch Linux, use [the AUR package](https://aur.archlinux.org/packages/localcommand).

### Troubleshooting

For Firefox, I had to edit `handlers.json` and add the following under the `schemes` key:

```json
"cmd": {
  "action": 2,
  "handlers": [
    {
      "name": "localcommand",
      "path": "/usr/bin/localcommand"
    }
  ]
},
```

## Isn't this dangerous?

Maybe, but I couldn't find out how. Do tell me if you find an issue!
