package helpers

import (
	"reflect"
	"strings"
	"time"
)

func DateFormat(d string) (f string) {
	s := strings.Split(d, "/")
	f = s[2] + "-" + s[1] + "-" + s[0]
	return
}

func DateFormatYear(d string) (f string) {
	s := strings.Split(d, "/")
	f = "20" + s[2] + "-" + s[1] + "-" + s[0]
	return
}

func MonthsCount(startDate time.Time, endDate time.Time) int {
	months := 0
	month := startDate.Month()
	for startDate.Before(endDate) {
		startDate = startDate.Add(time.Hour * 24)
		nextMonth := startDate.Month()
		if nextMonth != month {
			months++
		}
		month = nextMonth
	}
	return months + 1
}

func LongName(n string) (l string) {
	if n == "brad" {
		n = "bradley"
	}
	if n == "nat" {
		n = "natalie"
	}
	if n == "matt" {
		n = "matthew"
	}
	return n
}

func ChangeName(stylist string) (n string) {
	if stylist == "Michelle Stephenson" {
		n = "Michelle Railton"
	} else if stylist == "Jo Mahoney" {
		n = "Jo Birchall"
	} else if stylist == "Laura Crumplin" {
		n = "Laura Hall"
	} else if stylist == "Bradley Ryan" {
		n = "Brad Ryan"
	} else {
		n = stylist
	}
	return n
}

func GetBankAcc(salon string) (acc string) {
	switch salon {
	case "jakata":
		acc = "06517160"
	case "pk":
		acc = "02017546"
	case "base":
		acc = "17623364"
	}
	return acc
}

//func AddLinearRegressionPoints(data interface{}, fieldName string) {
//	v := reflect.ValueOf(data)
//	n := float64(v.Len())
//
//	var sumX, sumY, sumXY, sumX2 float64
//
//	// Calculate sums
//	for i := 0; i < v.Len(); i++ {
//		y := v.Index(i).FieldByName(fieldName).Float()
//		x := float64(i + 1)
//		sumX += x
//		sumY += y
//		sumXY += x * y
//		sumX2 += x * x
//	}
//
//	// Calculate slope (b1) and y-intercept (b0)
//	b1 := (n*sumXY - sumX*sumY) / (n*sumX2 - sumX*sumX)
//	b0 := (sumY - b1*sumX) / n
//
//	// Calculate corresponding points for the regression line
//	for i := 0; i < v.Len(); i++ {
//		x := float64(i + 1)
//		linearTotal := b0 + b1*x
//		v.Index(i).FieldByName("LinearTotal").SetFloat(linearTotal)
//	}
//}
func AddLinearRegressionPoints(data interface{}, fieldNames []string) {
	for _, fieldName := range fieldNames {
		v := reflect.ValueOf(data)
		n := float64(v.Len())

		var sumX, sumY, sumXY, sumX2 float64

		for i := 0; i < v.Len(); i++ {
			y := v.Index(i).FieldByName(fieldName).Float()
			x := float64(i + 1)
			sumX += x
			sumY += y
			sumXY += x * y
			sumX2 += x * x
		}

		b1 := (n*sumXY - sumX*sumY) / (n*sumX2 - sumX*sumX)
		b0 := (sumY - b1*sumX) / n

		linearFieldName := "Linear" + fieldName
		for i := 0; i < v.Len(); i++ {
			x := float64(i + 1)
			linearTotal := b0 + b1*x
			v.Index(i).FieldByName(linearFieldName).SetFloat(linearTotal)
		}
	}
}
