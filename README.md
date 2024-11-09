# Pillar 
It's a simple panel over ssh for Minecraft server with rcon, built entirely in Go with [charm_](https://charm.sh/).


> [!IMPORTANT]
> Pillar is currently under development.

## Features (WIP)
- [ ] Access control
- [ ] Config file
- [ ] Auto install
- Server Managment
    - [x] stop server 
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
- Import your server rcon password in env.
  
  Example env file:
  ```env
  HOSTRCON="localhost:25575"
  PASSWORDRCON=""
  ADMINCOMMAND="lp user %user% permission set admin true" # %user% will be used as username in panel
  ```
- Connect to panel
  
  <img alt="Demo panel" src="https://raw.githubusercontent.com/alozoBack/mcpanel/refs/heads/main/demo.gif" width="600" />

## License
[![WTFPL](http://www.wtfpl.net/download/wtfpl-badge-3/)](https://github.com/alozoBack/mcpanel/raw/main/LICENSE)

---

Build on [Charm](https://charm.sh).

<a href="https://charm.sh/"><img alt="The Charm logo" src="https://stuff.charm.sh/charm-badge.jpg" width="400"></a>

Charm热爱开源 • Charm loves open source

