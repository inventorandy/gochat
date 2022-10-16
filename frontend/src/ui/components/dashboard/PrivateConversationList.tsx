import { library } from '@fortawesome/fontawesome-svg-core';
import { faLock } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import * as React from 'react';
import { useEffect } from 'react';
import { GetPrivateConversations } from '../../../app/actions/conversation';
import { useAppDispatch, useAppSelector } from '../../../app/hooks';
import { RootState } from '../../../app/store';

// Add the Required Icons to the Library
library.add(faLock);

const PrivateConversationList: React.FC = () => {
  const dispatch = useAppDispatch();

  // Get the List of Private Conversations
  const conversations = useAppSelector(
    (state: RootState) => state.conversation.privateConversations
  );

  // useEffect Hook
  useEffect(() => {
    if (!conversations) {
      dispatch(GetPrivateConversations((err) => {}));
    }
  }, [dispatch, conversations]);

  // Render
  return (
    <div className='channel-list'>
      <h4>Private Channels</h4>
      <ul>
        {conversations &&
          conversations.map((conversation) => {
            return (
              <li key={conversation.id}>
                <a href={'/channels/' + conversation.id}>
                  <FontAwesomeIcon icon='lock' />
                  {conversation.label}
                </a>
              </li>
            );
          })}
      </ul>
    </div>
  );
};

export default PrivateConversationList;
