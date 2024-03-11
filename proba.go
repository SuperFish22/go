package main // исполняемый файл, а не библиотека

// импорт библиотек
import "fmt" // библиотека ввода и вывода данных

const pi = 3.14

-sfsf
sffsff
sfsfsfsfsfs

func main() {
	a := 19                    // автоматическое типилизация переменной
	var b string = "Пока дура" // ручное типилизацие переменой
	fmt.Println(b)
	fmt.Print(a)                 // вывод пез перевода стоки
	fmt.Println("Привет мир!!!") // вывод с переносом строки
	fmt.Printf("Мне: %d, а если в виде двойного числа то: %g", 19, 19.0)

}

func match() {
	fmt.Println("Я работаю ого вау")
	// Цикл со счётчиком
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
	// Цикл с условием
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i = +1
	}
}
