import { Injectable } from '@nestjs/common';
import { CreateIntentInput } from './dto/create-intent.input';
import axios from 'axios';
import { ProductInterface } from '@core/domain/interfaces/product.interface';
import { GraphQLClient } from 'graphql-request';
import { GraphqlService } from 'src/graphql/graphql.service';

@Injectable()
export class IntentService {

  constructor(private graphqlService: GraphqlService) {}

  async create(createIntentInput: CreateIntentInput) {
    const products = await Promise.all(
      createIntentInput.product_ids.map(async (product_id) => {
        return (
          await axios.get<ProductInterface>(
            'https://fakestoreapi.com/products/' + product_id,
          )
        ).data;
      }),
    );

    const intent = {
      user_id: createIntentInput.user_id,
      products,
    };

    const query = this.formatQueryCreateIntent(products, createIntentInput.user_id)    

    await this.graphqlService.request(
      query,
    )

    return intent;
  }

  private formatQueryCreateIntent(products: ProductInterface[], user_id: string) {
    const query = `
      mutation CreateIntent {
        createIntent(
          input: {
            user_id: "${user_id}",
            items: [
              ${products.map((product) => {
                return `{image: "${product.image}", title: "${product.title}", description: "${product.description}", price: ${product.price}, category: "${product.category}"}`
              }).join(',')}
            ]
          }
        ) {
          user_id
        }
      }
    `;
  
    return query;
  }
  
}
