import { Module } from "@nestjs/common";
import { AppController } from "./app.controller";
import { AppService } from "./app.service";
import { ConfigModule, ConfigService } from "@nestjs/config";
import { TypeOrmModule } from "@nestjs/typeorm";
import { BookModule } from "./book/book.module";

@Module({
	imports: [
		ConfigModule.forRoot({ isGlobal: true }),
		TypeOrmModule.forRootAsync({
			imports: [ConfigModule],
			inject: [ConfigService],
			useFactory: async (configService: ConfigService) => ({
				type: "mysql",
				host: configService.get<string>("DATABASE_HOST"),
				port: configService.get<number>("DATABASE_PORT"),
				username: configService.get<string>("DATABASE_USERNAME"),
				password: configService.get<string>("DATABASE_PASSWORD"),
				database: configService.get<string>("DATABASE_NAME"),
				autoLoadEntities: true,
			}),
		}),
		BookModule,
	],
	controllers: [AppController],
	providers: [AppService],
})
export class AppModule {}
