import * as React from 'react';
import { useState } from 'react';
import { useNavigate } from 'react-router';
import { Login } from '../../app/actions/user';
import { useAppDispatch } from '../../app/hooks';
import { LoginRequest } from '../../app/types/user';
import gopher from '../scss/images/gopher.png';
import '../scss/login.scss';

const LoginPage: React.FC = () => {
  const dispatch = useAppDispatch();

  // Set the Browser History
  const navigate = useNavigate();

  // Form Element Value States
  const [email, setEmail] = useState<string>('');
  const [password, setPassword] = useState<string>('');

  // Input Handlers
  const handleEmail = (e: React.ChangeEvent<HTMLInputElement>) => {
    setEmail(e.target.value);
  };
  const handlePassword = (e: React.ChangeEvent<HTMLInputElement>) => {
    setPassword(e.target.value);
  };

  // Login Form Handler
  const onLogin = (e: React.FormEvent<HTMLFormElement>) => {
    // Prevent Default Behaviour
    e.preventDefault();

    // Build the Login Request
    const req: LoginRequest = {
      email: email,
      password: password,
    };

    // Call the Login Method
    dispatch(
      Login(
        req,
        (res) => {
          // Navigate to the Dashboard
          navigate('/', { replace: true });
        },
        (err) => {
          // Handle Error
        }
      )
    );
  };
  // Render
  return (
    <form className='login-form' onSubmit={onLogin}>
      <h1>Welcome to GoChat</h1>
      <img src={gopher} alt='GoChat Gopher' />
      <div>
        <input
          type='email'
          id='email'
          placeholder='Email Address'
          value={email}
          onChange={handleEmail}
          autoComplete='Off'
        />
      </div>
      <div>
        <input
          type='password'
          id='password'
          placeholder='Password'
          value={password}
          onChange={handlePassword}
          autoComplete='Off'
        />
      </div>
      <div>
        <button>Log In</button>
      </div>
    </form>
  );
};

export default LoginPage;
