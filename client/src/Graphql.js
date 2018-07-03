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
}

export default GraphQL;