package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	listPlayer := make(map[int]map[string]any)

	fmt.Print("Masukkan jumlah pemain: ")
	scanner.Scan()
	inputPlayers := scanner.Text()

	fmt.Print("Masukkan jumlah dadu: ")
	scanner.Scan()
	inputDices := scanner.Text()

	players, _ := strconv.Atoi(inputPlayers)
	dices, _ := strconv.Atoi(inputDices)

	for i := 0; i < players; i++ {
		listPlayer[i] = map[string]any{
			"score":      0,
			"dices":      []int{},
			"count_dice": dices,
		}
	}

	fmt.Printf("Pemain = %d, Dadu = %d \n", players, dices)
	fmt.Println("================")

	counter := 1
	for true {
		if !checkContinuePlay(listPlayer) {
			break
		}

		fmt.Printf("Giliran %d lempar dadu :\n", counter)
		listPlayer = throwDice(listPlayer)
		for i, player := range listPlayer {
			fmt.Printf("Pemain #%d (%d) : %v \n", i+1, player["score"].(int), player["dices"])
		}

		fmt.Println("Setelah evaluasi :")
		listPlayer = evaluate(listPlayer)
		for i, player := range listPlayer {
			fmt.Printf("Pemain #%d (%d) : %v \n", i+1, player["score"].(int), player["dices"])
		}

		counter++
		fmt.Println("================")
	}

	winner := 0
	highestScore := 0
	lastPlayer := 0
	for i, player := range listPlayer {
		if player["score"].(int) > highestScore {
			highestScore = player["score"].(int)
			winner = 1
		}

		if len(player["dices"].([]int)) > 0 {
			lastPlayer = i
		}
	}

	fmt.Printf("Game berakhir karena hanya pemain #%d yang memiliki dadu.\n", lastPlayer+1)
	fmt.Printf("Game dimenangkan oleh pemain #%d karena memiliki poin lebih banyak dari pemain lainnya. \n", winner+1)
}

func getDiceNumber(n int) []int {
	data := []int{}

	for i := 0; i < n; i++ {
		data = append(data, rand.Intn(7-1)+1)
	}

	return data
}

func checkContinuePlay(listPlayer map[int]map[string]any) bool {
	counter := 0

	for _, player := range listPlayer {
		if player["count_dice"].(int) > 0 {
			counter++
		}
	}

	if counter == 1 {
		return false
	}

	return true
}

func throwDice(listPlayer map[int]map[string]any) map[int]map[string]any {
	for _, player := range listPlayer {
		player["dices"] = getDiceNumber(player["count_dice"].(int))
	}
	return listPlayer
}

func evaluate(listPlayer map[int]map[string]any) map[int]map[string]any {
	var (
		ok    bool
		score int
	)

	for i, player := range listPlayer {
		// check user has number 6
		player["dices"], score, ok = hasNumber6(player["dices"].([]int))
		if ok {
			player["count_dice"] = player["count_dice"].(int) - score
			player["score"] = player["score"].(int) + score
		}
		// check user has number 1
		if i == len(listPlayer)-1 {
			player["dices"], listPlayer[0]["dices"], score, ok = hasNumber1(player["dices"].([]int), listPlayer[0]["dices"].([]int))
			if ok {
				player["count_dice"] = player["count_dice"].(int) - score
				listPlayer[0]["count_dice"] = listPlayer[0]["count_dice"].(int) + score
			}
		} else {
			player["dices"], listPlayer[i+1]["dices"], score, ok = hasNumber1(player["dices"].([]int), listPlayer[i+1]["dices"].([]int))
			if ok {
				player["count_dice"] = player["count_dice"].(int) - score
				listPlayer[i+1]["count_dice"] = listPlayer[i+1]["count_dice"].(int) + score
			}
		}
	}

	return listPlayer
}

func hasNumber1(player1Dices []int, player2Dices []int) ([]int, []int, int, bool) {
	status := false
	k := 0
	counter := 0
	for _, dice := range player1Dices {
		if dice == 1 {
			status = true
			counter++
		} else {
			player1Dices[k] = dice
			k++
		}
	}

	if status {
		for i := 0; i < counter; i++ {
			player2Dices = append(player2Dices, 1)
		}
	}

	return player1Dices, player2Dices, counter, status
}

func hasNumber6(dices []int) ([]int, int, bool) {
	status := false
	k := 0
	counter := 0
	for _, dice := range dices {
		if dice == 6 {
			status = true
			counter++
		} else {
			dices[k] = dice
			k++
		}
	}
	dices = dices[:k]

	return dices, counter, status
}
