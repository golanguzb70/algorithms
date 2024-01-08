package main

import (
	"fmt"
	"net/http"
)

func main() {
	// arr := []int{7, 7, 7, 7} // [1, 2, 3] [1, 2, 4] [1, 2, 3, 4] [2, 3, 4]
	// arr := []int{2,4,6,8,10}
	// arr := []int{1,2,3,4}
	// fmt.Println(leetcode.NumberOfArithmeticSlices(arr))
	Send("Hello")
}

func Send(text string) {
	client := &http.Client{}

	var botUrl = fmt.Sprintf("https://api.telegram.org/bot"+"6338261043:AAGp1mgNQ9VYlR0oIqLbl5K9KYc1VGD5Y5Y"+"/sendMessage?chat_id="+"694596449"+"&text=%s", text)
	request, err := http.NewRequest("GET", botUrl, nil)
	if err != nil {
		return
	}
	resp, err := client.Do(request)
	if err != nil {
		return
	}

	defer resp.Body.Close()
}
