package management

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/auth0/go-auth0"
)

func TestEmailManager_Create(t *testing.T) {
	emailProvider := &Email{
		Name:               auth0.String("smtp"),
		Enabled:            auth0.Bool(true),
		DefaultFromAddress: auth0.String("no-reply@example.com"),
		Credentials: &EmailCredentials{
			SMTPHost: auth0.String("smtp.example.com"),
			SMTPPort: auth0.Int(587),
			SMTPUser: auth0.String("user"),
			SMTPPass: auth0.String("pass"),
		},
	}

	err := m.Email.Create(emailProvider)
	assert.NoError(t, err)

	defer cleanupEmailProvider(t)
}

func TestEmailManager_Read(t *testing.T) {
	expectedEmailProvider := givenAnEmailProvider(t)
	defer cleanupEmailProvider(t)

	actualEmailProvider, err := m.Email.Read()

	assert.NoError(t, err)
	assert.Equal(t, expectedEmailProvider.GetName(), actualEmailProvider.GetName())
	assert.Equal(t, expectedEmailProvider.GetEnabled(), actualEmailProvider.GetEnabled())
	assert.Equal(t, expectedEmailProvider.GetDefaultFromAddress(), actualEmailProvider.GetDefaultFromAddress())
	assert.Equal(
		t,
		expectedEmailProvider.GetCredentials().GetSMTPUser(),
		actualEmailProvider.GetCredentials().GetSMTPUser(),
	)
	assert.Equal(
		t,
		"",
		actualEmailProvider.GetCredentials().GetSMTPPass(),
	) // Passwords are not returned from the Auth0 API.
}

func TestEmailManager_Update(t *testing.T) {
	emailProvider := givenAnEmailProvider(t)
	defer cleanupEmailProvider(t)

	emailProvider.Enabled = auth0.Bool(false)
	emailProvider.DefaultFromAddress = auth0.String("info@example.com")

	err := m.Email.Update(emailProvider)
	assert.NoError(t, err)

	actualEmailProvider, err := m.Email.Read()
	assert.NoError(t, err)

	assert.False(t, actualEmailProvider.GetEnabled())
	assert.Equal(t, "info@example.com", actualEmailProvider.GetDefaultFromAddress())
}

func TestEmailManager_Delete(t *testing.T) {
	givenAnEmailProvider(t)

	err := m.Email.Delete()
	assert.NoError(t, err)

	_, err = m.Email.Read()
	assert.Error(t, err)
	assert.Implements(t, (*Error)(nil), err)
	assert.Equal(t, http.StatusNotFound, err.(Error).Status())
}

func givenAnEmailProvider(t *testing.T) *Email {
	emailProvider := &Email{
		Name:               auth0.String("smtp"),
		Enabled:            auth0.Bool(true),
		DefaultFromAddress: auth0.String("no-reply@example.com"),
		Credentials: &EmailCredentials{
			SMTPHost: auth0.String("smtp.example.com"),
			SMTPPort: auth0.Int(587),
			SMTPUser: auth0.String("user"),
			SMTPPass: auth0.String("pass"),
		},
	}

	err := m.Email.Create(emailProvider)
	require.NoError(t, err)

	return emailProvider
}

func cleanupEmailProvider(t *testing.T) {
	err := m.Email.Delete()
	require.NoError(t, err)
}
