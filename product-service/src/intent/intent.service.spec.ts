import { Test, TestingModule } from '@nestjs/testing';
import { IntentService } from './intent.service';
import { CreateIntentInput } from './dto/create-intent.input';

describe('IntentService', () => {
  let service: IntentService;

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      providers: [IntentService],
    }).compile();

    service = module.get<IntentService>(IntentService);
  });

  it('should be defined', () => {
    expect(service).toBeDefined();
  });

  it("Should Create a intent", async () => {
    const createIntent: CreateIntentInput = {
      product_ids: [1, 2, 3],
      user_id: 'here-user-id'
    }
    const response = await service.create(createIntent)

    expect(response).toBeDefined()
    expect(response.user_id).toBe(createIntent.user_id)
    expect(response.products.length).toBe(createIntent.product_ids.length)
  })
});
