import jwtDecode from 'jwt-decode';
import { Dispatch } from 'redux';
import { appAPI } from '../apiConn';
import { setCookie } from '../cookie';
import { setCurrentUser } from '../features/user';
import { APIError } from '../types/error';
import { jwt, LoginRequest, LoginResponse } from '../types/user';

export const Login =
  (
    req: LoginRequest,
    onSuccess: (res: LoginResponse) => void,
    onError: (err: any) => void
  ) =>
  async (dispatch: Dispatch) => {
    appAPI
      .post('/auth/login', req)
      .then((res) => {
        // Get the Response
        let response: LoginResponse = res.data;

        // Decode the Token
        const token = jwtDecode<jwt>(response.auth_token);

        // Create the Expiry Date
        const expiry = new Date();
        expiry.setTime(token.exp * 1000);

        // Set the Cookie
        setCookie('token', response.auth_token, expiry);

        // Set the Current User
        dispatch(setCurrentUser(response.user));

        // Call the Success Method
        onSuccess(response);
      })
      .catch((err) => {
        // Get the Error Message
        let error: APIError = err.response.data;

        // Call the Error Method
        onError(error);
      });
  };
