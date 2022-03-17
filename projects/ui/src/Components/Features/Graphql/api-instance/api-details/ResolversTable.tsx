import { Global } from '@emotion/core';
import { getDirective, MapperKind, mapSchema } from '@graphql-tools/utils';
import { Collapse } from 'antd';
import { useGetGraphqlApiDetails } from 'API/hooks';
import { ReactComponent as GraphQLIcon } from 'assets/graphql-icon.svg';
import { SoloInput } from 'Components/Common/SoloInput';
import { SoloModal } from 'Components/Common/SoloModal';
import {
  EnumTypeDefinitionNode,
  EnumValueDefinitionNode,
  FieldDefinitionNode,
  GraphQLSchema,
  Kind,
  ObjectTypeDefinitionNode,
} from 'graphql';
import gql from 'graphql-tag';
import { ClusterObjectRef } from 'proto/github.com/solo-io/skv2/api/core/v1/core_pb';
import React, { useEffect, useMemo, useState } from 'react';
import { EnumResolver } from './resolver-wizard/EnumResolver';
import { ResolverItem } from './ResolverItem';
import { resolversTableStyles } from './ResolversTable.style';
import { ResolverWizard } from './ResolverWizard';

function defineResolveDirective() {
  let directiveName = 'resolve';
  return {
    mockedDirectiveTypeDefs: `directive @${directiveName}(name: String) on FIELD_DEFINITION  `,
    mockedDirectiveTransformer: (schema: GraphQLSchema) =>
      mapSchema(schema, {
        [MapperKind.OBJECT_FIELD]: fieldConfig => {
          const mockedDirective = getDirective(
            schema,
            fieldConfig,
            directiveName
          )?.[0];
          if (mockedDirective) {
            fieldConfig.deprecationReason = mockedDirective['name'];
            return fieldConfig;
          }
        },
        [MapperKind.ENUM_VALUE]: enumValueConfig => {
          const mockedDirective = getDirective(
            schema,
            enumValueConfig,
            directiveName
          )?.[0];
          if (mockedDirective) {
            enumValueConfig.deprecationReason = mockedDirective['name'];
            return enumValueConfig;
          }
        },
      }),
  };
}

