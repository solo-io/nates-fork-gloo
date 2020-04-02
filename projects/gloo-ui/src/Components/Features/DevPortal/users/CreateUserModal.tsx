import React from 'react';
import {
  Tabs,
  TabList,
  Tab,
  TabPanels,
  TabPanel,
  TabPanelProps
} from '@reach/tabs';
import { ReactComponent as VewIcon } from 'assets/view-icon-blue.svg';
import { css } from '@emotion/core';
import { Formik } from 'formik';
import {
  SectionContainer,
  SectionHeader,
  SectionSubHeader
} from '../apis/CreateAPIModal';
import {
  SoloFormInput,
  SoloFormCheckbox
} from 'Components/Common/Form/SoloFormField';
import {
  SoloButtonStyledComponent,
  SoloCancelButton
} from 'Styles/CommonEmotions/button';
import { SoloTransfer, ListItemType } from 'Components/Common/SoloTransfer';
import useSWR from 'swr';
import { apiDocApi, portalApi, userApi, groupApi } from '../api';
import { Loading } from 'Components/Common/DisplayOnly/Loading';
import { ObjectRef } from 'proto/dev-portal/api/dev-portal/v1/common_pb';
import { User } from 'proto/dev-portal/api/grpc/admin/user_pb';
import { configAPI } from 'store/config/api';
import { Group } from 'proto/dev-portal/api/grpc/admin/group_pb';

const StyledTab = (
  props: {
    disabled?: boolean | undefined;
  } & TabPanelProps & {
      isSelected?: boolean | undefined;
    }
) => {
  const { isSelected, children } = props;
  return (
    <Tab
      {...props}
      className={`p-1 text-left w-48 text-white  pl-6 mb-2 focus:outline-none ${
        isSelected
          ? ' bg-blue-500 border-r-8 border-blue-300  '
          : 'bg-blue-600 '
      }`}>
      {children}
    </Tab>
  );
};

type GeneralSectionProps = {
  isEdit: boolean;
};

const GeneralSection = ({ isEdit }: GeneralSectionProps) => {
  const [showPassword, setShowPassword] = React.useState(true);
  return (
    <SectionContainer>
      <SectionHeader> {isEdit ? 'Edit User' : 'Create a User'}</SectionHeader>
      <div className='p-3 mb-2 text-gray-700 bg-gray-100 rounded-lg'>
        {isEdit ? 'Edit an' : 'Create a new'} API User
      </div>
      <div className='grid grid-flow-col grid-cols-5 grid-rows-3 gap-2'>
        <div className='col-span-3 mr-4'>
          <SoloFormInput
            name='name'
            title='Name'
            placeholder='Username'
            disabled={isEdit}
            hideError
          />
        </div>
        <div className='col-span-3 mr-4'>
          <SoloFormInput
            name='email'
            title='Email'
            placeholder='email@domain.com'
            disabled={isEdit}
            hideError
          />
        </div>
        <div className='relative col-span-3 mr-4'>
          <span
            className='absolute cursor-pointer bottom-2 right-4'
            onClick={() => setShowPassword(!showPassword)}>
            <VewIcon />
          </span>
          <SoloFormInput
            type={showPassword ? 'password' : 'text'}
            name='password'
            title='Password'
            placeholder='password'
            disabled={isEdit}
            hideError
          />
        </div>
      </div>
    </SectionContainer>
  );
};

const getGroups = (
  groupList: Group.AsObject[],
  { namespace, name }: ObjectRef.AsObject
): ListItemType[] => {
  return groupList.reduce((acc: ListItemType[], g) => {
    let userExists = !!g.status?.usersList.find(
      ref => ref.namespace === namespace && ref.name === name
    );
    if (userExists) {
      return [
        {
          namespace: g.metadata!.namespace,
          name: g.metadata!.name,
          displayValue: g.spec?.displayName || g.metadata?.name
        },
        ...acc
      ];
    }
    return acc;
  }, []);
};

type CreateUserValues = {
  name: string;
  email: string;
  password: string;
  chosenAPIs: ListItemType[];
  chosenPortals: ListItemType[];
  chosenGroups: ListItemType[];
};

