package cron_job

import (
	"context"
	"os"
	"strconv"
	"yourapp/internal/domain/service"

	"github.com/robfig/cron/v3"
)

// ScheduleCleanup schedules the user cleanup task
func ScheduleCleanup(userService service.UserService) *cron.Cron {
	c := cron.New()

	// Run every day at midnight
	_, err := c.AddFunc("0 0 * * *", func() {
		// Default to 30 days if not specified
		daysStr := os.Getenv("UNVERIFIED_USER_RETENTION_DAYS")
		days, err := strconv.Atoi(daysStr)
		if err != nil || days <= 0 {
			days = 30
		}

		ctx := context.Background()
		userService.CleanupUnverifiedUsers(ctx, days)
	})

	if err != nil {
		// Handle error
		return nil
	}

	c.Start()
	return c
}
