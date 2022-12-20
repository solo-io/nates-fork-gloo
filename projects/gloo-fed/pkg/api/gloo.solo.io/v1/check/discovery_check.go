// Code generated by skv2. DO NOT EDIT.

package check

import (
	"context"

	"github.com/solo-io/go-utils/stringutils"
	sk_sets "github.com/solo-io/skv2/contrib/pkg/sets/v2"
	corev1 "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
	"github.com/solo-io/skv2/pkg/ezkube"
	gloo_solo_io_v1 "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1"
	types2 "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1"
	"github.com/solo-io/solo-projects/projects/gloo-fed/pkg/api/fed.solo.io/v1/types"
	"github.com/solo-io/solo-projects/projects/gloo-fed/pkg/discovery/translator/summarize"
)

func GetUpstreamSummary(ctx context.Context, set sk_sets.ResourceSet[*gloo_solo_io_v1.Upstream], watchedNamespaces []string, cluster string) *types.GlooInstanceSpec_Check_Summary {
	summary := &types.GlooInstanceSpec_Check_Summary{}

	for _, upstreamIter := range set.List() {
		upstream := upstreamIter

		// If the resource is not in the right cluster, continue
		if ezkube.GetClusterName(upstream) != cluster {
			continue
		}

		// If the resource is not in a watched namespace, continue
		if len(watchedNamespaces) > 0 && !stringutils.ContainsString(upstream.Namespace, watchedNamespaces) {
			continue
		}

		summary.Total += 1

		if upstream.Status.GetState() == types2.UpstreamStatus_Rejected {
			summary.Errors = append(summary.Errors, &types.GlooInstanceSpec_Check_Summary_ResourceReport{
				Ref: &corev1.ObjectRef{
					Name:      upstream.Name,
					Namespace: upstream.Namespace,
				},
				Message: upstream.Status.Reason,
			})
		}

		if upstream.Status.GetState() == types2.UpstreamStatus_Warning {
			summary.Warnings = append(summary.Warnings, &types.GlooInstanceSpec_Check_Summary_ResourceReport{
				Ref: &corev1.ObjectRef{
					Name:      upstream.Name,
					Namespace: upstream.Namespace,
				},
				Message: upstream.Status.Reason,
			})
		}

	}

	summarize.SortLists(summary)
	return summary
}

func GetUpstreamGroupSummary(ctx context.Context, set sk_sets.ResourceSet[*gloo_solo_io_v1.UpstreamGroup], watchedNamespaces []string, cluster string) *types.GlooInstanceSpec_Check_Summary {
	summary := &types.GlooInstanceSpec_Check_Summary{}

	for _, upstreamGroupIter := range set.List() {
		upstreamGroup := upstreamGroupIter

		// If the resource is not in the right cluster, continue
		if ezkube.GetClusterName(upstreamGroup) != cluster {
			continue
		}

		// If the resource is not in a watched namespace, continue
		if len(watchedNamespaces) > 0 && !stringutils.ContainsString(upstreamGroup.Namespace, watchedNamespaces) {
			continue
		}

		summary.Total += 1

		if upstreamGroup.Status.GetState() == types2.UpstreamGroupStatus_Rejected {
			summary.Errors = append(summary.Errors, &types.GlooInstanceSpec_Check_Summary_ResourceReport{
				Ref: &corev1.ObjectRef{
					Name:      upstreamGroup.Name,
					Namespace: upstreamGroup.Namespace,
				},
				Message: upstreamGroup.Status.Reason,
			})
		}

		if upstreamGroup.Status.GetState() == types2.UpstreamGroupStatus_Warning {
			summary.Warnings = append(summary.Warnings, &types.GlooInstanceSpec_Check_Summary_ResourceReport{
				Ref: &corev1.ObjectRef{
					Name:      upstreamGroup.Name,
					Namespace: upstreamGroup.Namespace,
				},
				Message: upstreamGroup.Status.Reason,
			})
		}

	}

	summarize.SortLists(summary)
	return summary
}

func GetSettingsSummary(ctx context.Context, set sk_sets.ResourceSet[*gloo_solo_io_v1.Settings], watchedNamespaces []string, cluster string) *types.GlooInstanceSpec_Check_Summary {
	summary := &types.GlooInstanceSpec_Check_Summary{}

	for _, settingsIter := range set.List() {
		settings := settingsIter

		// If the resource is not in the right cluster, continue
		if ezkube.GetClusterName(settings) != cluster {
			continue
		}

		// If the resource is not in a watched namespace, continue
		if len(watchedNamespaces) > 0 && !stringutils.ContainsString(settings.Namespace, watchedNamespaces) {
			continue
		}

		summary.Total += 1

		if settings.Status.GetState() == types2.SettingsStatus_Rejected {
			summary.Errors = append(summary.Errors, &types.GlooInstanceSpec_Check_Summary_ResourceReport{
				Ref: &corev1.ObjectRef{
					Name:      settings.Name,
					Namespace: settings.Namespace,
				},
				Message: settings.Status.Reason,
			})
		}

		if settings.Status.GetState() == types2.SettingsStatus_Warning {
			summary.Warnings = append(summary.Warnings, &types.GlooInstanceSpec_Check_Summary_ResourceReport{
				Ref: &corev1.ObjectRef{
					Name:      settings.Name,
					Namespace: settings.Namespace,
				},
				Message: settings.Status.Reason,
			})
		}

	}

	summarize.SortLists(summary)
	return summary
}

func GetProxySummary(ctx context.Context, set sk_sets.ResourceSet[*gloo_solo_io_v1.Proxy], watchedNamespaces []string, cluster string) *types.GlooInstanceSpec_Check_Summary {
	summary := &types.GlooInstanceSpec_Check_Summary{}

	for _, proxyIter := range set.List() {
		proxy := proxyIter

		// If the resource is not in the right cluster, continue
		if ezkube.GetClusterName(proxy) != cluster {
			continue
		}

		// If the resource is not in a watched namespace, continue
		if len(watchedNamespaces) > 0 && !stringutils.ContainsString(proxy.Namespace, watchedNamespaces) {
			continue
		}

		summary.Total += 1

		if proxy.Status.GetState() == types2.ProxyStatus_Rejected {
			summary.Errors = append(summary.Errors, &types.GlooInstanceSpec_Check_Summary_ResourceReport{
				Ref: &corev1.ObjectRef{
					Name:      proxy.Name,
					Namespace: proxy.Namespace,
				},
				Message: proxy.Status.Reason,
			})
		}

		if proxy.Status.GetState() == types2.ProxyStatus_Warning {
			summary.Warnings = append(summary.Warnings, &types.GlooInstanceSpec_Check_Summary_ResourceReport{
				Ref: &corev1.ObjectRef{
					Name:      proxy.Name,
					Namespace: proxy.Namespace,
				},
				Message: proxy.Status.Reason,
			})
		}

	}

	summarize.SortLists(summary)
	return summary
}
