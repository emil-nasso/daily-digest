import React from 'react';

const Subscription = ({subscription}) => {
    return (
      <div>
        {subscription.id} - {subscription.source.name}
      </div>
    );
};

export default Subscription;

