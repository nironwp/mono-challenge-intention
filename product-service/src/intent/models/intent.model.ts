import { IntentInterface } from '@core/domain/interfaces/intent.interface';
import { ObjectType, Field, ID } from '@nestjs/graphql';
import { ProductModel } from '../../products/models/product.model';

@ObjectType()
export class IntentModel implements IntentInterface {
  @Field(() => ID, { description: 'The ID of the intent' })
  user_id: string;

  @Field(() => [ProductModel], {
    description: 'The products associated with the intent',
  })
  products: ProductModel[];
}
