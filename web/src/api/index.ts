import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

import { Agents } from './agents';

@Injectable({ providedIn: 'root' })
export class Api {
  readonly agents: Agents;

  constructor(private readonly _http: HttpClient) {
    this.agents = new Agents(this._http);
  }
}
