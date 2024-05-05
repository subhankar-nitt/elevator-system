package addStudent

type Student interface {
	GetFname() string
	GetLname() string
	GetClass() string
	GetRoll() string
}
type StudentImpl struct {
	fname string
	lname string
	class string
	roll  string
}

func NewStudent(fname string, lname string, class string, roll string) Student {
	student := new(StudentImpl)
	student.fname = fname
	student.lname = lname
	student.class = class
	student.roll = roll

	return student
}

func (student *StudentImpl) GetFname() string {
	return student.fname
}
func (student *StudentImpl) GetLname() string {
	return student.lname
}
func (student *StudentImpl) GetClass() string {
	return student.class
}
func (student *StudentImpl) GetRoll() string {
	return student.roll
}
