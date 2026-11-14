package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	personaldata.Personal
	Steps    int
	Duration time.Duration
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")

	if len(parts) != 2 {
		return fmt.Errorf("ожидалось 2 поля для дневной прогулки, получено %d", len(parts))
	}

	steps, err := strconv.Atoi(parts[0])

	if err != nil {
		return fmt.Errorf("не удалось преобразовать количество шагов в число")
	}

	if steps <= 0 {
		return fmt.Errorf("количество шагов должно быть больше нуля: %d", steps)
	}

	duration, err := time.ParseDuration(parts[1])

	if err != nil {
		return fmt.Errorf("не удалось преобразовать продолжительность прогулки в число")
	}

	if duration <= 0 {
		return fmt.Errorf("продолжительность прогулки должна быть больше нуля: %s", duration)
	}

	ds.Steps = steps
	ds.Duration = duration

	return
}

func (ds DaySteps) ActionInfo() (string, error) {
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	distance := spentenergy.Distance(ds.Steps, ds.Height)

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, distance, calories), nil
}
