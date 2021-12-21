package graphql_test

import (
	"fmt"

	"google.golang.org/protobuf/types/known/structpb"

	openapi "github.com/getkin/kin-openapi/openapi3"
	"github.com/ghodss/yaml"
	"github.com/golang/protobuf/ptypes"
	gographql "github.com/graphql-go/graphql"
	v2 "github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/extensions/filters/http/graphql/v2"
	. "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/graphql/v1alpha1"
	"github.com/solo-io/gloo/projects/gloo/pkg/utils"
	"github.com/solo-io/solo-kit/test/matchers"
	graphql2 "github.com/solo-io/solo-projects/projects/discovery/pkg/fds/discoveries/openapi/graphqlschematranslation"
	"github.com/solo-io/solo-projects/projects/discovery/pkg/fds/discoveries/openapi/printer"
	schemas "github.com/solo-io/solo-projects/projects/gloo/pkg/plugins/graphql"

	envoy_config_route_v3 "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/plugins"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

var _ = Describe("Graphql plugin", func() {
	var (
		plugin        *schemas.Plugin
		params        plugins.Params
		vhostParams   plugins.VirtualHostParams
		virtualHost   *v1.VirtualHost
		route         *v1.Route
		routeAction   *v1.Route_GraphqlSchemaRef
		httpListener  *v1.HttpListener
		gqlSchemaSpec *GraphQLSchema
	)

	BeforeEach(func() {
		routeAction = &v1.Route_GraphqlSchemaRef{
			GraphqlSchemaRef: &core.ResourceRef{
				Name:      "gql",
				Namespace: "gloo-system",
			},
		}
		route = &v1.Route{
			Action: routeAction,
		}

		gqlSchemaSpec = &GraphQLSchema{
			Metadata: &core.Metadata{
				Name:      "gql",
				Namespace: "gloo-system",
			},
			ExecutableSchema: &ExecutableSchema{
				Executor: &Executor{
					Executor: &Executor_Local_{
						Local: &Executor_Local{},
					},
				},
			},
		}
	})

	JustBeforeEach(func() {
		virtualHost = &v1.VirtualHost{
			Name:    "virt1",
			Domains: []string{"*"},
			Routes:  []*v1.Route{route},
		}

		httpListener = &v1.HttpListener{
			VirtualHosts: []*v1.VirtualHost{virtualHost},
		}
		proxy := &v1.Proxy{
			Metadata: &core.Metadata{
				Name:      "http-proxy",
				Namespace: "default",
			},
			Listeners: []*v1.Listener{{
				Name: "default",
				ListenerType: &v1.Listener_HttpListener{
					HttpListener: httpListener,
				},
			}},
		}

		params.Snapshot = &v1.ApiSnapshot{
			Proxies: v1.ProxyList{proxy},
			GraphqlSchemas: GraphQLSchemaList{
				gqlSchemaSpec,
			},
			Upstreams: v1.UpstreamList{
				{
					Metadata: &core.Metadata{
						Name:      "us",
						Namespace: "gloo-system",
					},
				},
			},
		}
		vhostParams = plugins.VirtualHostParams{
			Params:   params,
			Proxy:    proxy,
			Listener: proxy.Listeners[0],
		}
	})

	Context("process snapshot", func() {
		var (
			outRoute   envoy_config_route_v3.Route
			outFilters []plugins.StagedHttpFilter
		)

		var translateRoute = func() *v2.GraphQLRouteConfig {
			goTpfc := outRoute.TypedPerFilterConfig[schemas.FilterName]
			Expect(goTpfc).NotTo(BeNil())
			var perRouteGql v2.GraphQLRouteConfig
			err := ptypes.UnmarshalAny(goTpfc, &perRouteGql)
			Expect(err).NotTo(HaveOccurred())
			return &perRouteGql
		}

		JustBeforeEach(func() {
			outRoute = envoy_config_route_v3.Route{}
			routesParams := plugins.RouteParams{
				VirtualHostParams: vhostParams,
				VirtualHost:       virtualHost,
			}
			// run it like the translator:
			err := plugin.ProcessRoute(routesParams, route, &outRoute)
			Expect(err).NotTo(HaveOccurred())
			outFilters, err = plugin.HttpFilters(params, httpListener)
			Expect(err).NotTo(HaveOccurred())
		})

		BeforeEach(func() {
			plugin = schemas.NewPlugin()
			err := plugin.Init(plugins.InitParams{})
			Expect(err).NotTo(HaveOccurred())
		})

		Context("empty extensions", func() {
			It("can create the proper filters", func() {
				Expect(outFilters).To(HaveLen(1))
				gqlFilter := outFilters[0]
				Expect(gqlFilter.HttpFilter.Name).To(Equal(schemas.FilterName))
				Expect(gqlFilter.Stage).To(Equal(plugins.BeforeStage(plugins.RouteStage)))
				st := gqlFilter.HttpFilter.GetTypedConfig()
				// graphql is always added to HCM, only routes with graphql route config will use the empty config
				Expect(st).NotTo(BeNil())
			})
		})

		Context("per route/vhost", func() {

			Context("nil", func() {
				BeforeEach(func() {
					route.Action = &v1.Route_RouteAction{
						RouteAction: nil,
					}
				})

				It("is disabled on routes by default", func() {
					pfc := outRoute.TypedPerFilterConfig[schemas.FilterName]
					Expect(pfc).To(BeNil())
				})
			})

			Context("translate route config", func() {
				BeforeEach(func() {
					gqlSchemaSpec.ExecutableSchema.Executor.GetLocal().EnableIntrospection = true
				})

				It("sets enable introspection", func() {
					perRouteGql := translateRoute()
					Expect(perRouteGql.GetExecutableSchema().GetExecutor().GetLocal().GetEnableIntrospection()).To(BeTrue())
				})

				Context("translate resolutions", func() {

					BeforeEach(func() {
						body := `{"k1": {"k2": "val"}}`
						bodyStruct := &structpb.Value{}
						err := yaml.Unmarshal([]byte(body), bodyStruct)
						Expect(err).NotTo(HaveOccurred())
						gqlSchemaSpec.ExecutableSchema.Executor.GetLocal().Resolutions = []*Resolution{
							{
								Matcher: &QueryMatcher{
									Match: &QueryMatcher_FieldMatcher_{
										FieldMatcher: &QueryMatcher_FieldMatcher{
											Type:  "type",
											Field: "field",
										},
									},
								},
								Resolver: &Resolution_RestResolver{
									RestResolver: &RESTResolver{
										UpstreamRef: &core.ResourceRef{
											Name:      "us",
											Namespace: "gloo-system",
										},
										Request: &RequestTemplate{
											Headers: map[string]string{
												"header": "7.89",
											},
											QueryParams: map[string]string{
												"qp": "true",
											},
											Body: bodyStruct,
										},
										SpanName: "span",
									},
								},
							},
						}
					})

					It("sets resolvers", func() {
						perRouteGql := translateRoute()
						resolutions := perRouteGql.GetExecutableSchema().GetExecutor().GetLocal().GetResolutions()
						Expect(resolutions[0].Matcher.GetFieldMatcher().GetType()).To(Equal("type"))
						Expect(resolutions[0].Matcher.GetFieldMatcher().GetField()).To(Equal("field"))

						any := resolutions[0].GetResolver()
						Expect(any).NotTo(BeNil())
						msg, err := utils.AnyToMessage(any.TypedConfig)
						Expect(err).NotTo(HaveOccurred())
						restResolver := msg.(*v2.RESTResolver)

						Expect(restResolver.GetSpanName()).To(Equal("span"))
						Expect(restResolver.GetRequestTransform().GetHeaders()["header"].GetProviders()["ARBITRARY_PROVIDER_NAME"].GetTypedProvider().GetValue()).To(Equal("7.89"))
						Expect(restResolver.GetRequestTransform().GetQueryParams()["qp"].GetProviders()["ARBITRARY_PROVIDER_NAME"].GetTypedProvider().GetValue()).To(Equal("true"))

						// testing the recursive translation
						Expect(restResolver.GetRequestTransform().GetOutgoingBody().GetNode().GetKeyValues()[0].GetKey()).To(Equal("k1"))
						Expect(restResolver.GetRequestTransform().GetOutgoingBody().GetNode().GetKeyValues()[0].GetValue().GetNode().GetKeyValues()[0].GetKey()).To(Equal("k2"))
						Expect(restResolver.GetRequestTransform().GetOutgoingBody().GetNode().GetKeyValues()[0].GetValue().GetNode().GetKeyValues()[0].GetValue().GetValueProvider().GetProviders()["ARBITRARY_PROVIDER_NAME"].GetTypedProvider().GetValue()).To(Equal("val"))
					})

				})

			})

		})
		Context("graphql translation", func() {

			var (
				graphqlSchema *gographql.Schema
				resolutions   []*Resolution
			)
			AfterEach(func() {
				graphqlSchema = nil
				resolutions = nil
			})
			translateToSchema := func(openapiSchema string) {
				t := graphql2.NewOasToGqlTranslator(&v1.Upstream{
					Metadata: &core.Metadata{
						Name:      "hi",
						Namespace: "default",
					},
				})
				l := openapi.NewLoader()
				l.IsExternalRefsAllowed = true
				spec, err := l.LoadFromData([]byte(openapiSchema))
				ExpectWithOffset(1, err).NotTo(HaveOccurred())
				oass := []*openapi.T{spec}
				graphqlSchema, resolutions = t.CreateGraphqlSchema(oass)
				fmt.Println(printer.PrintFilteredSchema(graphqlSchema))
				crd := &GraphQLSchema{
					ExecutableSchema: &ExecutableSchema{
						Executor: &Executor{
							Executor: &Executor_Local_{
								Local: &Executor_Local{
									Resolutions: resolutions,
								},
							},
						},
						SchemaDefinition: printer.PrintFilteredSchema(graphqlSchema),
					},
				}
				b, err := yaml.Marshal(crd)
				ExpectWithOffset(1, err).NotTo(HaveOccurred())
				fmt.Println(string(b))
			}

			It("Handles links", func() {
				schemaToTranslate := schemas.GetSimpleLinkSchema()
				translateToSchema(schemaToTranslate)

				fields := graphqlSchema.QueryType().Fields()
				// Expect getEmployeeById query field to be created
				Expect(fields).To(HaveKey("getEmployeeById"))
				employeeByIdQueryField := fields["getEmployeeById"]
				Expect(employeeByIdQueryField.Args).To(HaveLen(1))
				Expect(employeeByIdQueryField.Args[0].Type).To(Equal(gographql.NewNonNull(gographql.String)))
				Expect(employeeByIdQueryField.Type.String()).To(Equal("Employee"))

				// Employee type should have fields, and userManager link should be resolved.
				employeeType := graphqlSchema.Type("Employee").(*gographql.Object)
				Expect(employeeType).To(Not(BeNil()))
				Expect(employeeType.Fields()).To(HaveLen(4))
				Expect(employeeType.Fields()).To(HaveKey("userManager"))
				userManagerField := employeeType.Fields()["userManager"]
				Expect(userManagerField.Type.String()).To(Equal("Employee"))

				// Resolvers should exist for Query.getEmployeeById and Employee.userManager
				Expect(resolutions).To(HaveLen(2))
				Expect(resolutions).To(ContainElement(matchers.MatchProto(&Resolution{
					Matcher: &QueryMatcher{
						Match: &QueryMatcher_FieldMatcher_{
							FieldMatcher: &QueryMatcher_FieldMatcher{
								Type:  "Query",
								Field: "getEmployeeById",
							},
						},
					},
					Resolver: &Resolution_RestResolver{
						RestResolver: &RESTResolver{
							UpstreamRef: &core.ResourceRef{
								Name:      "hi",
								Namespace: "default",
							},
							Request: &RequestTemplate{
								Headers: map[string]string{
									":method": "GET",
									":path":   "/employees/{$args.id}",
								},
							},
						},
					},
				})))
				Expect(resolutions).To(ContainElement(matchers.MatchProto(&Resolution{
					Matcher: &QueryMatcher{
						Match: &QueryMatcher_FieldMatcher_{
							FieldMatcher: &QueryMatcher_FieldMatcher{
								Type:  "Employee",
								Field: "userManager",
							},
						},
					},
					Resolver: &Resolution_RestResolver{
						RestResolver: &RESTResolver{
							UpstreamRef: &core.ResourceRef{
								Name:      "hi",
								Namespace: "default",
							},
							Request: &RequestTemplate{
								Headers: map[string]string{
									":method": "GET",
									":path":   "/employees/{$parent.managerId}",
								},
							},
						},
					},
				})))
			})

		})

	})
})
