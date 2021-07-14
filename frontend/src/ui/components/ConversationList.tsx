import * as React from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { SetCurrentConversation } from '../../app/actions/conversation';
import { Conversation } from '../../app/types/conversation';
import { library } from '@fortawesome/fontawesome-svg-core';
import { faHashtag, faLock } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { AppState } from '../../app/rootReducer';

// Add the Required Icons to the Library
library.add( faHashtag, faLock );

type ConversationListProps = {
  conversations?: Conversation[];
}

export const ConversationList: React.FC<ConversationListProps> = (props: ConversationListProps) => {
  // Set the Dispatcher
  const dispatch = useDispatch();

  // Define the API States
  const conversationState = useSelector((state: AppState) => state.conversationState);

  // Method for Selecting a Conversation Channel
  const selectConversation = (e: React.MouseEvent<HTMLElement>,conversation: Conversation) => {
    // Set the Current Conversation
    dispatch(SetCurrentConversation(conversation.id));

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
          let cssClass = "channel";
          if (conversation.id === conversationState.currentConversation?.id) {
            cssClass += " selected";
          }
          return(
            <li key={conversation.id}>
              <button
                className={cssClass}
                title={conversation.label}
                onClick={(e: React.MouseEvent<HTMLElement>) => { selectConversation(e, conversation)}}>
                  {conversation.is_public === true &&
                    <FontAwesomeIcon icon="hashtag" />
                  }
                  {conversation.is_public !== true &&
                    <FontAwesomeIcon icon="lock" />
                  }
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