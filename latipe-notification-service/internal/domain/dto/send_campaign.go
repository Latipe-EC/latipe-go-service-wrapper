package dto

type SendCampaignNotificationRequest struct {
	CampaignTopic   string `json:"campaign_topic" validate:"required"`
	Title           string `json:"title" validate:"required"`
	Body            string `json:"body" validate:"required"`
	Image           string `json:"image" validate:"required"`
	ScheduleDisplay string `json:"schedule_display" validate:"required"`
}

type SendCampaignNotificationResponse struct {
	ID string `json:"id"`
}
