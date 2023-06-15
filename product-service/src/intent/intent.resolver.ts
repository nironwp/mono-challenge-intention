import { Resolver, Mutation, Args, } from '@nestjs/graphql';
import { IntentService } from './intent.service';
import { IntentModel } from './models/intent.model';
import { CreateIntentInput } from './dto/create-intent.input';

@Resolver(() => IntentModel)
export class IntentResolver {
  constructor(private readonly intentService: IntentService) {}

  @Mutation(() => IntentModel)
  createIntent(
    @Args('createIntentInput') createIntentInput: CreateIntentInput,
  ) {
    return this.intentService.create(createIntentInput);
  }
}
