import React, { Component } from 'react';
import bus from './eventService';

class Digests extends Component {
  render() {
    return (
      <div className="flex">
        <div className="pr-4">
          <h2>Digest</h2>
          <ul>
          <li><button onClick={() => {bus.emit('select-digests','2018-01-01')}}>2018-01-01</button></li>
          <li><button onClick={() => {bus.emit('select-digests','2018-01-02')}}>2018-01-02</button></li>
          <li><button onClick={() => {bus.emit('select-digests','2018-01-03')}}>2018-01-03</button></li>
          </ul>
        </div>
        <div>
          <h3>{this.props.selectedDigestsDate}</h3>
          {this.props.digests.map((digest)=> (<Digest key={digest.subscription.id} digest={digest}/>))}
        </div>
      </div>
    );
  }
}

function Digest({digest}) {
  return (
    <div>
      <h4>{digest.subscription.id}: {digest.subscription.source.name}</h4>
      {
        digest.entries.length === 0 ?
          <p>Nothing today!</p>
        :
          digest.entries.map((entry) => (<Entry key={entry.id} entry={entry}/>))
      }
    </div>
  );
}

function Entry({entry}) {
  return (
    <div>
      <a className="cursor-pointer" href={entry.url} target='_blank'>{entry.title}</a>
      <div>
        {entry.excerpt}
      </div>
    </div>
  );
}

export default Digests;
