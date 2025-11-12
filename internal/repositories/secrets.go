package repositories

// Imports
import (
	context "context" // Context

	"github.com/google/uuid"              // Google package for UUID generation
	"joshgill.dev/pulse/internal/clients" // Internal Clients package
	"joshgill.dev/pulse/internal/models"  // Internal Models package
)


// Define the Secrets Repository Interface and corresponding structure
type SecretsRepositoryInterface interface {
    Create(ctx context.Context, environment string, value string, comment string) (*models.Secret, error)
    Get(ctx context.Context, key string, environment string) (*models.Secret, error)
    Update(ctx context.Context, key string, environment string, newValue string) (*models.Secret, error)
    Delete(ctx context.Context, key string, environment string) (error)
}


type secretsRepository struct {
    SecretsClient clients.InfisicalClientInterface
}


// Constructor
func NewSecretsRepository(secretsClient clients.InfisicalClientInterface) SecretsRepositoryInterface {
    return &secretsRepository{ SecretsClient: secretsClient }
}


// Method for creating a Secret, given the relevant information
func (s *secretsRepository) Create(ctx context.Context, environment string, value string, comment string) (*models.Secret, error) {
    // Create a "Secret" with the relevant information, return the result of the Client Method
    secret := models.Secret{
        Key: uuid.New().String(),
        ProjectID: "pulse",
        Environment: environment,
        Value: value,
        Comment: comment,
    }; return s.SecretsClient.CreateSecret(ctx, &secret)
}

// Method for getting a Secret, given a Key and Environment
func (s *secretsRepository) Get(ctx context.Context, key string, environment string) (*models.Secret, error) {
    // Create a "Secret" with the known information, return the result of the Client Method
    secret := models.Secret{
        Key: key,
        ProjectID: "pulse",
        Environment: environment,
    }; return s.SecretsClient.GetSecret(ctx, &secret)
}

// Method for updating a Secret, given the relevant information and a new Value
func (s *secretsRepository) Update(ctx context.Context, key string, environment string, newValue string) (*models.Secret, error) {
    // Create a "Secret" with the known information, return the result of the Client Method
    secret := models.Secret{
        Key: key,
        Environment: environment,
        ProjectID: "pulse",
        Value: newValue,
    }; return s.SecretsClient.UpdateSecret(ctx, &secret)
}

// Method for deleting a Secret, given the relevant information
func (s *secretsRepository) Delete(ctx context.Context, key string, environment string) (error) {
    // Create a "Secret" with the known information, return the result of the Client Method
    secret := models.Secret{
        Key: key,
        Environment: environment,
        ProjectID: "pulse",
    }; return s.SecretsClient.DeleteSecret(ctx, &secret)
}
