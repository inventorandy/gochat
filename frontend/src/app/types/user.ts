export interface LoginRequest {
  email: string;
  password: string;
}

export interface LoginResponse {
  auth_token: string;
  user: User;
}

export interface Jwt {
  user_id?: number;
  exp: number;
  iat: number;
}

export interface User {
  id?: string;
  email: string;
  password: string;
  first_name: string;
  last_name: string;
}

export enum UserActions {
  LOGIN = 'LOGIN',
}

interface LoginAction {
  type: typeof UserActions.LOGIN;
  user: User;
}

export type UserActionTypes = LoginAction;
