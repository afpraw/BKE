package main

import (
    "fmt"
    "bufio"
    "os"
)


// Functie om het boord te printen
func printBoord(boord *[3][3]rune) {
    fmt.Println("---")
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            fmt.Printf("%c ", boord[i][j])
        }
        fmt.Println()
    }
    fmt.Println("---")
}

// Functie om de spelers te informeren wie aan de beurt is.
func printSpelerMsg(even bool) {
    var speler rune
    if even {
        speler = rune('A')
    } else {
        speler = rune('B')
    }
    fmt.Printf("Speler %c mag nu een zet doen\n", speler)
}


func isInputValid(char byte) bool {
    return char >= '0' && char <= '8'
}

// Functie dat kijkt of de zet mag
func isLegaleZet(boord *[3][3]rune, input byte, even bool) bool {
    if !isInputValid(input) {
        return false
    }

    digit := int(input - '0')
    i := digit / 3
    j := digit % 3

    if boord[i][j] != '.' {
        return false
    }

    if even {
        boord[i][j] = 'X'
    } else {
        boord[i][j] = 'O'
    }
    return true
}

// States enum
const (
    A = iota
    B
    Draw
    Playing
)

// State (enum) -> String
func boordStateToString(state int) string {
    switch state {
    case A:
        return "A heeft gewonnen"
    case B:
        return "B heeft gewonnen"
    case Draw:
        return "Draw"
    case Playing:
        return "Playing"
    default:
        return "WTF"
    }

}

func horizontal(boord *[3][3]rune, input rune) bool {
    for i := 0; i < 3; i++ {
        row := true
        for j := 0; j < 3; j++ {
            if boord[i][j] != input {
                row = false
                break
            }
        }
        if row {
            return true
        }
    }
    return false
}


func vertical(boord *[3][3]rune, input rune) bool {
    for i := 0; i < 3; i++ {
        column := true
        for j := 0; j < 3; j++ {
            if boord[j][i] != input {
                column = false
                break
            }
        }
        if column {
            return true
        }
    }
    return false
}


func diagonal(boord *[3][3]rune, input rune) bool {
    var diag bool = true
    for i := 0; i < 3; i++ {
        if boord[i][i] != input {
            diag = false
        }
    }
    if diag {
        return true
    }

    diag = true
    for i := 0; i < 3; i++ {
        if boord[i][2 - i] != input {
            diag = false
        }
    }

    if diag {
        return true
    }

    return false
}

func keepPlaying(boord *[3][3]rune) bool {
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if boord[i][j] == '.' {
                return true
            }
        }
    }
    return false
}


// Check the state of the bord
func checkState(boord *[3][3]rune) int {
    if horizontal(boord, 'X') || vertical(boord, 'X') || diagonal(boord, 'X') {
        return A
    }

    if horizontal(boord, 'O') || vertical(boord, 'O') || diagonal(boord, 'O') {
        return B
    }

    if keepPlaying(boord) {
        return Playing
    }

    return Draw
}


func main() {

    // reader := bufio.NewReader(os.Stdin)
    even := true

    // Het boord waarop we spelen
    var boord [3][3]rune

    // Maak het boord leeg
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            boord[i][j] = rune('.')
        }
    }

    // Game loop
    for {

        printSpelerMsg(even)
        printBoord(&boord)
        // Lees de zet in van de speler
        reader := bufio.NewReader(os.Stdin)
        char, err := reader.ReadByte()

        if err != nil {
            fmt.Println("Error reading input:", err)
            return
        }

        // Doet de speler mongool?
        if !isInputValid(char) {
           fmt.Print("Ga maar iets goeds intypen vriend")
           continue
        }

        // Plaatst zet als het kan anders niet
        if !isLegaleZet(&boord, char, even) {
           fmt.Print("Is geen legale zet mongool")
           continue
        }

        // Andere speler mag
        even = !even

        state := checkState(&boord)
        if state != Playing {
            fmt.Printf("Het spel is voorbij: %s\n", boordStateToString(state))
            printBoord(&boord)
            break
        }


    }

}
