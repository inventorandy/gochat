import * as React from 'react';
import { useState } from 'react';
import { useDispatch } from 'react-redux';
import { SetLoggedInUser } from '../../app/actions/user';
import { appAPI } from '../../app/apiConn';
import { ErrorMessage } from '../../app/types/error';
import { LoginRequest, LoginResponse } from '../../app/types/user';

const LoginPage: React.FC = () => {
  // Set the States
  const [email, setEmail] = useState<string>();
  const [password, setPassword] = useState<string>();
  const [error, setError] = useState<string>();

  // Login Method
  const login = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const req: LoginRequest = {
      email: email,
      password: password
    };

    await appAPI.post("/auth/login", req).then(res => {
      // Handle the Login Response
      let loginResponse: LoginResponse = res.data;

      // Set the Current User
      SetLoggedInUser(loginResponse.user);

      // Set the JWT
      localStorage.setItem("authToken", loginResponse.auth_token);
    }).catch(error => {
      // Check if we got an error response
      if (error.response) {
        // Handle the error in here
        let errorBox = document.getElementById("login-error") as HTMLElement;
        let errorMsg: ErrorMessage = error.response.data;
        errorBox.innerText = errorMsg.message;
        errorBox.classList.remove("hidden");
      }
    });
  }

  // Handle Email Method
  const handleEmail = (e: React.ChangeEvent<HTMLInputElement>) => {
    const email = (e.target as HTMLInputElement).value;
    setEmail(email);
  }

  // Handle Password Method
  const handlePassword = (e: React.ChangeEvent<HTMLInputElement>) => {
    const password = (e.target as HTMLInputElement).value;
    setPassword(password);
  }

  // Render Method
  return(
    <div className="login-page">
      <form className="login-form" action="/auth/login" method="post" onSubmit={login}>
        <h1>Log In</h1>
        <p id="login-error" className="notification error hidden">{error}</p>
        <p>
          <label htmlFor="email">Email Address</label>
          <input type="email" id="email" name="email" onChange={handleEmail} autoComplete="Off" />
        </p>
        <p>
          <label htmlFor="password">Password</label>
          <input type="password" id="password" name="password" onChange={handlePassword} />
        </p>
        <p>
          <button className="primary">Log In</button>
        </p>
        <p>
          <a href="/auth/create-account" title="Register Account">Register Account</a>
        </p>
      </form>
    </div>
  );
}

export default LoginPage;