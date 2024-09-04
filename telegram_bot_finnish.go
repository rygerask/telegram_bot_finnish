package main

import (
	"math/rand"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const TOKEN = //token

var bot *tgbotapi.BotAPI
var chatId int64

var finnishTeacher = [4]string{"1", "2", "3", "4"}

var answersOne = []string{
	"päärynä - груша",
	"meloni - дыня",
	"vesi - арбуз",
	"sitruuna - лимон",
	"luumi - слива",
	"viinirypäle - виноград",
	"omena - яблоко",
	"kaali - капуста",
	"kesäkurpitsa - кабачок",
	"kurpitsa - тыква",
	"kukkakaali - цветная капуста",
	"porkkana - морковь",
	"valkosipuli - чеснок",
	"munakoiso - баклажан",
	"sieni - гриб",
	"peruna - картофель",
	"papu - бобы",
	"maissi - кукуруза",
	"kurkku - огурец",
	"herme - горох",
	"sipuli - лук",
	"mustikka - черника",
	"puolukka - брусника",
	"kursikka - вишня/черешня",
	"viinimarja/kerokka - смородина",
	"mansikka - клубника",
	"lakka - морошка",
	"vadelma - малина",
	"porsaanliha - свинина",
	"naudanliha - говядина",
	"kana - курица",
	"kala - рыба",
	"lohi - лосось",
	"nakki - сосиска",
	"makkara - колбаса",
	"katkarapu - креветка",
	"riisi - рис",
	"keitto - суп",
	"hampurilainen - гамбургер",
	"makaronilaatikko - запеканка",
	"pihvi - котлета",
	"kastike - соус",
	"jäätelö - мороженое",
	"kakku - торт",
	"suklaa - шоколад",
	"keksi - печенье",
	"piirakka - пирог",
	"karkki - карамелька",
	"pulla - плюшка",
	"korvapuusti - улитка с корицей",
	"pippuri - перец",
	"soija - соевый соус",
	"suola - соль",
	"öljy - оливковое масло",
	"etikka - уксус",
	"ketsuppi - кетчуп",
	"sinappu - горчица",
	"limu - лимонад",
	"maito - молоко",
	"siideri - сидр",
	"olut - пиво",
	"kivennäisvesi - минералка",
	"piimä - кисломолочное",
	"mehu - сок",
	"viini - вино",
	"kauramaito - овсяное молоко",
	"murot - хлопья",
	"puuro - каша",
	"leipä - хлеб",
	"karjalanpiirakka - карельские калитки",
	"voi - масло",
	"sämpylä - булочка для сендвича",
	"hunaja - мёд",
	"juusto - сыр",
	"muna - яйцо",
	"hillo - джем",
}

var answersTwo = []string{
	"tunturi - сопка",
	"ranta - берег",
	"mäki - холм",
	"saari - остров",
	"järvi - озеро",
	"meri - море",
	"salvi - пролив",
	"rannikko - побережье",
	"metsä - лес",
	"kivi - камень",
	"kallio - скала",
	"lahvi - залив",
	"vuori - гора",
	"joki - река",
	"niemi - мыс",
	"lohi - лосось",
	"hauki - щука",
	"ahven - окунь",
	"silakka - салака",
	"hirvi - лось",
	"karhu - медведь",
	"susi - волк",
	"kettu - лиса",
	"jänis - заяц",
	"orava - белка",
	"siili - ёж",
	"hiiri - мышь",
	"kääme - змея",
	"samakko - лягушка",
	"varis - ворона",
	"harakka - сорока",
	"lokki - чайка",
	"pulu - голубь",
	"pääskynen - ласточка",
	"sorsa - утка",
	"pöllö - сова",
	"joutsen - лебедь",
	"kärpänen - муха",
	"hyttynen - комар",
	"muurahainen - муравей",
	"perhonen - бабочка",
	"hämähäkki - паук",
	"mehiläinen - пчела",
	"kuusi - ель",
	"koivu - береза",
	"pihlaja - рябина",
	"mänty - сосна",
	"vaahtera - клён",
	"tammi - дуб",
	"lehmus - липа",
	"heinä - сено",
	"pensas - куст",
	"jäkälä - ягерь",
	"sammal - мох",
	"ruoho - трава",
}

var answersThree = []string{
	"animal_one",
	"animal_two",
	"animal_three",
	"animal_four",
	"animal_five",
}

var answersFour = []string{
	"insect_one",
	"insect_two",
	"insect_three",
	"insect_four",
	"insect_five",
}

func connectWithTelegram() {
	var err error
	if bot, err = tgbotapi.NewBotAPI(TOKEN); err != nil {
		panic("Cannot connect to Telegram")
	}
}

func sendMessage(msg string) {
	msgConfig := tgbotapi.NewMessage(chatId, msg)
	bot.Send(msgConfig)
}

func isMessageForFinnishTeacher(update *tgbotapi.Update) bool {
	if update.Message == nil || update.Message.Text == "" {
		return false
	}

	msgInLowerCase := strings.ToLower(update.Message.Text)
	for _, name := range finnishTeacher {
		if strings.Contains(msgInLowerCase, name) {
			return true
		}
	}
	return false
}

func sendMessageToBot(update *tgbotapi.Update) int {
	msgInLowerCase := strings.ToLower(update.Message.Text)
	for i, name := range finnishTeacher {
		if strings.Contains(msgInLowerCase, name) {
			return i
		}
	}
	return 100
}

func randomPart(name []string) string {
	index_one := rand.Intn(len(name))
	index_two := rand.Intn(len(name))
	index_three := rand.Intn(len(name))
	for i := 0; i < len(name); i++ {
		for j := 0; j < len(name); j++ {
			if index_two == index_one {
				index_two = rand.Intn(len(name))
			} else {
				j = len(name)
			}
		}

		if index_three == index_one || index_three == index_two {
			index_three = rand.Intn(len(name))
		} else {
			i = len(name)
		}
	}
	return string(name[index_one] + "\n\n" + name[index_two] + 
	"\n\n" + name[index_three])
}

func getFinnishTeacherAnswer(i int) string {
	switch finnishTeacher[i] {
	case "1":
		m := randomPart(answersOne)
		return m
	case "2":
		m := randomPart(answersTwo)
		return m
	case "3":
		m := randomPart(answersThree)
		return m
	case "4":
		m := randomPart(answersFour)
		return m
	default:
		return "ошибка"
	}
}

func sendAnswer(update *tgbotapi.Update, i int) {
	msg := tgbotapi.NewMessage(chatId, getFinnishTeacherAnswer(i))
	msg.ReplyToMessageID = update.Message.MessageID
	bot.Send(msg)
}

func main() {
	connectWithTelegram()

	updateConfig := tgbotapi.NewUpdate(0)
	for update := range bot.GetUpdatesChan(updateConfig) {
		if update.Message != nil && update.Message.Text == "/start" {
			chatId = update.Message.Chat.ID
			sendMessage("Какой топик ты хочешь повторить?\n\n1 - еда и напитки\n" +
				"2 - животные, природа\n3 - todo\n4 - todo")
		}

		if isMessageForFinnishTeacher(&update) {
			i := sendMessageToBot(&update)
			if i != 100 {
				sendAnswer(&update, i)
			}
		}
	}
}
