import axios from 'axios';
import { deleteCookie, getCookie } from './cookie';

export const appAPI = axios.create({
  baseURL: `${process.env.REACT_APP_API_URL}`,
});

export const isLoggedIn = (): boolean => {
  let jwt = getCookie('token');
  if (jwt === '' || jwt === null || jwt === undefined) {
    return false;
  } else {
    return true;
  }
};

export const clearSession = () => {
  deleteCookie('token');
  localStorage.clear();
};

export const getHeaderConfig = () => {
  return {
    headers: {
      // TODO: Include Bearer when we move to OAPI3
      Authorization: '' + getCookie('token'),
    },
  };
};
