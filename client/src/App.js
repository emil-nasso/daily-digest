import React, { Component } from 'react';
import GraphQL from './Graphql';
import Sources from './Sources';
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
        <header className="bg-grey-darker p-8 mb-4 font-xl">
          <h1 className="App-title">Daily-Digest</h1>
        </header>
        <main>
          <div className="flex border-b">
            <div className="mr-8">
              <Sources sources={this.state.sources} addSourceCallback={this.addSource.bind(this)}/>
            </div>
            <div>
              <Subscriptions subscriptions={this.state.subscriptions}/>
            </div>
          </div>
          <Digests graphQL={this.graphQL}/>
        </main>
      </div>
    );
  }
}

export default App;
