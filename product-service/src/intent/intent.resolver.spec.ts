import { Test, TestingModule } from '@nestjs/testing';
import { IntentResolver } from './intent.resolver';
import { IntentService } from './intent.service';

describe('IntentResolver', () => {
  let resolver: IntentResolver;

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      providers: [IntentResolver, IntentService],
    }).compile();

    resolver = module.get<IntentResolver>(IntentResolver);
  });

  it('should be defined', () => {
    expect(resolver).toBeDefined();
  });

  it('Should Create a intent', async () => {
    const createIntent = {
      product_ids: [1, 2, 3],
      user_id: 'here-user-id',
    };
    const response = await resolver.createIntent(createIntent);

    expect(response).toBeDefined();
    expect(response.user_id).toBe(createIntent.user_id);
    expect(response.products.length).toBe(createIntent.product_ids.length);
  });
});
