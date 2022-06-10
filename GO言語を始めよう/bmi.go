package main

import (
	"fmt"
	"math"
)

//体重(kg)と身長(cm)の定数を定義
const WEIGHT = 60
const HEIGHT = 165

func main() {
	//BMIの計算
	var hm = HEIGHT / 100.0
	var bmi = WEIGHT / math.Pow(hm, 2)
	var bestW = math.Pow(hm, 2) * 22.0
	var per = WEIGHT / bestW * 100

	//結果表示
	fmt.Printf("BMI=%f 肥満度=%f\n", bmi, per)
	fmt.Printf("WEIGHT = %T\n", WEIGHT)
	fmt.Printf("HEIGHT = %T\n", HEIGHT) 
	fmt.Printf("hm = %T\n", hm)
	fmt.Printf("bmi = %T\n", bmi)
	fmt.Printf("bestW = %T\n", bestW)
	fmt.Printf("per = %T\n", per)

	a := 10/3
	fmt.Printf("%T\n", a)
	fmt.Printf("%d\n", a)
	b := 10/3.0
	fmt.Printf("%T\n", b)
	fmt.Printf("%f\n", b)
	var c float64 = 10/3
	fmt.Printf("%T\n", c)
	fmt.Printf("%f\n", c)
}