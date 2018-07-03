import React, { Component } from 'react';
import Source from './Source'

class SourceAdder extends Component {
  render() {
    if (this.props.sources === undefined) {
      return <div>Loading...</div>
    }

    return (
      <div>
        <h2>Sources</h2>
        <ul>
          {this.props.sources.map((s) => (<Source key={s.id} source={s} addSourceCallback={this.props.addSourceCallback}/>))}
        </ul>
      </div>
    );
  }
}

export default SourceAdder;
