package main

func NQueen(n int ) int {
	if n > 32 {
		return -1 
	}
	// 1 代表可以放
	limit := 1 << n - 1
	return process(limit,0,0,0)
}

func process(limit int , colLim,LeftLim,RightLim int) int {
	
}