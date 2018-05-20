package main

import (
	"bufio"
	"fmt"
	"os"
	//"time"
)

//	***************
//	Блок переменных
//	***************
var currentPc PcInfoStruct

//	************
//	Блок функций
//	************

//	Свич действий
func doFunction(doThis string) {

	//	свич
	switch doThis {
	//	Информация о компьютере
	case "pc":
		//	дернули pc_info как рутину
		go pcInfo()
		// прервали
		break

	//	Ждем сообщений
	case "listener":
		//	дернули pc_info как рутину
		go listener(LISTENER_PORT)
		// прервали
		break

	//	Ждем сообщений
	case "listener2":
		//	дернули pc_info как рутину
		go listener(LISTENER_PORT2)
		// прервали
		break

	//	Ждем сообщений
	case "sender":
		// приглашение
		fmt.Println("Введите текст")
		// читаем текст
		var message string
		fmt.Scanln(&message)
		// 	шифруем
		ciphertext := encrypt(message)
		//	дернули pc_info как рутину
		go sender(LISTENER_PORT2, ciphertext)
		// прервали
		break

	//	Ждем сообщений
	case "sender2":
		// 	приглашение
		fmt.Println("Введите текст")
		// 	читаем текст
		var message string
		fmt.Scanln(&message)
		// 	шифруем
		ciphertext := encrypt(message)
		//	дернули pc_info как рутину
		go sender(LISTENER_PORT, ciphertext)
		// 	прервали
		break

	//	По умолчанию просто инфо выводим
	default:
		//	Информация ни о чем
		fmt.Println("Данный функционал не реализован")

	}
}

//	Инитим основной функционал
func initial() {
	//
}

//	Точка входа
func main() {
	//	Инитим
	initial()

	// Включаем листенер
	// listener(LISTENER_PORT)

	//	Передача управления на интерфейс управления
	userInterface()
}

//	Информация о компьютере
func pcInfo() {
	//	Вывод информации о текущей машине
	fmt.Println("Информация о системе")
	//timer := time.NewTimer(time.Second * 2)
	//<- timer.C
	//	отработка uname
	commandUname(&currentPc)
	//println("123")
	//	отработка lshw
	commandIp(&currentPc)

	fmt.Println(currentPc)
}

//	управление пользовательским интерфейсом
func userInterface() {

	//	Приглашение
	fmt.Print("Интерфейс управления: ")
	//
	for {
		//	Перехват Стдин
		in := bufio.NewReader(os.Stdin)
		doThis, _ := in.ReadString('\n')
		//	вызвали коллектор
		doFunction(doThis[:len(doThis)-1])

	}
	fmt.Scan()
}
