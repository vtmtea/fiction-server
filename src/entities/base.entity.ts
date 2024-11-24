import { Column, Entity, PrimaryGeneratedColumn } from "typeorm";
import * as dayjs from "dayjs";

@Entity()
export class BaseEntity {
	@PrimaryGeneratedColumn()
	id: number;

	@Column({
		transformer: {
			from: (value: string) => dayjs(value).format("YYYY-MM-DD HH:mm:ss"),
			to: (value: Date) => dayjs(value).format("YYYY-MM-DD HH:mm:ss"),
		},
	})
	createdAt: Date;

	@Column({
		transformer: {
			from: (value: string) => dayjs(value).format("YYYY-MM-DD HH:mm:ss"),
			to: (value: Date) => dayjs(value).format("YYYY-MM-DD HH:mm:ss"),
		},
	})
	updatedAt: Date;

	@Column({ nullable: true })
	deletedAt: Date;
}
