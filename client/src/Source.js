import React from 'react';

const Source = ({source, addSourceCallback}) => {
    return (
      <li>
        <strong>{source.name}</strong> - 
        {source.tags.map(t => (<span key={t}>[{t}]</span>))} -
        <button onClick={() => {addSourceCallback(source.id)}}>
          Add {source.name}
        </button>
        <p>
          {source.description}
        </p>
      </li>
    );
};

export default Source;

