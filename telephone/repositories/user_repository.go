package repositories

type UserRepository interface {
	Create(user *models.User) error
}

type userRepository struct {
	// MongoDB session or client
}

func NewUserRepository() UserRepository {
	// Initialize MongoDB client or session
	return &userRepository{}
}

func (r *userRepository) Create(user *models.User) error {
	// Insert user into MongoDB
	return nil
}
