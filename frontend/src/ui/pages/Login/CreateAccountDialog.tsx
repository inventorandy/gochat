import * as React from 'react';
import { useState } from 'react';
import { useDispatch } from 'react-redux';
import { CreateAccount } from '../../../app/actions/user';
import { User } from '../../../app/types/user';
import FormError from '../../components/Forms/FormError';
import FormInput from '../../components/Forms/TextInput';
import {
  Modal,
  ModalHeader,
  ModalContent,
  closeDialog,
} from '../../components/Modal/Modal';

const CreateAccountDialog: React.FC = () => {
  // Set the Dispatcher
  const dispatch = useDispatch();

  // Component States
  const [email, setEmail] = useState<string>('');
  const [password, setPassword] = useState<string>('');
  const [firstName, setFirstName] = useState<string>('');
  const [lastName, setLastName] = useState<string>('');
  const [error, setError] = useState<string>('');

  // Input Handlers
  const handleEmail = (e: React.ChangeEvent<HTMLInputElement>) => {
    setEmail(e.target.value);
  };
  const handlePassword = (e: React.ChangeEvent<HTMLInputElement>) => {
    setPassword(e.target.value);
  };
  const handleFirstName = (e: React.ChangeEvent<HTMLInputElement>) => {
    setFirstName(e.target.value);
  };
  const handleLastName = (e: React.ChangeEvent<HTMLInputElement>) => {
    setLastName(e.target.value);
  };

  // Cancel Handler
  const onCancel = (e: React.MouseEvent<HTMLButtonElement>) => {
    // Prevent Default Behaviour
    e.preventDefault();

    // Close the Dialog
    closeDialog('create-account');
  };

  // Create Account Handler
  const onCreateAccount = (e: React.FormEvent<HTMLFormElement>) => {
    // Prevent Default Behaviour
    e.preventDefault();

    // Build the User Object
    const user: User = {
      email: email,
      password: password,
      first_name: firstName,
      last_name: lastName,
    };

    // Call the API Method
    dispatch(
      CreateAccount(
        user,
        (user) => {
          // Close the Dialog
          closeDialog('create-account');
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
    <Modal id='create-account' hideOnLoad={true}>
      <ModalHeader dialog='create-account' title='Create Account' />
      <ModalContent>
        <form onSubmit={onCreateAccount}>
          <FormError error={error} />
          <div className='inputs box-sizing'>
            <FormInput
              type='text'
              id='cr-first-name'
              label='First Name'
              value={firstName}
              onChange={handleFirstName}
              required={true}
            />
            <FormInput
              type='text'
              id='cr-last-name'
              label='Last Name'
              value={lastName}
              onChange={handleLastName}
              required={true}
            />
            <FormInput
              type='email'
              id='cr-email'
              label='Email'
              value={email}
              onChange={handleEmail}
              required={true}
            />
            <FormInput
              type='password'
              id='cr-password'
              label='Password'
              value={password}
              onChange={handlePassword}
              required={true}
            />
          </div>
          <div className='button-options'>
            <button className='btn'>Create Account</button>
            <button className='btn' onClick={onCancel}>
              Cancel
            </button>
          </div>
        </form>
      </ModalContent>
    </Modal>
  );
};

export default CreateAccountDialog;
