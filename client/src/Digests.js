import React, { Component } from 'react';
import Digest from './Digest'

class Digests extends Component {

  render() {
    if (this.props.digests === undefined) {
      return <div>Loading digests</div>
    }

    return (
      <div>
        <h2>Digests</h2>
        {this.props.digests.map((d) => (<Digest key={d.id} digest={d}/>))}
      </div>
    );
  }
}

export default Digests;
