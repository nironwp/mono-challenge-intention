import { Injectable } from '@nestjs/common';
import { ProductModel } from './models/product.model';
import axios from 'axios';
import { ProductInterface } from '@core/domain/interfaces/product.interface';
import { RecipesArgs } from './dto/products.args';

@Injectable()
export class ProductsService {
  async findAll(args: RecipesArgs): Promise<ProductModel[]> {
    const products = (
      await axios.get<ProductInterface[]>(
        'https://fakestoreapi.com/products?limit=' + args.take,
      )
    ).data;

    return products.slice(args.skip).map((product) => {
      const product_model: ProductModel = {
        ...product,
      };
      return product_model;
    });
  }
}
