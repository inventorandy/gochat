import * as React from 'react';
import { Conversation } from '../../app/types/conversation';
import { ChatMessage } from './ChatMessage';

type MessageListProps = {
  conversation?: Conversation;
}

export const MessageList: React.FC<MessageListProps> = (props: MessageListProps) => {
  return(
    <div className="message-list">
      {props.conversation?.messages !== undefined && props.conversation?.messages !== null &&
        props.conversation?.messages.map((message, i) => {
          return(
            <ChatMessage message={message} />
          );
        })
      }
      {props.conversation?.messages === undefined &&
        <p>No messages yet</p>
      }
    </div>
  );
}