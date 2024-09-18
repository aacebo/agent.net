import cytoscape from 'cytoscape';
import { map } from 'rxjs';

import { Agent } from '../models';
import { StateProperty } from './state-property';

export interface _MeState {
  agents?: Agent[];
}

export class MeState extends StateProperty<_MeState> {
  constructor(value: _MeState = { }) {
    super(value);
  }

  get nodes$() {
    return this.pipe(
      map(state => state.agents || []),
      map(agents => agents.map(a => ({
        group: 'nodes',
        selectable: true,
        data: {
          ...a,
          content: a.description,
          weight: 50,
          size: (agents.filter(v => v.parent_id === a.id).length * 2) || 1,
          fontSize: 15,
          outgoingEdges: agents.filter(v => v.id === a.parent_id).length,
        }
      } as cytoscape.NodeDefinition)))
    );
  }

  get edges$() {
    return this.pipe(
      map(state => state.agents || []),
      map(agents => {
        const edges: cytoscape.EdgeDefinition[] = [];

        for (const agent of agents) {
          for (const child of agents.filter(a => a.parent_id === agent.id)) {
            edges.push({
              group: 'edges',
              selectable: false,
              data: {
                id: `${child.id} -> ${agent.id}`,
                source: child.id,
                target: agent.id,
                weight: 1,
              }
            });
          }
        }

        return edges;
      })
    );
  }
}
