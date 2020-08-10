package config_test

import (
	"context"
	"errors"
	"time"

	mock_token_validation "github.com/solo-io/ext-auth-service/pkg/config/oauth/token_validation/mocks"
	user_info_mocks "github.com/solo-io/ext-auth-service/pkg/config/oauth/user_info/mocks"
	"github.com/solo-io/ext-auth-service/pkg/config/oidc"

	configapi "github.com/solo-io/ext-auth-service/pkg/config"

	pbtypes "github.com/gogo/protobuf/types"

	"github.com/solo-io/solo-projects/projects/extauth/pkg/config/chain"

	"github.com/golang/mock/gomock"
	"github.com/solo-io/ext-auth-service/pkg/config/apr"
	chainmocks "github.com/solo-io/solo-projects/projects/extauth/pkg/config/chain/mocks"
	"github.com/solo-io/solo-projects/projects/extauth/pkg/plugins/mocks"

	"github.com/solo-io/ext-auth-plugins/api"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	extauthv1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/extauth/v1"
	"github.com/solo-io/solo-projects/projects/extauth/pkg/config"
)

var _ = Describe("Config Generator", func() {

	var (
		ctrl             *gomock.Controller
		generator        config.Generator
		pluginLoaderMock *mocks.MockLoader
		userInfoClient   *user_info_mocks.MockClient
		tokenValidator   *mock_token_validation.MockValidator
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		pluginLoaderMock = mocks.NewMockLoader(ctrl)
		userInfoClient = user_info_mocks.NewMockClient(ctrl)
		tokenValidator = mock_token_validation.NewMockValidator(ctrl)
		generator = config.NewGenerator(context.Background(), nil, "test-user-id-header", pluginLoaderMock, func(_ time.Duration, _ config.OAuthIntrospectionEndpoints) *config.OAuthIntrospectionClients {
			return &config.OAuthIntrospectionClients{
				TokenValidator: tokenValidator,
				UserInfoClient: userInfoClient,
			}
		})
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Context("plugin loading panics", func() {

		var (
			panicPlugin = &extauthv1.AuthPlugin{Name: "Panic"}
			ctx         context.Context
		)

		BeforeEach(func() {
			pluginLoaderMock.EXPECT().LoadAuthPlugin(gomock.Any(), panicPlugin).DoAndReturn(
				func(argCtx context.Context, _ *extauthv1.AuthPlugin) (api.AuthService, error) {
					ctx = argCtx
					panic("test load panic")
				},
			)
		})

		It("recovers from panic", func() {
			cfg, err := generator.GenerateConfig([]*extauthv1.ExtAuthConfig{
				{
					AuthConfigRefName: "default.test-authconfig",
					Configs: []*extauthv1.ExtAuthConfig_Config{
						{
							AuthConfig: &extauthv1.ExtAuthConfig_Config_PluginAuth{PluginAuth: panicPlugin},
						},
					},
				},
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(cfg).NotTo(BeNil())
			Expect(cfg.GetConfigCount()).To(Equal(0))
			Expect(ctx.Err()).To(Equal(context.Canceled))
		})
	})

	Context("all ext auth configs are valid", func() {
		It("correctly loads configs", func() {
			var okPlugin = &extauthv1.AuthPlugin{Name: "ThisOneWorks"}
			getAuthService := func(cfg configapi.Config, authConfigName string) api.AuthService {
				authService := cfg.AuthService(authConfigName)
				Expect(authService).NotTo(BeNil())
				authServiceChain, ok := authService.(chain.AuthServiceChain)
				Expect(ok).To(BeTrue())
				Expect(authServiceChain).NotTo(BeNil())
				services := authServiceChain.ListAuthServices()
				Expect(services).To(HaveLen(1))
				return services[0]
			}
			authServiceMock := chainmocks.NewMockAuthService(ctrl)
			authServiceMock.EXPECT().Start(gomock.Any()).Return(nil).AnyTimes()
			authServiceMock.EXPECT().Authorize(gomock.Any(), gomock.Any()).Times(0)

			pluginLoaderMock.EXPECT().LoadAuthPlugin(gomock.Any(), okPlugin).Return(authServiceMock, nil).Times(1)

			resources := []*extauthv1.ExtAuthConfig{
				{
					AuthConfigRefName: "default.plugin-authconfig",
					Configs: []*extauthv1.ExtAuthConfig_Config{
						{
							AuthConfig: &extauthv1.ExtAuthConfig_Config_PluginAuth{
								PluginAuth: okPlugin,
							},
						},
					},
				},
				{
					AuthConfigRefName: "default.basic-auth-authconfig",
					Configs: []*extauthv1.ExtAuthConfig_Config{
						{
							AuthConfig: &extauthv1.ExtAuthConfig_Config_BasicAuth{
								BasicAuth: &extauthv1.BasicAuth{
									Realm: "my-realm",
									Apr: &extauthv1.BasicAuth_Apr{
										Users: map[string]*extauthv1.BasicAuth_Apr_SaltedHashedPassword{
											"user": {
												Salt:           "salt",
												HashedPassword: "pwd",
											},
										},
									},
								},
							},
						},
					},
				},
				{
					AuthConfigRefName: "default.api-keys-authconfig",
					Configs: []*extauthv1.ExtAuthConfig_Config{
						{
							AuthConfig: &extauthv1.ExtAuthConfig_Config_ApiKeyAuth{
								ApiKeyAuth: &extauthv1.ExtAuthConfig_ApiKeyAuthConfig{
									ValidApiKeys: map[string]*extauthv1.ExtAuthConfig_ApiKeyAuthConfig_KeyMetadata{
										"key-1": {
											Username: "foo",
										},
										"key-2": {
											Username: "bar",
											Metadata: map[string]string{
												"user-id": "123",
											},
										},
									},
									HeaderName: "x-api-key",
									HeadersFromKeyMetadata: map[string]string{
										"x-user-id": "user-id",
									},
								},
							},
						},
					},
				},
				{
					AuthConfigRefName: "default.oauth-authconfig",
					Configs: []*extauthv1.ExtAuthConfig_Config{
						{
							AuthConfig: &extauthv1.ExtAuthConfig_Config_Oauth{
								Oauth: &extauthv1.ExtAuthConfig_OAuthConfig{
									IssuerUrl: "test",
								},
							},
						},
					},
				},
				{
					AuthConfigRefName: "default.ldap-authconfig",
					Configs: []*extauthv1.ExtAuthConfig_Config{
						{
							AuthConfig: &extauthv1.ExtAuthConfig_Config_Ldap{
								Ldap: &extauthv1.Ldap{
									Address:                 "my.server.com:389",
									UserDnTemplate:          "uid=%s,ou=people,dc=solo,dc=io",
									MembershipAttributeName: "someName",
									AllowedGroups: []string{
										"cn=managers,ou=groups,dc=solo,dc=io",
										"cn=developers,ou=groups,dc=solo,dc=io",
									},
									Pool: &extauthv1.Ldap_ConnectionPool{
										MaxSize: &pbtypes.UInt32Value{
											Value: uint32(5),
										},
										InitialSize: &pbtypes.UInt32Value{
											Value: uint32(0), // Set to 0, otherwise it will try to connect to the dummy address
										},
									},
								},
							},
						},
					},
				},
				{
					AuthConfigRefName: "default.oauth2-authconfig",
					Configs: []*extauthv1.ExtAuthConfig_Config{
						{
							AuthConfig: &extauthv1.ExtAuthConfig_Config_Oauth2{
								Oauth2: &extauthv1.ExtAuthConfig_OAuth2Config{
									OauthType: &extauthv1.ExtAuthConfig_OAuth2Config_AccessTokenValidation{
										AccessTokenValidation: &extauthv1.AccessTokenValidation{
											ValidationType: &extauthv1.AccessTokenValidation_IntrospectionUrl{
												IntrospectionUrl: "introspection-url",
											},
											UserinfoUrl:  "user-info-url",
											CacheTimeout: nil,
										},
									},
								},
							},
						},
					},
				},
			}

			pluginCfg, err := generator.GenerateConfig(resources)
			Expect(err).NotTo(HaveOccurred())
			Expect(pluginCfg).NotTo(BeNil())
			Expect(pluginCfg.GetConfigCount()).To(Equal(len(resources)))

			service := getAuthService(pluginCfg, resources[0].AuthConfigRefName)
			_, ok := service.(*chainmocks.MockAuthService)
			Expect(ok).To(BeTrue())

			service = getAuthService(pluginCfg, resources[1].AuthConfigRefName)
			aprConfig, ok := service.(*apr.Config)
			Expect(ok).To(BeTrue())
			Expect(aprConfig.Realm).To(Equal("my-realm"))
			Expect(aprConfig.SaltAndHashedPasswordPerUsername).To(BeEquivalentTo(
				map[string]apr.SaltAndHashedPassword{
					"user": {Salt: "salt", HashedPassword: "pwd"},
				}),
			)

			service = getAuthService(pluginCfg, resources[2].AuthConfigRefName)
			Expect(service).NotTo(BeNil())

			// Test that the Issuer Url always appends a trailing slash
			service = getAuthService(pluginCfg, resources[3].AuthConfigRefName)
			oidcConfig, ok := service.(*oidc.IssuerImpl)
			Expect(ok).To(BeTrue())
			Expect(oidcConfig.IssuerUrl).To(Equal("test/"))

			ldapService := getAuthService(pluginCfg, resources[4].AuthConfigRefName)
			Expect(ldapService).NotTo(BeNil())
		})

		It("uses the correct cache TTL in the presence or absence of user configuration", func() {
			ranClientsBuilder := false
			generator := config.NewGenerator(context.Background(), nil, "test-user-id-header", pluginLoaderMock, func(ttl time.Duration, _ config.OAuthIntrospectionEndpoints) *config.OAuthIntrospectionClients {
				ranClientsBuilder = true
				Expect(ttl).To(Equal(config.DefaultOAuthCacheTtl))

				return &config.OAuthIntrospectionClients{
					TokenValidator: tokenValidator,
					UserInfoClient: userInfoClient,
				}
			})

			_, err := generator.GenerateConfig([]*extauthv1.ExtAuthConfig{{
				AuthConfigRefName: "default.oauth2-authconfig",
				Configs: []*extauthv1.ExtAuthConfig_Config{
					{
						AuthConfig: &extauthv1.ExtAuthConfig_Config_Oauth2{
							Oauth2: &extauthv1.ExtAuthConfig_OAuth2Config{
								OauthType: &extauthv1.ExtAuthConfig_OAuth2Config_AccessTokenValidation{
									AccessTokenValidation: &extauthv1.AccessTokenValidation{
										ValidationType: &extauthv1.AccessTokenValidation_IntrospectionUrl{
											IntrospectionUrl: "introspection-url",
										},
										UserinfoUrl:  "user-info-url",
										CacheTimeout: nil, // not user-configured
									},
								},
							},
						},
					},
				},
			}})
			Expect(err).NotTo(HaveOccurred())
			Expect(ranClientsBuilder).To(BeTrue(), "Should have run the clients builder the first time")

			ranClientsBuilder = false
			generator = config.NewGenerator(context.Background(), nil, "test-user-id-header", pluginLoaderMock, func(ttl time.Duration, _ config.OAuthIntrospectionEndpoints) *config.OAuthIntrospectionClients {
				ranClientsBuilder = true
				Expect(ttl).To(Equal(time.Second))

				return &config.OAuthIntrospectionClients{
					TokenValidator: tokenValidator,
					UserInfoClient: userInfoClient,
				}
			})

			oneSecond := time.Second

			_, err = generator.GenerateConfig([]*extauthv1.ExtAuthConfig{{
				AuthConfigRefName: "default.oauth2-authconfig",
				Configs: []*extauthv1.ExtAuthConfig_Config{
					{
						AuthConfig: &extauthv1.ExtAuthConfig_Config_Oauth2{
							Oauth2: &extauthv1.ExtAuthConfig_OAuth2Config{
								OauthType: &extauthv1.ExtAuthConfig_OAuth2Config_AccessTokenValidation{
									AccessTokenValidation: &extauthv1.AccessTokenValidation{
										ValidationType: &extauthv1.AccessTokenValidation_IntrospectionUrl{
											IntrospectionUrl: "introspection-url",
										},
										UserinfoUrl:  "user-info-url",
										CacheTimeout: &oneSecond,
									},
								},
							},
						},
					},
				},
			}})
			Expect(err).NotTo(HaveOccurred())
			Expect(ranClientsBuilder).To(BeTrue(), "Should have run the clients builder the second time")
		})
	})

	Describe("fine-grained configuration updates", func() {

		When("the generator is invoked with the same exact resource multiple times", func() {

			var (
				plugin  = &extauthv1.AuthPlugin{Name: "SomePlugin"}
				ctxChan = make(chan context.Context)
			)

			BeforeEach(func() {
				authServiceMock := chainmocks.NewMockAuthService(ctrl)
				// Start is called only one time
				authServiceMock.EXPECT().Start(gomock.Any()).DoAndReturn(func(ctx context.Context) error {
					// Start functions are called asynchronously by the generator, so we need to wait for them to run
					// We also need to check this context, so with this channel we kill two birds with one stone
					ctxChan <- ctx
					return nil
				}).Times(1)
				authServiceMock.EXPECT().Authorize(gomock.Any(), gomock.Any()).Times(0)

				pluginLoaderMock.EXPECT().LoadAuthPlugin(gomock.Any(), plugin).Return(authServiceMock, nil).Times(1)
			})

			It("loads and starts the server configuration only once", func() {
				resources := []*extauthv1.ExtAuthConfig{
					{
						AuthConfigRefName: "default.my-auth-config",
						Configs: []*extauthv1.ExtAuthConfig_Config{
							{
								AuthConfig: &extauthv1.ExtAuthConfig_Config_PluginAuth{
									PluginAuth: plugin,
								},
							},
						},
					},
				}
				cfg, err := generator.GenerateConfig(resources)
				Expect(err).NotTo(HaveOccurred())
				Expect(cfg).NotTo(BeNil())
				Expect(cfg.GetConfigCount()).To(Equal(1))

				// Wait for start function to be called
				var ctx context.Context
				select {
				case <-time.After(time.Second):
					Fail("timed out waiting for Start function to be called")
				case ctx = <-ctxChan:
					// Verify that the context was not cancelled
					Expect(ctx.Err()).To(BeNil())
				}

				cfg, err = generator.GenerateConfig(resources)
				Expect(err).NotTo(HaveOccurred())
				Expect(cfg).NotTo(BeNil())
				Expect(cfg.GetConfigCount()).To(Equal(1))

				// Use another object with the same content just to make sure that we only care about structural equality
				resources2 := []*extauthv1.ExtAuthConfig{
					{
						AuthConfigRefName: "default.my-auth-config",
						Configs: []*extauthv1.ExtAuthConfig_Config{
							{
								AuthConfig: &extauthv1.ExtAuthConfig_Config_PluginAuth{
									PluginAuth: plugin,
								},
							},
						},
					},
				}

				cfg, err = generator.GenerateConfig(resources2)
				Expect(err).NotTo(HaveOccurred())
				Expect(cfg).NotTo(BeNil())
				Expect(cfg.GetConfigCount()).To(Equal(1))
				// Verify that the context was still not cancelled
				Expect(ctx).NotTo(BeNil())
				if ctx != nil {
					Expect(ctx.Err()).To(BeNil())
				}
			})
		})

		When("the generator is invoked with an updated version of an existing config", func() {

			var (
				plugin  = &extauthv1.AuthPlugin{Name: "SomePlugin"}
				ctxChan = make(chan context.Context)
			)

			BeforeEach(func() {
				authServiceMock := chainmocks.NewMockAuthService(ctrl)
				authServiceMock.EXPECT().Start(gomock.Any()).DoAndReturn(func(ctx context.Context) error {
					ctxChan <- ctx
					return nil
				}).Times(2)
				authServiceMock.EXPECT().Authorize(gomock.Any(), gomock.Any()).Times(0)
				pluginLoaderMock.EXPECT().LoadAuthPlugin(gomock.Any(), plugin).Return(authServiceMock, nil).Times(2)
			})

			It("loads and starts the server configuration twice, terminating the first instance", func() {
				resources := []*extauthv1.ExtAuthConfig{
					{
						AuthConfigRefName: "default.my-auth-config",
						Configs: []*extauthv1.ExtAuthConfig_Config{
							{
								AuthConfig: &extauthv1.ExtAuthConfig_Config_PluginAuth{
									PluginAuth: plugin,
								},
							},
						},
					},
				}
				cfg, err := generator.GenerateConfig(resources)
				Expect(err).NotTo(HaveOccurred())
				Expect(cfg).NotTo(BeNil())
				Expect(cfg.GetConfigCount()).To(Equal(1))

				// Wait for start function to be called
				var firstCtx context.Context
				select {
				case <-time.After(time.Second):
					Fail("timed out waiting for first Start function to be called")
				case firstCtx = <-ctxChan:
					// Verify that the context was not cancelled
					Expect(firstCtx.Err()).To(BeNil())
				}

				// Update the existing config
				plugin.PluginFileName = "plugin.so"

				cfg, err = generator.GenerateConfig(resources)
				Expect(err).NotTo(HaveOccurred())
				Expect(cfg).NotTo(BeNil())
				Expect(cfg.GetConfigCount()).To(Equal(1))

				// Wait for start function to be called
				var secondCtx context.Context
				select {
				case <-time.After(time.Second):
					Fail("timed out waiting for second Start function to be called")
				case secondCtx = <-ctxChan:
					// Verify that the context was not cancelled
					Expect(secondCtx.Err()).To(BeNil())
				}

				// Verify that the previous context was cancelled
				Expect(firstCtx).NotTo(BeNil())
				if firstCtx != nil {
					Expect(firstCtx.Err()).To(Equal(context.Canceled))
				}
			})
		})

		When("a currently existing config is no longer present in an update", func() {

			var (
				plugin  = &extauthv1.AuthPlugin{Name: "SomePlugin"}
				ctxChan = make(chan context.Context)
			)

			BeforeEach(func() {
				authServiceMock := chainmocks.NewMockAuthService(ctrl)
				authServiceMock.EXPECT().Start(gomock.Any()).DoAndReturn(func(ctx context.Context) error {
					ctxChan <- ctx
					return nil
				}).Times(1)
				authServiceMock.EXPECT().Authorize(gomock.Any(), gomock.Any()).Times(0)
				pluginLoaderMock.EXPECT().LoadAuthPlugin(gomock.Any(), plugin).Return(authServiceMock, nil).Times(1)
			})

			It("terminates the orphaned configuration", func() {
				resources := []*extauthv1.ExtAuthConfig{
					{
						AuthConfigRefName: "default.my-auth-config",
						Configs: []*extauthv1.ExtAuthConfig_Config{
							{
								AuthConfig: &extauthv1.ExtAuthConfig_Config_PluginAuth{
									PluginAuth: plugin,
								},
							},
						},
					},
				}
				cfg, err := generator.GenerateConfig(resources)
				Expect(err).NotTo(HaveOccurred())
				Expect(cfg).NotTo(BeNil())
				Expect(cfg.GetConfigCount()).To(Equal(1))

				// Wait for start function to be called
				var ctx context.Context
				select {
				case <-time.After(time.Second):
					Fail("timed out waiting for Start function to be called")
				case ctx = <-ctxChan:
					// Verify that the context was not cancelled
					Expect(ctx.Err()).To(BeNil())
				}

				// Send no config
				cfg, err = generator.GenerateConfig(nil)
				Expect(err).NotTo(HaveOccurred())
				Expect(cfg).NotTo(BeNil())
				// Resulting config is empty
				Expect(cfg.GetConfigCount()).To(Equal(0))

				// Verify that the previous context was cancelled
				Expect(ctx).NotTo(BeNil())
				if ctx != nil {
					Expect(ctx.Err()).To(Equal(context.Canceled))
				}
			})
		})

		When("we receive an invalid update for an existing config", func() {

			var (
				plugin  = &extauthv1.AuthPlugin{Name: "SomePlugin"}
				ctxChan = make(chan context.Context, 1)
			)

			BeforeEach(func() {
				authServiceMock := chainmocks.NewMockAuthService(ctrl)
				authServiceMock.EXPECT().Start(gomock.Any()).DoAndReturn(func(ctx context.Context) error {
					ctxChan <- ctx
					return nil
				}).Times(1)
				authServiceMock.EXPECT().Authorize(gomock.Any(), gomock.Any()).Times(0)
				pluginLoaderMock.EXPECT().LoadAuthPlugin(gomock.Any(), plugin).
					DoAndReturn(func(ctx context.Context, authPlugin *extauthv1.AuthPlugin) (api.AuthService, error) {
						if authPlugin.PluginFileName != "" {
							ctxChan <- ctx
							return nil, errors.New("some error")
						}
						return authServiceMock, nil
					}).Times(2)
			})

			It("keeps the previous valid configuration", func() {
				resources := []*extauthv1.ExtAuthConfig{
					{
						AuthConfigRefName: "default.my-auth-config",
						Configs: []*extauthv1.ExtAuthConfig_Config{
							{
								AuthConfig: &extauthv1.ExtAuthConfig_Config_PluginAuth{
									PluginAuth: plugin,
								},
							},
						},
					},
				}
				firstCfg, err := generator.GenerateConfig(resources)
				Expect(err).NotTo(HaveOccurred())
				Expect(firstCfg).NotTo(BeNil())
				Expect(firstCfg.GetConfigCount()).To(Equal(1))

				// Wait for start function to be called
				var firstCtx context.Context
				select {
				case <-time.After(time.Second):
					Fail("timed out waiting for Start function to be called")
				case firstCtx = <-ctxChan:
					// Verify that the context was not cancelled
					Expect(firstCtx.Err()).To(BeNil())
				}

				// Update the existing config
				plugin.PluginFileName = "updated.so"

				newCfg, err := generator.GenerateConfig(resources)
				Expect(err).NotTo(HaveOccurred())
				Expect(newCfg).NotTo(BeNil())
				Expect(newCfg.GetConfigCount()).To(Equal(1))
				Expect(newCfg).To(Equal(firstCfg))

				// Wait for start function to be called
				var secondCtx context.Context
				select {
				case <-time.After(time.Second):
					Fail("timed out waiting for second Start function to be called")
				case secondCtx = <-ctxChan:
					// Verify that the context was cancelled
					Expect(secondCtx.Err()).To(Equal(context.Canceled))
				}

				// Verify that the previous context was not cancelled
				Expect(firstCtx).NotTo(BeNil())
				if firstCtx != nil {
					Expect(firstCtx.Err()).To(BeNil())
				}
			})
		})
	})
})
