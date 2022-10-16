import * as React from 'react';
import { useEffect } from 'react';
import { GetPublicConversations } from '../../../app/actions/conversation';
import { useAppDispatch, useAppSelector } from '../../../app/hooks';
import { RootState } from '../../../app/store';

const PublicConversationList: React.FC = () => {
  const dispatch = useAppDispatch();

  // Get the List of Public Conversations
  const conversations = useAppSelector(
    (state: RootState) => state.conversation.publicConversations
  );

  // useEffect Hook
  useEffect(() => {
    if (!conversations) {
      dispatch(
        GetPublicConversations(
          (conversations) => {},
          (err) => {}
        )
      );
    }
  }, [dispatch, conversations]);

  // Render
  return (
    <div className='channel-list'>
      <h4>Public Channels</h4>
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

export default PublicConversationList;
