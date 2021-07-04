import { User, UserActions, UserActionTypes } from '../types/user';

// UserState
interface UserState {
  loggedInUser?: User,
  error?: string | null
}

// Set the initial (empty) state
const initialState: UserState = {
  loggedInUser: undefined,
  error: null
}

// User Reducer
const userReducer = (state: UserState = initialState, action: UserActionTypes): UserState => {
  switch(action.type) {
    case UserActions.SET_LOGGED_IN_USER:
      return {
        loggedInUser: action.user,
      }
    default:
      return state;
  }
}

export default userReducer;