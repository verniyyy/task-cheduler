package src

func CronUsecaseFactory() (CronUsecase, error) {
	scheduler := NewScheduler()
	jobCreator := NewJobCreator()
	cron := CronUsecase{
		Schedule:  scheduler,
		JobCreate: jobCreator,
	}
	return cron, nil
}
