import { InputType, ID, Field } from '@nestjs/graphql';

@InputType()
export class CreateIntentInput {
  @Field(() => ID, { description: 'The ID of the intent' })
  user_id: string;

  @Field((type) => [Number])
  product_ids: number[];
}
