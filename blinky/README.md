# Blinky

Source: <https://tinygo.org/docs/tutorials/blinky/>

We use the `machine` package which is provided by TinyGo, but not recognized
by VS Code, so we don't get the autocompletion and type hints. To solve this
issue, follow the instructions about [IDE integration](https://tinygo.org/docs/guides/ide-integration/).

For example, on my laptop, `tinygo info arduino` outputs :

```
$ tinygo info arduino
LLVM triple:       avr
GOOS:              linux
GOARCH:            arm
GOARM:             6
build tags:        avr baremetal linux arm atmega328p atmega avr5 arduino tinygo math_big_pure_go gc.conservative scheduler.none serial.uart
garbage collector: conservative
scheduler:         none
cached GOROOT:     /home/boisgera/.cache/tinygo/goroot-1a3665987356bd3f26671cbe3d70be39fe7f2f9bf3a11bacdd5b79ba2096c3c9
```

So my `.vscode/settings.json` should be:

```
{
    "go.toolsEnvVars": {
        "GOROOT": "/home/boisgera/.cache/tinygo/goroot-1a3665987356bd3f26671cbe3d70be39fe7f2f9bf3a11bacdd5b79ba2096c3c9",
        "GOFLAGS": "-tags=avr,baremetal,linux,arm,atmega328p,atmega,avr5,arduino,tinygo,math_big_pure_go,gc.conservative,scheduler.none,serial.uart"
    }
}
```