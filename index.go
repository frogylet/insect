package main

import (
    "fmt"
    "strings"
    "time"

    "github.com/nsf/termbox-go"
)

var (
    options     = []string{"Well, ¿and you?", "Very good", "Nice"}
    selected    = -1 // none selected :p
    cursorIndex = 0
    inputText   = ""
    progressVal = 0
)


// ANSI escape codes 
const (
    Red    = "\033[31m"
    Green  = "\033[32m"
    Yellow = "\033[33m"
    Gray   = "\033[90m"
    Reset  = "\033[0m"
)

// Printf with color
func Printf(text string, color string) {
    fmt.Println(color + text + Reset)
}



func drawSelector() {
    for i, opt := range options {
        prefix := "[ ]"
        if selected == i {
            prefix = "[x]"
        }
        style := termbox.ColorWhite
        if i == cursorIndex {
            style = termbox.ColorGreen | termbox.AttrBold
        }
        drawText(2, 2+i, fmt.Sprintf("%s %s", prefix, opt), style, termbox.ColorDefault)
    }
}

func drawInput(placeholderInp string) {
    text := inputText
    style := termbox.ColorYellow
    if text == "" {
        text = placeholderInp
        style = termbox.ColorBlack
    }
    drawText(2, 6+len(options), "Input: "+text, style, termbox.ColorDefault)
}




func drawProgressBar(percent int) {
    blocks := percent / 2
    bar := ""
    for i := 0; i < blocks; i++ {
        bar += "█"
    }
    bar = strings.TrimSpace(bar)
    drawText(2, 15+len(options), fmt.Sprintf("Progress: [%-20s] %d%%", bar, percent), termbox.ColorMagenta, termbox.ColorDefault)
}

func animateProgressBar() {
    for i := 0; i <= 100; i++ {
        progressVal = i
        termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
        drawHeader()
        drawSelector()
        drawInput("writting..")
        drawProgressBar(progressVal)
        termbox.Flush()
        time.Sleep(30 * time.Millisecond)
    }
}

func drawHeader() {
    drawText(2, 0, "q = exit", termbox.ColorBlack|termbox.AttrBold, termbox.ColorDefault)
}

func drawText(x, y int, text string, fg, bg termbox.Attribute) {
    for i, r := range text {
        termbox.SetCell(x+i, y, r, fg, bg)
    }
}




func main() {
    err := termbox.Init()
    if err != nil {
        panic(err)
    }
    defer termbox.Close()

    go animateProgressBar()

    for {
        termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
        drawHeader()
        drawSelector()
        drawInput("writting..")
        drawProgressBar(progressVal)
        termbox.Flush()

        switch ev := termbox.PollEvent(); ev.Type {
        case termbox.EventKey:
            switch ev.Key {
            case termbox.KeyArrowUp:
                if cursorIndex > 0 {
                    cursorIndex--
                }
            case termbox.KeyArrowDown:
                if cursorIndex < len(options)-1 {
                    cursorIndex++
                }
            case termbox.KeyEnter:
                selected = cursorIndex // selection
            case termbox.KeyEsc:
                return
            default:
                if ev.Ch == 'q' {
                    return
                }
                if ev.Ch != 0 {
                    inputText += string(ev.Ch)
                }
            }
        }
    }
}

