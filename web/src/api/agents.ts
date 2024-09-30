import { HttpClient } from '@angular/common/http';
import { firstValueFrom } from 'rxjs';

import { Agent } from '../models';

interface CreateParams {
  readonly name: string;
  readonly description?: string;
  readonly instructions?: string;
  readonly position?: {
    readonly x: number;
    readonly y: number;
  };
  readonly settings: {
    readonly api_key: string;
    readonly model: string;
    readonly frequency_penalty?: number;
    readonly logit_bias?: Record<string, any>;
    readonly logprobs?: boolean;
  };
}

interface UpdateParams {
  readonly description?: string;
  readonly instructions?: string;
  readonly position?: {
    readonly x: number;
    readonly y: number;
  };
  readonly settings?: {
    readonly api_key: string;
    readonly model: string;
    readonly frequency_penalty?: number;
    readonly logit_bias?: Record<string, any>;
    readonly logprobs?: boolean;
  };
}

export class Agents {
  constructor(private readonly _http: HttpClient) { }

  get() {
    return firstValueFrom(this._http.get<Agent[]>('/agents'));
  }

  getChildren(id: string) {
    return firstValueFrom(this._http.get<Agent[]>(`/agents/${id}/agents`));
  }

  create(params: CreateParams): Promise<Agent>;
  create(parentId: string, params: CreateParams): Promise<Agent>;
  create(idOrParams: string | CreateParams, params?: CreateParams) {
    const id = typeof idOrParams === 'string' ? idOrParams : undefined;
    const body = typeof idOrParams === 'string' ? params : idOrParams;
    let url = '/agents';

    if (id) {
      url = `/agents/${id}/agents`;
    }

    return firstValueFrom(this._http.post<Agent>(url, body));
  }

  update(id: string, params: UpdateParams) {
    return firstValueFrom(this._http.patch<Agent>(`/agents/${id}`, params));
  }

  delete(id: string) {
    return firstValueFrom(this._http.delete<void>(`/agents/${id}`));
  }

  start(id: string) {
    return firstValueFrom(this._http.post<void>(`/agents/${id}/start`, null));
  }

  stop(id: string) {
    return firstValueFrom(this._http.post<void>(`/agents/${id}/stop`, null));
  }
}
