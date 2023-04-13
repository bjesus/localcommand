<h1 align="center">
  🖥️ Local Command 🖥️
</h1>

<p align="center">
  <b>Run local commands directly from the web!</b>
</p>



https://user-images.githubusercontent.com/55081/231873510-6f81c265-5e18-4462-844b-3214de348822.mp4



<p align="center">
  <img src="https://img.shields.io/badge/License-MIT-green.svg">
</p>

Local Command is an interface for running local commands directly from the web. It lets you use links with the `cmd://` scheme, like `cmd://ls -l` or any other command you want to run. It can be useful for triggering actions directly from internal dashboards.

## Usage

The `localcommand` program expects a string argument string starting with `cmd://` followed by a URL-encoded command, for example, `localcommand cmd://uname` or `localcommand cmd://ls%20-l`. After registering it as a scheme handler for `cmd://`, you can use `cmd://` links to invoke commands form your browser. Local Command asks for confirmation before running a command, unless you define this command as pre-approved. You can configure almost everything about the way it runs: whether to run commands in a terminal (and which), how to approve running commands, and what to do with their output.

You can test Local Command by trying to run `ps -aux` [here](https://cmd.bjesus.workers.dev/ps%20-aux).

## Configuration

Local Command requires a configuration file to run. On Linux, this file will be placed under `~/.config/localcommand/config.toml`. If you're unsure about the path in your OS, check out the [xdg](https://github.com/adrg/xdg) library, or just run Local Command and it will tell you where it expects the file.

See [example.toml](example.toml) for configuration reference.

## Installation

### Linux

1. Place the binary in your path
2. Place the [desktop file](localcommand.desktop) in `/usr/share/applications`

For Arch Linux, use [the AUR package](https://aur.archlinux.org/packages/localcommand).

### Troubleshooting

For Firefox, you might need to edit `handlers.json` and add the following under the `schemes` key:

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

Maybe, but I couldn't find out how. Do tell me if you find an issue! It's always a good idea to be cautious when running commands from the web. I recommend that you only use Local Command with trusted sources.

## License

This project is licensed under the MIT License.
