import { Column, Entity } from "typeorm";
import { BaseEntity } from "./base.entity";

@Entity("category")
export class CategoryEntity extends BaseEntity {
	@Column()
	name: string;
}
