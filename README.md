
# l2 tasks repository

## 1. Базовая задача
Создать программу печатающую точное время с использованием
NTP -библиотеки. Инициализировать как go module. Использовать
библиотеку `github.com/beevik/ntp.` Написать программу
печатающую текущее время / точное время с использованием этой
библиотеки.
Требования:
1. Программа должна быть оформлена как go module
2. Программа должна корректно обрабатывать ошибки
библиотеки: выводить их в STDERR и возвращать ненулевой
код выхода в OS
## 2. Задача на распаковку
Создать Go-функцию, осуществляющую примитивную распаковку
строки, содержащую повторяющиеся символы/руны, например:<br>
● "a4bc2d5e" => "aaaabccddddde"<br>
● "abcd" => "abcd"<br>
● "45" => "" (некорректная строка)<br>
● "" => ""<br>
Дополнительно
Реализовать поддержку escape-последовательностей.

## 3. Утилита sort
Отсортировать строки в файле по аналогии с консольной
утилитой sort <br> Реализовать поддержку утилитой следующих ключей:<br>
-k — указание колонки для сортировки (слова в строке могут
выступать в качестве колонок, по умолчанию разделитель —
пробел)<br>
-n — сортировать по числовому значению<br>
-r — сортировать в обратном порядке<br>
-u — не выводить повторяющиеся строки<br>

## 4. Поиск анаграмм по словарю
Написать функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,<br>
'листок', 'слиток' и 'столик' - другому.<br>

## 5. Утилита grep
Реализовать утилиту фильтрации по аналогии с консольной
утилитой (man grep — смотрим описание и основные параметры).
Реализовать поддержку утилитой следующих ключей:<br>
-A - "after" печатать +N строк после совпадения<br>
-B - "before" печатать +N строк до совпадения<br>
-C - "context" (A+B) печатать ±N строк вокруг совпадения<br>
-c - "count" (количество строк)<br>
-i - "ignore-case" (игнорировать регистр)<br>
-v - "invert" (вместо совпадения, исключать)<br>
-F - "fixed", точное совпадение со строкой, не паттерн<br>
-n - "line num", напечатать номер строки<br>
## 6. Утилита cut
Реализовать утилиту аналог консольной команды cut (man cut).
Утилита должна принимать строки через STDIN, разбивать по
разделителю (TAB) на колонки и выводить запрошенные.
Реализовать поддержку утилитой следующих ключей:<br>
-f - "fields" - выбрать поля (колонки)<br>
-d - "delimiter" - использовать другой разделитель<br>
-s - "separated" - только строки с разделителем<br>
## 7. Or channel
Реализовать функцию, которая будет объединять один или более
done-каналов в single-канал, если один из его составляющих каналов
закроется.
Очевидным вариантом решения могло бы стать выражение при
использованием select, которое бы реализовывало эту связь, однако
иногда неизвестно общее число done-каналов, с которыми вы
работаете в рантайме. В этом случае удобнее использовать вызов
единственной функции, которая, приняв на вход один или более
or-каналов, реализовывала бы весь функционал.
### Определение функции:
    var or func(channels ...<- chan interface{}) <- chan interface{}
### Пример использования функции:
    sig := func(after time.Duration) <- chan interface{} {
      c := make(chan interface{})
      go func() {
        defer close(c)
        time.Sleep(after)
      }()
      return c
    } 
    start := time.Now()
    <-or (
          sig(2*time.Hour),
          sig(5*time.Minute),
          sig(1*time.Second),
          sig(1*time.Hour),
          sig(1*time.Minute),
    ) 
    fmt.Printf(“fone after %v”, time.Since(start))
## 8. Взаимодействие с ОС
Необходимо реализовать свой собственный UNIX-шелл-утилиту с
поддержкой ряда простейших команд:
## 9. Утилита wget
Реализовать утилиту wget с возможностью скачивать сайты
целиком.
## 10. Утилита telnet
Реализовать простейший telnet-клиент.
Примеры вызовов:<br>
go-telnet --timeout=10s host port go-telnet mysite.ru 8080<br>
go-telnet --timeout=3s 1.1.1.1 123<br>

## 11. HTTP-сервер
Реализовать HTTP-сервер для работы с календарем. В рамках
задания необходимо работать строго со стандартной
HTTP-библиотекой.
Методы API:<br>
● POST /create_event<br>
● POST /update_event<br>
● POST /delete_event<br>
● GET /events_for_day<br>
● GET /events_for_week<br>
● GET /events_for_month<br>
Параметры передаются в виде www-url-form-encoded (т.е.
обычные user_id=3&date=2019-09-09). В GET методах параметры
передаются через queryString, в POST через тело запроса.

# Чтение и понимание кода
## 1. Что выведет программа? Объяснить вывод программы.package main
    import "fmt"
    func main() {
      a := [5]int{76, 77, 78, 79, 80}
      var b []int = a[1:4]
      fmt.Println(b)
    }
## 2. Что выведет программа? Объяснить вывод программы. Объяснить как работают defer’ы и порядок их вызовов.
    package main
    import "fmt"
    func test() (x int) {
      defer func() {
        x++
      }()
      x = 1
      return
      }
    func anotherTest() int {
      var x int
      defer func() {
        x++
      }()
      x = 1
      return x
    }
    func main() {
      fmt.Println(test())fmt.Println(anotherTest())
    }
## 3. Что выведет программа? Объяснить вывод программы. Объяснить внутреннее устройство интерфейсов и их отличие от пустых интерфейсов.
    package main
    import (
      "fmt"
      "os"
    )
    func Foo() error {
      var err *os.PathError = nil
      return err
    }
    func main() {
      err := Foo()
      fmt.Println(err)
      fmt.Println(err == nil)
    }
## 4. Что выведет программа? Объяснить вывод программы.
    package main
    func main() {
      ch := make(chan int)
      go func() {
        for i := 0; i < 10; i++ {
          ch <- i
        }
      }()
      for n := range ch {
      println(n)
      }
    }
## 5. Что выведет программа? Объяснить вывод программы.
    package main
    type customError struct {
      msg string
    }
    func (e *customError) Error() string {
      return e.msg
    }
    func test() *customError {
      {
        // do something
      }
      return nil
    }
    func main() {
      var err error
      err = test()
      if err != nil {
        println("error")
        return
      }
      println("ok")
    }
## 6. Что выведет программа? Объяснить вывод программы. Рассказать про внутреннее устройство слайсов и что происходит при передаче их в качестве аргументов функции.
    package main
    import "fmt"
    func main() {
      var s = []string{"1", "2", "3"}
      modifySlice(s)
      fmt.Println(s)
    }
    func modifySlice(i []string) {i[0] = "3"
      i = append(i, "4")
      i[1] = "5"
      i = append(i, "6")
    }
## 7. Что выведет программа? Объяснить вывод программы.
    package main
    import (
      "fmt"
      "math/rand"
      "time"
    )
    func asChan(vs ...int) <-chan int {
      c := make(chan int)
      go func() {
        for _, v := range vs {
          c <- v
          time.Sleep(time.Duration(rand.Intn(1000)) *
          time.Millisecond)
        }
        close(c)
      }()
      return c
    }
    func merge(a, b <-chan int) <-chan int {
      c := make(chan int)
      go func() {
        for {
          select {
            case v := <-a:
            c <- v
            case v := <-b:
            c <- v
          }
        }
      }()
      return c
    }
    func main() {
      a := asChan(1, 3, 5, 7)
      b := asChan(2, 4 ,6, 8)
      c := merge(a, b )
      for v := range c {
        fmt.Println(v)
      }
    }
