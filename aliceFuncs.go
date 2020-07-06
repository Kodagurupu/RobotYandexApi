package main

import "github.com/azzzak/alice"

func helpFunction(response alice.Response) *alice.Response {
	return response.Text("Для управления промороботом используйте следующие команды: Вперед, Назад, В лево, В право")
}
