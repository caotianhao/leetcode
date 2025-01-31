package main

import "fmt"

func minOperationsMaxProfit(customers []int, boardingCost, runningCost int) int {
	ans := -1
	maxProfit1599, totalProfit, operations, customersCount := 0, 0, 0, 0
	for _, c := range customers {
		operations++
		customersCount += c
		curCustomers := min(customersCount, 4)
		customersCount -= curCustomers
		totalProfit += boardingCost*curCustomers - runningCost
		if totalProfit > maxProfit1599 {
			maxProfit1599 = totalProfit
			ans = operations
		}
	}
	if customersCount == 0 {
		return ans
	}
	profitEachTime := boardingCost*4 - runningCost
	if profitEachTime <= 0 {
		return ans
	}
	if customersCount > 0 {
		fullTimes := customersCount / 4
		totalProfit += profitEachTime * fullTimes
		operations += fullTimes
		if totalProfit > maxProfit1599 {
			maxProfit1599 = totalProfit
			ans = operations
		}
		remainingCustomers := customersCount % 4
		remainingProfit := boardingCost*remainingCustomers - runningCost
		totalProfit += remainingProfit
		if totalProfit > maxProfit1599 {
			maxProfit1599 = totalProfit
			ans++
		}
	}
	return ans
}

func main() {
	fmt.Println(minOperationsMaxProfit([]int{8, 3}, 5, 6))            //3
	fmt.Println(minOperationsMaxProfit([]int{10, 9, 6}, 6, 4))        //7
	fmt.Println(minOperationsMaxProfit([]int{3, 4, 0, 5, 1}, 1, 92))  //-1
	fmt.Println(minOperationsMaxProfit([]int{10, 10, 6, 4, 7}, 3, 8)) //9
	fmt.Println(minOperationsMaxProfit([]int{0, 43, 37, 9, 23, 35, 18, 7, 45, 3, 8, 24,
		1, 6, 37, 2, 38, 15, 1, 14, 39, 27, 4, 25, 27, 33, 43, 8, 44, 30, 38, 40, 20,
		5, 17, 27, 43, 11, 6, 2, 30, 49, 30, 25, 32, 3, 18, 23, 45, 43, 30, 14, 41, 17,
		42, 42, 44, 38, 18, 26, 32, 48, 37, 5, 37, 21, 2, 9, 48, 48, 40, 45, 25, 30, 49,
		41, 4, 48, 40, 29, 23, 17, 7, 5, 44, 23, 43, 9, 35, 26, 44, 3, 26, 16, 31, 11,
		9, 4, 28, 49, 43, 39, 9, 39, 37, 7, 6, 7, 16, 1, 30, 2, 4, 43, 23, 16, 39, 5, 30,
		23, 39, 29, 31, 26, 35, 15, 5, 11, 45, 44, 45, 43, 4, 24, 40, 7, 36, 10, 10, 18,
		6, 20, 13, 11, 20, 3, 32, 49, 34, 41, 13, 11, 3, 13, 0, 13, 44, 48, 43, 23, 12,
		23, 2}, 43, 54)) //993
}