type ResolversTableType = {
  apiRef: ClusterObjectRef.AsObject;
};
const ResolversTable: React.FC<ResolversTableType> = props => {
  const { apiRef } = props;
  const {
    data: graphqlApi,
    error: graphqlApiError,
    mutate,
  } = useGetGraphqlApiDetails({
    name: apiRef.name,
    namespace: apiRef.namespace,
    clusterName: apiRef.clusterName,
  });

  const [currentResolver, setCurrentResolver] = React.useState<any>();
  const [currentResolverName, setCurrentResolverName] = React.useState('');
  const [currentResolverType, setCurrentResolverType] = React.useState('');
  const [hasDirective, setHasDirective] = React.useState(false);
  const [fieldWithDirective, setFieldWithDirective] = React.useState('');
  const [fieldWithoutDirective, setFieldWithoutDirective] = React.useState('');
  const [modalOpen, setModalOpen] = React.useState(false);
  const openModal = () => setModalOpen(true);
  const closeModal = () => setModalOpen(false);

  const listRef = React.useRef<HTMLDivElement>(null);

  const [fieldTypesMap, setFieldTypesMap] = React.useState<
    [string, FieldDefinitionNode[]][]
  >([]);
  const [enumTypesMap, setEnumTypesMap] = React.useState<
    [string, EnumValueDefinitionNode[]][]
  >([]);

  React.useEffect(() => {
    if (graphqlApi) {
      let query = gql`
        ${graphqlApi.spec?.executableSchema?.schemaDefinition}
      `;
      if (query) {
        let objectTypeDefs = query.definitions.filter(
          (def: any) => def.kind === 'ObjectTypeDefinition'
        ) as ObjectTypeDefinitionNode[];
        let enumTypeDefs = query.definitions.filter(
          (def: any) => def.kind === 'EnumTypeDefinition'
        ) as EnumTypeDefinitionNode[];

        const enumFieldDefinitions = enumTypeDefs?.map(ot => [
          `Enum ${ot.name.value}`,
          ot.values?.filter(
            f => f?.kind === 'EnumValueDefinition'
          ) as EnumValueDefinitionNode[],
        ]) as [string, EnumValueDefinitionNode[]][];

        let fieldDefinitions = (
          objectTypeDefs.map(ot => [
            ot.name.value,
            ot.fields?.filter(
              f => f?.kind === 'FieldDefinition'
            ) as FieldDefinitionNode[],
          ]) as [string, FieldDefinitionNode[]][]
        )?.sort(([aTypeName], [bTypeName]) => {
          // Ordering: Query, mutation, Everything else.
          if (aTypeName === 'Query') {
            return -1;
          } else if (bTypeName === 'Query') {
            return 1;
          }
          if (aTypeName === 'Mutation') {
            return -1;
          } else if (bTypeName === 'Mutation') {
            return 1;
          }
          return 0;
        });
        setFieldTypesMap(fieldDefinitions);
        setEnumTypesMap(enumFieldDefinitions);
        // setEnumTypesMap(
        //   // This example can be used to test Enums
        //   [
        //     [
        //       'Enum',
        //       [
        //         {
        //           kind: Kind.ENUM_VALUE_DEFINITION,
        //           description: {
        //             kind: Kind.STRING,
        //             value: 'test description 1',
        //           },
        //           name: { kind: Kind.NAME, value: 'test name 1' },
        //         },
        //         {
        //           kind: Kind.ENUM_VALUE_DEFINITION,
        //           description: {
        //             kind: Kind.STRING,
        //             value: 'test description 2',
        //           },
        //           name: { kind: Kind.NAME, value: 'test name 2' },
        //         },
        //       ],
        //     ],
        //   ]
        // );
      }
    }
  }, [graphqlApi]);

  function handleResolverConfigModal(
    resolverName: string,
    resolverType: string
  ) {
    let [currentResolverName, currentResolver] =
      graphqlApi?.spec?.executableSchema?.executor?.local?.resolutionsMap.find(
        ([rName, resolver]) => rName.includes(resolverName)
      ) ?? ['', ''];
    setCurrentResolver(currentResolver);
    setCurrentResolverName(resolverName);
    setCurrentResolverType(resolverType);

    let isListType = Object.fromEntries(fieldTypesMap)[resolverType]?.some(
      f => f.name.value === resolverName && f.type.kind === 'ListType'
    );

    let fieldType = isListType
      ? fieldTypesMap
          .find(([t, f]) => t === resolverType)?.[1]
          //@ts-ignore
          ?.find(f => f.name.value === resolverName)?.type?.type?.name?.value
      : fieldTypesMap
          .find(([t, f]) => t === resolverType)?.[1]
          //@ts-ignore
          ?.find(f => f.name.value === resolverName)?.type?.name?.value;

    let fieldWithDirective = '';
    let fieldWithoutDirective = '';
    if (isListType) {
      fieldWithoutDirective = `${resolverName}: [${fieldType}]`;
      fieldWithDirective = `${resolverName}: [${fieldType}] @resolve(name: "${resolverName}")`;
    } else {
      fieldWithoutDirective = `${resolverName}: ${fieldType}`;
      fieldWithDirective = `${resolverName}: ${fieldType} @resolve(name: "${resolverName}")`;
    }

    setHasDirective(
      !!graphqlApi?.spec?.executableSchema?.schemaDefinition.includes(
        fieldWithDirective
      )
    );
    setFieldWithDirective(fieldWithDirective);
    setFieldWithoutDirective(fieldWithoutDirective);

    openModal();
  }

  // --- SEARCH LOGIC --- //
  // Field + enums should probably be combined.
  const [fieldTypesMapCopy, setFieldTypesMapCopy] = useState(fieldTypesMap);
  const [enumTypesMapCopy, setEnumTypesMapCopy] = useState(enumTypesMap);
  const [searchText, setSearchText] = useState('');
  useEffect(() => {
    const lstext = searchText.toLowerCase();
    // field types map
    const newFieldTypesMapCopy = [] as typeof fieldTypesMapCopy;
    fieldTypesMap.forEach(([typename, fields]) => {
      if (
        typeof typename === 'string' &&
        typename.toLowerCase().includes(lstext)
      ) {
        newFieldTypesMapCopy.push([typename, fields]);
      } else {
        const newFields = fields.filter(f =>
          f.name.value.toLowerCase().includes(lstext)
        );
        if (newFields.length > 0)
          newFieldTypesMapCopy.push([typename, newFields]);
      }
    });
    setFieldTypesMapCopy(newFieldTypesMapCopy);
    // enum types map
    const newEnumTypesMapCopy = [] as typeof enumTypesMapCopy;
    enumTypesMap.forEach(([typename, fields]) => {
      if (
        typeof typename === 'string' &&
        typename.toLowerCase().includes(lstext)
      ) {
        newEnumTypesMapCopy.push([typename, fields]);
      } else {
        const newFields = fields.filter(f =>
          f.name.value.toLowerCase().includes(lstext)
        );
        if (newFields.length > 0)
          newEnumTypesMapCopy.push([typename, newFields]);
      }
    });
    setEnumTypesMapCopy(newEnumTypesMapCopy);
  }, [searchText, enumTypesMap, fieldTypesMap]);

  const defaultActivePanelKey = useMemo(() => {
    return fieldTypesMapCopy?.length > 0
      ? [`${apiRef.namespace}-${apiRef.name}-${fieldTypesMapCopy[0][0]}`]
      : enumTypesMapCopy?.length > 0
      ? [enumTypesMapCopy[0][0]]
      : undefined;
  }, [fieldTypesMapCopy, enumTypesMapCopy]);

  return (
    <div ref={listRef} className='relative'>
      <Global styles={resolversTableStyles} />

      <div className='max-w-[500px] mb-5'>
        <SoloInput
          placeholder='Filter by name...'
          value={searchText}
          onChange={s => setSearchText(s.target.value)}
        />
      </div>

      {defaultActivePanelKey !== undefined && (
        <Collapse defaultActiveKey={defaultActivePanelKey}>
          {fieldTypesMapCopy.map(([typeName, fields]) => {
            return (
              <Collapse.Panel
                key={`${apiRef.namespace}-${apiRef.name}-${typeName}`}
                header={
                  <div className='inline font-medium text-gray-900 whitespace-nowrap'>
                    <GraphQLIcon className='w-4 h-4 fill-current inline' />
                    &nbsp;&nbsp;
                    {typeName}
                  </div>
                }>
                <ResolverItem
                  resolverType={typeName}
                  fields={fields}
                  handleResolverConfigModal={handleResolverConfigModal}
                />
              </Collapse.Panel>
            );
          })}
          {enumTypesMapCopy.map(([typeName, fields]) => {
            return (
              <Collapse.Panel
                key={typeName}
                header={
                  <div className='inline font-medium text-gray-900 whitespace-nowrap'>
                    <GraphQLIcon className='w-4 h-4 fill-current inline' />
                    &nbsp;&nbsp;
                    {typeName}
                  </div>
                }>
                <EnumResolver fields={fields} resolverType={typeName} />
              </Collapse.Panel>
            );
          })}
        </Collapse>
      )}

      <SoloModal visible={modalOpen} width={750} onClose={closeModal}>
        <ResolverWizard
          resolver={currentResolver}
          hasDirective={hasDirective}
          fieldWithDirective={fieldWithDirective}
          fieldWithoutDirective={fieldWithoutDirective}
          resolverName={currentResolverName}
          onClose={() => {
            closeModal();
            mutate();
          }}
        />
      </SoloModal>
    </div>
  );
};

export default ResolversTable;
