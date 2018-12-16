package main

import (
	"fmt"
	"math/rand"
	"time"
	"net/http"
	"encoding/json"
	"strconv"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	http.HandleFunc("/v1/lottery", apiRequestHandler)
	http.ListenAndServe(":9000", nil)
}

type ResponseStruct struct {
	Result []int	`json:"result"`
}

func apiRequestHandler(w http.ResponseWriter, r *http.Request) {
	// GETリクエストかどうか判定
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)	// 405
		error := make(map[string]string)
		error["Error"] = "Method Not Allowed"
		json.NewEncoder(w).Encode(error)
		return
	}

	// クエリパラメータの取得
	requestParams := r.URL.Query()

	// ガチャ実行
	results := runLottery(requestParams)

	// レスポンス作成
	response := ResponseStruct{}
	for _, result := range results {
		response.Result = append(response.Result, result)
		fmt.Printf("%d\n", result)
	}
	json.NewEncoder(w).Encode(response)
}

func runLottery(params map[string][]string) []int {
	const TYPE_NUM= 10;
	probabilities := parseProbability(params, TYPE_NUM)
	lotteryNum := parseLotteryNum(params)

	results := []int{}
	for i := 0; i < lotteryNum; i++ {
		lotteryResult := runLotteryOnce(probabilities)
		results = append(results, lotteryResult)
	}

	return results
}

func parseProbability(params map[string][]string, maxNum int) []float64 {
	probabilities := [] float64{}
	for i := 0; i < maxNum; i++ {
		key := "type" + strconv.Itoa(i + 1)
		if len(params[key]) > 0 {
			probability, _ := strconv.ParseFloat(params[key][0], 64)	
			probabilities = append(probabilities, probability)
		} else {
			probabilities = append(probabilities, 0.0)
		}
	}

	return probabilities
}

func parseLotteryNum(params map[string][]string) int {
	key := "num"
	num, _ := strconv.Atoi(params[key][0])
	return num
}

func runLotteryOnce(probabilities []float64) int {
	random := rand.Float64()
	total := 0.0
	for index, probability := range probabilities {
		if total <= random && total + probability > random {
			return index
		}
		total += probability
	}

	return -1;
}
