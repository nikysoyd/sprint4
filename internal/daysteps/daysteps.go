package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-4-sprint-final/internal/spentcalories"
)

var (
	StepLength = 0.65 // длина шага в метрах
)

func parsePackage(data string) (steps int, duration time.Duration, err error) {
	// ваш код ниже
	parts := strings.Split(data, " ")

	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("Мало данных")
	}
	steps, err = strconv.Atoi(parts[0])

	if steps < 0 {
		return 0, 0, fmt.Errorf("Ошибка в шагах")
	}

	if err != nil {
		return 0, 0, fmt.Errorf("ошибка преобразования шагов: %v", err)
	}

	duration, err = time.ParseDuration(parts[1])

	if duration < 0 {
		return 0, 0, fmt.Errorf("Ошибка в дурации")
	}

	if err != nil {
		return 0, 0, fmt.Errorf("ошибка преобразования дурации: %v", err)
	}

	return steps, duration, nil

}

// DayActionInfo обрабатывает входящий пакет, который передаётся в
// виде строки в параметре data. Параметр storage содержит пакеты за текущий день.
// Если время пакета относится к новым суткам, storage предварительно
// очищается.
// Если пакет валидный, он добавляется в слайс storage, который возвращает
// функция. Если пакет невалидный, storage возвращается без изменений.
func DayActionInfo(data string, weight, height float64) string {
	// ваш код ниже
	steps, duration, err := parsePackage(data)

	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		return ""
	}

	if steps < 0 {
		fmt.Errorf("Ошибка: %w\n", err)
		return ""
	}

	if duration < 0 {
		fmt.Errorf("Ошибка: %w\n", err)
		return ""
	}

	distance := (float64(steps) * StepLength) / 1000

	WalkingSpentCalories := spentcalories.WalkingSpentCalories(steps, weight, height, duration)

	counterSteps := fmt.Sprintf("Количество шагов: %d.", steps)
	counterDist := fmt.Sprintf("Дистанция составила %.2f км.", distance)
	counterCal := fmt.Sprintf("Вы сожгли %.2f ккал.", WalkingSpentCalories)
	message := counterSteps + counterDist + counterCal
	return message
}
