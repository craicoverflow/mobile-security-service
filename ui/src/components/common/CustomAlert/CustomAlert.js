import React from 'react';
import './CustomAlert.css';
import { Alert, AlertActionCloseButton } from '@patternfly/react-core';

/**
 * A custom alert component that wraps the Patternfly 4 Alert.
 *
 * @param {Object} visible - The visible state of the alert.
 * @param {Object} onClose - The callback function to execute on close.
 * @param {Object} children - Any child props - usually the description.
 * @param {Object} props - The component props
 */
export const CustomAlert = ({ visible, onClose, children, ...props }) => {
  if (!visible) {
    return null;
  }

  return (
    <div className="Alert">
      <Alert
        {...props}
        action={<AlertActionCloseButton onClose={onClose} />}>
        {children}
      </Alert>
    </div>
  );
};
