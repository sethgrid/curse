## Curse

A utility for manipulating the terminal cursor. Current feature set:
- Get terminal cursor position
- Move cursor
- Move up, down n-lines
- Clear line
- Clear Screen (up, down, all)
- Set Color

Example usage:


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
