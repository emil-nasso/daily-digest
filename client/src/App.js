import React, { Component } from 'react';
import './App.css';
import GraphQL from './Graphql';
import SourceAdder from './SourceAdder';
import Digests from './Digests'

class App extends Component {

  constructor(props) {
    super(props);
    this.graphQL = new GraphQL();

    this.state = {
      message: 'hello world',
      sources: undefined,
      digests: undefined
    };

    this.loadSources();
    this.loadDigests();
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
  
  loadDigests() {
    this.graphQL.query(`query AllDigests {
      digests {
        id
        source{
          id
          name
          description
          tags
        }
      }
    }`).then((data) => this.setState({digests: data.data.digests}));
  }

  addSource(id) {
    this.graphQL.query(`mutation CreateDigest ($id: String!) {
      newDigest(input: {
        sourceId: $id
      }) {
        id
        source{
          id
          name
          description
          tags
        }
      }
    }`, { id }).then(() => {this.loadDigests()})
  }

  render() {
    return (
      <div className="App">
        <header className="App-header">
          <h1 className="App-title">Daily-Digest</h1>
        </header>
        <main>
          <SourceAdder sources={this.state.sources} addSourceCallback={this.addSource.bind(this)}/>
          <Digests digests={this.state.digests}/>
        </main>
      </div>
    );
  }
}

export default App;
