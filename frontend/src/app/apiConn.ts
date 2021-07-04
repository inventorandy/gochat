import axios from 'axios';

export const appAPI = axios.create({
  baseURL: `${process.env.REACT_APP_API_URL}`
})

export const isLoggedIn = (): boolean => {
  let jwt = localStorage.getItem("jwt");
  if (jwt === "" || jwt === null) {
    return false;
  } else {
    return true;
  }
}