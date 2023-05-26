import { Module } from '@nestjs/common';
import { AppController } from './app.controller';
import { ApolloDriver } from '@nestjs/apollo';
import { AppService } from './app.service';
import { ThrottlerModule } from '@nestjs/throttler';
import { ConfigModule } from '@nestjs/config';
import * as Joi from 'joi';
import { GraphQLModule } from '@nestjs/graphql';
import { DirectiveLocation, GraphQLDirective } from 'graphql';
import { ProductsModule } from './products/products.module';
import { IntentModule } from './intent/intent.module';
import { GraphqlService } from './graphql/graphql.service';

@Module({
  imports: [
    ThrottlerModule.forRoot({
      ttl: 60,
      limit: 10,
    }),
    GraphQLModule.forRoot({
      driver: ApolloDriver,
      autoSchemaFile: 'schema.gql',
      installSubscriptionHandlers: true,
      buildSchemaOptions: {
        directives: [
          new GraphQLDirective({
            name: 'upper',
            locations: [DirectiveLocation.FIELD_DEFINITION],
          }),
        ],
      },
    }),
    ProductsModule,
    IntentModule,
    ConfigModule.forRoot({
      isGlobal: true,
      validationSchema: Joi.object({
        PORT: Joi.number().required(),
        GRAPHQL_ENDPOINT: Joi.string().required(),
      })
    })
  ],
  controllers: [AppController],
  providers: [AppService, GraphqlService],
})
export class AppModule {}
