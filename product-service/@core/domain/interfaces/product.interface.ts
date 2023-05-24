import { RootAgregateinterface } from "./root-agreggate.interface";

export interface ProductInterface extends RootAgregateinterface {
    id: number,
    title:string,
    price:number,
    category:string,
    description:string,
    image:string
}