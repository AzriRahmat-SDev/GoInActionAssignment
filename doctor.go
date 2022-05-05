package main

import "fmt"

type doctor struct {
	id   int
	name string
}

var doctorList []*doctor

func initDoctors() {
	doctorList = []*doctor{
		{1, "Dr Vickram"},
		{2, "Dr Idris"},
		{3, "Dr Lim"},
	}
}
func incrementDoctor() int {
	max := 0
	for _, doctor := range doctorList {
		if doctor.id > max {
			max = doctor.id
		}
	}
	return max + 1
}

func getDoctorById(id int) *doctor {
	for _, value := range doctorList {
		if value.id == id {
			return value
		}
	}
	return nil
}

func addDoctor(value *doctor) {
	value.id = incrementDoctor()
	doctorList = append(doctorList, value)
}

func deleteDoctor(id int) error {
	for i, value := range doctorList {
		if value.id == id {
			doctorList = append(doctorList[:i], doctorList[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("index of error not found")
}
