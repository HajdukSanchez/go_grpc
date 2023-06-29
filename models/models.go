package models

type Student struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  int32  `json:"age"`
}

type Test struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Question struct {
	Id       string `json:"id"`
	Question string `json:"question"`
	Answers  string `json:"answers"`
	TestId   string `json:"test_id"`
}

type Enrollment struct {
	StudentId string `json:"student_id"`
	TestId    string `json:"test_id"`
}
