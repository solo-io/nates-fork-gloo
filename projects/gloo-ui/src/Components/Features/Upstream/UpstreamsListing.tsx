import * as React from 'react';
/** @jsx jsx */
import { jsx } from '@emotion/core';

import styled from '@emotion/styled/macro';
import { withRouter, RouteComponentProps } from 'react-router';
import { colors } from 'Styles';
import {
  ListingFilter,
  StringFilterProps,
  TypeFilterProps,
  CheckboxFilterProps,
  RadioFilterProps
} from '../../Common/ListingFilter';
import { CatalogTableToggle } from 'Components/Common/CatalogTableToggle';
import { Breadcrumb } from 'Components/Common/Breadcrumb';

const StringFilters: StringFilterProps[] = [
  {
    displayName: 'Filter By Name...',
    placeholder: 'Filter by name...',
    value: ''
  }
];

const CheckboxFilters: CheckboxFilterProps[] = [
  {
    displayName: 'AWS',
    value: false
  },
  {
    displayName: 'Azure',
    value: false
  },
  {
    displayName: 'REST',
    value: false
  },
  {
    displayName: 'gRPC',
    value: false
  },
  {
    displayName: 'Kubernetes',
    value: false
  },
  {
    displayName: 'Static',
    value: false
  }
];

const Heading = styled.div`
  display: flex;
  justify-content: space-between;
  margin-bottom: 20px;
`;

interface Props extends RouteComponentProps {
  //... eg, virtualservice?: string
}

export const UpstreamsListing = (props: Props) => {
  const [catalogNotTable, setCatalogNotTable] = React.useState<boolean>(true);

  const listDisplay = (
    strings: StringFilterProps[],
    types: TypeFilterProps[],
    checkboxes: CheckboxFilterProps[],
    radios: RadioFilterProps[]
  ) => {
    return (
      <div>
        {strings.map(fil => {
          return (
            <div key={fil.displayName}>
              <span>{fil.displayName}</span>
              <span>{fil.value}</span>
            </div>
          );
        })}
        {checkboxes.map(fil => {
          return (
            <div key={fil.displayName}>
              <span>{fil.displayName}</span>
              <span>{!!fil.value ? 'true' : 'false'}</span>
            </div>
          );
        })}
      </div>
    );
  };

  return (
    <div>
      <Heading>
        <Breadcrumb />
        <CatalogTableToggle
          listIsSelected={!catalogNotTable}
          onToggle={() => {
            setCatalogNotTable(cNt => !cNt);
          }}
        />
      </Heading>
      <ListingFilter
        strings={StringFilters}
        checkboxes={CheckboxFilters}
        filterFunction={listDisplay}
      />
    </div>
  );
};
