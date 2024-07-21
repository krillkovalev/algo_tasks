package main
import "fmt"


func main() {
    prices :=  []int{1, 2}
    maxProfit(prices)
}

func maxProfit(prices []int) int {
    max := 0
    var chunk []int
    min := prices[0]
    for i := 0; i < len(prices); i++ {
        if prices[i] <= min && prices[i] >= 0 && i < len(prices) - 1 {
            min = prices[i]
            chunk = prices[i+1:]
        }
    }
    for _, v := range chunk {
        if v > max {
            max = v
        }
    }
    fmt.Printf("%d %d", min, max)
    if max <= min {
        return 0
    }
    return max - min
    
}