import { Injectable } from '@angular/core';

import { MeState } from './me.state';

@Injectable({ providedIn: 'root' })
export class State {
  readonly me$ = new MeState();
}
