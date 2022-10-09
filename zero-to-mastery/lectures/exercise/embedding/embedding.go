//--Summary:
//  Create a system monitoring dashboard using the existing dashboard
//  component structures. Each array element in the components represent
//  a 1-second sampling.
//
//--Requirements:

package main

import "fmt"

type Bytes int
type Celcius float32

type BandwidthUsage struct {
	amount []Bytes
}

func (m BandwidthUsage) AverageBandwidth() int {
	sum := 0
	for i := 0; i < len(m.amount); i++ {
		sum += int(m.amount[i])
	}
	return sum / len(m.amount)
}

type CpuTemp struct {
	temp []Celcius
}

func (c CpuTemp) AverageCpuTemp() float32 {
	var sum float32 = 0.0
	for i := 0; i < len(c.temp); i++ {
		sum += float32(c.temp[i])
	}
	return sum / float32(len(c.temp))
}

type MemoryUsage struct {
	amount []Bytes
}

func (m MemoryUsage) AverageMemoryUsage() int {
	sum := 0
	for i := 0; i < len(m.amount); i++ {
		sum += int(m.amount[i])
	}
	return sum / len(m.amount)
}

type Dashboard struct {
	BandwidthUsage
	CpuTemp
	MemoryUsage
}

func (d Dashboard) printAverages() {
	fmt.Println(d.AverageBandwidth())
	fmt.Println(d.AverageCpuTemp())
	fmt.Println(d.AverageMemoryUsage())
}

//* Create functions to calculate averages for each dashboard component
//* Using struct embedding, create a Dashboard structure that contains
//  each dashboard component

//* Print out a 5-second average from each component using promoted
//  methods on the Dashboard

func main() {
	bandwidth := BandwidthUsage{[]Bytes{50000, 100000, 130000, 80000, 90000}}
	temp := CpuTemp{[]Celcius{50, 51, 53, 51, 52}}
	memory := MemoryUsage{[]Bytes{800000, 800000, 810000, 820000, 800000}}

	dashboard := Dashboard{
		bandwidth,
		temp,
		memory,
	}

	dashboard.printAverages()
}
