import * as React from 'react';
import { useEffect } from 'react';
import { GetPrivateConversations } from '../../../app/actions/conversation';
import { useAppDispatch, useAppSelector } from '../../../app/hooks';
import { RootState } from '../../../app/store';

const PrivateConversationList: React.FC = () => {
  const dispatch = useAppDispatch();

  // Get the List of Private Conversations
  const conversations = useAppSelector(
    (state: RootState) => state.conversation.privateConversations
  );

  // useEffect Hook
  useEffect(() => {
    if (!conversations) {
      dispatch(
        GetPrivateConversations(
          (conversations) => {},
          (err) => {}
        )
      );
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
                  {'#' + conversation.label}
                </a>
              </li>
            );
          })}
      </ul>
    </div>
  );
};

export default PrivateConversationList;
