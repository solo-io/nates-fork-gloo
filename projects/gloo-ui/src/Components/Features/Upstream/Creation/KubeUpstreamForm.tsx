import {
  SoloFormInput,
  SoloFormTypeahead
} from 'Components/Common/Form/SoloFormField';
import {
  InputRow,
  SoloFormTemplate
} from 'Components/Common/Form/SoloFormTemplate';
import * as React from 'react';
import { configAPI } from 'store/config/api';
import useSWR from 'swr';
import * as yup from 'yup';

// TODO: handle servicespec and subset spec
export interface KubeValuesType {
  kubeServiceName: string;
  kubeServiceNamespace: string;
  kubeServicePort: number;
}

export const kubeInitialValues: KubeValuesType = {
  kubeServiceName: '',
  kubeServiceNamespace: 'gloo-system',
  kubeServicePort: 8080
};

interface Props {}

// TODO: figure out which fields are required
export const kubeValidationSchema = yup.object().shape({
  kubeServiceName: yup.string(),
  kubeServiceNamespace: yup.string(),
  kubeServicePort: yup.number()
});

export const KubeUpstreamForm: React.FC<Props> = () => {
  const { data: namespacesList, error: listNamespacesError } = useSWR(
    'listNamespaces',
    configAPI.listNamespaces
  );
  const { data: podNamespace, error: podNamespaceError } = useSWR(
    'getPodNamespace',
    configAPI.getPodNamespace
  );

  if (!podNamespace) {
    return <div>Loading...</div>;
  }
  return (
    <SoloFormTemplate formHeader='Kubernetes Upstream Settings'>
      <InputRow>
        <div>
          <SoloFormInput
            name='kubeServiceName'
            title='Service Name'
            placeholder='Service Name'
          />
        </div>
        <div>
          <SoloFormTypeahead
            name='kubeServiceNamespace'
            title='Service Namespace'
            defaultValue={podNamespace}
            presetOptions={(namespacesList ?? []).map(ns => {
              return { value: ns };
            })}
          />
        </div>
        <div>
          <SoloFormInput
            name='kubeServicePort'
            title='Service Port'
            placeholder='Service Port'
            type='number'
          />
        </div>
      </InputRow>
    </SoloFormTemplate>
  );
};
