// cmd/artificialeaselcoreultra/main.go
package main

import (
"flag"
"log"
"os"

"artificialeaselcoreultra/internal/artificialeaselcoreultra"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := artificialeaselcoreultra.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
