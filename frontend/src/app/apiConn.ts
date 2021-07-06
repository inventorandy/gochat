import axios from 'axios';
import { Dispatch } from 'redux';

export const appAPI = axios.create({
  baseURL: `${process.env.REACT_APP_API_URL}`
})

export const isLoggedIn = (): boolean => {
  let authToken = localStorage.getItem("authToken");
  if (authToken === "" || authToken === null) {
    return false;
  } else {
    return true;
  }
}