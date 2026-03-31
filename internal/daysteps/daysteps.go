package daysteps

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/commonerrors"
	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	if len(data) == 0 {
		return 0, 0, commonerrors.ErrEmptyData
	}

	stringSlice := strings.Split(data, ",")

	if len(stringSlice) != 2 {
		return 0, 0, errors.New("Что-то пошло не так. В строке нет 2 нужных элементов.")
	}

	countOfSteps, err := strconv.Atoi(stringSlice[0])

	if err != nil {
		return 0, 0, commonerrors.ErrWithConvertStringToInt
	}

	if countOfSteps <= 0 {
		return 0, 0, commonerrors.ErrCountOfStepsBelowZero
	}

	duration, err := time.ParseDuration(stringSlice[1])

	if err != nil {
		return 0, 0, commonerrors.ErrParseDurationFromStr
	}

	if duration <= 0 {
		return 0, 0, commonerrors.ErrParseDurationFromStr
	}

	return countOfSteps, duration, nil
}

func calcDistance(stepCount int) float64 {
	return float64(stepCount) * stepLength
}

/* DayActionInfo возращает дистанцию в километрах и количество потраченных калорий в виде стринги */
func DayActionInfo(data string, weight, height float64) string {
	steps, duration, err := parsePackage(data)

	if err != nil {
		log.Println(err)
		return ""
	}

	if steps <= 0 {
		return ""
	}

	distanceInMeters := calcDistance(steps)
	distanceInKm := distanceInMeters / mInKm
	spentCalories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)

	if err != nil {
		log.Println(err)
		return ""
	}

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, distanceInKm, spentCalories)
}
