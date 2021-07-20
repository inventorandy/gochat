import * as React from 'react';
import { useEffect, useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { CreateConversation, GetPrivateConversations } from '../../app/actions/conversation';
import { GetAllUsers } from '../../app/actions/user';
import { AppState } from '../../app/rootReducer';
import { Conversation } from '../../app/types/conversation';
import { ErrorMessage } from '../../app/types/error';
import { User } from '../../app/types/user';
import ModalDialog, { closeDialog } from './ModalDialog';

const CreatePrivateChannelDialog: React.FC = () => {
  // Set the Dispatcher
  const dispatch = useDispatch();

  // Define the API States
  const conversationState = useSelector((state: AppState) => state.conversationState);
  const userState = useSelector((state: AppState) => state.userState);

  // Set the States
  const [channelName, setChannelName] = useState<string>();
  const [users, setUsers] = useState<User[]>();
  const [conversation, setConversation] = useState<Conversation>();
  const [error, setError] = useState<string>();

  // Load the Required API Data
  useEffect(() => {
    dispatch(GetAllUsers());
    let newUserList = new Array<User>();
    if (userState.loggedInUser !== undefined) {
      newUserList.push(userState.loggedInUser);
    }
    setUsers(newUserList);
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  // Handle Channel Name Method
  const handleChannelName = (e: React.ChangeEvent<HTMLInputElement>) => {
    const channelName = (e.target as HTMLInputElement).value;
    setChannelName(channelName);
  }

  // Handle Toggling a User
  const toggleUser = (e: React.ChangeEvent<HTMLInputElement>, user: User) => {
    let userList = users;
    let didRemove = false;
    if (userList !== undefined) {
      for (var i = 0; i < userList.length; i++) {
        if (userList[i].id === user.id) {
          userList.splice(i, 1);
          didRemove = true;
          e.target.checked = false;
        }
      }
    }
    if (!didRemove) {
      userList?.push(user);
      e.target.checked = true;
    }
    setUsers(userList);
  }

  // Determine whether to check one of the user checkboxes
  const isChecked = (user: User) => {
    if (users !== undefined) {
      for (var i = 0; i < users.length; i++) {
        if (users[i].id === user.id) {
          return true;
        }
      }
    }
    return false;
  }

  const createChannel = (e: React.FormEvent<HTMLFormElement>) => {
    // Prevent Default Behaviour
    e.preventDefault();

    // Create the Conversation Object
    let conversation: Conversation = {
      label: channelName,
      is_public: false,
      messages: [],
      participants: users
    }

    // Send the API Request
    dispatch(CreateConversation(conversation, channelCreated, channelCreateFailed));
  }

  // Callback for Successful Channel Creation
  const channelCreated = (conversation: Conversation) => {
    // Close the Dialog
    closeDialog("create-private-channel-dialog");

    // Call the Fetch Conversations Method
    dispatch(GetPrivateConversations);
  }

  // Callback for Failed Channel Creation
  const channelCreateFailed = (error: ErrorMessage) => {
    // Get the Error Box Element
    let errorBox = document.getElementById("create-private-channel-error") as HTMLElement;
    setError(error.message);
    errorBox.classList.remove("hidden");
  }

  // Render Method
  return(
    <ModalDialog id="create-private-channel-dialog" title="Create Private Channel" hideOnLoad={true} className="create-private-channel-dialog">
      <form action="/create-private-channel" method="post" onSubmit={createChannel}>
      <p id="create-private-channel-error" className="notification error hidden">{error}</p>
        <p>
          <label htmlFor="private-channel-name">Channel Name</label>
          <input type="text" id="private-channel-name" name="private_channel_name" autoComplete="Off" onChange={handleChannelName} required />
        </p>
        <p>Click User to add to the Channel</p>
        <ul>
          {userState.users?.map((user, i) => {
            return(
              <li key={user.id}>
                <input defaultChecked={isChecked(user)} id={"checkbox-" + user.id} type="checkbox" onChange={(e: React.ChangeEvent<HTMLInputElement>) => { toggleUser(e, user); }} />
                <label htmlFor={"checkbox-" + user.id}>{user.first_name} {user.last_name}</label>
              </li>
            )
          })}
        </ul>
        <p>
          <button className="primary">Create Channel</button>
        </p>
      </form>
    </ModalDialog>
  );
}

export default CreatePrivateChannelDialog;