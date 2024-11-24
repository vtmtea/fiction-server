import { Module } from "@nestjs/common";
import { TypeOrmModule } from "@nestjs/typeorm";
import { BookEntity } from "../entities/book.entity";
import { BookService } from "./book.service";
import { BookController } from "./book.controller";
import { CategoryEntity } from "../entities/category.entity";
import { SourceEntity } from "../entities/source.entity";

@Module({
	imports: [TypeOrmModule.forFeature([BookEntity, CategoryEntity, SourceEntity])],
	providers: [BookService],
	controllers: [BookController],
})
export class BookModule {}
