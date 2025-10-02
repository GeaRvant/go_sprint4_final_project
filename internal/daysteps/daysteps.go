package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	strParts := strings.Split(data, ",")
	if len(strParts) != 2 {
		return 0, 0, fmt.Errorf("Ошибка преобразования строки: ожидалось 2 части, получено %d", len(strParts))
	}
	elementToInt, err := strconv.Atoi(strings.TrimSpace(strParts[0]))
	if err != nil {
		return 0, 0, fmt.Errorf("Ошибка преобразования в int: %w", err)
	}
	if elementToInt < 0 {
		return 0, 0, fmt.Errorf("Ошибка: отрицательное количество шагов (%d)", elementToInt)
	}
	duration, err := time.ParseDuration(strings.TrimSpace(strParts[1]))
	if err != nil {
		return 0, 0, fmt.Errorf("Ошибка парсинга длительности: %w", err)
	}
	return elementToInt, duration, nil

}

func DayActionInfo(data string, weight, height float64) string {
	walkSteps, walkDuration, err := parsePackage(data)
	if err != nil {
		return fmt.Sprintf("Ошибка при получении данных: %v", err)
	}
	if walkSteps < 0 {
		return fmt.Sprintf("Ошибка: отрицательное количество шагов (%d)", walkSteps)
	}
	if walkDuration < 0 {
		return fmt.Sprintf("Ошибка: отрицательная длительность (%d)", walkDuration)
	}
	distance := (float64(walkSteps) * stepLength) / float64(mInKm)
	calories := WalkingSpentCalories(walkSteps, weight, height, walkDuration)
	return fmt.Sprintf(
		"Количество шагов: %d.\nДистанция составила: %.2f км.\nВы сожгли: %.2f ккал.\n",
		walkSteps, distance, calories)
}
