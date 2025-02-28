package env

import (
	"testing"
	"time"

	. "github.com/onsi/gomega"
)

func Test_GetConfig(t *testing.T) {
	g := NewGomegaWithT(t)
	envs := map[string]string{
		// required
		"CLIENT_ID":              "CLIENT_ID",
		"CLIENT_SECRET":          "CLIENT_SECRET",
		"TOKEN_ENDPOINT":         "TOKEN_ENDPOINT",
		"WEBHOOK_TOKEN_ENDPOINT": "WEBHOOK_TOKEN_ENDPOINT",
		"DOMAIN":                 "DOMAIN",
		"EVENT_TYPE_PREFIX":      "EVENT_TYPE_PREFIX",
		// optional
		"BEB_API_URL":                "BEB_API_URL",
		"BEB_NAMESPACE":              "/test",
		"WEBHOOK_ACTIVATION_TIMEOUT": "60s",
		"ENABLE_NEW_CRD_VERSION":     "true",
	}

	for k, v := range envs {
		t.Setenv(k, v)
	}
	config := GetConfig()
	// Ensure required variables can be set
	g.Expect(config.ClientID).To(Equal(envs["CLIENT_ID"]))
	g.Expect(config.ClientSecret).To(Equal(envs["CLIENT_SECRET"]))
	g.Expect(config.TokenEndpoint).To(Equal(envs["TOKEN_ENDPOINT"]))
	g.Expect(config.WebhookTokenEndpoint).To(Equal(envs["WEBHOOK_TOKEN_ENDPOINT"]))
	g.Expect(config.Domain).To(Equal(envs["DOMAIN"]))
	g.Expect(config.EventTypePrefix).To(Equal(envs["EVENT_TYPE_PREFIX"]))
	g.Expect(config.BEBNamespace).To(Equal(envs["BEB_NAMESPACE"]))
	// Ensure optional variables can be set
	g.Expect(config.BEBAPIURL).To(Equal(envs["BEB_API_URL"]))

	webhookActivationTimeout, err := time.ParseDuration(envs["WEBHOOK_ACTIVATION_TIMEOUT"])
	g.Expect(err).ShouldNot(HaveOccurred())
	g.Expect(config.WebhookActivationTimeout).To(Equal(webhookActivationTimeout))
	g.Expect(config.EnableNewCRDVersion).To(Equal(true))
}
