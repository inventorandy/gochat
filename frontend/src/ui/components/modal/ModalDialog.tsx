import * as React from 'react';
import '../../scss/modal.scss';

type ModalDialogProps = {
  id: string;
  title: string;
  hidden: boolean;
  onClickClose: (e: React.MouseEvent<HTMLButtonElement>) => void;
} & React.PropsWithChildren;

const ModalDialog: React.FC<ModalDialogProps> = (props: ModalDialogProps) => {
  // Render
  return (
    <>
      <div
        id={props.id + '_overlay'}
        className={'modal-overlay' + (props.hidden ? ' hidden' : '')}
      />
      <div
        id={props.id}
        className={'modal-dialog' + (props.hidden ? ' hidden' : '')}
      >
        <div className='header'>
          <h4>{props.title}</h4>
          <button onClick={props.onClickClose}>&times;</button>
        </div>
        <div className='content'>{props.children}</div>
        <div className='footer' />
      </div>
    </>
  );
};

export default ModalDialog;
