/// <reference types="vite/client" />
import { ResultType as DataType, ListResultType as PageDataType } from "./api/universal_types"

declare module "url";

declare global {
    type ResultType<T> =  DataType<T>
    type ListResultType<T> =  PageDataType<T>
}