import { ChangeDetectionStrategy, Component, ViewEncapsulation } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterOutlet } from '@angular/router';
import { BehaviorSubject } from 'rxjs';

import { Api } from '../api';
import { State } from '../state';
import { CytoscapeModule, NodeData } from '../components/cytoscape';
import { IconModule } from '../components/icon';
import { DrawerModule } from '../components/drawer';
import { Agent } from '../models';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [
    CommonModule,
    RouterOutlet,
    CytoscapeModule,
    IconModule,
    DrawerModule
  ],
  host: { class: 'app-root' },
  templateUrl: './app.component.html',
  styleUrl: './app.component.scss',
  encapsulation: ViewEncapsulation.None,
  changeDetection: ChangeDetectionStrategy.OnPush
})
export class AppComponent {
  readonly agent$ = new BehaviorSubject<Agent | undefined>(undefined);

  constructor(
    readonly state: State,
    private readonly _api: Api
  ) { }

  async ngOnInit() {
    const agents = await this._api.agents.get();
    this.state.me$.set('agents', agents);
  }

  onSelect(nodes: NodeData[]) {
    if (nodes.length === 0) {
      this.agent$.next(undefined);
      return;
    }

    const node = nodes[0];
    const agents = this.state.me$.get('agents') || [];
    const agent = agents.find(a => a.id === node.id);
    this.agent$.next(agent);
  }

  async onPositionChange(node: NodeData) {
    const agent = await this._api.agents.update(node.id, {
      position: node.position
    });

    const agents = this.state.me$.get('agents') || [];
    const i = agents.findIndex(v => v.id === agent.id);

    if (i === -1) return;

    agents[i] = agent;
    this.state.me$.set('agents', agents);
  }
}
