export interface NodeData {
  readonly id: string;
  readonly name: string;
  readonly version: string;
  readonly content: string;
  readonly size?: number;
  readonly outgoingEdges?: number;
  readonly weight: number;
  readonly fontSize: number;
}
