import { Component, OnInit, ViewEncapsulation, ChangeDetectionStrategy, ElementRef, Input, OnDestroy, Output, EventEmitter } from '@angular/core';
import { coerceNumberProperty } from '@angular/cdk/coercion';
import cytoscape from 'cytoscape';

import { debounce } from '../../utils';

import { LAYOUT } from './layout';
import { NodeData } from './node-data';
import { STYLES } from './styles';

@Component({
  selector: 'cytoscape',
  template: '',
  styleUrls: ['./cytoscape.component.scss'],
  host: { class: 'cytoscape' },
  encapsulation: ViewEncapsulation.None,
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class CytoscapeComponent implements OnInit, OnDestroy {
  @Input()
  get nodes() { return this._nodes; }
  set nodes(v) {
    this._nodes = v;
    this.render();
  }
  private _nodes: cytoscape.NodeDefinition[] = [];

  @Input()
  get edges() { return this._edges; }
  set edges(v) {
    this._edges = v;
    this.render();
  }
  private _edges: cytoscape.EdgeDefinition[] = [];

  @Input()
  get zoom() { return this._zoom; }
  set zoom(v) {
    this._zoom = coerceNumberProperty(v);
    this.zoomChange.emit(v);
  }
  private _zoom = 0.6;

  @Output() zoomChange = new EventEmitter<number>();
  @Output() nodesSelect = new EventEmitter<NodeData[]>();

  private _graph?: cytoscape.Core;
  private readonly _runZoom = debounce(() => this.zoom = this._graph?.zoom() || 0, 500);

  constructor(private readonly _el: ElementRef<HTMLElement>) { }

  ngOnInit() {
    this._graph = cytoscape({
      container: this._el.nativeElement,
      style: STYLES,
      layout: LAYOUT,
      selectionType: 'single',
      zoom: this._zoom,
      boxSelectionEnabled: true,
      elements: {
        nodes: this.nodes,
        edges: this.edges,
      }
    });

    setTimeout(() => {
      this._graph?.layout(LAYOUT).run();
    }, 500);

    this._graph.on('zoom', () => {
      this._runZoom();
    });

    this._graph.on('select', debounce((e: cytoscape.EventObject) => {
      this.nodesSelect.emit(e.cy.nodes(':selected').map(n => n.data()));
    }, 100));

    this._graph.on('unselect', debounce((e: cytoscape.EventObject) => {
      this.nodesSelect.emit(e.cy.nodes(':selected').map(n => n.data()));
    }, 100));
  }

  ngOnDestroy() {
    this._graph?.destroy();
  }

  center() {
    this._graph?.center().fit();
  }

  clear() {
    this._graph?.elements().remove();
  }

  layout() {
    this._graph?.layout(LAYOUT).run();
  }

  image() {
    return this._graph?.png({ output: 'base64' });
  }

  goTo(id: string) {
    const node = this._graph?.$id(id);
    this._graph?.center(node).fit(node);
  }

  render() {
    this._graph?.json({
      elements: {
        nodes: this._nodes,
        edges: this._edges,
      }
    });
  }
}
