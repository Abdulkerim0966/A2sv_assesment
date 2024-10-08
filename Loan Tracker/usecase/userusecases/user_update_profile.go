package userusecase

import (
	"loanTracker/config"
	"loanTracker/domain"
)

func (u *UserUsecase) UpdateProfile(user *domain.User, claims *domain.LoginClaims) error {
	// Get existing user details
	existingUser, err := u.UserRepo.GetUserByUsernameorEmail(claims.Username)
	if err != nil {
		return err
	}

	// Check if the first name is present
	if user.FirstName != "" {
		existingUser.FirstName = user.FirstName
	}

	// Check if the last name is present
	if user.LastName != "" {
		existingUser.LastName = user.LastName
	}

	// Check if the bio is present
	if user.Bio != "" {
		existingUser.Bio = user.Bio
	}

	// Check if the address is present
	if user.Address != "" {
		existingUser.Address = user.Address
	}

	// Check if the email is present and is unique and valid
	if user.Email != "" {
		// Validate the email
		err = config.IsValidEmail(user.Email)
		if err != nil {
			return err
		}

		// Check if the email is unique
		err = u.UserRepo.CheckUsernameAndEmail(existingUser.Username, user.Email)
		if err != nil {
			return err
		}

		existingUser.Email = user.Email
	}

	// Check if the username is present and is unique
	if user.Username != "" {
		err = config.IsValidUsername(user.Username)
		if err != nil {
			return err
		}

		// Check if the username is unique
		err = u.UserRepo.CheckUsernameAndEmail(user.Username, existingUser.Email)
		if err != nil {
			return err
		}

		existingUser.Username = user.Username
	
	}

	// Check if the password is present and is strong
	if user.Password != "" {
		// Validate the password
		err = config.IsStrongPassword(user.Password)
		if err != nil {
			return err
		}

		// Hash the new password
		hashedPassword, err := config.HashPassword(user.Password)
		if err != nil {
			return err
		}

		existingUser.Password = hashedPassword
	}

	// Check if the role is present
	if user.Role != "" {
		return config.ErrUpdateRole
	}

	// Check if the joined date is present
	if !user.JoinedDate.IsZero() {
		return config.ErrUpdateJoined
	}

	// Update the user profile in the repository
	err = u.UserRepo.UpdateProfile(claims.Username, existingUser)
	if err != nil {
		return err
	}

	return nil
}
