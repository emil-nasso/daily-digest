import React, { Component } from 'react';

class Sources extends Component {
  render() {
    if (this.props.sources === undefined) {
      return <div>Loading...</div>
    }

    return (
      <div>
        <h2 className="mb-4">Sources</h2>
        <div>
          {this.props.sources.map((s) => (<Source key={s.id} source={s} addSourceCallback={this.props.addSourceCallback}/>))}
        </div>
      </div>
    );
  }
}

function Source({source, addSourceCallback}){
  return (
    <div className="mb-4">
      <strong>{source.name}</strong> - 
      {source.tags.map(t => (<SourceTag tag={t}/>))} - <AddSourceButton source={source} addSourceCallback={addSourceCallback}/>
      <p>
        {source.description}
      </p>
    </div>
  );
};


function SourceTag({tag}){
  return (
    <span className="bg-blue-light border-blue border-1 rounded p-1 m-1">{tag}</span>
  );
}

function AddSourceButton({source, addSourceCallback}){
  return (
    <button className="border shadow rounded p-1" onClick={() => {addSourceCallback(source.id)}}>
        Add {source.name}
    </button>
  );
}

export default Sources;
