import { Controller, Get, Param } from "@nestjs/common";
import { BookEntity } from "../entities/book.entity";
import { BookService } from "./book.service";

@Controller("book")
export class BookController {
	constructor(private readonly bookService: BookService) {}

	@Get(":id")
	async findById(@Param("id") id: number): Promise<BookEntity> {
		return await this.bookService.findById(id);
	}
}
