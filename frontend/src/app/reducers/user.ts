import { User, UserActions, UserActionTypes } from '../types/user';

// UserState
interface UserState {
  loggedInUser?: User;
  workspaceUsers?: User[];
}

// Set the initial (empty) state
const initialState: UserState = {
  loggedInUser: undefined,
  workspaceUsers: undefined,
};

// User Reducer
const userReducer = (
  state: UserState = initialState,
  action: UserActionTypes
): UserState => {
  switch (action.type) {
    case UserActions.LOGIN:
      return {
        ...state,
        loggedInUser: action.user,
      }
    default:
      return state;
  }
};

export default userReducer;
