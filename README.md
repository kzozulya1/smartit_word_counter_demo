Задача:
Реализовать приложение, которое будет выдавать список из 10-ти самых часто используемых слов, длина которых превышает заданную. Слова должны подсчитываться в множестве текстовых файлов в указанной папке. Приложение должно производить подсчёт в многопоточном режиме.

P.S: Ожидается реализация на Thread, для анализа слов использовать Regex.
При выполнении задачи не забывать о принципах SOLID.
При оценки выполненного задания будет учитываться эффективность работы приложения

Запуск аппликации:
1) Открыть файл cmd\smartit_word_counterd\main.go и изменить константы (настройки), при необходимости:

wordsLimit = 10  // длина списка  самых часто используемых слов 

minWordLenFilter = 5 	// критерий длины слова

dir = "/wcounter/files" // директория текстовых файлов


2) Запустить тест
go test ./...

3) Запустить аппликацию:
go run ./cmd/smartit_word_counterd

В результате должен быть вывод списка из 10-ти самых часто используемых слов, длина которых превышает 5 символов.
