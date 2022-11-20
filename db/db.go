package db

import "fmt"

type Man struct {
	Name     string
	LastName string
	Age      int
	Gender   string
	Crimes   int
}

//Эти две функции, к сожалению, мне не удалось переместить в пакет Serch, хотя
//мне очень хотелось. Придется оставить их здесь  
func (m Man) CountCrimes() int {
	return m.Crimes
}

func (m Man) Info() string {
	return fmt.Sprintf("Name: %s\nLast Name: %s\nAge: %d\nGender: %s\nCrimes: %d\n",
		m.Name,
		m.LastName,
		m.Age,
		m.Gender,
		m.Crimes)
}

var Peoples = map[string]Man{
	"Roman":    {"Roman", "Segregeev", 23, "male", 10},
	"Sergey":   {"Sergey", "Vitaliev", 45, "male", 2},
	"Maxim":    {"Maxim", "Pevica", 38, "famale", 1},
	"Sasha":    {"Sasha", "Seriy", 25, "male", 5},
	"Mikhael":  {"Mikhael", "Palich-Terentiev", 12, "male", 1},
	"Simen":    {"Simen", "Upravdom", 78, "male", 0},
	"Bobr":     {"Bobr", "Dobr", 2, "animal", 17928},
	"Serebeda": {"Serebeda", "Jumashvili", 45, "male", 1},
}
