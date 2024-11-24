import { NestFactory } from "@nestjs/core";
import { AppModule } from "./app.module";
import { ConfigService } from "@nestjs/config";
import { Single } from "./pkg/crawler/single";

async function bootstrap() {
	const app = await NestFactory.create(AppModule);
	const configService = app.get(ConfigService);
	await app.listen(configService.get<number>("PORT") ?? 3000);
}

bootstrap();

Single("https://www.beqege.cc/16647/");
