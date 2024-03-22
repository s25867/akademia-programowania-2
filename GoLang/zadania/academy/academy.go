package academy

type Student struct {
	Name       string
	Grades     []int
	Project    int
	Attendance []bool
}

// AverageGrade returns an average grade given a
// slice containing all grades received during a
// semester, rounded to the nearest integer.
func AverageGrade(grades []int) int {

	var sum float64
	var avg float64
	var result int

	for counter := 0; counter < len(grades); counter++ {
		sum = sum + float64(grades[counter])
	}

	if len(grades) == 0 {
		avg = sum
	} else {
		avg = sum / float64(len(grades))
	}

	a := int(avg*10) % 10
	if a >= 5 {
		result = int(avg) + 1
	} else {
		result = int(avg)
	}

	return result
}

// AttendancePercentage returns a percentage of class
// attendance, given a slice containing information
// whether a student was present (true) of absent (false).
//
// The percentage of attendance is represented as a
// floating-point number ranging from 0 to 1.
func AttendancePercentage(attendance []bool) float64 {

	var avg float64

	for counter := 0; counter < len(attendance); counter++ {
		if attendance[counter] {
			avg++
		}
	}
	result := avg / float64(len(attendance))
	return result
}

// FinalGrade returns a final grade achieved by a student,
// ranging from 1 to 5.
//
// The final grade is calculated as the average of a project grade
// and an average grade from the semester, with adjustments based
// on the student's attendance. The final grade is rounded
// to the nearest integer.

// If the student's attendance is below 80%, the final grade is
// decreased by 1. If the student's attendance is below 60%, average
// grade is 1 or project grade is 1, the final grade is 1.
func FinalGrade(s Student) int {

	var avg float64

	if AttendancePercentage(s.Attendance) < 0.6 || AverageGrade(s.Grades) == 1 || s.Project == 1 {
		return 1
	} else if AttendancePercentage(s.Attendance) < 0.8 {
		avg = (float64(AverageGrade(s.Grades)) + float64(s.Project)) / 2
		a := int(avg*10) % 10
		if a >= 5 {
			return int(avg)
		} else {
			return int(avg) - 1
		}
	} else {
		avg = (float64(AverageGrade(s.Grades)) + float64(s.Project)) / 2
		a := int(avg*10) % 10
		if a >= 5 {
			return int(avg) + 1
		} else {
			return int(avg)
		}

	}
}

// GradeStudents returns a map of final grades for a given slice of
// Student structs. The key is a student's name and the value is a
// final grade.
func GradeStudents(students []Student) map[string]uint8 {

	result := make(map[string]uint8)

	for counter := 0; counter < len(students); counter++ {
		result[students[counter].Name] = uint8(FinalGrade(students[counter]))
	}

	return result
}
