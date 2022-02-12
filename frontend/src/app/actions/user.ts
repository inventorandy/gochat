import jwtDecode from 'jwt-decode';
import { Dispatch } from 'redux';
import { appAPI } from '../apiConn';
import { setCookie } from '../helpers/cookie';
import { APIError } from '../types/generic';
import {
  Jwt,
  LoginRequest,
  LoginResponse,
  User,
  UserActions,
} from '../types/user';

/**
 * Creates a new user on the system.
 * @param account Input account object with user info.
 * @param onSuccess Called when account is created.
 * @param onError Called for any API error.
 * @returns
 */
export const CreateAccount =
  (
    account: User,
    onSuccess: (user: User) => void,
    onError: (err: APIError) => void
  ) =>
  async (dispatch: Dispatch) => {
    // Call the API
    await appAPI
      .post('/account', account)
      .then((res) => {
        // Get the Account Info
        let user: User = res.data;

        // Call the onSuccess Method
        onSuccess(user);
      })
      .catch((err) => {
        // Get the Error Object
        let error: APIError = err.response.data;

        // Handle Error
        onError(error);
      })
      .finally(() => {
        // Create an Error Object
        let error: APIError = {
          code: 500,
          message: 'Severe API Error.',
        };

        // Handle Error
        onError(error);
      });
  };

/**
 * Login method called the authentication API endoint and returns an access token and user.
 * @param req Login request payload.
 * @param onSuccess Called when login is successful.
 * @param onError Called for any API error.
 * @returns
 */
export const Login =
  (
    req: LoginRequest,
    onSuccess: (res: LoginResponse) => void,
    onError: (err: APIError) => void
  ) =>
  async (dispatch: Dispatch) => {
    // Call the API
    await appAPI
      .post('/auth/login', req)
      .then((res) => {
        // Get the Login Response
        let response: LoginResponse = res.data;

        // Decode the Token
        const token = jwtDecode<Jwt>(response.auth_token);

        // Create the Expiry Date
        const expiry = new Date();
        expiry.setTime(token.exp * 1000);

        // Set a Cookie
        setCookie('token', response.auth_token, expiry);

        // Call the Dispatcher
        dispatch({
          type: UserActions.LOGIN,
          user: response.user,
        });

        // Call the onSuccess Method
        onSuccess(response);
      })
      .catch((err) => {
        // Get the Error Object
        let error: APIError = err.response.data;

        // Handle Error
        onError(error);
      })
      .finally(() => {
        // Create an Error Object
        let error: APIError = {
          code: 500,
          message: 'Severe API Error.',
        };

        // Handle Error
        onError(error);
      });
  };
