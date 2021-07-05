import * as React from 'react';
import { useDispatch } from 'react-redux';
import { SetCurrentConversation } from '../../app/actions/conversation';
import { Conversation } from '../../app/types/conversation';

type ConversationListProps = {
  conversations?: Conversation[];
}

export const ConversationList: React.FC<ConversationListProps> = (props: ConversationListProps) => {
  // Set the Dispatcher
  const dispatch = useDispatch();

  // Method for Selecting a Conversation Channel
  const selectConversation = (e: React.MouseEvent<HTMLElement>,conversation: Conversation) => {
    // Set the Current Conversation
    dispatch(SetCurrentConversation(conversation));

    // Update some ClassNames...
    let buttons: HTMLCollectionOf<Element> = document.getElementsByClassName("channel");
    for (var i = 0; i < buttons.length; i++) {
      let button: Element = buttons[i];
      button.classList.remove("selected");
    }
    let selectedButton: HTMLElement = e.target as HTMLElement;
    selectedButton.classList.add("selected");
  }
  return(
    <>
    {props.conversations !== undefined &&
      <ul>
        {props.conversations.map((conversation, i) => {
          return(
            <li key={conversation.id}>
              <button
                className="channel"
                title={conversation.label}
                onClick={(e: React.MouseEvent<HTMLElement>) => { selectConversation(e, conversation)}}>
                  {conversation.label}
                </button>
            </li>
          )
        })}
      </ul>
    }
    </>
  );
}