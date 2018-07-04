import React, { Component } from 'react';

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

function Subscription({subscription}){
  return (
    <div>
      {subscription.id} - {subscription.source.name}
    </div>
  );
};

export default Subscriptions;
