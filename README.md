<h1 align="center">
  🖥️ Local Command 🖥️
</h1>

<p align="center">
  <b>Run local commands directly from the web!</b>
</p>
<p align="center">
  <img src="https://img.shields.io/badge/License-MIT-green.svg">
</p>



https://user-images.githubusercontent.com/55081/231964667-556f9308-22d4-42bc-b091-a7072853bfb1.mp4



---
Local Command is an interface for running local commands directly from the web. It lets you use links with the `cmd://` scheme, like `cmd://ls -l` or any other command you want to run. It can be useful for triggering actions directly from internal dashboards.

## Usage

The `localcommand` program expects a string argument string starting with `cmd://` followed by a URL-encoded command, for example, `localcommand cmd://uname` or `localcommand cmd://ls%20-l`. After registering it as a scheme handler for `cmd://`, you can use `cmd://` links to invoke commands form your browser. Local Command asks for confirmation before running a command, unless you define this command as pre-approved. You can configure almost everything about the way it runs: whether to run commands in a terminal (and which), how to approve running commands, and what to do with their output.

You can test Local Command by trying to run `echo it worked!` [here](https://cmd.bjesus.workers.dev/echo%20it%20worked!).

## Configuration

Local Command requires a configuration file to run. On Linux, this file will be placed under `~/.config/localcommand/config.toml`. On macOS, it's `~/Library/Application Support/localcommand/config.toml`. If you're unsure about the path in your OS, check out the [xdg](https://github.com/adrg/xdg) library, or just run Local Command and it will tell you where it expects the file.

See [example.toml](example.toml) for configuration reference.

## Installation

Obtain the `localcommand` binary from the [Releases](https://github.com/bjesus/localcommand/releases) page or compile it yourself.

### Linux

1. Place the binary in your path
2. Place the [desktop file](localcommand.desktop) in `/usr/share/applications`

For Arch Linux, use [the AUR package](https://aur.archlinux.org/packages/localcommand).


### macOS

1. Place the binary under `/Applications`
2. `git clone` the repo and copy to `Local Command.app` folder to `/Applications`
3. If using the default config.toml, make sure to `brew install zenity` first. You might need to set your config to use `/usr/loca/bin/zenity` instead of `zenity`
4. If the `cmd://` scheme still isn't working, try `brew install duti && duti -s org.bjesus.LocalCommand cmd`

I'm not a macOS user, so if you think there's an easier way to set this up, let me know!

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

Maybe, but I couldn't find out how. Do tell me if you find an issue! It's always a good idea to be cautious when running commands from the web. I recommend that you only use Local Command with trusted sources. If you're running the command in a terminal, I highly recommend not using `sh` because it might allow running unapproved commands as approved, e.g. by calling commands like `approved && unapproved`.

## Related

- [Zenity](https://help.gnome.org/users/zenity/3.32/) - used in the example config to confirm and show commands output
- [Platypus](https://github.com/sveinbjornt/Platypus/) - used to wrap the binary with a macOS bundle app (whatever that is!) and register the scheme
- [duti](https://github.com/moretension/duti/) - a useful tool for selecting default applications for file types and URI schemes

## License

This project is licensed under the MIT License.
