import { requestClient } from '#/api/request';
import type { Category } from '#/api/core/workorder_category';
// 表单设计相关类型
export interface ListFormDesignReq {
  page: number;
  size: number;
  category_id?: number;
  status?: number;
  search?: string;
}

export interface DetailFormDesignReq {
  id: number;
}

export interface PublishFormDesignReq {
  id: number;
}

export interface CloneFormDesignReq {
  id: number;
  name: string;
}

export interface FormFieldOption {
  label: string;
  value: any;
}

export interface FormFieldValidation {
  min_length?: number;
  max_length?: number;
  min?: number;
  max?: number;
  pattern?: string;
  message?: string;
}

export interface FormField {
  id: string;
  type: string;
  label: string;
  name: string;
  required: boolean;
  placeholder?: string;
  default_value?: any;
  options?: FormFieldOption[];
  validation?: FormFieldValidation;
  props?: Record<string, any>;
  sort_order?: number;
}

export interface FormSchema {
  fields: FormField[];
  layout?: string;
  style?: string;
}

export interface FormDesignReq {
  id?: number;
  name: string;
  description: string;
  schema: FormSchema;
  category_id?: number;
}

export interface FormDesignResp {
  id: number;
  name: string;
  description: string;
  schema: FormSchema;
  version: number;
  status: number;
  category_id?: number;
  category?: Category;
  creator_id: number;
  creator_name: string;
  created_at: string;
  updated_at: string;
}

export interface FormDesignItem {
  id: number;
  name: string;
  description: string;
  version: number;
  status: number;
  category_id?: number;
  category?: Category;
  creator_id: number;
  creator_name: string;
  created_at: string;
  updated_at: string;
}

export interface PreviewFormDesignResp {
  id: number;
  schema: FormSchema;
}

export interface ValidateFormDesignResp {
  is_valid: boolean;
  errors?: string[];
}

// 表单设计相关接口
export async function createFormDesign(data: FormDesignReq) {
  return requestClient.post('/workorder/form-design/create', data);
}

export async function updateFormDesign(data: FormDesignReq) {
  return requestClient.put(`/workorder/form-design/update/${data.id}`, data);
}

export async function deleteFormDesign(data: DetailFormDesignReq) {
  return requestClient.delete(`/workorder/form-design/delete/${data.id}`);
}

export async function listFormDesign(data: ListFormDesignReq) {
  return requestClient.get('/workorder/form-design/list', { params: data });
}

export async function detailFormDesign(data: DetailFormDesignReq) {
  return requestClient.get(`/workorder/form-design/detail/${data.id}`);
}

export async function publishFormDesign(data: PublishFormDesignReq) {
  return requestClient.post(`/workorder/form-design/publish/${data.id}`);
}

export async function cloneFormDesign(data: CloneFormDesignReq) {
  return requestClient.post(`/workorder/form-design/clone/${data.id}`, data);
}

export async function previewFormDesign(data: DetailFormDesignReq) {
  return requestClient.get(`/workorder/form-design/preview/${data.id}`);
}
