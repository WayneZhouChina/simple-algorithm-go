package main

func heightChecker(heights []int) int {
    if heights == nil || len(heights) == 0 {
        return 0
    }
    
    heightsCopy := make([]int, len(heights))
    copy(heightsCopy, heights)
    sort.Ints(heights)
    
    var count int
    for i:=0; i<len(heights);i++ {
        if heightsCopy[i] != heights[i] {
            count++
        }
    }
    return count
}
