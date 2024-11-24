import { Column, Entity } from "typeorm";
import { BaseEntity } from "./base.entity";

@Entity("source")
export class SourceEntity extends BaseEntity {
	@Column()
	name: string;

	@Column()
	domain: string;
}
