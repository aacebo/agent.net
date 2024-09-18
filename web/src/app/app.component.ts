import { ChangeDetectionStrategy, Component, ViewEncapsulation } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterOutlet } from '@angular/router';

import { CytoscapeModule } from '../components/cytoscape';
import { Api } from '../api';
import { State } from '../state';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [
    CommonModule,
    RouterOutlet,
    CytoscapeModule
  ],
  host: { class: 'app-root' },
  templateUrl: './app.component.html',
  styleUrl: './app.component.scss',
  encapsulation: ViewEncapsulation.None,
  changeDetection: ChangeDetectionStrategy.OnPush
})
export class AppComponent {
  constructor(
    readonly state: State,
    private readonly _api: Api
  ) { }

  async ngOnInit() {
    const agents = await this._api.agents.get();
    this.state.me$.set('agents', agents);
  }

  onSelect(nodes: any[]) {
    console.log(nodes);
  }
}
