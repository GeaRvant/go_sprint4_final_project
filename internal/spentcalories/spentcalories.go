package spentcalories

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	strParts := strings.Split(data, ",")
	if len(strParts) != 3 {
		return 0, "", 0, fmt.Errorf("Ошибка преобразования строки: ожидалось 3 части, получено %d", len(strParts))
	}
	elementToInt, err := strconv.Atoi(strings.TrimSpace(strParts[0]))
	if err != nil {
		return 0, "", 0, fmt.Errorf("Ошибка преобразования в int: %w", err)
	}
	duration, err := time.ParseDuration(strings.TrimSpace(strParts[2]))
	if err != nil {
		return 0, "", 0, fmt.Errorf("Ошибка парсинга длительности: %w", err)
	}
	return elementToInt, strParts[1], duration, nil
}

func distance(steps int, height float64) float64 {
	stepsLength := height * stepLengthCoefficient
	return (stepsLength * float64(steps)) / float64(mInKm)
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}
	dist := distance(steps, height)
	durationHours := duration.Hours()
	if durationHours == 0 {
		return 0
	}
	return dist / durationHours
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	steps, training, duration, err := parseTraining(data)
	if err != nil {
		log.Println("Ошибка при получении данных", err)
		return "", err
	}
	switch training {
	case "Ходьба":
		dist := distance(steps, height)
		speed := meanSpeed(steps, height, duration)
		calories, err := WalkingSpentCalories(steps, weight, height, duration)
		if err != nil {
			log.Println("Ошибка при получении данных", err)
			return "", err
		}
		return fmt.Sprintf(
			"Тип тренировки: %s.\nДлительность: %v,\nДистанция составила: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f ккал.\n",
			training, duration, dist, speed, calories), nil
	case "Бег":
		dist := distance(steps, height)
		speed := meanSpeed(steps, height, duration)
		calories, err := RunningSpentCalories(steps, weight, height, duration)
		if err != nil {
			log.Println("Ошибка при получении данных", err)
			return "", err
		}
		return fmt.Sprintf(
			"Тип тренировки: %s.\nДлительность: %v,\nДистанция составила: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f ккал.\n",
			training, duration, dist, speed, calories), nil
	default:
		return "", fmt.Errorf("Ошибка: неизвестный тип тренировки (%v)", training)
	}
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, fmt.Errorf("Ошибка: некорректное количество шагов (%d)", steps)
	}
	if weight <= 0 {
		return 0, fmt.Errorf("Ошибка: некорректное значение веса (%.2f)", weight)
	}
	if height <= 0 {
		return 0, fmt.Errorf("Ошибка: некорректное значение роста (%.2f)", height)
	}
	if duration <= 0 {
		return 0, fmt.Errorf("Ошибка: некорректное значение длительности (%v)", duration)
	}
	speed := meanSpeed(steps, height, duration)
	durationMin := duration.Minutes()
	return (weight * speed * durationMin) / minInH, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, fmt.Errorf("Ошибка: некорректное количество шагов (%d)", steps)
	}
	if weight <= 0 {
		return 0, fmt.Errorf("Ошибка: некорректное значение веса (%.2f)", weight)
	}
	if height <= 0 {
		return 0, fmt.Errorf("Ошибка: некорректное значение роста (%.2f)", height)
	}
	if duration <= 0 {
		return 0, fmt.Errorf("Ошибка: некорректное значение длительности (%v)", duration)
	}
	speed := meanSpeed(steps, height, duration)
	durationMin := duration.Minutes()
	return (weight * speed * durationMin) / minInH, nil
}
