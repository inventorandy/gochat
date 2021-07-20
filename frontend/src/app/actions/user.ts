import { Dispatch } from 'redux';
import { appAPI } from '../apiConn';
import { User, UserActions } from '../types/user';

export const SetLoggedInUser = (user: User) => (dispatch: Dispatch) => {
  dispatch({
    type: UserActions.SET_LOGGED_IN_USER,
	  user: user
  });
}

export const GetLoggedInUser = () => (dispatch: Dispatch) => {
  appAPI.get("/account", { headers: { "Authorization": localStorage.getItem("authToken") } }).then(res => {
    // Get the User from the Response
    let user: User = res.data;

    // Do the Dispatch
    dispatch({
      type: UserActions.GET_LOGGED_IN_USER,
	    user: user,
    });
  }).catch(error => {
    // Handle the Error Response
  })
}

export const GetAllUsers = () => (dispatch: Dispatch) => {
  appAPI.get("/account/get-all", { headers: { "Authorization": localStorage.getItem("authToken") } }).then(res => {
    // Get the Users from the Response
    let users: User[] = res.data;

    // Do the Dispatch
    dispatch({
      type: UserActions.GET_ALL_USERS,
      users: users,
    });
  }).catch(error => {
    // Handle the Error Response
  })
}