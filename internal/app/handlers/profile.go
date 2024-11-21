package handlers

import (
	"context"
	"log"
	"strings"

	"github.com/jinzhu/copier"

	"github.com/mayuka-c/leaderboard-system-go/internal/app/models"
	db "github.com/mayuka-c/leaderboard-system-go/internal/pkg/db/sqlc"
)

func (h *service) CreateProfile(ctx context.Context, input models.CreateProfile) (map[string]int64, error) {
	id, err := h.dbQueries.CreateProfile(ctx, db.CreateProfileParams{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Age:       input.Age,
		Gender:    db.GenderT(input.Gender),
		PlayerID:  input.PlayerID,
	})
	if err != nil {
		log.Fatalf("Failed to create profile for playerID: %d into DB. Err: %s", input.PlayerID, err.Error())
		return map[string]int64{}, err
	}

	return map[string]int64{"id": id}, err
}

func (h *service) GetPlayerProfile(ctx context.Context, player_id int64) (models.Profile, error) {
	profile := models.Profile{}

	dbProfile, err := h.dbQueries.PlayerProfile(ctx, player_id)
	if err != nil {
		if strings.Contains(err.Error(), "sql: no rows in result set") {
			return models.Profile{}, nil
		}
		log.Fatalf("Failed to get the profile details for player with id: %d. Err: %s", player_id, err.Error())
		return profile, err
	}

	if err := copier.Copy(&profile, dbProfile); err != nil {
		log.Fatalf("Error while doing the copy to profile model. Err: %s", err.Error())
		return profile, err
	}

	return profile, err
}

func updateInputFields(getDbProfile db.Profile, input *models.UpdateProfile) {
	if input.FirstName == "" {
		input.FirstName = getDbProfile.FirstName
	}
	if input.LastName == "" {
		input.LastName = getDbProfile.LastName
	}
	if input.Email == "" {
		input.Email = getDbProfile.Email
	}
	if input.Age == 0 {
		input.Age = getDbProfile.Age
	}
	if input.Gender == "" {
		input.Gender = string(getDbProfile.Gender)
	}
}

func (h *service) UpdateProfile(ctx context.Context, input models.UpdateProfile) (models.Profile, error) {
	getDbProfile, err := h.dbQueries.GetProfile(ctx, input.ID)
	if err != nil {
		if strings.Contains(err.Error(), "sql: no rows in result set") {
			return models.Profile{}, nil
		}
		log.Fatalf("Failed to get the profile details with id: %d. Err: %s", input.ID, err.Error())
		return models.Profile{}, err
	}

	updateInputFields(getDbProfile, &input)

	dbProfile, err := h.dbQueries.UpdateProfile(ctx, db.UpdateProfileParams{
		ID:        input.ID,
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Age:       input.Age,
		Gender:    db.GenderT(input.Gender),
		PlayerID:  getDbProfile.PlayerID,
	})
	if err != nil {
		if strings.Contains(err.Error(), "sql: no rows in result set") {
			return models.Profile{}, nil
		}
		log.Fatalf("Failed to update profile with id:%d record into DB. Err: %s", input.ID, err.Error())
		return models.Profile{}, err
	}

	profile := models.Profile{}
	if err := copier.Copy(&profile, dbProfile); err != nil {
		log.Fatalf("Error while doing the copy to profile update model. Err: %s", err.Error())
		return profile, err
	}

	return profile, err
}
