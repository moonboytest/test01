package search

import "module10/db"

//Функция поиска подозреваемого с самым большим кол-вом преступлений.
//Создаем  массив для подозреваемых, которые есть в нашей базе
//И переменную для подозреваемого с самым большим кол-вом преступлений
//Затем в цикле проходимся по списку.
//Если ни одного подозреваемого нет в нашей базе - возврааем строку
//и выходим из функции

func FindTheCriminal(d map[string]db.Man, s []string) string {
	//функция принимает на вход мапу из пакета db и список подозреваемых
	//и в случае успеха возвращает функцию GetInfo из пакета db
	//которая возвращает информацию о преступнике
	var criminals []string
	var mostCriminal string

	for _, v := range s {
		_, ok := db.Peoples[v]
		if ok {
			criminals = append(criminals, v)
		}
	}
	if len(criminals) == 0 {
		return "В базе данных нет информации по запрошенным подозреваемым"
	}
	counter := 0
	for _, v := range s {

		if db.Peoples[v].CountCrimes() > counter {
			counter = db.Peoples[v].CountCrimes()
			mostCriminal = v

		}
	}
	return db.Peoples[mostCriminal].Info()

}


