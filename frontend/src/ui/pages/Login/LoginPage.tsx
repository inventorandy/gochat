import * as React from 'react';
import { useState } from 'react';
import { useDispatch } from 'react-redux';
import { useNavigate } from 'react-router';
import { Login } from '../../../app/actions/user';
import { LoginRequest } from '../../../app/types/user';
import FormError from '../../components/Forms/FormError';
import { openDialog } from '../../components/Modal/Modal';
import CreateAccountDialog from './CreateAccountDialog';

const LoginPage: React.FC = () => {
  // Set the Dispatcher
  const dispatch = useDispatch();

  // Set the Navigator
  const navigate = useNavigate();

  // Set Initial States
  const [email, setEmail] = useState<string>('');
  const [password, setPassword] = useState<string>('');
  const [error, setError] = useState<string>('');

  // Input Handlers
  const handleEmail = (e: React.ChangeEvent<HTMLInputElement>) => {
    setEmail(e.target.value);
  };
  const handlePassword = (e: React.ChangeEvent<HTMLInputElement>) => {
    setPassword(e.target.value);
  };

  // Create Account Dialog Opener
  const onCreateAccount = (e: React.MouseEvent<HTMLButtonElement>) => {
    // Prevent Default Behaviour
    e.preventDefault();

    // Open the Dialog
    openDialog('create-account');
  };

  // Login Handler
  const onLogin = (e: React.FormEvent<HTMLFormElement>) => {
    // Prevent Default Behaviour
    e.preventDefault();

    // Build the Login Request
    const req: LoginRequest = {
      email: email,
      password: password,
    };

    // Call the API Method
    dispatch(
      Login(
        req,
        (res) => {
          // Navigate to the Main App
          navigate('/', { replace: true });
        },
        (err) => {
          // Set the Error
          setError(err.message);
        }
      )
    );
  };

  // Render
  return (
    <>
      <form className='login' onSubmit={onLogin}>
        <h1>Log In</h1>
        <FormError error={error} />
        <div className='login-gopher' />
        <input
          type='email'
          className='btl-rounded btr-rounded bdr-t-solid bdr-l-solid bdr-r-solid bdr-b-none text-input box-sizing full-width bg-primary-overlay lg'
          id='email'
          placeholder='Email Address'
          required={true}
          autoComplete='Off'
          autoCapitalize='Off'
          onChange={handleEmail}
          value={email}
        />
        <input
          type='password'
          className='bbl-rounded bbr-rounded bdr-solid text-input box-sizing full-width bg-primary-overlay lg'
          id='password'
          placeholder='Password'
          required={true}
          onChange={handlePassword}
          value={password}
        />
        <p className='flex-right'>
          <button className='btn'>Log In</button>
          <button className='txt-btn' onClick={onCreateAccount}>
            Create Account
          </button>
        </p>
      </form>
      <CreateAccountDialog />
    </>
  );
};

export default LoginPage;
