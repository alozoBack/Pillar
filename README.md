# MCPanel 
It's a simple panel over ssh for Minecraft server with rcon, built entirely in Go with [charm_](https://charm.sh/).


> [!IMPORTANT]
> MCPanel is currently under development.

## Features (WIP)
- [ ] Access control
- [ ] Config file
- [ ] Auto install
- Server Managment
    - [x] top server 
    - [x] Execute commands 
        - [x] Check TPS
        - [x] Check players
        - [x] Give admin permission
        - [ ] LuckPerms editor open ( idk how to do this )
    - [ ] Check logs
    - [ ] Plugins managment
    - [ ] Config managment

## Usage
### Build
```bash
go build
```
### Run
Import your server rcon password in export.
Example with sh script:
```bash
#/bin/bash
export HOSTRCON="localhost:25575"
export PASSWORDRCON=""
export ADMINCOMMAND=""
./mcpanel
```

## License
[![WTFPL](http://www.wtfpl.net/download/wtfpl-badge-3/)](https://github.com/alozoBack/mcpanel/raw/main/LICENSE)

---

Build on [Charm](https://charm.sh).

<a href="https://charm.sh/"><img alt="The Charm logo" src="https://stuff.charm.sh/charm-badge.jpg" width="400"></a>

Charm热爱开源 • Charm loves open source

