import { ProductInterface } from "./product.interface"

export interface IntentInterface {
    user_id: string
    products: ProductInterface[]
}