import { Resolver, Query, Mutation, Args, Int } from '@nestjs/graphql';
import { ProductsService } from './products.service';
import { ProductModel } from './models/product.model';
import { RecipesArgs } from './dto/products.args';

@Resolver(() => ProductModel)
export class ProductsResolver {
  constructor(private readonly productsService: ProductsService) {}

  @Query(() => [ProductModel], { name: 'products' })
  findAll(@Args() recipesArgs: RecipesArgs) {
    return this.productsService.findAll(recipesArgs);
  }
}
