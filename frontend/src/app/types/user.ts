export interface User {
  id?: string;
  created_at?: string;
  updated_at?: string;
  email: string;
  password?: string;
  salt?: string;
  first_name?: string;
  last_name?: string;
}

export interface LoginRequest {
  email: string;
  password: string;
}

export interface LoginResponse {
  auth_token: string;
  user: User;
}

export type jwt = {
  account_id?: string;
  exp: number;
  iat: number;
};
