package main

const t0 = 0.0 
const T = 1
const eps = 0.0001
const tau0 = 0.5
const epsM = 0.0001




func y1d (t float64, y1 float64, y2 float64) float64 {
	 return t + y1*y2
 }

func y2d (t float64, y1 float64, y2 float64) float64 {
	return t - y1*y1
}

func main() {

}
