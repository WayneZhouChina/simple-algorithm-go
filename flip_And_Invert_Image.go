package main

func flipAndInvertImage(A [][]int) [][]int {
    for _, row := range A {
        row = reverse(row)
        for j, v := range row {
            row[j] = v ^ 1
        }
    }
    
    return A
}

func reverse(s []int) []int {
        for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
        s[i], s[j] = s[j], s[i]
    }

    return s
}
