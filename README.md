# docker-selector [![Build Status](https://travis-ci.org/ujiro99/docker-selector.svg?branch=master)](https://travis-ci.org/ujiro99/docker-selector)
Docker container selector using peco.

![demo](https://github.com/ujiro99/docker-selector/blob/master/demo.gif)

# Requirements
* docker
* peco

# Installation

You can get binary from github release page.

[-> Release Page](https://github.com/ujiro99/docker-selector/releases)

or, use `go get`:

```bash
$ go get -u github.com/ujiro99/docker-selector
```

# Usage
Use this with key bindings.  
An example of calling with `Ctrl + d`.

* bash
```bash
# add this .bashrc
peco-docker-selector() {
    local l=$(\docker-selector -a)
    READLINE_LINE="${READLINE_LINE:0:$READLINE_POINT}${l}${READLINE_LINE:$READLINE_POINT}"
    READLINE_POINT=$(($READLINE_POINT + ${#l}))
}
bind -x '"\C-d": peco-docker-selector'
```

* nyagos
```lua
-- add this .nyagos
nyagos.bindkey("C_D", function(this)
    local result = nyagos.eval('docker-selector.exe -a')
    this:call("CLEAR_SCREEN")
    return result
end)
```
