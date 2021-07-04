import * as React from 'react';
import { Route, Redirect } from "react-router-dom";
import { isLoggedIn } from '../../app/apiConn';

type PrivateRouteProps = {
  component: React.FC;
  path: string;
  exact?: boolean;
}

const PrivateRoute: React.FC<React.PropsWithChildren<PrivateRouteProps>> = (props: React.PropsWithChildren<PrivateRouteProps>) => {
  return isLoggedIn() ? (<Route path={props.path} exact={props.exact} component={props.component} />) : (<Redirect to="/auth/login" />);
}

export default PrivateRoute;