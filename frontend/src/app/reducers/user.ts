import { User, UserActions, UserActionTypes } from '../types/user';

// UserState
interface UserState {
  loggedInUser?: User,
  users?: User[],
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
        ...state,
        loggedInUser: action.user,
      }
    case UserActions.GET_LOGGED_IN_USER:
      return {
        ...state,
        loggedInUser: action.user,
      }
    case UserActions.GET_ALL_USERS:
      return {
        ...state,
        users: action.users,
      }
    default:
      return state;
  }
}

export default userReducer;