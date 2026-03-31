package spentcalories

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	/* пакет с ошибками */
	"github.com/Yandex-Practicum/tracker/internal/commonerrors"
)

const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

var (
	ErrLenSliceMoreThree = errors.New("Неверный формат входных данных")
	ErrWeightBelowZero   = errors.New("Вес меньше или равен 0")
	ErrHeightBelowZero   = errors.New("Рост меньше или равен 0")
	ErrDurationBelowZero = errors.New("Продолжительность меньше или равна 0")
)

func parseTraining(data string) (int, string, time.Duration, error) {
	if len(data) == 0 {
		return 0, "", 0, commonerrors.ErrEmptyData
	}

	/* ["678", "Ходьба", "0h50m"] */
	stringSlice := strings.Split(data, ",")

	if len(stringSlice) == 3 {
		countOfSteps, err := strconv.Atoi(stringSlice[0])

		if err != nil {
			return 0, "", 0, commonerrors.ErrWithConvertStringToInt
		}

		if countOfSteps <= 0 {
			return 0, "", 0, commonerrors.ErrCountOfStepsBelowZero
		}

		typeOfTraining := stringSlice[1]

		duration, err := time.ParseDuration(stringSlice[2])

		if err != nil {
			return 0, "", 0, commonerrors.ErrParseDurationFromStr
		}

		if duration <= 0 {
			return 0, "", 0, ErrDurationBelowZero
		}

		return countOfSteps, typeOfTraining, duration, nil
	}

	return 0, "", 0, ErrLenSliceMoreThree
}

func distance(steps int, height float64) float64 {
	return float64(steps) * height * stepLengthCoefficient / mInKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}

	distanceInKm := distance(steps, height)
	durationInHours := duration.Hours()

	return distanceInKm / durationInHours
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	steps, typeOfTraining, duration, err := parseTraining(data)

	if err != nil {
		log.Println(err)
		return "", err
	}

	switch typeOfTraining {
	case "Бег":
		distanceValue := distance(steps, height)
		meanSpeedValue := meanSpeed(steps, height, duration)
		spentCalories, err := RunningSpentCalories(steps, weight, height, duration)

		if err != nil {
			return "", err
		}

		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", typeOfTraining, duration.Hours(), distanceValue, meanSpeedValue, spentCalories), nil

	case "Ходьба":
		distanceValue := distance(steps, height)
		meanSpeedValue := meanSpeed(steps, height, duration)
		spentCalories, err := WalkingSpentCalories(steps, weight, height, duration)

		if err != nil {
			return "", err
		}

		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", typeOfTraining, duration.Hours(), distanceValue, meanSpeedValue, spentCalories), nil
	}

	return "", errors.New("неизвестный тип тренировки")
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, commonerrors.ErrCountOfStepsBelowZero
	}

	if weight <= 0 {
		return 0, ErrWeightBelowZero
	}

	if height <= 0 {
		return 0, ErrHeightBelowZero
	}

	if duration <= 0 {
		return 0, ErrDurationBelowZero
	}

	meanSpeedValue := meanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	spentCalories := (weight * meanSpeedValue * durationInMinutes) / minInH

	return spentCalories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, commonerrors.ErrCountOfStepsBelowZero
	}

	if weight <= 0 {
		return 0, ErrWeightBelowZero
	}

	if height <= 0 {
		return 0, ErrHeightBelowZero
	}

	if duration <= 0 {
		return 0, ErrDurationBelowZero
	}

	meanSpeedValue := meanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	spentCalories := (weight * meanSpeedValue * durationInMinutes) / minInH
	walkingSpentCalories := spentCalories * walkingCaloriesCoefficient

	return walkingSpentCalories, nil
}
