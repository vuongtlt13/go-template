package cron_job

import (
	"context"
	"os"
	"strconv"
	"time"
	"yourapp/internal/domain/service"

	"github.com/robfig/cron/v3"
)

// ScheduleCleanup schedules the user cleanup task
func ScheduleCleanup(service service.RoleService) *cron.Cron {
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
		service.CleanupUnverifiedUsers(ctx, days)
	})

	if err != nil {
		// Handle error
		return nil
	}

	c.Start()
	return c
}

// CleanupUnverifiedUsers implementation in ServiceImpl
func (s *service.ServiceImpl) CleanupUnverifiedUsers(ctx context.Context, olderThanDays int) error {
	cutoffTime := time.Now().AddDate(0, 0, -olderThanDays)

	// Find users to be deleted for logging purposes
	users, err := s.repo.FindUnverifiedUsersCreatedBefore(ctx, cutoffTime)
	if err != nil {
		return err
	}

	// Log users to be deleted
	for _, user := range users {
		// Log user details
	}

	// Delete the users
	return s.repo.DeleteUnverifiedUsers(ctx, olderThanDays)
}
