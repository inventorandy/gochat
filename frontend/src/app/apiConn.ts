import axios from 'axios';
import { deleteCookie, getCookie } from './helpers/cookie';

export const appAPI = axios.create({
  baseURL: `${process.env.REACT_APP_API_URL}`,
});

//export const headers = { headers: { Authorization: getCookie('token') || '' } }
export const headers = () => {
  return { headers: { Authorization: getCookie('token') || '' } };
}

/**
 * Checks if a user is logged in.
 * @returns boolean
 */
export const isLoggedIn = (): boolean => {
  let jwt = getCookie('token');
  if (jwt === '' || jwt === null || jwt === undefined) {
    return false;
  } else {
    return true;
  }
};

/**
 * Clears the Session and any Storage
 */
export const clearSession = () => {
  deleteCookie('token');
  localStorage.clear();
};
