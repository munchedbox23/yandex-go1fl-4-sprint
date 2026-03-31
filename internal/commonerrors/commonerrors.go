package commonerrors

import "errors"

/* Так как переменные дублировались в internal/spentcalories и в daysteps - решил вынести их в отдельный internal файлик :) */
var (
	ErrEmptyData              = errors.New("Пришла пустая строка")
	ErrWithConvertStringToInt = errors.New("Не получилось преобразовать стрингу в int")
	ErrCountOfStepsBelowZero  = errors.New("Странно, кол-во шагов должно быть больше нуля")
	ErrParseDurationFromStr   = errors.New("К сожалению, не удалось распарсить продолжительность прогулки")
)
