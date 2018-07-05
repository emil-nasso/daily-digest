class GraphQL {
    url =  'http://localhost:8080/graphql'

    query(query, variables){
        return fetch(this.url, {
            method: 'POST',
            headers: {
              'Content-Type': 'application/json',
              'Accept': 'application/json',
            },
            body: JSON.stringify({query, variables})
          })
            .then(r => r.json());
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
}

export default GraphQL;