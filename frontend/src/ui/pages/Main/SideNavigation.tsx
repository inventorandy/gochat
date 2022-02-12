import * as React from 'react';
import { useEffect, useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { useNavigate } from 'react-router';
import {
  GetConversation,
  GetPublicConversations,
} from '../../../app/actions/conversation';
import { clearSession } from '../../../app/apiConn';
import combineCSS from '../../../app/helpers/combineCSS';
import { AppState } from '../../../app/rootReducer';
import { Conversation } from '../../../app/types/conversation';
import CreatePublicChannelDialog from '../../components/Dialogs/CreatePublicChannelDialog';
import { openDialog } from '../../components/Modal/Modal';

const SideNavigation: React.FC = () => {
  // Set the Dispatcher
  const dispatch = useDispatch();

  // Set the Navigator
  const navigate = useNavigate();

  // Get the App States
  const convState = useSelector((state: AppState) => state.conversationState);

  // Set the State Variables
  const [publicConvos, setPublicConvos] = useState<Conversation[]>(
    convState.publicConversations || []
  );

  // useEffect Hook
  useEffect(() => {
    dispatch(
      GetPublicConversations(
        (conversations) => {
          // Set the Conversation List
          setPublicConvos(conversations);
        },
        (err) => {}
      )
    );
  }, [dispatch]);

  // Logout Button Handler
  const onLogout = (e: React.MouseEvent<HTMLButtonElement>) => {
    // Prevent Default Behaviour
    e.preventDefault();

    // Clear the Session
    clearSession();

    // Go Back to the Welcome Screen
    navigate('/welcome', { replace: true });
  };

  // Select Channel Button Handler
  const onSelectChannel = (
    e: React.MouseEvent<HTMLElement>,
    channelID: string
  ) => {
    // Prevent Default Behaviour
    e.preventDefault();

    // Call the API Method
    dispatch(
      GetConversation(channelID, (err) => {
        // Do some error handling shit
      })
    );
  };

  // Add Public Channel Button Handler
  const onAddChannel = (e: React.MouseEvent<HTMLButtonElement>) => {
    // Prevent Default Behaviour
    e.preventDefault();

    // Open the Dialog
    openDialog('create-public-channel');
  }

  // Render
  return (
    <div className='app-drawer box-sizing'>
      <h2 className='center-align'>GoChat!</h2>
      <h3>
        <span>Public Channels</span>
        <button>+</button>
      </h3>
      {publicConvos?.map((conversation) => {
        return (
          <button
            className={combineCSS(
              'left-align full-width box-sizing bdr-none bdr-rounded',
              convState.currentConversation?.id === conversation.id
                ? 'selected'
                : ''
            )}
            key={conversation.id}
            onClick={(e) => onSelectChannel(e, conversation.id)}
          >
            {'# ' + conversation.label}
          </button>
        );
      })}
      <button onClick={onLogout}>Log Out</button>
      <CreatePublicChannelDialog />
    </div>
  );
};

export default SideNavigation;
