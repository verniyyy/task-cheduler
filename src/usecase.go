package src

type CronUsecase struct {
	Schedule  Scheduler
	JobCreate JobCreator
	runners   []func()
}

func (u CronUsecase) Run() error {
	for _, runner := range u.runners {
		runner()
	}
	return nil
}

type SendGoogleChatInput struct {
	Webhook    string `json:"webhook_url"`
	Message    string `json:"message"`
	EveryDayAt string `json:"every_day_at"`
}

func (u *CronUsecase) AddSendGoocleChatJob(in SendGoogleChatInput) error {
	task := Task{
		Job: u.JobCreate.GoogleChatJob(in.Webhook, in.Message),
	}
	if in.EveryDayAt != "" {
		u.runners = append(u.runners, func() {
			u.Schedule.RunEveryDayAt(task, in.EveryDayAt)
		})
	}
	return nil
}
