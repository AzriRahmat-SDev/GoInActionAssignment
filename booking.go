package main

type booking struct {
	bookingId  int
	customerId int
	doctorId   int
	date       string
}

var bookingList map[int]booking
var bookingId int

func init() {
	bookingList = make(map[int]booking)
	list := []booking{
		{1, 1, 1, "2022-06-06"},
		{2, 1, 1, "2022-11-07"},
		{3, 1, 1, "2022-08-31"},
		{3, 1, 1, "2022-10-01"},
		{1, 1, 1, "2022-05-10"},
		{2, 1, 1, "2022-07-23"},
	}

	for _, value := range list {
		newBooking(value)
	}

}

func bookingIsAvail(doctorId int, date string) bool {
	for _, value := range bookingList {
		if value.doctorId == doctorId && value.date == date {
			return false
		}
	}
	return true
}

func newBooking(value booking) int {
	bookingId++
	value.bookingId = bookingId
	bookingList[bookingId] = value

	return bookingId
}

func (b *booking) getDoctor() string {
	return getDoctorById(b.bookingId).Name
}

func deleteBookingFromBookingList(id int) error {
	for result, value := range bookingList {
		if value.doctorId == id {
			delete(bookingList, result)
		}
	}
	return nil
}
