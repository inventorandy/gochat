import { library } from '@fortawesome/fontawesome-svg-core';
import { faHashtag } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import * as React from 'react';
import { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import {
  CreateConversation,
  GetPublicConversations,
} from '../../../app/actions/conversation';
import { useAppDispatch, useAppSelector } from '../../../app/hooks';
import { RootState } from '../../../app/store';
import { Conversation } from '../../../app/types/conversation';
import ModalDialog from '../modal/ModalDialog';

// Add the Required Icons to the Library
library.add(faHashtag);

const PublicConversationList: React.FC = () => {
  // Get the Conversation ID
  const { id } = useParams();

  const dispatch = useAppDispatch();

  // Get the List of Public Conversations
  const conversations = useAppSelector(
    (state: RootState) => state.conversation.publicConversations
  );

  // useEffect Hook
  useEffect(() => {
    if (!conversations) {
      dispatch(GetPublicConversations((err) => {}));
    }
  }, [dispatch, conversations]);

  // Flag for Dialog Visibility
  const [hidden, setHidden] = useState<boolean>(true);

  // Open Add Channel Dialog
  const onOpenDialog = (e: React.MouseEvent<HTMLButtonElement>) => {
    // Prevent Default Behaviour
    e.preventDefault();

    // Display the Dialog
    setHidden(false);
  };

  // Close Add Channel Dialog
  const onCloseDialog = (e: React.MouseEvent<HTMLButtonElement>) => {
    // Prevent Default Behaviour
    e.preventDefault();

    // Hide the Dialog
    setHidden(true);
  };

  // Channel Name Handler
  const [chanName, setChanName] = useState<string>('');
  const handleChannelName = (e: React.ChangeEvent<HTMLInputElement>) => {
    setChanName(e.target.value);
  };

  // Create New Channel Handler
  const onCreateChannel = (e: React.FormEvent<HTMLFormElement>) => {
    // Prevent Default Behaviour
    e.preventDefault();

    // Build the Conversation Object
    let conv: Conversation = {
      label: chanName,
      is_public: true,
      messages: [],
    };

    dispatch(
      CreateConversation(conv, (err) => {
        //something
      })
    );

    // Hide the Dialog
    setHidden(true);

    // Clear the Input
    setChanName('');
  };

  // Render
  return (
    <div className='channel-list'>
      <h4>
        <span>Public Channels</span>
        <button onClick={onOpenDialog}>+</button>
      </h4>
      <ul>
        {conversations &&
          conversations.map((conversation) => {
            let css = id && id === conversation.id ? 'selected' : '';
            return (
              <li key={conversation.id}>
                <a href={'/channels/' + conversation.id} className={css}>
                  <FontAwesomeIcon icon='hashtag' />
                  {conversation.label}
                </a>
              </li>
            );
          })}
      </ul>
      <ModalDialog
        id='create-public-channel'
        title='Create Public Channel'
        hidden={hidden}
        onClickClose={onCloseDialog}
      >
        <form className='channel-dialog' onSubmit={onCreateChannel}>
          <input
            type='text'
            id='channel-name'
            placeholder='Channel Name'
            value={chanName}
            onChange={handleChannelName}
          />
          <div className='center'>
            <button>Create Channel</button>
          </div>
        </form>
      </ModalDialog>
    </div>
  );
};

export default PublicConversationList;
