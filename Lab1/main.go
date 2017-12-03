package main

import (
	"math"
)

const t0 = -1.0
const T = 2*math.Pi - 1
const y0 = 8.0
const eps = 0.000001
const tau0 = 0.2
const epsM = 0.0001

func main() {
	method := RungeKuttaMethod{}

	method.Start(t0, T, y0, eps, tau0, epsM)

	println("FINISHED!")
}

// u'(t) = f(t,u) = (e^t0)*(e^(1-cos(t+1))) + sin(t+1)*u
func f(t float64, u float64) float64 {
	return (1.0/math.E)*math.Pow(math.E, (1-math.Cos(t+1))) + math.Sin(t+1)*u
}

// Analytic solution of this differential equation found in Maple
func u(t float64) float64 {
	return math.Pow(math.E, -math.Cos(t+1))*t - math.E*(math.Pow(math.E, -math.Cos(t+1))*(-8.0-1.0/math.E))
}
