import * as React from 'react';

// Dialog Properties
type CreateChannelDialogProps = {
	isPublic: boolean;
}

// Dialog Component
const CreateChannelDialog: React.FC<CreateChannelDialogProps> = (props: CreateChannelDialogProps) => {
  return(
    <div className="modal-dialog">
      <div className="header">
        <h3>{props.isPublic ? "Create Channel" : "Create Private Chat"}</h3>
      </div>
    </div>
  );
}