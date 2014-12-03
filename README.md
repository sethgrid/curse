## Curse

A utility for manipulating the terminal cursor. Current feature set:
- Get terminal cursor position
- Move cursor
- Move up, down n-lines
- Clear line
- Clear Screen (up, down, all)
- Set Color

Basic Example usage (see below for an inline-progress bar):

```go
    package main

    import (
        "fmt"
        "log"

        "github.com/sethgrid/curse"
    )

    func main() {

        c, err := curse.New()
        if err != nil {
            log.Fatal(err)
        }

        c.SetColorBold(curse.RED).SetBackgroundColor(curse.BLACK)
        fmt.Println("Position: ", c.Position)
        c.SetDefaultStyle()
        fmt.Println("something to be erased")
        c.MoveUp(1).EraseCurrentLine().MoveDown(1)
    }
```

Progress Bar Example:

```go
    package main

    import (
        "fmt"
        "strings"
        "time"

        "github.com/sethgrid/curse"
    )

    func main() {
        fmt.Println("Progress Bar")
        total := 150
        progressBarWidth := 80
        c, _ := curse.New()

        // give some buffer space on the terminal
        fmt.Println()

        // display a progress bar
        for i := 0; i <= total; i++ {
            c.MoveUp(1)
            c.EraseCurrentLine()
            fmt.Printf("%d/%d ", i, total)

            c.MoveDown(1)
            c.EraseCurrentLine()
            fmt.Printf("%s", progressBar(i, total, progressBarWidth))

            time.Sleep(time.Millisecond * 25)
        }
        // end the previous last line of output
        fmt.Println()
        fmt.Println("Complete")
    }

    func progressBar(progress, total, width int) string {
        bar := make([]string, width)
        for i, _ := range bar {
            if float32(progress)/float32(total) > float32(i)/float32(width) {
                bar[i] = "*"
            } else {
                bar[i] = " "
            }
        }
        return "[" + strings.Join(bar, "") + "]"
    }
```