export const CreateUserModal: React.FC<{
  onClose: () => void;
  existingUser?: User.AsObject;
}> = props => {
  const { data: portalsList, error: portalsListError } = useSWR(
    'listPortals',
    portalApi.listPortals
  );
  const { data: apiDocsList, error: apiDocsError } = useSWR(
    'listApiDocs',
    apiDocApi.listApiDocs
  );
  const { data: groupsList, error: groupsError } = useSWR(
    'listGroups',
    groupApi.listGroups
  );

  const { data: podNamespace, error: podNamespaceError } = useSWR(
    'getPodNamespace',
    configAPI.getPodNamespace
  );

  const [tabIndex, setTabIndex] = React.useState(0);
  const handleTabsChange = (index: number) => {
    setTabIndex(index);
  };

  // TODO make this a handle write that knows what to do if an existing user is here
  const handleWriteUser = async (values: CreateUserValues) => {
    const {
      chosenAPIs,
      name,
      chosenPortals,
      email,
      chosenGroups,
      password
    } = values;

    if (!props.existingUser) {
      let newUser = new User().toObject();

      await userApi.createUser({
        user: {
          ...newUser!,
          metadata: {
            ...newUser.metadata!,
            name,
            namespace: podNamespace!
          },
          spec: {
            email,
            username: name
          }
        },
        password,
        portalsList: chosenPortals,
        groupsList: chosenGroups,
        apiDocsList: chosenAPIs,
        userOnly: false
      });
    } else {
      userApi.updateUser(
        {
          namespace: props.existingUser.metadata!.namespace,
          name: props.existingUser.metadata!.name
        },
        chosenAPIs,
        chosenGroups,
        chosenPortals
      );
    }
    props.onClose();
  };

  if (!apiDocsList || !portalsList || !groupsList) {
    return <Loading center>Loading...</Loading>;
  }

  const getInitialGroups = (): ListItemType[] =>
    !!props.existingUser
      ? getGroups(groupsList, {
          namespace: props.existingUser.metadata!.namespace,
          name: props.existingUser.metadata!.name
        })
      : [];

  const getInitialApis = (): ListItemType[] => {
    if (!props.existingUser || !props.existingUser.status?.accessLevel) {
      return [];
    }

    return props.existingUser.status.accessLevel.apiDocsList.map(
      ({ namespace, name }) => {
        let matchingDoc = apiDocsList.find(
          doc =>
            doc.metadata?.namespace === namespace && doc.metadata.name === name
        );
        return {
          namespace,
          name,
          displayValue: matchingDoc?.status?.displayName
        };
      }
    );
  };

  const getInitialPortals = (): ListItemType[] => {
    if (!props.existingUser || !props.existingUser.status?.accessLevel) {
      return [];
    }

    return props.existingUser.status.accessLevel.portalsList.map(
      ({ namespace, name }) => {
        let matchingPortal = portalsList.find(
          portal =>
            portal.metadata?.namespace === namespace &&
            portal.metadata.name === name
        );

        return {
          namespace,
          name,
          displayValue: matchingPortal?.spec?.displayName
        };
      }
    );
  };

  return (
    <>
      <div
        css={css`
          width: 750px;
        `}
        className='bg-white rounded-lg shadow '>
        <Formik<CreateUserValues>
          initialValues={{
            name: props.existingUser?.spec?.username || '',
            email: props.existingUser?.spec?.email || '',
            password: !!props.existingUser ? '*********' : '',
            chosenGroups: getInitialGroups(),
            chosenAPIs: getInitialApis(),
            chosenPortals: getInitialPortals()
          }}
          onSubmit={handleWriteUser}>
          {formik => (
            <>
              <Tabs
                className='bg-blue-600 rounded-lg h-96'
                index={tabIndex}
                onChange={handleTabsChange}
                css={css`
                  display: grid;
                  height: 450px;
                  grid-template-columns: 190px 1fr;
                `}>
                <TabList className='flex flex-col mt-6'>
                  <StyledTab>General</StyledTab>
                  <StyledTab>Groups</StyledTab>

                  <StyledTab>APIs</StyledTab>
                  <StyledTab>Portals</StyledTab>
                </TabList>

                <TabPanels className='bg-white rounded-r-lg'>
                  <TabPanel className='relative flex flex-col justify-between h-full focus:outline-none'>
                    <GeneralSection isEdit={!!props.existingUser} />
                    <div className='flex items-end justify-between h-full px-6 mb-4 '>
                      <button
                        className='text-blue-500 cursor-pointer'
                        onClick={props.onClose}>
                        cancel
                      </button>
                      <SoloButtonStyledComponent
                        onClick={() => setTabIndex(tabIndex + 1)}>
                        Next Step
                      </SoloButtonStyledComponent>
                    </div>
                  </TabPanel>
                  <TabPanel className='relative flex flex-col justify-between h-full focus:outline-none'>
                    <SectionContainer>
                      <SectionHeader>
                        {!!props.existingUser ? 'Edit' : 'Create a'} User:
                        Groups
                      </SectionHeader>
                      <div className='p-3 mb-2 text-gray-700 bg-gray-100 rounded-lg'>
                        Select a Group to which you'd like to add this user
                      </div>
                      <SoloTransfer
                        allOptionsListName='Available Groups'
                        allOptions={groupsList
                          .sort((a, b) =>
                            a.metadata?.name === b.metadata?.name
                              ? 0
                              : a.metadata!.name > b.metadata!.name
                              ? 1
                              : -1
                          )
                          .map(group => {
                            return {
                              name: group.metadata?.name!,
                              namespace: group.metadata?.namespace!,
                              displayValue: group.spec?.displayName
                            };
                          })}
                        chosenOptionsListName='Selected Groups'
                        chosenOptions={formik.values.chosenGroups.map(group => {
                          return {
                            name: group.name,
                            namespace: group.namespace,
                            displayValue: group.displayValue
                          };
                        })}
                        onChange={newChosenOptions => {
                          console.log('newChosenOptions', newChosenOptions);
                          formik.setFieldValue(
                            'chosenGroups',
                            newChosenOptions
                          );
                        }}
                      />
                    </SectionContainer>
                    <div className='flex items-end justify-between h-full px-6 mb-4 '>
                      <button
                        className='text-blue-500 cursor-pointer'
                        onClick={props.onClose}>
                        cancel
                      </button>
                      <div>
                        <SoloCancelButton
                          onClick={() => handleTabsChange(0)}
                          className='mr-2'>
                          Back
                        </SoloCancelButton>
                        <SoloButtonStyledComponent
                          onClick={() => setTabIndex(tabIndex + 1)}>
                          Next Step
                        </SoloButtonStyledComponent>
                      </div>
                    </div>
                  </TabPanel>
                  <TabPanel className='relative flex flex-col justify-between h-full focus:outline-none'>
                    <SectionContainer>
                      <SectionHeader>
                        {!!props.existingUser ? 'Edit' : 'Create a'} User: APIs
                      </SectionHeader>
                      <div className='p-3 mb-2 text-gray-700 bg-gray-100 rounded-lg'>
                        Select the APIs you'd like to make available to this
                        user
                      </div>
                      <SoloTransfer
                        allOptionsListName='Available APIs'
                        allOptions={apiDocsList
                          .sort((a, b) =>
                            a.metadata?.name === b.metadata?.name
                              ? 0
                              : a.metadata!.name > b.metadata!.name
                              ? 1
                              : -1
                          )
                          .map(apiDoc => {
                            return {
                              name: apiDoc.metadata?.name!,
                              namespace: apiDoc.metadata?.namespace!,
                              displayValue: apiDoc.status?.displayName
                            };
                          })}
                        chosenOptionsListName='Selected APIs'
                        chosenOptions={formik.values.chosenAPIs}
                        onChange={newChosenOptions => {
                          console.log('newChosenOptions', newChosenOptions);
                          formik.setFieldValue('chosenAPIs', newChosenOptions);
                        }}
                      />
                    </SectionContainer>
                    <div className='flex items-end justify-between h-full px-6 mb-4 '>
                      <button
                        className='text-blue-500 cursor-pointer'
                        onClick={props.onClose}>
                        cancel
                      </button>
                      <div>
                        <SoloCancelButton
                          onClick={() => handleTabsChange(0)}
                          className='mr-2'>
                          Back
                        </SoloCancelButton>
                        <SoloButtonStyledComponent
                          onClick={() => setTabIndex(tabIndex + 1)}>
                          Next Step
                        </SoloButtonStyledComponent>
                      </div>
                    </div>
                  </TabPanel>

                  <TabPanel className='relative flex flex-col justify-between h-full focus:outline-none'>
                    <SectionContainer>
                      <SectionHeader>
                        {!!props.existingUser ? 'Edit' : 'Create a'} User:
                        Portal
                      </SectionHeader>
                      <div className='p-3 mb-2 text-gray-700 bg-gray-100 rounded-lg'>
                        Select the portals you'd like to make available to this
                        user
                      </div>
                      <SoloTransfer
                        allOptionsListName='Available Portals'
                        allOptions={portalsList
                          .sort((a, b) =>
                            a.metadata?.name === b.metadata?.name
                              ? 0
                              : a.metadata!.name > b.metadata!.name
                              ? 1
                              : -1
                          )
                          .map(portal => {
                            return {
                              name: portal.metadata?.name!,
                              namespace: portal.metadata?.namespace!,
                              displayValue: portal.spec?.displayName
                            };
                          })}
                        chosenOptionsListName='Selected Portal'
                        chosenOptions={formik.values.chosenPortals}
                        onChange={newChosenOptions => {
                          console.log('newChosenOptions', newChosenOptions);
                          formik.setFieldValue(
                            'chosenPortals',
                            newChosenOptions
                          );
                        }}
                      />
                    </SectionContainer>
                    <div className='flex items-end justify-between h-full px-6 mb-4 '>
                      <button
                        className='text-blue-500 cursor-pointer'
                        onClick={props.onClose}>
                        cancel
                      </button>
                      <div>
                        <SoloCancelButton
                          onClick={() => handleTabsChange(1)}
                          className='mr-2'>
                          Back
                        </SoloCancelButton>
                        <SoloButtonStyledComponent
                          onClick={formik.handleSubmit}>
                          {!!props.existingUser ? 'Update User' : 'Create User'}
                        </SoloButtonStyledComponent>
                      </div>
                    </div>
                  </TabPanel>
                </TabPanels>
              </Tabs>
            </>
          )}
        </Formik>
      </div>
    </>
  );
};
