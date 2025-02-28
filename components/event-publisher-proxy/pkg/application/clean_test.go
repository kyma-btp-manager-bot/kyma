package application

import (
	"testing"

	"github.com/stretchr/testify/require"

	applicationv1alpha1 "github.com/kyma-project/kyma/components/application-operator/pkg/apis/applicationconnector/v1alpha1"

	"github.com/kyma-project/kyma/components/event-publisher-proxy/pkg/application/applicationtest"
)

func TestCleanName(t *testing.T) {
	testCases := []struct {
		givenApplication *applicationv1alpha1.Application
		wantName         string
	}{
		// application type label is missing, then use the application name
		{
			givenApplication: applicationtest.NewApplication("alphanumeric0123", nil),
			wantName:         "alphanumeric0123",
		},
		{
			givenApplication: applicationtest.NewApplication("alphanumeric0123", map[string]string{"ignore-me": "value"}),
			wantName:         "alphanumeric0123",
		},
		{
			givenApplication: applicationtest.NewApplication("with.!@#none-$%^alphanumeric_&*-characters", nil),
			wantName:         "withnonealphanumericcharacters",
		},
		{
			givenApplication: applicationtest.NewApplication("with.!@#none-$%^alphanumeric_&*-characters", map[string]string{"ignore-me": "value"}),
			wantName:         "withnonealphanumericcharacters",
		},
		// application type label is available, then use it instead of the application name
		{
			givenApplication: applicationtest.NewApplication("alphanumeric0123", map[string]string{TypeLabel: "apptype"}),
			wantName:         "apptype",
		},
		{
			givenApplication: applicationtest.NewApplication("with.!@#none-$%^alphanumeric_&*-characters", map[string]string{TypeLabel: "apptype"}),
			wantName:         "apptype",
		},
		{
			givenApplication: applicationtest.NewApplication("alphanumeric0123", map[string]string{TypeLabel: "apptype=with.!@#none-$%^alphanumeric_&*-characters"}),
			wantName:         "apptypewithnonealphanumericcharacters",
		},
		{
			givenApplication: applicationtest.NewApplication("with.!@#none-$%^alphanumeric_&*-characters", map[string]string{TypeLabel: "apptype=with.!@#none-$%^alphanumeric_&*-characters"}),
			wantName:         "apptypewithnonealphanumericcharacters",
		},
	}

	for _, tc := range testCases {
		if gotName := GetCleanTypeOrName(tc.givenApplication); tc.wantName != gotName {
			t.Errorf("Clean application name:[%s] failed, want:[%v] but got:[%v]", tc.givenApplication.Name, tc.wantName, gotName)
		}
	}
}

func Test_GetTypeOrName(t *testing.T) {
	testCases := []struct {
		name             string
		givenApplication *applicationv1alpha1.Application
		wantName         string
	}{
		// application type label is missing, then use the application name
		{
			name:             "Should return application name if no labels",
			givenApplication: applicationtest.NewApplication("alphanumeric0123", nil),
			wantName:         "alphanumeric0123",
		},
		{
			name:             "Should return application name if label with right key does not exists",
			givenApplication: applicationtest.NewApplication("alphanumeric0123", map[string]string{"ignore-me": "value"}),
			wantName:         "alphanumeric0123",
		},
		{
			name:             "Should return application name as unclean",
			givenApplication: applicationtest.NewApplication("with.!@#none-$%^alphanumeric_&*-characters", nil),
			wantName:         "with.!@#none-$%^alphanumeric_&*-characters",
		},
		// application type label is available, then use it instead of the application name
		{
			name:             "Should return application label instead of name",
			givenApplication: applicationtest.NewApplication("alphanumeric0123", map[string]string{TypeLabel: "apptype"}),
			wantName:         "apptype",
		},
		{
			name:             "Should return application label as unclean",
			givenApplication: applicationtest.NewApplication("alphanumeric0123", map[string]string{TypeLabel: "apptype=with.!@#none-$%^alphanumeric_&*-characters"}),
			wantName:         "apptype=with.!@#none-$%^alphanumeric_&*-characters",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			require.Equal(t, tc.wantName, GetTypeOrName(tc.givenApplication))
		})
	}
}

func TestIsCleanName(t *testing.T) {
	testCases := []struct {
		givenName string
		wantClean bool
	}{
		{givenName: "with.dot", wantClean: false},
		{givenName: "with-dash", wantClean: false},
		{givenName: "with_underscore", wantClean: false},
		{givenName: "with#special$characters", wantClean: false},
		{givenName: "alphabetical", wantClean: true},
		{givenName: "alphanumeric0123", wantClean: true},
	}

	for _, tc := range testCases {
		if gotClean := IsCleanName(tc.givenName); tc.wantClean != gotClean {
			t.Errorf("Is clean application name:[%s] failed, want:[%v] but got:[%v]", tc.givenName, tc.wantClean, gotClean)
		}
	}
}
