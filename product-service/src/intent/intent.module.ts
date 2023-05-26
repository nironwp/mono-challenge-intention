import { Module } from '@nestjs/common';
import { IntentService } from './intent.service';
import { IntentResolver } from './intent.resolver';
import { GraphqlService } from 'src/graphql/graphql.service';

@Module({
  providers: [IntentResolver, IntentService, GraphqlService],
})
export class IntentModule {}
