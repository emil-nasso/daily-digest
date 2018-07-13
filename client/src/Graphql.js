class GraphQL {
    constructor(sessionKey) {
      this.sessionKey = sessionKey;
    }
    url =  'http://localhost:8080/graphql'

    query(query, variables){
        return fetch(this.url, {
            method: 'POST',
            headers: {
              'Content-Type': 'application/json',
              'Accept': 'application/json',
              'Authorization': `Bearer ${this.sessionKey}`
            },
            body: JSON.stringify({query, variables})
          })
            .then(r => r.json())
            .then(r => {
                if (r.errors) {
                  let error = new Error("GraphQL error")
                  error.details = r.errors
                  throw error;
                }
                return r;
            })
          ;
    }

    loadSources(){
      return this.query(`query AllSources{
        sources {
          id
          name
          description
          tags
        }
      }`);
    }

    loadSubscriptions() {
      return this.query(`query AllSubscriptions {
        subscriptions {
          id
          source{
            id
            name
            description
            tags
          }
        }
      }`);
    }

    loadDigests(date) {
      return this.query(
        `query GetDigest($date: String!) {
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
        `,
        {date}
      );
    }

    createSource(id) {
      return this.query(`mutation CreateDigest ($id: String!) {
        newSubscription(input: {
          sourceId: $id
        }) {
          id
        }
      }`, { id })
    }

    registerUser(username, password) {
      return this.query(`mutation Register ($username: String!, $password: String!){
        register(input:{username: $username, password: $password})
      }`, { username, password })
    }

    login(username, password) {
      return this.query(`mutation Login ($username: String!, $password: String!){
        login(input:{username: $username, password: $password})
      }`, { username, password })
    }
}

export default GraphQL;