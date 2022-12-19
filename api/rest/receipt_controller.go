package rest

import (
	"encoding/json"
	"math"
	"net/http"
	"receipt-processor-challenge/dtos"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

var idPointsMap map[string]int

func getUUID() string {
	return uuid.New().String()
}

func stripString(name string) string {

	var result strings.Builder
	for i := 0; i < len(name); i++ {
		x := name[i]
		if ('a' <= x && x <= 'z') || ('A' <= x && x <= 'Z') ||
			('0' <= x && x <= '9') {
			result.WriteByte(x)
		}
	}
	return result.String()
}

func process(receipt dtos.Receipt) {

	totalPoints := 0
	name := stripString(receipt.Retailer)
	totalPoints += len(name)
	total, _ := strconv.ParseFloat(receipt.Total, 64)
	if total == float64(int(total)) {
		totalPoints += 50
	}
	if total*4 == float64(int(total*4)) {
		totalPoints += 25
	}
	totalPoints += (len(receipt.Items) / 2) * 5
	for _, item := range receipt.Items {
		descLength := len(strings.TrimSpace(item.ShortDescription))
		if descLength%3 == 0 {
			price, _ := strconv.ParseFloat(item.Price, 64)
			price *= 0.2
			totalPoints += int(math.Ceil(price))
		}
	}
	purchaseDate := receipt.PurchaseDate
	date := strings.Split(purchaseDate, "-")[2]
	if dt, _ := strconv.Atoi(date); dt%2 == 1 {
		totalPoints += 6
	}
	purchaseTime := receipt.PurchaseTime
	time := strings.Split(purchaseTime, ":")[0]
	if t, _ := strconv.Atoi(time); t >= 14 && t <= 16 {
		totalPoints += 10
	}
	idPointsMap[receipt.Id] = totalPoints
}

func ProcessReceipts(w http.ResponseWriter, r *http.Request) {

	var receipt dtos.Receipt
	_ = json.NewDecoder(r.Body).Decode(&receipt)
	receipt.Id = getUUID()
	process(receipt)
	resp := dtos.ProcessRequestResponse{
		Id: receipt.Id,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func GetPointsForReceipt(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id := params["id"]
	resp := dtos.PointsResponse{
		Points: idPointsMap[id],
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func init() {
	idPointsMap = make(map[string]int)
}
