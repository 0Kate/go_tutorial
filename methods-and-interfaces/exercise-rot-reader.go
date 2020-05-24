package main

import (
    "io"
    "os"
    "strings"
)

type rot13Reader struct {
    r io.Reader
}

func (rot rot13Reader) Read(p []byte) (n int, err error) {
    n, err = rot.r.Read(p)
    for i := 0; i < len(p); i++ {
        if 'A' <= p[i] && p[i] <= 'Z' {
            p[i] = byte(int('A') + (((int(p[i]) - int('A')) + 13) % 26))
        } else if 'a' <= p[i] && p[i] <= 'z' {
            p[i] = byte(int('a') + (((int(p[i]) - int('a')) + 13) % 26))
        }
    }
    return
}

func main() {
    s := strings.NewReader("Lbh penpxrq gur pbqr!")
    r := rot13Reader{s}
    io.Copy(os.Stdout, &r)
}
