package Avito

import (
	"fmt"
	"reflect"
)

// Условие задачи
// Мы в Авито любим проводить соревнования, — недавно мы устроили
// чемпионат по шагам. И вот настало время подводить итоги!

// Необходимо определить userIds участников, которые прошли
// наибольшее количество шагов steps за все дни, не пропустив
// ни одного дня соревнований.

// Пример
// # Пример 1
// # ВВОД
// statistics = [
//     [{ userId: 1, steps: 1000 }, { userId: 2, steps: 1500 }],
//     [{ userId: 2, steps: 1000 }]
// ]

// # ВЫВОД
// champions = { userIds: [2], steps: 2500 }

// # Пример 2
// statistics = [
//     [{ userId: 1, steps: 2000 }, { userId: 2, steps: 1500 }],
//     [{ userId: 2, steps: 4000 }, { userId: 1, steps: 3500 }]
// ]

// # ВЫВОД
// champions = { userIds: [1, 2], steps: 5500 }

type Statistic struct {
	UserID int
	Steps  int
}

type Result struct {
	UserIDs []int
	Steps   int
}

func getChampions(statistics [][]Statistic) Result {

}

func main() {
	fmt.Println("example 1", reflect.DeepEqual(
		getChampions(
			[][]Statistic{
				{{UserID: 1, Steps: 2000}, {UserID: 2, Steps: 1500}},
				{{UserID: 2, Steps: 1000}},
			},
		),
		Result{
			UserIDs: []int{2},
			Steps:   2500,
		},
	))

	fmt.Println("example 2", reflect.DeepEqual(
		getChampions(
			[][]Statistic{
				{{UserID: 1, Steps: 2000}, {UserID: 2, Steps: 1500}},
				{{UserID: 2, Steps: 4000}, {UserID: 1, Steps: 3500}},
			},
		),
		Result{
			UserIDs: []int{1, 2},
			Steps:   5500,
		},
	))
}

// Решение
//func getChampions(statistics [][]Statistic) Result {
//	totalDays := len(statistics)
//	resSteps := make(map[int]int)
//	participation := make(map[int]int)
//
//	for _, day := range statistics {
//		dayBeen := make(map[int]bool)
//		for _, static := range day {
//			resSteps[static.UserID] += static.Steps
//			if !dayBeen[static.UserID] {
//				participation[static.UserID]++
//				dayBeen[static.UserID] = true
//			}
//		}
//	}
//
//	maxStep := 0
//	champions := []int{}
//
//	for user, day := range participation {
//		if totalDays == day {
//			if resSteps[user] > maxStep {
//				maxStep = resSteps[user]
//				champions = []int{user}
//			} else if resSteps[user] == maxStep {
//				champions = append(champions, user)
//			}
//		}
//	}
//
//	return Result{UserIDs: champions, Steps: maxStep}
//}
