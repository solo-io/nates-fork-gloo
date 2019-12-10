import { SoloWarning } from 'Components/Common/SoloWarningContent';
import {
  CreateUpstreamRequest,
  DeleteUpstreamRequest,
  UpdateUpstreamRequest
} from 'proto/github.com/solo-io/solo-projects/projects/grpcserver/api/v1/upstream_pb';
import { Dispatch } from 'redux';
import { guardByLicense } from 'store/config/actions';
import { MessageAction, SuccessMessageAction } from 'store/modal/types';
import { upstreams } from './api';
import {
  CreateUpstreamAction,
  DeleteUpstreamAction,
  ListUpstreamsAction,
  UpstreamAction,
  UpdateUpstreamAction
} from './types';

export const listUpstreams = () => {
  return async (dispatch: Dispatch) => {
    try {
      const response = await upstreams.listUpstreams();
      dispatch<ListUpstreamsAction>({
        type: UpstreamAction.LIST_UPSTREAMS,
        payload: response.upstreamDetailsList
      });
    } catch (error) {}
  };
};

export const deleteUpstream = (
  deleteUpstreamRequest: DeleteUpstreamRequest.AsObject
) => {
  return async (dispatch: Dispatch) => {
    // dispatch(showLoading());

    try {
      guardByLicense();
      const response = await upstreams.deleteUpstream(deleteUpstreamRequest);
      dispatch<DeleteUpstreamAction>({
        type: UpstreamAction.DELETE_UPSTREAM,
        payload: deleteUpstreamRequest
      });
      // dispatch(hideLoading());
    } catch (error) {
      SoloWarning('There was an error deleting the upstream.', error);
    }
  };
};

export const createUpstream = (
  createUpstreamRequest: CreateUpstreamRequest.AsObject
) => {
  return async (dispatch: Dispatch) => {
    console.log('createUpstreamRequest', createUpstreamRequest);
    // dispatch(showLoading());

    guardByLicense();
    try {
      const response = await upstreams.createUpstream(createUpstreamRequest);
      dispatch<CreateUpstreamAction>({
        type: UpstreamAction.CREATE_UPSTREAM,
        payload: response.upstreamDetails!
      });
      // dispatch(hideLoading());
      dispatch<SuccessMessageAction>({
        type: MessageAction.SUCCESS_MESSAGE,
        message: `Upstream ${
          response.upstreamDetails!.upstream!.metadata!.name
        } successfully created.`
      });
    } catch (error) {
      SoloWarning('There was an error creating the upstream.', error);
    }
  };
};

export const updateUpstream = (
  updateUpstreamRequest: UpdateUpstreamRequest.AsObject
) => {
  return async (dispatch: Dispatch) => {
    // dispatch(showLoading());
    guardByLicense();
    try {
      const response = await upstreams.updateUpstream(updateUpstreamRequest);
      dispatch<UpdateUpstreamAction>({
        type: UpstreamAction.UPDATE_UPSTREAM,
        payload: response.upstreamDetails!
      });
    } catch (error) {
      SoloWarning('There was an error updating the upstream.', error);
    }
  };
};
