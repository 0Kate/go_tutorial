package main

import "golang.org/x/tour/pic"

func f(x, y uint8) uint8 {
    return x ^ y
}

func Pic(dx, dy, int) [][]uint8 {
    picture := make([][]uint8, dy)

    for i := 0; i < dy; i++ {
        pictureRow := make([]uint8, dx)
        for j := 0; j < dx; j++ {
            pictureRow[j] = f(uint8(i), uint8(j))
        }
        picture[i] = pictureRow
    }

    return picture
}

func main() {
    pic.Show(Pic)
}
