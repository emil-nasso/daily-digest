import React from 'react';

const Digest = ({digest,}) => {
    return (
      <div>
        {digest.id} - {digest.source.name}
      </div>
    );
};

export default Digest;

