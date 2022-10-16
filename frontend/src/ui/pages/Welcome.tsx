import * as React from 'react';
import { currentUser } from '../../app/features/user';
import { useAppSelector } from '../../app/hooks';

const WelcomePage: React.FC = () => {
  // Get the User
  const user = useAppSelector(currentUser);

  // Render
  return (
    <div className='welcome'>
      <h1>Welcome, {user?.first_name}</h1>
      <p>Select a channel on the left to begin chatting.</p>
    </div>
  );
};

export default WelcomePage;
