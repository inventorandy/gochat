import { Dispatch } from 'redux';
import { User, UserActions } from '../types/user';

export const SetLoggedInUser = (user: User) => async(dispatch: Dispatch) => {
  dispatch({
    type: UserActions.SET_LOGGED_IN_USER,
	  user: user
  });
}