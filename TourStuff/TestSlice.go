package main

func Pic(dx, dy int) [][]uint8 {
    pic := make([][]uint8, dy)

    for y := 0; y < dy; y++ {
        pic[y] = make([]uint8, dx)

        for x := 0; x < dx; x++ {
            pic[y][x] = uint8(x * y)
        }
    }

    return pic
}

func main() {
    result := Pic(5, 5)

    for _, row := range result {
        for _, value := range row {
            print(value, " ")
        }
        println()
    }
}
