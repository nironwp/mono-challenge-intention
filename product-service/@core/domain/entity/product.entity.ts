import { ProductInterface } from "../interfaces/product.interface";

export class ProductEntity implements ProductInterface {
    constructor(
      public id: number,
      public title: string,
      public price: number,
      public category: string,
      public description: string,
      public image: string
    ) {}
  
    getId(): number {
      return this.id;
    }

    getTitle(): string {
        return this.title;
    }

    getPrice(): number {
        return this.price;
    }

    getCategory(): string {
        return this.category;
    }

    getDescription(): string {
        return this.description;
    }

    getImage(): string {
        return this.image;
    }
}