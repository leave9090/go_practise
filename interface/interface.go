package _interface

//Makeer
type Peosener interface {
	Name() string
	Study()
}
type Student struct {
}

func (this *Student) Name() string {
	return "interface"
}

func (this *Student) Study() {

}

//Student

type ManStudent struct {
	Student
	Names string
}

func (this *ManStudent) Name() string {

	return "interface"
}
func (this *ManStudent) Study() {
	this.Student.Study()
}
