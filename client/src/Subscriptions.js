import React, { Component } from 'react';
import Subscription from './Subscription'

class Subscriptions extends Component {

  render() {
    if (this.props.subscriptions === undefined) {
      return <div>Loading subscriptions</div>
    }

    return (
      <div>
        <h2>Subscriptions</h2>
        {this.props.subscriptions.map((sub) => (<Subscription key={sub.id} subscription={sub}/>))}
      </div>
    );
  }
}

export default Subscriptions;
