import * as React from 'react';
import { useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { SendMessage } from '../../app/actions/conversation';
import { AppState } from '../../app/rootReducer';
import { Message } from '../../app/types/conversation';

export const MessageEditor: React.FC = () => {
  // Set the Dispatcher
  const dispatch = useDispatch();

  // Define the API States
  const conversationState = useSelector((state: AppState) => state.conversationState);

  // Set the States
  const [message, setMessage] = useState<string>();

  // Handle Input Changes in the Message Box
  const handleMessage = (e: React.ChangeEvent<HTMLTextAreaElement>) => {
    const message = (e.target as HTMLTextAreaElement).value;
    setMessage(message);
  }

  // Send the Message to the API
  const sendMessage = (e: React.MouseEvent<HTMLElement>) => {
    e.preventDefault();
    const chatMessage: Message = {
      conversation_id: conversationState.currentConversation?.id,
      message: message
    }
    dispatch(SendMessage(chatMessage));

    // Clear the Message Box
    let messageBox: HTMLTextAreaElement = document.getElementById("message-content") as HTMLTextAreaElement;
    messageBox.value = "";
  }

  // Render Method
  return(
    <form className="message-editor">
      <div className="inner">
        <textarea className="message-editor-input" id="message-content" name="message_content" onChange={handleMessage}></textarea>
        <button className="send-message-button primary" id="send-message" onClick={sendMessage}>Send</button>
      </div>
    </form>
  );
}