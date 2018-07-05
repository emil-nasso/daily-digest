import React, { Component } from 'react';
import GraphQL from './Graphql';
import Sources from './Sources';
import Subscriptions from './Subscriptions';
import Digests from './Digests';
import bus from './eventService'

class App extends Component {

  constructor(props) {
    super(props);
    this.graphQL = new GraphQL();

    this.state = {
      message: 'hello world',
      sources: undefined,
      subscriptions: undefined,
      digests: [],
      selectedDigestsDate: undefined,
    };

    this.loadSources();
    this.loadSubscriptions();
  }

  componentDidMount(){
    bus.on('add-source', (id) => {
      this.createSource(id).then(() => {this.loadSubscriptions()});
    });
    bus.on('select-digests', (date) => {
      this.loadDigests(date);
    });
  }

  loadSources(){
    this.graphQL.loadSources().then((data) => this.setState({sources: data.data.sources}));
  }

  loadDigests(date) {
    this.setState({selectedDigestsDate: date})
    this.graphQL.loadDigests(date).then((result) => {
      this.setState({digests: result.data.digests});
    });
  }
  
  loadSubscriptions() {
    this.graphQL.loadSubscriptions().then((data) => this.setState({subscriptions: data.data.subscriptions}));
  }

  createSource(id) {
    return this.graphQL.createSource(id);
  }

  render() {
    return (
      <div className="App">
        <header className="bg-grey-darker p-8 mb-4 font-xl">
          <h1 className="App-title">Daily-Digest</h1>
        </header>
        <main className="m-2">
          <div className="flex border-b">
            <div>
              <Subscriptions subscriptions={this.state.subscriptions}/>
            </div>
            <div className="ml-8">
              <Sources sources={this.state.sources}/>
            </div>
          </div>
          <Digests selectedDate={this.state.selectedDigestsDate} digests={this.state.digests}/>
        </main>
      </div>
    );
  }
}

export default App;
