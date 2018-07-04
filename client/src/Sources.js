import React, { Component } from 'react';

class Sources extends Component {

  constructor(props){
    super(props)
    this.state = {
      expanded : null
    };
  }

  toggleExpanded(id) {
    this.setState((prevState) => {
        if (id === prevState.expanded) {
          return {expanded: null};
        }
        return {expanded: id};
    })
  }

  render() {
    if (this.props.sources === undefined) {
      return <div>Loading...</div>
    }

    return (
      <div>
        <h2 className="mb-4">Sources</h2>
        <div>
          {this.props.sources.map((s) => (
              <Source
                key={s.id}
                source={s}
                expanded={s.id === this.state.expanded}
                addSourceCallback={this.props.addSourceCallback}
                expandCallback={this.toggleExpanded.bind(this)}
              />
          ))}
        </div>
      </div>
    );
  }
}

function Source({source, addSourceCallback, expanded, expandCallback}){
  let contents;
  if (expanded) {
    contents = (
      <div className="p-1">
        <p className="mb-2">
          {source.description}
        </p>
        <AddSourceButton source={source} addSourceCallback={addSourceCallback}/>
      </div>
    );
  }
  return (
    <div className="mb-4">
      <div className="border p-2 cursor-pointer" onClick={() => { expandCallback(source.id) }}>
        <strong>{source.name}</strong>
        {source.tags.map(t => (<SourceTag key={t} tag={t}/>))}
      </div>
      <div className="border">
        {contents}
      </div>
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
    <button className="border shadow rounded p-1 hover:bg-blue-dark hover:text-white" onClick={() => {addSourceCallback(source.id)}}>
        Add {source.name}
    </button>
  );
}

export default Sources;
