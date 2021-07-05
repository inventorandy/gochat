import * as React from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { SendMessage } from '../../app/actions/conversation';
import { AppState } from '../../app/rootReducer';
import { Message } from '../../app/types/conversation';

export const MessageEditor: React.FC = () => {
  // Set the Dispatcher
  const dispatch = useDispatch();

  // Define the API States
  const conversationState = useSelector((state: AppState) => state.conversationState);

  const sendMessage = (e: React.MouseEvent<HTMLElement>) => {
    e.preventDefault();
    const message: Message = {
      conversation_id: conversationState.currentConversation?.id,
      message: "Hello World!"
    }
    dispatch(SendMessage(message));
  }

  // Render Method
  return(
    <form className="message-editor">
      <div className="inner">
        <textarea className="message-editor-input" id="message-content" name="message_content"></textarea>
        <button className="send-message-button" id="send-message" onClick={sendMessage}>Send</button>
      </div>
    </form>
  );
}