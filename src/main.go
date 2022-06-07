package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/Reterer/card-memo/src/card"
)

var inviteMsg string = `Вас приветствует простой интерфейс использования этой прогой
	Команды отправляются в виде: <команда> <номер карточки>
	Список команд:
		chname - Изменить имя
		body   - Показать тело карточки
		chbody - Изменить тело карточки
		mkcard - Добавить карточку
		rmcard - Удалить карточку
`

func main() {
	fmt.Println(inviteMsg)
	for {
		cards := card.GetsCard()
		for i, card := range cards {
			fmt.Println(i, "|", card.Header, "|", card.CreationTime().Format(time.UnixDate))
		}
		cmd := readCommand()
		if cmd != nil {
			cmd.Exec(cards)
		}
	}
}

type executer interface {
	Exec(cards []card.Card)
}

type chnameCmd struct {
	id int
}

func (cmd *chnameCmd) Exec(cards []card.Card) {
	fmt.Print("Введите Новое имя карточки: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	cards[cmd.id].Header = text
	card.ReplaceCard(cmd.id, cards[cmd.id])
}

type bodyCmd struct {
	id int
}

func (cmd *bodyCmd) Exec(cards []card.Card) {
	fmt.Println("Тело карточки:")
	fmt.Println(cards[cmd.id].Body)
}

type chbodyCmd struct {
	id int
}

func (cmd *chbodyCmd) Exec(cards []card.Card) {
	fmt.Print("Введите Новое тело карточки: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	cards[cmd.id].Body = text
	card.ReplaceCard(cmd.id, cards[cmd.id])
}

type mkcardCmd struct {
}

func (cmd *mkcardCmd) Exec(cards []card.Card) {
	fmt.Print("Введите имя карточки: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	newCard := card.MakeCard(text, -1)
	card.AddCardIntoDB(newCard)
}

type rmcard struct {
	id int
}

func (cmd *rmcard) Exec(cards []card.Card) {
	card.RemoveCard(cmd.id)
}

func readCommand() executer {
	var command string
	var idCard int
	fmt.Scanf("%s", &command)

	switch {
	case command == "chname":
		fmt.Scanf("%d\n", &idCard)
		return &chnameCmd{idCard}
	case command == "body":
		fmt.Scanf("%d\n", &idCard)
		return &bodyCmd{idCard}
	case command == "chbody":
		fmt.Scanf("%d\n", &idCard)
		return &chbodyCmd{idCard}
	case command == "mkcard":
		fmt.Scanf("\n")
		return &mkcardCmd{}
	case command == "rmcard":
		fmt.Scanf("%d\n", &idCard)
		return &rmcard{idCard}
	}

	return nil
}
