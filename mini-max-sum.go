package main

import (
	"fmt"
	"math"
)

func miniMaxSum(arr []int32) {
    // Write your code here
    var min int64 = math.MaxInt64
    var max int64 = math.MinInt64
    
    for _, n := range arr {
        temp := int64(0)
        for _,  k := range arr {
            temp += int64(k)
        }
        temp -= int64(n)

        fmt.Println("#",temp)
        if temp > max {
            max = temp
        }
        if temp < min {
            min = temp
        }
    }

    fmt.Println(min, max)
}

func main() {
    miniMaxSum([]int32{769082435,210437958,673982045,375809214,380564127})
}
