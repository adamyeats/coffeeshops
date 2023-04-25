import { Client, cacheExchange, fetchExchange } from 'urql';

export const client = new Client({
  url:
    process.env.REACT_APP_GRAPHQL_ENDPOINT ?? 'http://localhost:8080/graphql',
  exchanges: [cacheExchange, fetchExchange]
});
