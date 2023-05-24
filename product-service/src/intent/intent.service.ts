import { Injectable } from '@nestjs/common';
import { CreateIntentInput } from './dto/create-intent.input';
import axios from 'axios';
import { ProductInterface } from '@core/domain/interfaces/product.interface';


@Injectable()
export class IntentService {
  async create(createIntentInput: CreateIntentInput) {
    const products =await Promise.all(createIntentInput.product_ids.map(async (product_id) => {
      return (await axios.get<ProductInterface>('https://fakestoreapi.com/products/'+product_id)).data
    }))

    const intent = {
      user_id: createIntentInput.user_id,
      products
    }

    // TODO - send intent to another service

    return intent
  }
}
