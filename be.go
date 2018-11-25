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
	fmt.Printf("hello, world\n")

	rand.Seed(time.Now().UnixNano())

	computeProbability(0.5)

	http.HandleFunc("/v1/lottery", apiRequestHandler)

	http.ListenAndServe(":8080", nil)
}

func computeProbability(probability float64) bool {
	random := rand.Float64()
	fmt.Println(random)
	if random < probability {
		return true
	} else {
		return false
	}
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
	runLottery(requestParams)

	// レスポンスを返還
	json.NewEncoder(w).Encode(requestParams)
}

func runLottery(params map[string][]string) {
	fmt.Printf("%v\n", params)

	probabilities := parseProbability(params)

	for _, v := range probabilities {
		fmt.Printf("%f\n", v)
	}
}

func parseProbability(params map[string][]string) []float64 {
	const TYPE_NUM= 10;
	probabilities := [] float64{}
	for i := 0; i < TYPE_NUM; i++ {
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