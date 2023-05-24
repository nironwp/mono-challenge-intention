import { Module } from '@nestjs/common';
import { IntentService } from './intent.service';
import { IntentResolver } from './intent.resolver';

@Module({
  providers: [IntentResolver, IntentService]
})
export class IntentModule {}
