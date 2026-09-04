package main

import (
	"fmt"
	"time"
)

func main() {
	// Текущая дата
	now := time.Now()
	
	// Дата Нового года (следующий год)
	var newYear time.Time
	if now.Month() == time.December && now.Day() > 31 {
		newYear = time.Date(now.Year()+2, time.January, 1, 0, 0, 0, 0, now.Location())
	} else if now.Month() == time.December && now.Day() == 31 {
		newYear = time.Date(now.Year()+1, time.January, 1, 0, 0, 0, 0, now.Location())
	} else {
		newYear = time.Date(now.Year()+1, time.January, 1, 0, 0, 0, 0, now.Location())
	}
	
	// Разница во времени
	diff := newYear.Sub(now)
	
	// Подсчет дней (целых дней)
	days := int(diff.Hours() / 24)
	
	// Оставшиеся часы, минуты, секунды
	hours := int(diff.Hours()) % 24
	minutes := int(diff.Minutes()) % 60
	seconds := int(diff.Seconds()) % 60
	
	// Вывод результата
	fmt.Printf("🎄 До Нового года осталось:\n")
	fmt.Printf("   %d дней, %d часов, %d минут, %d секунд\n", days, hours, minutes, seconds)
	fmt.Printf("   или %d полных дней\n", days)
	
	// Дополнительная информация
	if days == 0 {
		fmt.Println("🎅 С НОВЫМ ГОДОМ! 🎉")
	} else if days < 7 {
		fmt.Println("🎁 Осталась всего неделя! Готовьте подарки! 🎁")
	} else if days < 30 {
		fmt.Println("⛄ Создавайте новогоднее настроение! ⛄")
	} else {
		fmt.Println("❄️ Время планировать праздник! ❄️")
	}
}