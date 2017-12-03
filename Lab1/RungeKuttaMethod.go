package main

import (
	"math"
	"strconv"
)

// RungeKuttaMethod is the method that is imlemented in this lab work
type RungeKuttaMethod struct {
	t          float64
	T          float64
	y          float64
	tau        float64
	eMax       float64
	v          float64
	t1         float64
	w          float64
	k1, k2, k3 float64
	E          float64
	tauH       float64
	epsM       float64
}

func (m RungeKuttaMethod) Start(t0 float64, T float64, y0 float64, eps float64, tau0 float64, epsM float64) {
	m.Step1(t0, T, y0, eps, tau0, epsM)
	return
}

func (m RungeKuttaMethod) Step1(t0 float64, T float64, y0 float64, eps float64, tau0 float64, epsM float64) {
	m.t = t0
	m.T = T
	m.y = y0
	m.tau = tau0
	m.eMax = 0
	m.epsM = epsM

	m.Step2()
	return
}

func (m RungeKuttaMethod) Step2() {
	println("t = " + strconv.FormatFloat(m.t, 'f', 8, 64))
	println("y = " + strconv.FormatFloat(m.y, 'f', 8, 64))
	println("u(t) = " + strconv.FormatFloat(u(m.t), 'f', 8, 64))
	println("|y - u(t)| = " + strconv.FormatFloat(math.Abs(m.y-u(m.t)), 'f', 8, 64))

	m.Step3()
	return
}

func (m RungeKuttaMethod) Step3() {
	if math.Abs(m.T-m.t) < m.epsM || m.t > m.T {
		m.Step10()
		return
	} else {
		m.Step4()
	}
	return
}

func (m RungeKuttaMethod) Step4() {
	if m.t+m.tau > m.T {
		m.tau = m.T - m.t
	}

	m.Step5()
	return
}

func (m RungeKuttaMethod) Step5() {
	m.v = m.y
	m.t1 = m.t

	m.Step6()
	return
}

func (m RungeKuttaMethod) Step6() {
	m.k1 = f(m.t, m.y)

	m.Step7()
	return
}

func (m RungeKuttaMethod) Step7() {
	m.k2 = f(m.t+1*m.tau, m.y+m.tau*1*m.k1)
	m.k3 = f(m.t+0.5*m.tau, m.y+m.tau*0.25*m.k1+m.tau*0.25*m.k2)

	m.w = m.y + m.tau*(0.5*m.k1+0.5*m.k2+0*m.k3)
	m.y = m.y + m.tau*(1.0/6.0*m.k1+1.0/6.0*m.k2+4.0/6.0*m.k3)

	m.Step8()
	return
}

func (m RungeKuttaMethod) Step8() {
	m.E = math.Abs(m.y-m.w) / math.Max(1.0, math.Abs(m.y))

	m.tauH = m.tau * math.Min(5.0, math.Max(0.1, 0.9*math.Pow(eps/m.E, 1.0/(1.0+8.0))))

	m.Step9()
	return
}

func (m RungeKuttaMethod) Step9() {
	if m.E <= eps {

		if m.tau > 0 {
			m.t = m.t + m.tau
		} else {
			m.Step10()
			return
		}

		un := u(m.t)

		m.tau = m.tauH

		println("-----------------------------------")
		println("t = " + strconv.FormatFloat(m.t, 'f', 20, 64))
		println("y = " + strconv.FormatFloat(m.y, 'f', 20, 64))
		println("u(t) = " + strconv.FormatFloat(un, 'f', 20, 64))
		println("|y - u(t)| = " + strconv.FormatFloat(math.Abs(m.y-un), 'f', 20, 64))

		if m.eMax < math.Abs(m.y-un) {
			m.eMax = math.Abs(m.y - un)
		}

		m.Step3()
		return
	}

	m.y = m.v
	m.t = m.t1
	m.tau = m.tauH

	m.Step7()
	return
}

func (m RungeKuttaMethod) Step10() {
	println("Maximal error = " + strconv.FormatFloat(m.eMax, 'f', 8, 64))
	return
}
