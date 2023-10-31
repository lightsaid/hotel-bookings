// NOTE: 通用类型定义

// 分页数据基础信息
export type MetaType = {
    current_page: number;
    page_size: number;
    first_page: number;
    last_page: number;
    total_records: number;
}

// 非列表请求通用结构, NOTE: 在d.ts中设为全局类型
export type ResultType<T> = {
	result: T;
	msg: string;
	code: number;
}

// 列表请求通用结构， NOTE: 在d.ts中设为全局类型
export type ListResultType<T> = {
	list: T[];
	meta: MetaType;
	msg: string;
	code: number;
}

