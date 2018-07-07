import React, { Component } from 'react';
import GraphQL from './Graphql';
import Sources from './Sources';
import Subscriptions from './Subscriptions';
import LoginPage from './LoginPage';
import Digests from './Digests';
import bus from './eventService';

class App extends Component {

  constructor(props) {
    super(props);

    this.storage = window.localStorage
    let sessionKey = this.storage.getItem("sessionKey");
    this.graphQL = new GraphQL(sessionKey);

    this.state = {
      message: 'hello world',
      sources: undefined,
      subscriptions: undefined,
      digests: [],
      selectedDigestsDate: undefined,
      sessionKey,
      page: 'settings',
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
    bus.on('login', (credentials) => {
      console.log('login', credentials);
    });
    bus.on('register', (credentials) => {
      this.graphQL.registerUser(credentials.username, credentials.password).then((data) => {
        let sessionKey = data.data.register;
        this.graphQL = new GraphQL(sessionKey);
        this.loadSubscriptions();
        this.storage.setItem("sessionKey", sessionKey);
        this.setState({
          sessionKey,
          page: "settings"
        })
      }, this.handleError.bind(this))
    });
  }

  handleError(reason) {
    // Check if it's actually a access denied error
      console.log(reason)
      console.log(reason.details)
      if (reason.details[0] && reason.details[0].message === "Access denied"){
        this.setState({
          page: "registration"
        })
      } else {
        alert("Unknown error. Logged to console.");
      }
  }

  loadSources(){
    this.graphQL.loadSources()
    .then((data) => {
        this.setState(
          {sources: data.data.sources}
        )
    }, this.handleError.bind(this));
  }

  loadDigests(date) {
    this.setState({selectedDigestsDate: date})
    this.graphQL.loadDigests(date).then((result) => {
      this.setState({digests: result.data.digests});
    }, this.handleError.bind(this));
  }
  
  loadSubscriptions() {
    this.graphQL.loadSubscriptions().then(
      (data) => this.setState({subscriptions: data.data.subscriptions}),
      this.handleError.bind(this)
    );
  }

  createSource(id) {
    return this.graphQL.createSource(id);
  }

  render() {
    let contents = '';

    switch(this.state.page) {
      case "settings":
        contents = <SettingsPage subscriptions={this.state.subscriptions} sources={this.state.sources} selectedDigestsDate={this.state.selectedDigestsDate} digests={this.state.digests}/>;
        break;
        case "registration":
          contents = <LoginPage/>;
          break;
        default:
          break;
    }

    return (
      <div className="App">
        <header className="bg-grey-darker p-8 mb-4 font-xl">
          <h1 className="App-title">Daily-Digest</h1>
        </header>
        {contents}
      </div>
    );
  }
}

function SettingsPage({subscriptions,sources,selectedDigestsDate, digests}){
  return (
    <main className="m-2">
      <div className="flex border-b">
        <div>
          <Subscriptions subscriptions={subscriptions}/>
        </div>
        <div className="ml-8">
          <Sources sources={sources}/>
        </div>
      </div>
      <Digests selectedDate={selectedDigestsDate} digests={digests}/>
    </main>
  );
}

export default App;
