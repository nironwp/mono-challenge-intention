import { ProductInterface } from '@core/domain/interfaces/product.interface';
import { ObjectType, Field, ID } from '@nestjs/graphql';

@ObjectType()
export class ProductModel implements ProductInterface {
  @Field((type) => ID)
  id: number;

  @Field()
  title: string;

  @Field()
  price: number;

  @Field()
  category: string;

  @Field()
  description: string;

  @Field()
  image: string;
}
