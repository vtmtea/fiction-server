import { Column, Entity, JoinColumn, OneToOne } from "typeorm";
import { BaseEntity } from "./base.entity";
import { CategoryEntity } from "./category.entity";
import { SourceEntity } from "./source.entity";

@Entity("books")
export class BookEntity extends BaseEntity {
	@Column()
	name: string;

	@Column({ name: "category_id" })
	categoryId: number;

	@Column({ name: "author_id" })
	authorId: number;

	@Column({ name: "source_id" })
	sourceId: number;

	@Column({ name: "source_url" })
	sourceUrl: string;

	@Column({ name: "last_chapter_name" })
	lastChapterName: string;

	@Column({ name: "last_chapter_url" })
	lastChapterUrl: string;

	@Column({ name: "last_update_time" })
	lastUpdateTime: Date;

	@Column()
	description: string;

	@Column()
	cover: string;

	@Column({ name: "update_status" })
	updateStatus: string;

	@OneToOne(() => CategoryEntity, (category) => category.id)
	@JoinColumn({ name: "category_id" })
	category: CategoryEntity;

	@OneToOne(() => SourceEntity, (source) => source.id)
	@JoinColumn({ name: "source_id" })
	source: SourceEntity;
}
