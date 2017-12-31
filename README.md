# docker-selector
Docker container selector using peco.

# Requirements
* docker
* peco

# Installation
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
    local l=$(\docker-selector)
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
