import React, { Component } from 'react';
import bus from './eventService';

class LoginPage extends Component {

  constructor(props){
    super(props)
    this.state = {
      username: '',
      password: ''
    }
  }

  handleChange(field, event) {
    let state = {}
    state[field] = event.target.value
    this.setState(state)
  }

  render() {
    // TODO: extract inputs and buttons to basic elements
    return (
      <main className="m-2">
        Username: <input className="border" value={this.state.username} onChange={(e) => this.handleChange('username', e)} type="text"/><br/>
        Password: <input className="border" value={this.state.password} onChange={(e) => this.handleChange('password', e)} type="password"/><br/>
        <button onClick={() => bus.emit('login', {username: this.state.username, password: this.state.password})} className="border rounded shadow px-4 py-2 m-2">
          Log in
        </button>
        <button onClick={() => bus.emit('register', {username: this.state.username, password: this.state.password})} className="border rounded shadow px-4 py-2 m-2">
          Register
        </button>
      </main>
    );
  }
}

export default LoginPage;
