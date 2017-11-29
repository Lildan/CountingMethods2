package main

import "math"

const t0 = -1.0
const T = 2*math.Pi - 1
const u0 = 8
const eps = 0.000001
const tau0 = 0.5
const epsM = 0.000001

func main() {
	println("Hi Lildan!")
}

// u'(t) = f(t,u)
func f(t float64, u float64) float64 {
	return math.Pow(math.E, t0)*math.Pow(math.E, (1-math.Cos(t+1))) + math.Sin(t+1)*u
}

// Analytic solution of this differential equation found in Maple
func u(t float64) float64 {
	return 42.0
}
