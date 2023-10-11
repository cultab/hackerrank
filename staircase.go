package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

/*
 * Complete the 'staircase' function below.
 *
 * The function accepts INTEGER n as parameter.
 *
 *       _____#
 *       ____##
 *       ___###
 *       __####
 *       _#####
 *       ######
 *
 *
 *
 */

func staircase(n int32) {
    N := int(n)
    // Write your code here
    for i := 0; i < N; i++ {
        for j := 0; j < N; j++ {
            if j + 1 < N - i {
                fmt.Print(" ")
            } else {
                fmt.Print("#")
            }
        }
        fmt.Print("\n")
    }

}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    n := int32(nTemp)

    staircase(n)
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
