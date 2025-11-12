package clients

// Imports
import (
	context "context"						// Context
	"os"									// Environment Variable Access
	"time"									// Time

	infisical "github.com/infisical/go-sdk" // Infisical SDK for Secret Management
	"joshgill.dev/pulse/internal/models"    // Internal model definitions
)


type InfisicalClientInterface interface {
	CreateSecret(ctx context.Context, secret *models.Secret) (*models.Secret, error)
	GetSecret(ctx context.Context, secret *models.Secret) (*models.Secret, error)
	UpdateSecret(ctx context.Context, secret *models.Secret) (*models.Secret, error)
	DeleteSecret(ctx context.Context, secret *models.Secret) (error)
}

type infisicalClient struct {
	SDK infisical.InfisicalClientInterface
}

// Constructor
func NewInfisical() (InfisicalClientInterface, error) {
	// Create the Client
	client := infisical.NewInfisicalClient(
		context.Background(),
		infisical.Config{ AutoTokenRefresh: true },		// Let the SDK handle token refresh
	)

	// Attempt to Authenticate to Infisical using the Environment Variable Values, if the error is not nil return it
	_, err:= client.Auth().UniversalAuthLogin(
		os.Getenv("INFISICAL_DB_USER"),
		os.Getenv("INFISICAL_DB_PASSWORD"),
	); if err != nil { return nil, err }
	// Otherwise, return the client
	return &infisicalClient{SDK: client}, nil
}


// Function for creating a Secret
func (c *infisicalClient) CreateSecret(ctx context.Context, secret *models.Secret) (*models.Secret, error) {
	// Create the Secret, ignoring the returned Infisical Secret struct
	// If the error is not nil, return it
	secretMade, err := c.SDK.Secrets().Create(infisical.CreateSecretOptions{
		ProjectID: secret.ProjectID,
		Environment: secret.Environment,
		SecretKey: secret.Key,
		SecretValue: secret.Value,
		SecretComment: secret.Comment,
	}); if err != nil { return nil, err }

	// Otherwise, set the relevant information and return the secret
	secret.Version = secretMade.Version
	secret.Retrieved = time.Now()
	return secret, nil
}

// Function for getting a Secret
func (c *infisicalClient) GetSecret(ctx context.Context, secret *models.Secret) (*models.Secret, error) {
	// Get the Secret from Infisical
	// If the error returned is not nil, return it
	secretFound, err := c.SDK.Secrets().Retrieve(infisical.RetrieveSecretOptions{
		SecretKey: secret.Key,
		ProjectID: secret.ProjectID,
		Environment: secret.Environment,
	}); if err != nil { return nil, err }

	// Otherwise, populate relevant values and return the secret pointer
	secret.Value = secretFound.SecretValue
	secret.Version = secretFound.Version
	secret.Comment = secretFound.SecretComment
	secret.Retrieved = time.Now()
	return secret, nil
}

// Function for updating a Secret
func (c *infisicalClient) UpdateSecret(ctx context.Context, secret *models.Secret) (*models.Secret, error) {
	// Update the Secret
	// If the error returned is not nil, return it
	newSecret, err := c.SDK.Secrets().Update(infisical.UpdateSecretOptions{
		ProjectID: secret.ProjectID,
		Environment: secret.Environment,
		SecretKey: secret.Key,
		NewSecretValue: secret.Value,
	}); if err != nil { return nil, err }

	// Update the given reference with the new values and return it
	secret.Version = newSecret.Version
	return secret, nil
}

// Function for deleting a Secret
func (c *infisicalClient) DeleteSecret(ctx context.Context, secret *models.Secret) (error) {
	// Delete the secret, ignoring the returned secret value
	// Return the error value, nil or otherwise
	_, err := c.SDK.Secrets().Delete(infisical.DeleteSecretOptions{
		ProjectID: secret.ProjectID,
		Environment: secret.Environment,
		SecretKey: secret.Key,
	}); return err
}
