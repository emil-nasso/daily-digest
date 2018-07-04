import React, { Component } from 'react';
import Source from './Source'

class Digests extends Component {
  constructor(props) {
    super(props);
    this.state = {
      date: undefined,
      digests: []
    };
  }
  loadDigests(date) {
    let query = `query GetDigest($date: String!) {
      digests(date:$date){
        subscription{
          id
          source {
            id
            name
            description
          }
        }
        entries{
          id
          publishedAt
          title
          excerpt
          url
        }
      }
    }
    `;

    this.setState({date})
    this.props.graphQL.query(query, {date}).then((result) => {
      this.setState({digests: result.data.digests});
    });
  }
  render() {
    return (
      <div>
        <h2>Digest</h2>
        <ul>
         <li><button onClick={() => {this.loadDigests('2018-01-01')}}>2018-01-01</button></li>
         <li><button onClick={() => {this.loadDigests('2018-01-02')}}>2018-01-02</button></li>
         <li><button onClick={() => {this.loadDigests('2018-01-03')}}>2018-01-03</button></li>
        </ul>
        <h3>{this.state.date}</h3>
        {this.state.digests.map((digest)=> (<Digest key={digest.subscription.id} digest={digest}/>))}
      </div>
    );
  }
}

function Digest({digest}) {
  return (
    <div>
      <h4>{digest.subscription.id}: {digest.subscription.source.name}</h4>
      {
        digest.entries.length == 0 ?
          <p>Nothing today!</p>
        :
          digest.entries.map((entry) => (<Entry key={entry.id} entry={entry}/>))
      }
    </div>
  );
}

function Entry({entry}) {
  return (
    <div>
      <a href={entry.url} target='_blank'>{entry.title}</a>
      <div>
        {entry.excerpt}
      </div>
    </div>
  );
}

export default Digests;
