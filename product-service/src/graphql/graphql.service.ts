import { Injectable } from '@nestjs/common';
import { GraphQLClient } from 'graphql-request';

@Injectable()
export class GraphqlService {
  graphqlClient: GraphQLClient;
  constructor() {
    this.graphqlClient = new GraphQLClient(process.env.GRAPHQL_ENDPOINT);
  }

    async request(query: string, variables?: any) {
        return this.graphqlClient.request(query, variables);
    }
}
