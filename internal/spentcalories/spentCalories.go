package spentcalories

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep = 0.65 // средняя длина шага.
	mInKm   = 1000 // количество метров в километре.
	minInH  = 60   // количество минут в часе.
)

func parseTraining(data string) (steps int, trainingName string, duration time.Duration, err error) {
	// ваш код ниже
	infoTraining := strings.Split(data, ",")

	if len(infoTraining) != 3 {

		return 0, "", 0, fmt.Errorf("Error")
	}

	steps, err = strconv.Atoi(infoTraining[0])

	if steps < 0 {
		return 0, "", 0, fmt.Errorf("Ошибка в шагах")
	}

	if err != nil {
		return 0, "", 0, fmt.Errorf("ошибка преобразования шагов: %v", err)
	}

	trainingName = infoTraining[1]

	duration, err = time.ParseDuration(infoTraining[2])

	if duration < 0 {
		return 0, "", 0, fmt.Errorf("Ошибка в дурации")
	}

	if err != nil {
		return 0, "", 0, fmt.Errorf("ошибка преобразования дурации: %v", err)
	}

	return steps, trainingName, duration, nil
}

// distance возвращает дистанцию(в километрах), которую преодолел пользователь за время тренировки.
//
// Параметры:
//
// steps int — количество совершенных действий (число шагов при ходьбе и беге).
func distance(steps int) float64 {
	// ваш код ниже
	distance := float64(steps) * lenStep / mInKm

	return distance

}

// meanSpeed возвращает значение средней скорости движения во время тренировки.
//
// Параметры:
//
// steps int — количество совершенных действий(число шагов при ходьбе и беге).
// duration time.Duration — длительность тренировки.
func meanSpeed(steps int, duration time.Duration) float64 {
	// ваш код ниже
	if duration < 0 {
		return 0
	}

	distance := distance(steps)

	meanSpeed := distance / duration.Hours()

	return meanSpeed
}

// ShowTrainingInfo возвращает строку с информацией о тренировке.
//
// Параметры:
//
// data string - строка с данными.
// weight, height float64 — вес и рост пользователя.
func TrainingInfo(data string, weight, height float64) string {
	// ваш код ниже
	steps, trainingName, duration, err := parseTraining(data)

	if err != nil {
		return fmt.Sprintf("Ошибка: %v", err)
	}

	var calories float64
	var name string

	switch trainingName {
	case "Ходьба":
		calories = WalkingSpentCalories(steps, weight, height, duration)
		name = "Ходьба"
	case "Бег":
		calories = RunningSpentCalories(steps, weight, duration)
		name = "Бег"
	default:
		return "Неизвестный тип тренировки"
	}

	speed := meanSpeed(steps, duration)
	distance := distance(steps)

	message := fmt.Sprintf(
		"Тип тренировки: %s\n"+
			"Длительность: %.2f ч.\n"+
			"Дистанция: %.2f км.\n"+
			"Скорость: %.2f км/ч.\n"+
			"Сожгли калорий: %.2f",
		name,
		duration.Hours(), // преобразуем duration в часы
		distance,
		speed,
		calories,
	)

	return message
}

// Константы для расчета калорий, расходуемых при беге.
const (
	runningCaloriesMeanSpeedMultiplier = 18.0 // множитель средней скорости.
	runningCaloriesMeanSpeedShift      = 20.0 // среднее количество сжигаемых калорий при беге.
)

// RunningSpentCalories возвращает количество потраченных колорий при беге.
//
// Параметры:
//
// steps int - количество шагов.
// weight float64 — вес пользователя.
// duration time.Duration — длительность тренировки.
func RunningSpentCalories(steps int, weight float64, duration time.Duration) float64 {
	// ваш код здесь
	meanSpeed := meanSpeed(steps, duration)

	RunningSpentCalories := ((runningCaloriesMeanSpeedMultiplier * meanSpeed) - runningCaloriesMeanSpeedShift) * weight

	return RunningSpentCalories
}

// Константы для расчета калорий, расходуемых при ходьбе.
const (
	walkingCaloriesWeightMultiplier = 0.035 // множитель массы тела.
	walkingSpeedHeightMultiplier    = 0.029 // множитель роста.
)

// WalkingSpentCalories возвращает количество потраченных калорий при ходьбе.
//
// Параметры:
//
// steps int - количество шагов.
// duration time.Duration — длительность тренировки.
// weight float64 — вес пользователя.
// height float64 — рост пользователя.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) float64 {
	// ваш код здесь
	meanSpeed := meanSpeed(steps, duration)

	WalkingSpentCalories := ((walkingCaloriesWeightMultiplier * weight) + (meanSpeed*meanSpeed/height)*walkingSpeedHeightMultiplier) * duration.Hours() * float64(minInH)

	return WalkingSpentCalories
}
