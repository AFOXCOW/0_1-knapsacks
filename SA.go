package zero_one_backpack

import (
	"math"
	"math/rand"
	"time"
)

func policy(choose []bool, indexx int, indexy int, v []int, w []int) (deltaV int, deltaW int, Case int) {
	if indexx == indexy {
		if choose[indexx] == true {
			deltaV = -v[indexx]
			deltaW = -w[indexx]
			Case = 4
			return
		} else {
			deltaV = v[indexx]
			deltaW = w[indexx]
			Case = 5
			return
		}
	} else if choose[indexx] == false && choose[indexy] == true {

		deltaV = v[indexx] - v[indexy]
		deltaW = w[indexx] - w[indexy]
		Case = 0
		return
	} else if choose[indexx] == true && choose[indexy] == false {

		deltaV = v[indexy] - v[indexx]
		deltaW = w[indexy] - w[indexx]
		Case = 1
		return
	} else if choose[indexx] == false && choose[indexy] == false {

		deltaV = v[indexx] + v[indexy]
		deltaW = w[indexx] + w[indexy]
		Case = 2
		return
	} else {

		deltaV = -v[indexy]
		deltaW = -w[indexy]
		Case = 3
		return
	}
}
func doitbyCase(choose []bool, Case int, x int, y int) {
	switch Case {
	case 0:
		choose[x] = true
		choose[y] = false
	case 1:
		choose[x] = false
		choose[y] = true
	case 2:
		choose[x] = true
		choose[y] = true
	case 3:
		choose[y] = false
	case 4:
		choose[x] = false
	case 5:
		choose[x] = true
	}
}

func Possibility(curr_weight int, deltaV int, deltaW int, max_w int, T float64) (po float64) {
	if curr_weight+deltaW > max_w {
		po = 0
	} else if deltaV > 0 {
		po = 1
	} else {
		po = math.Exp(float64(deltaV) / T)
	}
	return
}
func randX(n int) (x int) {
	rand.Seed(time.Now().UnixNano())
	x = rand.Intn(n)
	return
}
func SA(w []int, v []int, capa int, num int) (things []thing, best int) {
	choose := make([]bool, num)
	curr_value := 0
	curr_weight := 0
	T := 100.0     //初始温度
	t_min := 0.1   //终止温度
	ratio := 0.999 //温度下降率
	var (
		deltaV int
		deltaW int
		Case   int
	)
	for T > t_min {
		x := randX(num)
		y := randX(num)
		deltaV, deltaW, Case = policy(choose, x, y, v, w)
		po := Possibility(curr_weight, deltaV, deltaW, capa, T)
		test := rand.Float64()
		if test < po {
			doitbyCase(choose, Case, x, y)
			curr_value += deltaV
			curr_weight += deltaW
		}
		T = ratio * T
	}
	for i := 0; i < len(choose); i++ {
		if choose[i] == true {
			things = append(things, thing{i, w[i], v[i]})
		}
	}
	best = curr_value
	return
}
