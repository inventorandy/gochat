// User API Object
export interface User {
  id?: string;
  email?: string;
  password?: string;
  first_name?: string;
  last_name?: string;
}

// LoginRequest API Object
export interface LoginRequest {
  email?: string;
  password?: string;
}

// LoginResponse API Object
export interface LoginResponse {
  auth_token: string;
  user: User;
}

// UserActions Definitions
export enum UserActions {
  SET_LOGGED_IN_USER = 1,
  GET_LOGGED_IN_USER = 2,
  CLEAR_LOGGED_IN_USER = 3,
  LOGIN = 4,
}

// Login User Action Type
interface SetLoggedInUser {
  type: typeof UserActions.SET_LOGGED_IN_USER,
  user: User,
}

// Get Logged In User Action Type
interface GetLoggedInUser {
  type: typeof UserActions.GET_LOGGED_IN_USER,
  user: User,
}

// Export the Action Type
export type UserActionTypes = SetLoggedInUser | GetLoggedInUser;