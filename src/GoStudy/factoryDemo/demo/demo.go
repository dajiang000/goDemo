package demo

type stu struct {
	Name  string
	score int
}

func NewStu(name string, score int) *stu {
	return &stu{name, score}
}

func (stu *stu) GetScore() int {
	return stu.score
}
