import * as React from 'react';
import { Conversation } from '../../app/types/conversation';

type ConversationListProps = {
  conversations?: Conversation[];
}

export const ConversationList: React.FC<ConversationListProps> = (props: ConversationListProps) => {
  return(
    <>
    {props.conversations !== undefined &&
      <ul>
        {props.conversations.map((item, i) => {
          return(
            <li key={item.id}>
              <a className="channel" href="#" title={item.label}>{item.label}</a>
            </li>
          )
        })}
      </ul>
    }
    </>
  );
}