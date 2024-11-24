import { InjectRepository } from "@nestjs/typeorm";
import { BookEntity } from "../entities/book.entity";
import { Repository } from "typeorm";
import { Injectable } from "@nestjs/common";

@Injectable()
export class BookService {
	constructor(@InjectRepository(BookEntity) private bookRepository: Repository<BookEntity>) {}

	findById(bookId: number): Promise<BookEntity | null> {
		return this.bookRepository.findOne({ where: { id: bookId }, relations: ["category", "source"] });
	}
}
