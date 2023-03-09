package payroll

type Developer struct {
	Individual        Employee
	HourlyRate        float64
	HoursWorkedInYear float64
	Review            map[string]interface{}
}

func (d Developer) Pay() (string, float64) {
	return d.Individual.FullName(), d.Calculate()
}
func (d Developer) Calculate() float64 {
	return d.HourlyRate * d.HoursWorkedInYear
}
func (d Developer) ReviewRating(reviewStandard map[string]int) float64 {
	var totalRating float64 = 0.0
	r := d.Review
	for _, v := range r {
		switch v.(type) {
		case int:
			intValue := v.(int)
			totalRating = totalRating + float64(intValue)
		case int32:
			intValue := v.(int32)
			totalRating = totalRating + float64(intValue)
		case int64:
			intValue := v.(int64)
			totalRating = totalRating + float64(intValue)
		case float32:
			intValue := v.(float32)
			totalRating = totalRating + float64(intValue)
		case float64:
			intValue := v.(float64)
			totalRating = totalRating + (intValue)
		case string:
			//fmt.Println(v)
			temp := reviewStandard[v.(string)]
			totalRating = totalRating + float64(temp)
		default:
			totalRating += 0
		}
	}
	return totalRating
}
