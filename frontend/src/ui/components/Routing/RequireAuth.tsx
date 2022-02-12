import { Navigate, useLocation } from "react-router-dom";
import { clearSession, isLoggedIn } from "../../../app/apiConn";

function RequireAuth({ children }: { children: JSX.Element }) {
  let location = useLocation();

  if (!isLoggedIn()) {
    // Clear the Stored State
    clearSession();

    // Redirect them to the /login page, but save the current location they were
    // trying to go to when they were redirected. This allows us to send them
    // along to that page after they login, which is a nicer user experience
    // than dropping them off on the home page.
    return <Navigate to='/welcome' state={{ from: location }} />;
  }

  return children;
}

export default RequireAuth;