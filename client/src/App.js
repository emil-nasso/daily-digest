import React, { Component } from 'react';
import './App.css';
import GraphQL from './Graphql';
import SourceAdder from './SourceAdder';
import Subscriptions from './Subscriptions';
import Digests from './Digests';

class App extends Component {

  constructor(props) {
    super(props);
    this.graphQL = new GraphQL();

    this.state = {
      message: 'hello world',
      sources: undefined,
      subscriptions: undefined
    };

    this.loadSources();
    this.loadSubscriptions();
  }

  loadSources(){
    this.graphQL.query(`query AllSources{
      sources {
        id
        name
        description
        tags
      }
    }`).then((data) => this.setState({sources: data.data.sources}));
  }
  
  loadSubscriptions() {
    this.graphQL.query(`query AllSubscriptions {
      subscriptions {
        id
        source{
          id
          name
          description
          tags
        }
      }
    }`).then((data) => this.setState({subscriptions: data.data.subscriptions}));
  }

  addSource(id) {
    this.graphQL.query(`mutation CreateDigest ($id: String!) {
      newSubscription(input: {
        sourceId: $id
      }) {
        id
      }
    }`, { id }).then(() => {this.loadSubscriptions()})
  }

  render() {
    return (
      <div className="App">
        <header className="App-header">
          <h1 className="App-title">Daily-Digest</h1>
        </header>
        <main>
          <SourceAdder sources={this.state.sources} addSourceCallback={this.addSource.bind(this)}/>
          <Subscriptions subscriptions={this.state.subscriptions}/>
          <Digests graphQL={this.graphQL}/>
        </main>
      </div>
    );
  }
}

export default App;
