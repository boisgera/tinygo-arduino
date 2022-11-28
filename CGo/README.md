

 ## ["How to call CGo from TinyGo"](https://github.com/tinygo-org/tinygo/issues/60)
 

> There is very limited support for CGo.

(2018)


> when you use CGo you'll have to output a .o file and 
> link it manually with the .o files from the C files you're using.

> The only thing "supported" at the moment is function calls. 


> CGo support was again improved a lot in the latest release. It is now starting to be usable.

(2019)

> Is there any doc or other information on the state of using C other than this issue?

> Not really. I would suggest trying things out and seeing whether it works or not.
> Known limitations:
>
> No C.GoString and related functions.  
> Limited support for C macros (for example, #define FOO (3 + 4) won't work, but #define FOO (3) would).

> Do you think it could be possible to use this to call existing sensor libraries written in C instead of having to rewrite them in TinyGo?
> 
> Maybe. The best way to know that is to try it out.

> On the other hand, the tinyfs project and @aykevl 's suggestion of moving the C code to the same package of the Go files were really useful. Simple C code is now compiling along with the TinyGo source code with no extra efforts, thanks!

TinyGo Bluetooth
--------------------------------------------------------------------------------

<https://github.com/tinygo-org/bluetooth>


TinyGo TinyFS
--------------------------------------------------------------------------------

<https://github.com/tinygo-org/tinyfs>


--------------------------------------------------------------------------------

```go
//#cgo CFLAGS: -I/home/ron/.gvm/pkgsets/go1.11/global/src/github.com/aykevl/tinygo/src/examples/cxp

package main

/*
#include "hello.h"
*/
import "C"

func main() {
    for {
        println(C.fortytwo())
    }
}
```

```bash
tinygo build -target=microbit -o=cxp.hex -no-debug examples/cxp
```

--------------------------------------------------------------------------------

So far I am testing an example with TinyGo on Linux with manual compilation
of the C code and I am having a link error:

    tinygo:ld.lld: error: undefined symbol: add
    >>> referenced by app.go:10 (/home/boisgera/VOYAGER/ENS/IDS/tinygo-arduino/CGo/sandbox/app.go:10)
    >>>               /home/boisgera/.cache/tinygo/thinlto/llvmcache-03103C6F40EF284187A4533805A1811F6198F876:(runtime.run$1$gowrapper)
    failed to run tool: ld.lld
    error: failed to link /tmp/tinygo3193446596/main: exit status 1

Update: removed manual compilation (I don't generate the .o, tinygo does),
same error.

--------------------------------------------------------------------------------

**TODO.** 

  - read/tweak more about tinyfs, see how it works

  - tinygo should manage the c file compilation, that's only the linking
    step that has issues.

  - is the lack of proper module declaration an issue here?

  - try Go version of CGo first? (would make sense)