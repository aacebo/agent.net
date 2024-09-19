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
        position: a.position,
        data: {
          id: a.id,
          parent: a.parent_id,
          name: a.name,
          description: a.description,
          content: a.name,
          position: a.position,
          status: a.status,
          address: a.address,
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
      map(_ => {
        const edges: cytoscape.EdgeDefinition[] = [];

        // for (const agent of agents) {
        //   for (const child of agents.filter(a => a.id !== agent.id && a.parent_id === agent.parent_id)) {
        //     edges.push({
        //       group: 'edges',
        //       selectable: false,
        //       data: {
        //         id: `${child.id} -> ${agent.id}`,
        //         source: child.id,
        //         target: agent.id,
        //         weight: 1,
        //       }
        //     });
        //   }
        // }

        return edges;
      })
    );
  }
}
