import { Dispatch } from 'redux';
import { LoginRequest, LoginResponse, User, UserActions } from '../types/user';
import { appAPI } from '../apiConn';

export const SetLoggedInUser = (user: User) => async(dispatch: Dispatch) => {
  dispatch({
    type: UserActions.SET_LOGGED_IN_USER,
	  user: user
  });
}