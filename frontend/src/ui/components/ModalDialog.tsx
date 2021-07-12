import * as React from 'react';
import { library } from '@fortawesome/fontawesome-svg-core';
import { faTimes } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';

// Add the Required Icons to the Library
library.add( faTimes );

// Open Dialog Method
export const openDialog = (id: string) => {
  let dialog = document.getElementById(id);
    let overlay = document.getElementById(id + "_overlay");
    dialog?.classList.remove("hidden");
    overlay?.classList.remove("hidden");
}

// Open Dialog Method
export const closeDialog = (id: string) => {
  let dialog = document.getElementById(id);
    let overlay = document.getElementById(id + "_overlay");
    dialog?.classList.add("hidden");
    overlay?.classList.add("hidden");
}

// Dialog Properties
type ModalDialogProps = {
  id: string;
  title: string;
  className?: string;
  hideOnLoad?: boolean;
}

// Dialog Component
const ModalDialog: React.FC<React.PropsWithChildren<ModalDialogProps>> = (props: React.PropsWithChildren<ModalDialogProps>) => {
  // Determine the CSS Class
  let cssClass = "modal-dialog";
  if (props.className !== undefined) {
    cssClass += " " + props.className;
  }

  // Determine whether the Dialog is hidden on page load
  if (props.hideOnLoad === true) {
    cssClass += " hidden";
  }

  // Render Method
  return(
    <>
    <div className={props.hideOnLoad === true ? "modal-overlay hidden" : "modal-overlay"} id={props.id + "_overlay"} />
    <div className={cssClass} id={props.id}>
      <div className="header">
        <h4>{props.title}</h4>
        <button onClick={(e) => closeDialog(props.id)}>
          <FontAwesomeIcon icon="times" />
        </button>
      </div>
      <div className="content">
        {props.children}
      </div>
      <div className="footer" />
    </div>
    </>
  );
}

export default ModalDialog;